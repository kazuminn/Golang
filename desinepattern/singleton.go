package singleton

type singleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	return instance
}
