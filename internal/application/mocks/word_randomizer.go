// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	domain "makly/hangman/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// WordRandomizer is an autogenerated mock type for the WordRandomizer type
type WordRandomizer struct {
	mock.Mock
}

type WordRandomizer_Expecter struct {
	mock *mock.Mock
}

func (_m *WordRandomizer) EXPECT() *WordRandomizer_Expecter {
	return &WordRandomizer_Expecter{mock: &_m.Mock}
}

// ChoiceWord provides a mock function with given fields: category, difficulty
func (_m *WordRandomizer) ChoiceWord(category *domain.Category, difficulty domain.Difficulty) (*domain.Word, error) {
	ret := _m.Called(category, difficulty)

	if len(ret) == 0 {
		panic("no return value specified for ChoiceWord")
	}

	var r0 *domain.Word
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Category, domain.Difficulty) (*domain.Word, error)); ok {
		return rf(category, difficulty)
	}
	if rf, ok := ret.Get(0).(func(*domain.Category, domain.Difficulty) *domain.Word); ok {
		r0 = rf(category, difficulty)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Word)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Category, domain.Difficulty) error); ok {
		r1 = rf(category, difficulty)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WordRandomizer_ChoiceWord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChoiceWord'
type WordRandomizer_ChoiceWord_Call struct {
	*mock.Call
}

// ChoiceWord is a helper method to define mock.On call
//   - category *domain.Category
//   - difficulty domain.Difficulty
func (_e *WordRandomizer_Expecter) ChoiceWord(category interface{}, difficulty interface{}) *WordRandomizer_ChoiceWord_Call {
	return &WordRandomizer_ChoiceWord_Call{Call: _e.mock.On("ChoiceWord", category, difficulty)}
}

func (_c *WordRandomizer_ChoiceWord_Call) Run(run func(category *domain.Category, difficulty domain.Difficulty)) *WordRandomizer_ChoiceWord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Category), args[1].(domain.Difficulty))
	})
	return _c
}

func (_c *WordRandomizer_ChoiceWord_Call) Return(word *domain.Word, err error) *WordRandomizer_ChoiceWord_Call {
	_c.Call.Return(word, err)
	return _c
}

func (_c *WordRandomizer_ChoiceWord_Call) RunAndReturn(run func(*domain.Category, domain.Difficulty) (*domain.Word, error)) *WordRandomizer_ChoiceWord_Call {
	_c.Call.Return(run)
	return _c
}

// NewWordRandomizer creates a new instance of WordRandomizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWordRandomizer(t interface {
	mock.TestingT
	Cleanup(func())
}) *WordRandomizer {
	mock := &WordRandomizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
