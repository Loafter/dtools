package main

import "log"
import "errors"
import "sync"

type JobDispatcher interface {
	Dispatch(interface{}) (error, interface{})
}

type ErrorDispatcher interface {
	DispatchError(*FailedJob) error
}

type JobBallancer struct {
	inJobChan       chan interface{}
	activeJob       map[string]interface{}
	errorDispatcher ErrorDispatcher
	jobDispatcher   JobDispatcher
	waitJobDone     sync.WaitGroup
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
			jobBallancer.addJob(job.JobId, job.JobData)
			go jobBallancer.startJob(job)
			log.Println("info: normal dispatch")
		case DoneJob:
			log.Println("info: try remove task id=" + job.JobId)
			err := jobBallancer.removeJob(job.JobId)
			if err == nil {
				log.Println("info: successul remove task id=" + job.JobId)
			} else {
				log.Println("error: faled remove task with err " + job.JobId)
			}
			jobBallancer.waitJobDone.Done()
		case FailedJob:
			err := jobBallancer.removeJob(job.JobId)
			if err == nil {
				log.Println("info: successul remove failed task id=" + job.JobId)
			} else {
				log.Println("error: faled remove failed task id=" + job.JobId)
			}
			jobBallancer.errorDispatcher.DispatchError(&job)
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

//add job
func (jobBallancer *JobBallancer) addJob(id string, job interface{}) error {
	if jobBallancer.activeJob == nil {
		return errors.New("error: job list is null")
	}
	jobBallancer.activeJob[id] = job
	return nil
}

//check if work confilct
func (jobBallancer *JobBallancer) isConflictedJob(taskData interface{}) bool {
	if _, ok := taskData.(IsVerifiable); !ok {
		errors.New("warning: this task date is not verifiable")
		return false
	}
	for _, job := range jobBallancer.activeJob {
		if ver, ok := job.(IsVerifiable); !ok {
			if ver.IsConflict(ver) {
				return true
			}
		}
	}
	return false
}

func (jobBallancer *JobBallancer) PushJob(job Job) error {
	if jobBallancer.inJobChan == nil {
		return errors.New("error: JobChan is not inited")
	}
	jobBallancer.inJobChan <- job
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

func (jobBallancer *JobBallancer) Init(jobDispatcher JobDispatcher, errorDispatcher ErrorDispatcher) {
	jobBallancer.errorDispatcher = errorDispatcher
	jobBallancer.jobDispatcher = jobDispatcher
	jobBallancer.activeJob = make(map[string]interface{})
	jobBallancer.inJobChan = make(chan interface{})
	go jobBallancer.takeJob()
}
