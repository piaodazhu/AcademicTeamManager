package model

type Summary struct {
	StudentNum int `json:"student_num"`
	ProjectNum int `json:"project_num"`
	OutputNum  int `json:"output_num"`
	StudentPan []*PanArg `json:"student_pan"`
	ProjectPan []*PanArg `json:"project_pan"`
	OutputPan []*PanArg `json:"output_pan"`
}

type PanArg struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}