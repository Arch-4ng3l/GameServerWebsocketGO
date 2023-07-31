package types

var RenderDistance float32 = 100.0

type Object struct {
	Name string  `json:"name"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}

type Setting struct {
	RenderDistance float32 `json:"renderdistance"`
}
