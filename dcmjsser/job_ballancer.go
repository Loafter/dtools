package main

import "log"
import "errors"
import "sync"
import "crypto/rand"
import "fmt"

type JobDispatcher interface {
	Dispatch(interface{}) (interface{}, error)
}

type ErrDispatcher interface {
	DispatchError(FaJob) error
}

type CompDispatcher interface {
	DispatchSuccess(CompJob) error
}

type JobBallancer struct {
	jChan    chan interface{}
	acJob    map[string]Job
	slJob    map[string]Job
	errDisp  ErrDispatcher
	jobDisp  JobDispatcher
	compDisp CompDispatcher
	JbDone   sync.WaitGroup
	aJobC    int
}

func (jbal JobBallancer) checkInit() bool {
	return (jbal.acJob == nil) && (jbal.jChan == nil) && (jbal.slJob == nil)
}

func (jbal *JobBallancer) startJob(jdat interface{}) {
	job := jdat.(Job)
	dispResult, err := jbal.jobDisp.Dispatch(job.Data)
	if err != nil {
		log.Println("info: failed job detected", err)
		jbal.jChan <- FaJob{Job: job, ErrorData: err}
	} else {
		log.Printf("info: completed job detected %v", dispResult)
		jbal.jChan <- CompJob{Job: job, ResultData: dispResult}
	}

}

func (jbal *JobBallancer) takeJob() {
	for {
		//extract job from queue
		recivedTask := <-jbal.jChan
		log.Println("info: job taken")
		switch job := recivedTask.(type) {
		case TermJob:
			//if we recive terminate signal need return
			log.Println("info: recive terminate dispatch singal")
			return
		case Job:
			//regular dispath
			if len(jbal.acJob) < jbal.aJobC {
				jbal.JbDone.Add(1)
				jbal.addActiveJob(job)
				go jbal.startJob(job)
				log.Println("info: normal dispatch")
			} else {
				jbal.addSleepJob(job)
				jbal.JbDone.Add(1)
				log.Println("info: attend maximum active job")
			}
		case CompJob:
			//notify about sucess
			if err := jbal.compDisp.DispatchSuccess(job); err != nil {
				log.Println("error: failed dispatch success" + job.Job.JobId)
			}
			//remove success compleated job
			jbal.removeJob(job.Job.JobId)
			jbal.JbDone.Done()
			jbal.resumeJobs()
		case FaJob:
			//notify about sucess
			if err := jbal.errDisp.DispatchError(job); err != nil {
				log.Println("error: failed dispatch error" + job.Job.JobId)
			}
			//remove success compleated job
			jbal.removeJob(job.Job.JobId)
			jbal.JbDone.Done()
			jbal.resumeJobs()
		default:
			log.Fatalln("error: unknown job type")
			jbal.JbDone.Done()
		}
	}
}

//remove successul complited job
func (jbal *JobBallancer) removeJob(jid string) error {
	if _, isFind := jbal.acJob[jid]; isFind {
		delete(jbal.acJob, jid)
	} else {
		return errors.New("error: can't remove job because job with id not found")
	}
	return nil
}

//remove successul complited job
func (jbal *JobBallancer) getJobByID(jid string) (*Job, error) {

	if val, isFind := jbal.acJob[jid]; isFind {
		return &val, nil
	} else {
		return nil, errors.New("error: can't find job with id")
	}
}

//add job
func (jbal *JobBallancer) addActiveJob(job Job) error {
	if jbal.checkInit() {
		return errors.New("error: job list not inited")
	}
	jbal.acJob[job.JobId] = job
	return nil
}
func (jbal *JobBallancer) addSleepJob(job Job) error {
	if jbal.checkInit() {
		return errors.New("error: job list not inited")
	}
	jbal.slJob[job.JobId] = job
	return nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func (jbal *JobBallancer) resumeJobs() {
	jbc := min(jbal.aJobC-len(jbal.acJob), len(jbal.slJob))
	for val := range jbal.slJob {
		if jbc > 0 {
			jb := jbal.slJob[val]
			delete(jbal.slJob, val)
			jbal.JbDone.Done()
			go jbal.PushJob(jb.Data)
		} else {
			return
		}
		jbc--
	}

}
func (jbal *JobBallancer) PushJob(jdat interface{}) error {
	if jbal.checkInit() {
		return errors.New("error: JobChan is not inited")
	}
	uid := genUid()
	job := Job{JobId: uid, Data: jdat}
	jbal.jChan <- job
	return nil

}

func (jbal *JobBallancer) TerminateTakeJob() error {
	if jbal.checkInit() {
		return errors.New("error: is not inited")
	}
	jbal.JbDone.Wait()
	jbal.jChan <- TermJob{}
	close(jbal.jChan)
	if len(jbal.acJob) > 0 {
		return errors.New("error: list job is not empty")
	}
	log.Println("info: greacefully terminate take job")
	return nil
}

func (jbal *JobBallancer) Init(jdis JobDispatcher, cmd CompDispatcher, erd ErrDispatcher) {
	jbal.errDisp = erd
	jbal.jobDisp = jdis
	jbal.compDisp = cmd
	jbal.acJob = make(map[string]Job)
	jbal.slJob = make(map[string]Job)
	jbal.aJobC = 10
	jbal.jChan = make(chan interface{})
	go jbal.takeJob()
	log.Println("info: job ballancer inited")
}

func (jbal JobBallancer) GetJobsList() ([]string, error) {
	if jbal.checkInit() {
		return nil, errors.New("error: job ballancer is not inited")
	}
	descr := []string{}

	for i := range jbal.acJob {
		if de, ok := jbal.acJob[i].Data.(Descriptable); ok {
			descr = append(descr, "active job: "+de.GetDescript())
		} else {
			descr = append(descr, "active job: this type job don't have description")
		}

	}
	for i := range jbal.slJob {
		if de, ok := jbal.slJob[i].Data.(Descriptable); ok {
			descr = append(descr, "sleeping job: "+de.GetDescript())
		} else {
			descr = append(descr, "sleeping job: this type job don't have description")
		}
	}
	return descr, nil
}

func genUid() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
