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

// Fat Task Data that could accumulate data from any state (everything that could be stored
type State struct {
    Comment string `json:",omitempty"`
    Type TaskType
    Resource string `json:",omitempty"`
    Next string `json:",omitempty"`
    End bool `json:",omitempty"`

    Retry Retry `json:",omitempty"`
    Catch Catch `json:",omitempty"`

    Branches []StateMachine

    // Data Related to Lambda Execution
    // Input path is a JSON from raw JSON that task process and pass to lambda
    InputPath interface{} `json:",omitempty"`
	// Output path is a JSON from raw JSON that task gets from Lambda
    OutputPath interface{} `json:",omitempty"`
    ResultPath interface{} `json:",omitempty"`
    Parameters interface{} `json:",omitempty"`
}

