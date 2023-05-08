// Code generated by mockery v2.14.0. DO NOT EDIT.

package mockrootcoord

import (
	context "context"

	model "github.com/milvus-io/milvus/internal/metastore/model"
	mock "github.com/stretchr/testify/mock"
)

// GarbageCollector is an autogenerated mock type for the GarbageCollector type
type GarbageCollector struct {
	mock.Mock
}

type GarbageCollector_Expecter struct {
	mock *mock.Mock
}

func (_m *GarbageCollector) EXPECT() *GarbageCollector_Expecter {
	return &GarbageCollector_Expecter{mock: &_m.Mock}
}

// GcCollectionData provides a mock function with given fields: ctx, coll
func (_m *GarbageCollector) GcCollectionData(ctx context.Context, coll *model.Collection) (uint64, error) {
	ret := _m.Called(ctx, coll)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Collection) uint64); ok {
		r0 = rf(ctx, coll)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Collection) error); ok {
		r1 = rf(ctx, coll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GarbageCollector_GcCollectionData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GcCollectionData'
type GarbageCollector_GcCollectionData_Call struct {
	*mock.Call
}

// GcCollectionData is a helper method to define mock.On call
//  - ctx context.Context
//  - coll *model.Collection
func (_e *GarbageCollector_Expecter) GcCollectionData(ctx interface{}, coll interface{}) *GarbageCollector_GcCollectionData_Call {
	return &GarbageCollector_GcCollectionData_Call{Call: _e.mock.On("GcCollectionData", ctx, coll)}
}

func (_c *GarbageCollector_GcCollectionData_Call) Run(run func(ctx context.Context, coll *model.Collection)) *GarbageCollector_GcCollectionData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Collection))
	})
	return _c
}

func (_c *GarbageCollector_GcCollectionData_Call) Return(ddlTs uint64, err error) *GarbageCollector_GcCollectionData_Call {
	_c.Call.Return(ddlTs, err)
	return _c
}

// GcPartitionData provides a mock function with given fields: ctx, pChannels, partition
func (_m *GarbageCollector) GcPartitionData(ctx context.Context, pChannels []string, partition *model.Partition) (uint64, error) {
	ret := _m.Called(ctx, pChannels, partition)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, []string, *model.Partition) uint64); ok {
		r0 = rf(ctx, pChannels, partition)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, *model.Partition) error); ok {
		r1 = rf(ctx, pChannels, partition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GarbageCollector_GcPartitionData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GcPartitionData'
type GarbageCollector_GcPartitionData_Call struct {
	*mock.Call
}

// GcPartitionData is a helper method to define mock.On call
//  - ctx context.Context
//  - pChannels []string
//  - partition *model.Partition
func (_e *GarbageCollector_Expecter) GcPartitionData(ctx interface{}, pChannels interface{}, partition interface{}) *GarbageCollector_GcPartitionData_Call {
	return &GarbageCollector_GcPartitionData_Call{Call: _e.mock.On("GcPartitionData", ctx, pChannels, partition)}
}

func (_c *GarbageCollector_GcPartitionData_Call) Run(run func(ctx context.Context, pChannels []string, partition *model.Partition)) *GarbageCollector_GcPartitionData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(*model.Partition))
	})
	return _c
}

func (_c *GarbageCollector_GcPartitionData_Call) Return(ddlTs uint64, err error) *GarbageCollector_GcPartitionData_Call {
	_c.Call.Return(ddlTs, err)
	return _c
}

// ReDropCollection provides a mock function with given fields: collMeta, ts
func (_m *GarbageCollector) ReDropCollection(collMeta *model.Collection, ts uint64) {
	_m.Called(collMeta, ts)
}

// GarbageCollector_ReDropCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReDropCollection'
type GarbageCollector_ReDropCollection_Call struct {
	*mock.Call
}

// ReDropCollection is a helper method to define mock.On call
//  - collMeta *model.Collection
//  - ts uint64
func (_e *GarbageCollector_Expecter) ReDropCollection(collMeta interface{}, ts interface{}) *GarbageCollector_ReDropCollection_Call {
	return &GarbageCollector_ReDropCollection_Call{Call: _e.mock.On("ReDropCollection", collMeta, ts)}
}

func (_c *GarbageCollector_ReDropCollection_Call) Run(run func(collMeta *model.Collection, ts uint64)) *GarbageCollector_ReDropCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Collection), args[1].(uint64))
	})
	return _c
}

func (_c *GarbageCollector_ReDropCollection_Call) Return() *GarbageCollector_ReDropCollection_Call {
	_c.Call.Return()
	return _c
}

// ReDropPartition provides a mock function with given fields: pChannels, partition, ts
func (_m *GarbageCollector) ReDropPartition(pChannels []string, partition *model.Partition, ts uint64) {
	_m.Called(pChannels, partition, ts)
}

// GarbageCollector_ReDropPartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReDropPartition'
type GarbageCollector_ReDropPartition_Call struct {
	*mock.Call
}

// ReDropPartition is a helper method to define mock.On call
//  - pChannels []string
//  - partition *model.Partition
//  - ts uint64
func (_e *GarbageCollector_Expecter) ReDropPartition(pChannels interface{}, partition interface{}, ts interface{}) *GarbageCollector_ReDropPartition_Call {
	return &GarbageCollector_ReDropPartition_Call{Call: _e.mock.On("ReDropPartition", pChannels, partition, ts)}
}

func (_c *GarbageCollector_ReDropPartition_Call) Run(run func(pChannels []string, partition *model.Partition, ts uint64)) *GarbageCollector_ReDropPartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].(*model.Partition), args[2].(uint64))
	})
	return _c
}

func (_c *GarbageCollector_ReDropPartition_Call) Return() *GarbageCollector_ReDropPartition_Call {
	_c.Call.Return()
	return _c
}

// RemoveCreatingCollection provides a mock function with given fields: collMeta
func (_m *GarbageCollector) RemoveCreatingCollection(collMeta *model.Collection) {
	_m.Called(collMeta)
}

// GarbageCollector_RemoveCreatingCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCreatingCollection'
type GarbageCollector_RemoveCreatingCollection_Call struct {
	*mock.Call
}

// RemoveCreatingCollection is a helper method to define mock.On call
//  - collMeta *model.Collection
func (_e *GarbageCollector_Expecter) RemoveCreatingCollection(collMeta interface{}) *GarbageCollector_RemoveCreatingCollection_Call {
	return &GarbageCollector_RemoveCreatingCollection_Call{Call: _e.mock.On("RemoveCreatingCollection", collMeta)}
}

func (_c *GarbageCollector_RemoveCreatingCollection_Call) Run(run func(collMeta *model.Collection)) *GarbageCollector_RemoveCreatingCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Collection))
	})
	return _c
}

func (_c *GarbageCollector_RemoveCreatingCollection_Call) Return() *GarbageCollector_RemoveCreatingCollection_Call {
	_c.Call.Return()
	return _c
}

// RemoveCreatingPartition provides a mock function with given fields: partition, ts
func (_m *GarbageCollector) RemoveCreatingPartition(partition *model.Partition, ts uint64) {
	_m.Called(partition, ts)
}

// GarbageCollector_RemoveCreatingPartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCreatingPartition'
type GarbageCollector_RemoveCreatingPartition_Call struct {
	*mock.Call
}

// RemoveCreatingPartition is a helper method to define mock.On call
//  - partition *model.Partition
//  - ts uint64
func (_e *GarbageCollector_Expecter) RemoveCreatingPartition(partition interface{}, ts interface{}) *GarbageCollector_RemoveCreatingPartition_Call {
	return &GarbageCollector_RemoveCreatingPartition_Call{Call: _e.mock.On("RemoveCreatingPartition", partition, ts)}
}

func (_c *GarbageCollector_RemoveCreatingPartition_Call) Run(run func(partition *model.Partition, ts uint64)) *GarbageCollector_RemoveCreatingPartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Partition), args[1].(uint64))
	})
	return _c
}

func (_c *GarbageCollector_RemoveCreatingPartition_Call) Return() *GarbageCollector_RemoveCreatingPartition_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewGarbageCollector interface {
	mock.TestingT
	Cleanup(func())
}

// NewGarbageCollector creates a new instance of GarbageCollector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGarbageCollector(t mockConstructorTestingTNewGarbageCollector) *GarbageCollector {
	mock := &GarbageCollector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
