package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID  int
	Job func() error
}

type Worker struct {
	ID        int
	TaskQueue chan Task
	wg        *sync.WaitGroup
}

func NewWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) Worker {
	return Worker{
		ID:        id,
		TaskQueue: taskQueue,
		wg:        wg,
	}
}

func (w Worker) Start() {
	go func() {
		for task := range w.TaskQueue {
			fmt.Printf("Worker %d starting task %d\n", w.ID, task.ID)
			err := task.Job()
			if err != nil {
				fmt.Printf("Worker %d failed to execute task %d: %s\n", w.ID, task.ID, err)
			} else {
				fmt.Printf("Worker %d finished task %d\n", w.ID, task.ID)
			}
			w.wg.Done()
		}
	}()
}

type WorkerPool struct {
	TaskQueue chan Task
	wg        *sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	taskQueue := make(chan Task)
	wg := &sync.WaitGroup{}

	wp := &WorkerPool{
		TaskQueue: taskQueue,
		wg:        wg,
	}

	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, taskQueue, wg)
		worker.Start()
	}

	return wp
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.wg.Add(1)
	wp.TaskQueue <- task
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.TaskQueue)
}

func main() {
	numWorkers := 5
	workerPool := NewWorkerPool(numWorkers)

	for i := 1; i <= 10; i++ {
		task := Task{
			ID: i,
			Job: func() error {
				time.Sleep(1 * time.Second)
				return nil
			},
		}
		workerPool.AddTask(task)
	}

	workerPool.Wait()
	fmt.Println("Все задачи выполнены и программа завершена")
}
