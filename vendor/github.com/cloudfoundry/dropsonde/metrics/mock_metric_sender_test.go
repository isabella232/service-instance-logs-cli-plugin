// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package metrics_test

import (
	"github.com/cloudfoundry/dropsonde/metric_sender"
	"github.com/cloudfoundry/sonde-go/events"
)

type mockMetricSender struct {
	SendCalled chan bool
	SendInput  struct {
		Event chan events.Event
	}
	SendOutput struct {
		Ret0 chan error
	}
	ValueCalled chan bool
	ValueInput  struct {
		Name  chan string
		Value chan float64
		Unit  chan string
	}
	ValueOutput struct {
		Ret0 chan metric_sender.ValueChainer
	}
	ContainerMetricCalled chan bool
	ContainerMetricInput  struct {
		AppID     chan string
		Instance  chan int32
		Cpu       chan float64
		Mem, Disk chan uint64
	}
	ContainerMetricOutput struct {
		Ret0 chan metric_sender.ContainerMetricChainer
	}
	CounterCalled chan bool
	CounterInput  struct {
		Name chan string
	}
	CounterOutput struct {
		Ret0 chan metric_sender.CounterChainer
	}
	SendValueCalled chan bool
	SendValueInput  struct {
		Name  chan string
		Value chan float64
		Unit  chan string
	}
	SendValueOutput struct {
		Ret0 chan error
	}
	IncrementCounterCalled chan bool
	IncrementCounterInput  struct {
		Name chan string
	}
	IncrementCounterOutput struct {
		Ret0 chan error
	}
	AddToCounterCalled chan bool
	AddToCounterInput  struct {
		Name  chan string
		Delta chan uint64
	}
	AddToCounterOutput struct {
		Ret0 chan error
	}
	SendContainerMetricCalled chan bool
	SendContainerMetricInput  struct {
		ApplicationId chan string
		InstanceIndex chan int32
		CpuPercentage chan float64
		MemoryBytes   chan uint64
		DiskBytes     chan uint64
	}
	SendContainerMetricOutput struct {
		Ret0 chan error
	}
}

func newMockMetricSender() *mockMetricSender {
	m := &mockMetricSender{}
	m.SendCalled = make(chan bool, 100)
	m.SendInput.Event = make(chan events.Event, 100)
	m.SendOutput.Ret0 = make(chan error, 100)
	m.ValueCalled = make(chan bool, 100)
	m.ValueInput.Name = make(chan string, 100)
	m.ValueInput.Value = make(chan float64, 100)
	m.ValueInput.Unit = make(chan string, 100)
	m.ValueOutput.Ret0 = make(chan metric_sender.ValueChainer, 100)
	m.ContainerMetricCalled = make(chan bool, 100)
	m.ContainerMetricInput.AppID = make(chan string, 100)
	m.ContainerMetricInput.Instance = make(chan int32, 100)
	m.ContainerMetricInput.Cpu = make(chan float64, 100)
	m.ContainerMetricInput.Mem = make(chan uint64, 100)
	m.ContainerMetricInput.Disk = make(chan uint64, 100)
	m.ContainerMetricOutput.Ret0 = make(chan metric_sender.ContainerMetricChainer, 100)
	m.CounterCalled = make(chan bool, 100)
	m.CounterInput.Name = make(chan string, 100)
	m.CounterOutput.Ret0 = make(chan metric_sender.CounterChainer, 100)
	m.SendValueCalled = make(chan bool, 100)
	m.SendValueInput.Name = make(chan string, 100)
	m.SendValueInput.Value = make(chan float64, 100)
	m.SendValueInput.Unit = make(chan string, 100)
	m.SendValueOutput.Ret0 = make(chan error, 100)
	m.IncrementCounterCalled = make(chan bool, 100)
	m.IncrementCounterInput.Name = make(chan string, 100)
	m.IncrementCounterOutput.Ret0 = make(chan error, 100)
	m.AddToCounterCalled = make(chan bool, 100)
	m.AddToCounterInput.Name = make(chan string, 100)
	m.AddToCounterInput.Delta = make(chan uint64, 100)
	m.AddToCounterOutput.Ret0 = make(chan error, 100)
	m.SendContainerMetricCalled = make(chan bool, 100)
	m.SendContainerMetricInput.ApplicationId = make(chan string, 100)
	m.SendContainerMetricInput.InstanceIndex = make(chan int32, 100)
	m.SendContainerMetricInput.CpuPercentage = make(chan float64, 100)
	m.SendContainerMetricInput.MemoryBytes = make(chan uint64, 100)
	m.SendContainerMetricInput.DiskBytes = make(chan uint64, 100)
	m.SendContainerMetricOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockMetricSender) Send(event events.Event) error {
	m.SendCalled <- true
	m.SendInput.Event <- event
	return <-m.SendOutput.Ret0
}
func (m *mockMetricSender) Value(name string, value float64, unit string) metric_sender.ValueChainer {
	m.ValueCalled <- true
	m.ValueInput.Name <- name
	m.ValueInput.Value <- value
	m.ValueInput.Unit <- unit
	return <-m.ValueOutput.Ret0
}
func (m *mockMetricSender) ContainerMetric(appID string, instance int32, cpu float64, mem, disk uint64) metric_sender.ContainerMetricChainer {
	m.ContainerMetricCalled <- true
	m.ContainerMetricInput.AppID <- appID
	m.ContainerMetricInput.Instance <- instance
	m.ContainerMetricInput.Cpu <- cpu
	m.ContainerMetricInput.Mem <- mem
	m.ContainerMetricInput.Disk <- disk
	return <-m.ContainerMetricOutput.Ret0
}
func (m *mockMetricSender) Counter(name string) metric_sender.CounterChainer {
	m.CounterCalled <- true
	m.CounterInput.Name <- name
	return <-m.CounterOutput.Ret0
}
func (m *mockMetricSender) SendValue(name string, value float64, unit string) error {
	m.SendValueCalled <- true
	m.SendValueInput.Name <- name
	m.SendValueInput.Value <- value
	m.SendValueInput.Unit <- unit
	return <-m.SendValueOutput.Ret0
}
func (m *mockMetricSender) IncrementCounter(name string) error {
	m.IncrementCounterCalled <- true
	m.IncrementCounterInput.Name <- name
	return <-m.IncrementCounterOutput.Ret0
}
func (m *mockMetricSender) AddToCounter(name string, delta uint64) error {
	m.AddToCounterCalled <- true
	m.AddToCounterInput.Name <- name
	m.AddToCounterInput.Delta <- delta
	return <-m.AddToCounterOutput.Ret0
}
func (m *mockMetricSender) SendContainerMetric(applicationId string, instanceIndex int32, cpuPercentage float64, memoryBytes uint64, diskBytes uint64) error {
	m.SendContainerMetricCalled <- true
	m.SendContainerMetricInput.ApplicationId <- applicationId
	m.SendContainerMetricInput.InstanceIndex <- instanceIndex
	m.SendContainerMetricInput.CpuPercentage <- cpuPercentage
	m.SendContainerMetricInput.MemoryBytes <- memoryBytes
	m.SendContainerMetricInput.DiskBytes <- diskBytes
	return <-m.SendContainerMetricOutput.Ret0
}
