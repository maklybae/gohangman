package domain

type Category struct {
	Name        string   `json:"name"`
	EasyWords   []string `json:"easy"`
	MediumWords []string `json:"medium"`
	HardWords   []string `json:"hard"`
}
