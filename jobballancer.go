package main

import "log"
import "errors"

type JobDispatcher interface {
	Dispatch(taskResponse TaskResponse)
}

type ErrorDispatcher interface {
	NotifyError(TaskError)
	DispatchError(error)
}

type JobBallancer struct {
	inJobChan       chan interface{}
	workedJob       map[string]Tasker
	errorDispatcher ErrorDispatcher
	jobDispatcher   JobDispatcher
}

func (jobBallancer *JobBallancer) takeJob() {
	for {
		//extract job from queue
		recivedTask := <-jobBallancer.inJobChan
		switch task := recivedTask.(type) {
		case TerminateDispatchJob:
			//if we recive recive terminate signal return
			log.Println("info: recive terminate dispatch singal")
			return
		case TaskResponse:
			//regular dispath
			jobBallancer.jobDispatcher.Dispatch(task)
			log.Println("info: recive terminate dispatch singal")
		case Tasker:
			log.Println("info: try remove task id=" + task.TaskId())
			err := jobBallancer.removeJob(task)
			if err != nil {
				log.Println("info: successul remove task id=" + task.TaskId())
			} else {
				log.Println(err)
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
	_, isFind := jobBallancer.workedJob[tasker.TaskId()]
	if isFind {
		delete(jobBallancer.workedJob, tasker.TaskId())
	} else {
		errors.New("error: can't remove task id=" + tasker.TaskId())
	}
	return nil
}

//check if work confilct
func (jobBallancer *JobBallancer) isConfilctedJob(tasker Tasker) bool {
	for _, job := range jobBallancer.workedJob {
		if job.(Tasker).IsConflict(tasker) {
			return false
		}

	}
	return true
}

func (jobBallancer *JobBallancer) PushJob(taskResponse TaskResponse) {
	jobBallancer.inJobChan <- taskResponse
}
