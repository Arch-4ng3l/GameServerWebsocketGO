package types

var RenderDistance float32 = 100.0

type Object struct {
	Name string  `json:"name"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}

type Setting struct {
	RenderDistance float32 `json:"nums"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ObjectAmount struct {
	Nums uint `json:"nums"`
}
