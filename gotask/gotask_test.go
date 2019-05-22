package gotask

import (
	"fmt"
	"testing"
)

var gotask = NewGoTask(1000, false)

func TestNewGoTask(t *testing.T) {

	gotask.Add(compute, Student{Name: "spw", Age: 26})

	gotask.Start()
}

type Student struct {
	Name string
	Age  int
}

func compute(data ...interface{}) {

	stu := gotask.GetParamter(0, data).(Student)

	fmt.Println(stu.Name)
	fmt.Println(stu.Age)

}
