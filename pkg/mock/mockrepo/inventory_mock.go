// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/inventory.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	models "ahava/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInventoryRepository is a mock of InventoryRepository interface.
type MockInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepositoryMockRecorder
}

// MockInventoryRepositoryMockRecorder is the mock recorder for MockInventoryRepository.
type MockInventoryRepositoryMockRecorder struct {
	mock *MockInventoryRepository
}

// NewMockInventoryRepository creates a new mock instance.
func NewMockInventoryRepository(ctrl *gomock.Controller) *MockInventoryRepository {
	mock := &MockInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryRepository) EXPECT() *MockInventoryRepositoryMockRecorder {
	return m.recorder
}

// AddInventory mocks base method.
func (m *MockInventoryRepository) AddInventory(inventory models.AddInventory, url string) (models.InventoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInventory", inventory, url)
	ret0, _ := ret[0].(models.InventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInventory indicates an expected call of AddInventory.
func (mr *MockInventoryRepositoryMockRecorder) AddInventory(inventory, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInventory", reflect.TypeOf((*MockInventoryRepository)(nil).AddInventory), inventory, url)
}

// CheckInventory mocks base method.
func (m *MockInventoryRepository) CheckInventory(pid int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInventory", pid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckInventory indicates an expected call of CheckInventory.
func (mr *MockInventoryRepositoryMockRecorder) CheckInventory(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInventory", reflect.TypeOf((*MockInventoryRepository)(nil).CheckInventory), pid)
}

// CheckPrice mocks base method.
func (m *MockInventoryRepository) CheckPrice(inventory_id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPrice", inventory_id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPrice indicates an expected call of CheckPrice.
func (mr *MockInventoryRepositoryMockRecorder) CheckPrice(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPrice", reflect.TypeOf((*MockInventoryRepository)(nil).CheckPrice), inventory_id)
}

// CheckStock mocks base method.
func (m *MockInventoryRepository) CheckStock(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStock", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckStock indicates an expected call of CheckStock.
func (mr *MockInventoryRepositoryMockRecorder) CheckStock(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStock", reflect.TypeOf((*MockInventoryRepository)(nil).CheckStock), inventory_id)
}

// DeleteInventory mocks base method.
func (m *MockInventoryRepository) DeleteInventory(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInventory", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInventory indicates an expected call of DeleteInventory.
func (mr *MockInventoryRepositoryMockRecorder) DeleteInventory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInventory", reflect.TypeOf((*MockInventoryRepository)(nil).DeleteInventory), id)
}

// UpdateInventory mocks base method.
func (m *MockInventoryRepository) UpdateInventory(id int, model models.UpdateInventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", id, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockInventoryRepositoryMockRecorder) UpdateInventory(id, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockInventoryRepository)(nil).UpdateInventory), id, model)
}

// ListProducts mocks base method.
func (m *MockInventoryRepository) ListProducts(page int) ([]models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", page)
	ret0, _ := ret[0].([]models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockInventoryRepositoryMockRecorder) ListProducts(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockInventoryRepository)(nil).ListProducts), page)
}

// ListProductsByCategory mocks base method.
func (m *MockInventoryRepository) ListProductsByCategory(id int) ([]models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductsByCategory", id)
	ret0, _ := ret[0].([]models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductsByCategory indicates an expected call of ListProductsByCategory.
func (mr *MockInventoryRepositoryMockRecorder) ListProductsByCategory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductsByCategory", reflect.TypeOf((*MockInventoryRepository)(nil).ListProductsByCategory), id)
}

// SearchProducts mocks base method.
func (m *MockInventoryRepository) SearchProducts(key string) ([]models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProducts", key)
	ret0, _ := ret[0].([]models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockInventoryRepositoryMockRecorder) SearchProducts(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockInventoryRepository)(nil).SearchProducts), key)
}

// ShowIndividualProducts mocks base method.
func (m *MockInventoryRepository) ShowIndividualProducts(id string) (models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowIndividualProducts", id)
	ret0, _ := ret[0].(models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowIndividualProducts indicates an expected call of ShowIndividualProducts.
func (mr *MockInventoryRepositoryMockRecorder) ShowIndividualProducts(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowIndividualProducts", reflect.TypeOf((*MockInventoryRepository)(nil).ShowIndividualProducts), id)
}

// UpdateInventory mocks base method.
// func (m *MockInventoryRepository) UpdateInventory(pid int , model models.UpdateInventory) (models.InventoryResponse, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateInventory", pid, model)
// 	ret0, _ := ret[0].(models.InventoryResponse)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// UpdateInventory indicates an expected call of UpdateInventory.
// func (mr *MockInventoryRepositoryMockRecorder) UpdateInventory(pid, stock interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockInventoryRepository)(nil).UpdateInventory), pid, stock)
// }

// UpdateProductImage mocks base method.
func (m *MockInventoryRepository) UpdateProductImage(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductImage indicates an expected call of UpdateProductImage.
func (mr *MockInventoryRepositoryMockRecorder) UpdateProductImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductImage", reflect.TypeOf((*MockInventoryRepository)(nil).UpdateProductImage), arg0, arg1)
}
