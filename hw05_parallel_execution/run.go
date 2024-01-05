package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChannel := make(chan Task)
	resultChannel := make(chan error)
	doneChannel := make(chan interface{})

	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go execTaskInWorker(wg, taskChannel, resultChannel, doneChannel)
	}

	isLimitExceeded := isErrorsLimitExceeded(tasks, taskChannel, resultChannel, m)

	close(taskChannel)
	close(doneChannel)
	wg.Wait()

	if isLimitExceeded {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func execTaskInWorker(wg *sync.WaitGroup, taskCh <-chan Task, resultCh chan<- error, doneCh <-chan interface{}) {
	defer wg.Done()
	for task := range taskCh {
		result := task()
		select {
		case <-doneCh:
		case resultCh <- result:
		}
	}
}

func isErrorsLimitExceeded(tasks []Task, taskCh chan<- Task, resultCh <-chan error, m int) bool {
	var countErrors = 0
	for i := 0; i < len(tasks); {
		task := tasks[i]

		select {
		case err := <-resultCh:
			if m > 0 && err != nil {
				countErrors++
				if countErrors > m {
					return true
				}
			}
		case taskCh <- task:
			i++
		}
	}

	return false
}
