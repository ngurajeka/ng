package ng

//Data standard structure of data from jsonapi.org
type Data struct {
	Data  interface{} `json:"data"`
	Links interface{} `json:"links"`
}
