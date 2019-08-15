// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/SmartBFT-Go/consensus/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// ApplicationMock is an autogenerated mock type for the ApplicationMock type
type ApplicationMock struct {
	mock.Mock
}

// Deliver provides a mock function with given fields: proposal, signature
func (_m *ApplicationMock) Deliver(proposal types.Proposal, signature []types.Signature) {
	_m.Called(proposal, signature)
}
