package gotask

import (
	"sync"
	"time"
)

/*
* GoTask is a high-performance concurrent goroutine manager written in golang.
* For this case, GoTask reduce your cpu rate and memory and make your project more efficency.
 */
type GoTask struct {
	wg         sync.WaitGroup //wait for all gorotines finished
	tasks      []GoTaskDetail //tasks
	maxTaskNum int            //max count of gorotine
	quickMode  bool           //if set to quick mode, jobs will be executed when added
	curTaskNum int
}

/*
* GoTaskDetail defines methods
 */
type GoTaskDetail struct {
	fn     func(...interface{})
	params interface{}
}

/*
* Generate GoTask manager
 */
func NewGoTask(maxTaskNum int, quickMode bool) *GoTask {
	ret := &GoTask{
		wg:         sync.WaitGroup{},
		tasks:      make([]GoTaskDetail, 0),
		maxTaskNum: maxTaskNum,
		quickMode:  quickMode,
		curTaskNum: 0,
	}
	return ret
}

/*
* Add tasks
 */
func (self *GoTask) Add(task func(...interface{}), params ...interface{}) {
	if !self.quickMode {
		self.tasks = append(self.tasks, GoTaskDetail{
			fn:     task,
			params: params,
		})
	} else {
		go func(v GoTaskDetail) {
			for {
				if self.curTaskNum < self.maxTaskNum {
					break
				}
				// time.Sleep(time.Millisecond * 50)
			}
			self.wg.Add(1)
			self.curTaskNum++
			defer func() {
				self.wg.Done()
				self.curTaskNum--
			}()
			v.fn(v.params)
		}(GoTaskDetail{fn: task, params: params})
	}
}

/*
* Get paramters from context
 */
func (self *GoTask) GetParamter(index int, params interface{}) interface{} {
	if p, ok := params.([]interface{}); ok {
		if len(p) > 0 {
			if t, ok := p[0].([]interface{}); ok {
				if len(t) > index {
					return t[index]
				}
			}
		}
	}
	return nil
}

/*
* Start concurrent tasks
 */
func (self *GoTask) Start() {
	if self.quickMode {
		self.wg.Wait()
		return
	}
	curTaskNum := 0
	for _, v := range self.tasks {
		self.wg.Add(1)
		curTaskNum++
		go func(v GoTaskDetail) {
			defer func() {
				self.wg.Done()
				curTaskNum--
			}()
			v.fn(v.params)
		}(v)
		for {
			if curTaskNum < self.maxTaskNum {
				break
			}
			time.Sleep(time.Millisecond * 50)
		}
	}
	self.wg.Wait()
}

/*
* if set quickMode == true, you must invoke Done() to finish manually. Deprecated
 */
func (self *GoTask) Done() {
	if self.quickMode == false {
		return
	}
	self.wg.Done()
	return
}
