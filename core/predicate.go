package core

func True(v interface{}) bool {
	return v != false
}

func False(v interface{}) bool {
	return v == false
}
