package language

import (
	//"fmt"
)

type TaskType string

const(
	Task TaskType = "Task"
	Pass = "Pass"
	Choice = "Choice"
	Wait = "Wait"
	Succeed = "Succeed"
	Fail = "Fail"
	Parallel = "Parallel"
	Map = "Map"
)

type TaskData struct {
    Comment *string
    Type TaskType
    Resource *string
    Next *string
    End bool

    // Data Related to Lambda Execution
    // Input path is a JSON from raw JSON that task process and pass to lambda
    InputPath interface{}
	// Output path is a JSON from raw JSON that task gets from Lambda
    OutputPath interface{}
    ResultPath interface{}
    Parameters interface{}
}

