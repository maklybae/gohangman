package domain

type Category struct {
	Name        string `json:"name"`
	EasyWords   []Word `json:"easy"`
	MediumWords []Word `json:"medium"`
	HardWords   []Word `json:"hard"`
}
