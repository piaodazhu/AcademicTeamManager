package model

type Summary struct {
	StudentNum  int       `json:"student_num"`
	ProjectNum  int       `json:"project_num"`
	OutputNum   int       `json:"output_num"`
	StudentPan  []*PanArg `json:"student_pan"`
	ProjectPan  []*PanArg `json:"project_pan"`
	OutputPan   []*PanArg `json:"output_pan"`
	ProgressBar []*BarArg `json:"progress_bar"`
}

type PanArg struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type BarArg struct {
	Name  string  `gorm:"column:name" json:"name"`
	Value float64 `gorm:"column:finishrate" json:"value"`
}
