package mocks

import (
	"mini-project/models"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockLeaderboardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderboardRepositoryMockRecorder
}

func NewMockLeaderboardRepository(ctrl *gomock.Controller) *MockLeaderboardRepository {
	mock := &MockLeaderboardRepository{ctrl: ctrl}
	mock.recorder = &MockLeaderboardRepositoryMockRecorder{mock}
	return mock
}

type MockLeaderboardRepositoryMockRecorder struct {
	mock *MockLeaderboardRepository
}

func (m *MockLeaderboardRepository) EXPECT() *MockLeaderboardRepositoryMockRecorder {
	return m.recorder
}

func (m *MockLeaderboardRepository) CreateLeaderboard(leaderboard *models.Leaderboard) error {
	ret := m.ctrl.Call(m, "CreateLeaderboard", leaderboard)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLeaderboardRepositoryMockRecorder) CreateLeaderboard(leaderboard interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLeaderboard", reflect.TypeOf((*MockLeaderboardRepository)(nil).CreateLeaderboard), leaderboard)
}

func (m *MockLeaderboardRepository) GetByID(id uint) (models.Leaderboard, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.Leaderboard) 
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaderboardRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLeaderboardRepository)(nil).GetByID), id)
}

func (m *MockLeaderboardRepository) GetAll() ([]models.Leaderboard, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Leaderboard) 
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaderboardRepositoryMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockLeaderboardRepository)(nil).GetAll))
}

func (m *MockLeaderboardRepository) GetByUserID(userID uint) (models.Leaderboard, error) {
	ret := m.ctrl.Call(m, "GetByUserID", userID)
	ret0, _ := ret[0].(models.Leaderboard) 
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaderboardRepositoryMockRecorder) GetByUserID(userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockLeaderboardRepository)(nil).GetByUserID), userID)
}

type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

func (m *MockUserRepository) GetByID(id uint) (models.User, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.User) 
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), id)
}

func (m *MockUserRepository) GetByEmail(email string) (models.User, error) {
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(models.User) 
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByEmail), email)
}

func (m *MockUserRepository) Register(user models.User) error {
    ret := m.ctrl.Call(m, "Register", user)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockUserRepositoryMockRecorder) Register(user interface{}) *gomock.Call {
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserRepository)(nil).Register), user)
}

func (m *MockUserRepository) Update(user models.User) error {
    ret := m.ctrl.Call(m, "Update", user)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockUserRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), user)
}


