package main

import "log"
import "errors"
import "sync"

type JobDispatcher interface {
	Dispatch(TaskResponse, chan interface{}) error
}

type ErrorDispatcher interface {
	NotifyError(TaskError) error
	DispatchError(error) error
}

type JobBallancer struct {
	inJobChan       chan interface{}
	activeJob       map[string]Tasker
	errorDispatcher ErrorDispatcher
	jobDispatcher   JobDispatcher
	waitJobDone     sync.WaitGroup
}

func (jobBallancer *JobBallancer) takeJob() {
	jobBallancer.waitJobDone.Add(1)
	defer jobBallancer.waitJobDone.Done()
	for {
		//extract job from queue
		log.Println("info: start wait take job")
		recivedTask := <-jobBallancer.inJobChan
		log.Println("info: job taken")
		switch task := recivedTask.(type) {
		case TerminateDispatchJob:
			//if we recive recive terminate signal return
			log.Println("info: recive terminate dispatch singal")
			return
		case TaskResponse:
			//regular dispath
			jobBallancer.addJob(task.Tasker)
			jobBallancer.jobDispatcher.Dispatch(task, jobBallancer.inJobChan)
			log.Println("info: normal dispatch")
		case Tasker:
			log.Println("info: try remove task id=" + task.TaskId())
			err := jobBallancer.removeJob(task)
			if err == nil {
				log.Println("info: successul remove task id=" + task.TaskId())
			} else {
				log.Println("error: faled remove task with err " + err.Error())
			}
		case TaskError:
			log.Println("error: " + task.Error.Error())
			jobBallancer.errorDispatcher.NotifyError(task)
		case error:
			log.Println("error: " + task.Error())
			jobBallancer.errorDispatcher.DispatchError(task)
		default:
			log.Println("error: unknown type")
		}
	}
}

//remove successul complited job
func (jobBallancer *JobBallancer) removeJob(tasker Tasker) error {
	if _, isFind := jobBallancer.activeJob[tasker.TaskId()]; isFind {
		delete(jobBallancer.activeJob, tasker.TaskId())
	} else {
		return errors.New("error: can't remove job because job with id not found")
	}
	return nil
}

//add job
func (jobBallancer *JobBallancer) addJob(tasker Tasker) error {
	if jobBallancer.activeJob == nil {
		return errors.New("error: job list is null")
	}
	jobBallancer.activeJob[tasker.TaskId()] = tasker
	return nil
}

//check if work confilct
func (jobBallancer *JobBallancer) isConflictedJob(tasker Tasker) bool {
	for _, job := range jobBallancer.activeJob {
		if job.(Tasker).IsConflict(tasker) {
			return true
		}

	}
	return false
}

func (jobBallancer *JobBallancer) PushJob(taskResponse TaskResponse) error {
	if jobBallancer.inJobChan == nil {
		return errors.New("error: is not inited")
	}
	jobBallancer.inJobChan <- taskResponse
	return nil
}

func (jobBallancer *JobBallancer) TerminateTakeJob() error {
	if jobBallancer.inJobChan == nil {
		return errors.New("error: is not inited")
	}
	jobBallancer.inJobChan <- TerminateDispatchJob{}
	if len(jobBallancer.activeJob) > 0 {
		log.Println("warning: list job is not empty")
	}
	jobBallancer.waitJobDone.Wait()
	log.Println("infor: greacefully terminate take job")
	return nil
}

func (jobBallancer *JobBallancer) Init(jobDispatcher JobDispatcher, errorDispatcher ErrorDispatcher) {
	jobBallancer.errorDispatcher = errorDispatcher
	jobBallancer.jobDispatcher = jobDispatcher
	jobBallancer.activeJob = make(map[string]Tasker)
	jobBallancer.inJobChan = make(chan interface{})
	go jobBallancer.takeJob()
}
