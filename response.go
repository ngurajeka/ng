package ng

//NewResponse create new response from an interface,
//return the interface or create an errorResponse.
//otherwise, not implemented yet
func NewResponse(v interface{}) interface{} {
	var response interface{}
	switch v.(type) {
	case Errors:
		response = ErrorResponse(v)
	case Data:
		response = v
		// default:
		//TODO We need to do something bud!
	}
	return response
}

func ErrorResponse(v interface{}) MapString {
	return MapString{
		"status": MapString{
			"error": v,
		},
	}
}
