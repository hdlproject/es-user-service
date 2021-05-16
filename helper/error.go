package helper

type RequestAlreadyProcessed struct {
	Message string
}

func (instance *RequestAlreadyProcessed) Error() string {
	return instance.Message
}
