// Code generated by mockery v2.32.4. DO NOT EDIT.

package segments

import (
	schemapb "github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
	mock "github.com/stretchr/testify/mock"

	segcorepb "github.com/milvus-io/milvus/internal/proto/segcorepb"
)

// MockCollectionManager is an autogenerated mock type for the CollectionManager type
type MockCollectionManager struct {
	mock.Mock
}

type MockCollectionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCollectionManager) EXPECT() *MockCollectionManager_Expecter {
	return &MockCollectionManager_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: collectionID
func (_m *MockCollectionManager) Get(collectionID int64) *Collection {
	ret := _m.Called(collectionID)

	var r0 *Collection
	if rf, ok := ret.Get(0).(func(int64) *Collection); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Collection)
		}
	}

	return r0
}

// MockCollectionManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCollectionManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - collectionID int64
func (_e *MockCollectionManager_Expecter) Get(collectionID interface{}) *MockCollectionManager_Get_Call {
	return &MockCollectionManager_Get_Call{Call: _e.mock.On("Get", collectionID)}
}

func (_c *MockCollectionManager_Get_Call) Run(run func(collectionID int64)) *MockCollectionManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCollectionManager_Get_Call) Return(_a0 *Collection) *MockCollectionManager_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCollectionManager_Get_Call) RunAndReturn(run func(int64) *Collection) *MockCollectionManager_Get_Call {
	_c.Call.Return(run)
	return _c
}

// PutOrRef provides a mock function with given fields: collectionID, schema, meta, loadMeta
func (_m *MockCollectionManager) PutOrRef(collectionID int64, schema *schemapb.CollectionSchema, meta *segcorepb.CollectionIndexMeta, loadMeta *querypb.LoadMetaInfo) {
	_m.Called(collectionID, schema, meta, loadMeta)
}

// MockCollectionManager_PutOrRef_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutOrRef'
type MockCollectionManager_PutOrRef_Call struct {
	*mock.Call
}

// PutOrRef is a helper method to define mock.On call
//   - collectionID int64
//   - schema *schemapb.CollectionSchema
//   - meta *segcorepb.CollectionIndexMeta
//   - loadMeta *querypb.LoadMetaInfo
func (_e *MockCollectionManager_Expecter) PutOrRef(collectionID interface{}, schema interface{}, meta interface{}, loadMeta interface{}) *MockCollectionManager_PutOrRef_Call {
	return &MockCollectionManager_PutOrRef_Call{Call: _e.mock.On("PutOrRef", collectionID, schema, meta, loadMeta)}
}

func (_c *MockCollectionManager_PutOrRef_Call) Run(run func(collectionID int64, schema *schemapb.CollectionSchema, meta *segcorepb.CollectionIndexMeta, loadMeta *querypb.LoadMetaInfo)) *MockCollectionManager_PutOrRef_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*schemapb.CollectionSchema), args[2].(*segcorepb.CollectionIndexMeta), args[3].(*querypb.LoadMetaInfo))
	})
	return _c
}

func (_c *MockCollectionManager_PutOrRef_Call) Return() *MockCollectionManager_PutOrRef_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCollectionManager_PutOrRef_Call) RunAndReturn(run func(int64, *schemapb.CollectionSchema, *segcorepb.CollectionIndexMeta, *querypb.LoadMetaInfo)) *MockCollectionManager_PutOrRef_Call {
	_c.Call.Return(run)
	return _c
}

// Ref provides a mock function with given fields: collectionID, count
func (_m *MockCollectionManager) Ref(collectionID int64, count uint32) bool {
	ret := _m.Called(collectionID, count)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, uint32) bool); ok {
		r0 = rf(collectionID, count)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCollectionManager_Ref_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ref'
type MockCollectionManager_Ref_Call struct {
	*mock.Call
}

// Ref is a helper method to define mock.On call
//   - collectionID int64
//   - count uint32
func (_e *MockCollectionManager_Expecter) Ref(collectionID interface{}, count interface{}) *MockCollectionManager_Ref_Call {
	return &MockCollectionManager_Ref_Call{Call: _e.mock.On("Ref", collectionID, count)}
}

func (_c *MockCollectionManager_Ref_Call) Run(run func(collectionID int64, count uint32)) *MockCollectionManager_Ref_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(uint32))
	})
	return _c
}

func (_c *MockCollectionManager_Ref_Call) Return(_a0 bool) *MockCollectionManager_Ref_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCollectionManager_Ref_Call) RunAndReturn(run func(int64, uint32) bool) *MockCollectionManager_Ref_Call {
	_c.Call.Return(run)
	return _c
}

// Unref provides a mock function with given fields: collectionID, count
func (_m *MockCollectionManager) Unref(collectionID int64, count uint32) bool {
	ret := _m.Called(collectionID, count)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, uint32) bool); ok {
		r0 = rf(collectionID, count)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCollectionManager_Unref_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unref'
type MockCollectionManager_Unref_Call struct {
	*mock.Call
}

// Unref is a helper method to define mock.On call
//   - collectionID int64
//   - count uint32
func (_e *MockCollectionManager_Expecter) Unref(collectionID interface{}, count interface{}) *MockCollectionManager_Unref_Call {
	return &MockCollectionManager_Unref_Call{Call: _e.mock.On("Unref", collectionID, count)}
}

func (_c *MockCollectionManager_Unref_Call) Run(run func(collectionID int64, count uint32)) *MockCollectionManager_Unref_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(uint32))
	})
	return _c
}

func (_c *MockCollectionManager_Unref_Call) Return(_a0 bool) *MockCollectionManager_Unref_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCollectionManager_Unref_Call) RunAndReturn(run func(int64, uint32) bool) *MockCollectionManager_Unref_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCollectionManager creates a new instance of MockCollectionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCollectionManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCollectionManager {
	mock := &MockCollectionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
