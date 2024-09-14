package domain

import "fmt"

type Category struct {
	Name        string `json:"name"`
	EasyWords   []Word `json:"easy"`
	MediumWords []Word `json:"medium"`
	HardWords   []Word `json:"hard"`
}

type BadCategoryError struct {
	Message string
}

func (e *BadCategoryError) Error() string {
	return fmt.Sprintf("bad category: %s", e.Message)
}
