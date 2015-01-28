package main

import "log"
import "errors"
import "sync"
import "crypto/rand"
import "fmt"

type JobDispatcher interface {
	Dispatch(interface{}) (interface{}, error)
}

type ErrorDispatcher interface {
	DispatchError(*FailedJob, interface{}) error
}

type CompletedDispatcher interface {
	DispatchSuccess(*CompletedJob, interface{}) error
}
type JobBallancer struct {
	inJobChan           chan interface{}
	activeJob           map[string]Job
	errorDispatcher     ErrorDispatcher
	jobDispatcher       JobDispatcher
	completedDispatcher CompletedDispatcher
	waitJobDone         sync.WaitGroup
}

func (jobBallancer *JobBallancer) startJob(jobd interface{}) {
	err, dispResult := jobBallancer.jobDispatcher.Dispatch(jobd)
	if err != nil {
		jobBallancer.inJobChan <- err
	} else {
		jobBallancer.inJobChan <- dispResult
	}

}

func (jobBallancer *JobBallancer) takeJob() {
	for {
		//extract job from queue
		recivedTask := <-jobBallancer.inJobChan
		log.Println("info: job taken")
		switch job := recivedTask.(type) {
		case TerminateDispatchJob:
			//if we recive terminate signal need return
			log.Println("info: recive terminate dispatch singal")
			return
		case Job:
			//regular dispath
			jobBallancer.waitJobDone.Add(1)
			jobBallancer.addJob(job)
			go jobBallancer.startJob(job)
			log.Println("info: normal dispatch")
		case CompletedJob:
			log.Println("info: try remove task id=" + job.JobId)
			err := jobBallancer.removeJob(job.JobId)
			if err == nil {
				log.Println("info: successul remove task id=" + job.JobId)
			} else {
				log.Println("error: faled remove task with err " + job.JobId)
			}
			getJob, err := jobBallancer.getJobByID(job.JobId)
			if err == nil {
				jobBallancer.completedDispatcher.DispatchSuccess(&job, getJob)
			} else {
				log.Println("error: failed dispatch success" + job.JobId)
			}
			jobBallancer.waitJobDone.Done()
		case FailedJob:
			err := jobBallancer.removeJob(job.JobId)
			if err == nil {
				log.Println("info: successul remove failed task id=" + job.JobId)
			} else {
				log.Println("error: faled remove failed task id=" + job.JobId)
			}
			getJob, err := jobBallancer.getJobByID(job.JobId)
			if err == nil {
				jobBallancer.errorDispatcher.DispatchError(&job, getJob)
			} else {
				log.Println("error: failef dispatch error" + job.JobId)
			}
			jobBallancer.waitJobDone.Done()

		default:
			log.Println("error: unknown job type")
		}
	}
}

//remove successul complited job
func (jobBallancer *JobBallancer) removeJob(jobId string) error {
	if _, isFind := jobBallancer.activeJob[jobId]; isFind {
		delete(jobBallancer.activeJob, jobId)
	} else {
		return errors.New("error: can't remove job because job with id not found")
	}
	return nil
}

//remove successul complited job
func (jobBallancer *JobBallancer) getJobByID(jobId string) (*Job, error) {
	if val, isFind := jobBallancer.activeJob[jobId]; isFind {
		return &val, nil
	} else {
		return nil, errors.New("error: can't find job with id")
	}
}

//add job
func (jobBallancer *JobBallancer) addJob(job Job) error {
	if jobBallancer.activeJob == nil {
		return errors.New("error: job list is null")
	}
	jobBallancer.activeJob[job.JobId] = job
	return nil
}

//check if work confilct
func (jobBallancer *JobBallancer) isConflictedJob(taskData interface{}) bool {
	/*if _, ok := taskData.(IsVerifiable); !ok {
		errors.New("warning: this task date is not verifiable")
		return false
	}
	for _, job := range jobBallancer.activeJob {
		if ver, ok := job.(IsVerifiable); !ok {
			if ver.IsConflict(ver) {
				return true
			}
		}
	}*/
	return false
}

func (jobBallancer *JobBallancer) PushJob(jobData Job, dataToDispatchSuccess interface{}, dataToDispatchError interface{}) error {
	if jobBallancer.inJobChan == nil {
		return errors.New("error: JobChan is not inited")
	}
	jobBallancer.inJobChan <- Job{JobId: genUid(), Data: jobData}
	return nil
}

func (jobBallancer *JobBallancer) TerminateTakeJob() error {
	if jobBallancer.inJobChan == nil {
		return errors.New("error: is not inited")
	}
	jobBallancer.waitJobDone.Wait()
	jobBallancer.inJobChan <- TerminateDispatchJob{}

	close(jobBallancer.inJobChan)
	if len(jobBallancer.activeJob) > 0 {
		return errors.New("error: list job is not empty")
	}

	log.Println("info: greacefully terminate take job")
	return nil
}

func (jobBallancer *JobBallancer) Init(jobDispatcher JobDispatcher, errorDispatcher ErrorDispatcher, completedDispatcher CompletedDispatcher) {
	jobBallancer.errorDispatcher = errorDispatcher
	jobBallancer.jobDispatcher = jobDispatcher
	jobBallancer.completedDispatcher = completedDispatcher
	jobBallancer.activeJob = make(map[string]Job)
	jobBallancer.inJobChan = make(chan interface{})
	go jobBallancer.takeJob()
}

func genUid() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}