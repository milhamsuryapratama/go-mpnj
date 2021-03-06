// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package categorytest

import (
	context "context"
	categories "go-mpnj/categories"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, category
func (_m *Service) Create(ctx context.Context, category *categories.Category) error {
	ret := _m.Called(ctx, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Category) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Service) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, category
func (_m *Service) Get(ctx context.Context, category *[]categories.Category) error {
	ret := _m.Called(ctx, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]categories.Category) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, category, id
func (_m *Service) Update(ctx context.Context, category *categories.Category, id int) error {
	ret := _m.Called(ctx, category, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Category, int) error); ok {
		r0 = rf(ctx, category, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
