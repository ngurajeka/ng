package ng

//Data standard structure of data from jsonapi.org
type Data struct {
	Data     interface{} `json:"data"`
	Included interface{} `json:"included,omitempty"`
	Links    interface{} `json:"links"`
	Meta     interface{} `json:"meta,omitempty"`
}
