package main

import (
	"fmt"
	"time"
)

func main() {
	metrics := &Metrics{}

	err := Execute(func() error {
		fmt.Println("Executing...")
		return nil
	}, metrics)

	fmt.Println(err)
}

func Execute(f func() error, metrics *Metrics) (err error) {

	metrics.StoreExecution()

	err1 := f()
	if err1 != nil {
		metrics.StoreFailure()
	} else {
		metrics.StoreSuccess()
	}
	return err1
}

type Metrics struct {
	execution []time.Time
	success   []time.Time
	failure   []time.Time
}

func (m *Metrics) StoreExecution() {
	m.execution = append(m.execution, time.Now())
}

func (m *Metrics) StoreSuccess() {
	m.success = append(m.success, time.Now())
}

func (m *Metrics) StoreFailure() {
	m.failure = append(m.failure, time.Now())
}
