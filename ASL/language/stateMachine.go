package language

type StateMachine struct {
	Comment string
	StartAt string
	Tasks []TaskData
}

type StateMachineExecutionResult struct {
	Status bool
	Error* error
	ResultData interface{}
}

// this is a micro machine that is used for store Branches Items
type MicroStateMachine struct {

}

////////////////////////////////////////////////////// PUBLIC FUNCTIONS ///////////////////////////////////////////////

type StateMachineExecuter interface{
	// Execute from JSON: get, validate and execute
    ExecuteViaJson(json *string) StateMachineExecutionResult
    // Execute already loaded state machine
    ExecuteViaSm(stateMachine *StateMachine) StateMachineExecutionResult
}

// We are passing path to json file
func Load(file *string) (*StateMachine, error) {
	stateMachine := StateMachine{}
	stateMachine.StartAt = "Empty"
	// todo: umv: impl
	return &stateMachine, nil
}

// validate JSON
func ValidateJson(json *string) (bool, []error) {
	return true, nil
}

// validate JSON
func ValidateSm(json *string) (bool, []error) {
	return true, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////// PRIVATE FUNCTIONS //////////////////////////////////////////////
func execute(stateMachine *StateMachine) StateMachineExecutionResult {
	result := StateMachineExecutionResult{}
	return result
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////