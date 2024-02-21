// Code generated by mockery. DO NOT EDIT.

package mocksvc

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// GroupProjectService is an autogenerated mock type for the GroupProjectService type
type GroupProjectService struct {
	mock.Mock
}

// AddAllTeamsToProject provides a mock function with given fields: ctx, groupID, teamRole
func (_m *GroupProjectService) AddAllTeamsToProject(ctx context.Context, groupID string, teamRole *[]admin.TeamRole) (*admin.PaginatedTeamRole, *http.Response, error) {
	ret := _m.Called(ctx, groupID, teamRole)

	if len(ret) == 0 {
		panic("no return value specified for AddAllTeamsToProject")
	}

	var r0 *admin.PaginatedTeamRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *[]admin.TeamRole) (*admin.PaginatedTeamRole, *http.Response, error)); ok {
		return rf(ctx, groupID, teamRole)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *[]admin.TeamRole) *admin.PaginatedTeamRole); ok {
		r0 = rf(ctx, groupID, teamRole)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedTeamRole)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *[]admin.TeamRole) *http.Response); ok {
		r1 = rf(ctx, groupID, teamRole)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *[]admin.TeamRole) error); ok {
		r2 = rf(ctx, groupID, teamRole)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteProjectLimit provides a mock function with given fields: ctx, limitName, projectID
func (_m *GroupProjectService) DeleteProjectLimit(ctx context.Context, limitName string, projectID string) (map[string]interface{}, *http.Response, error) {
	ret := _m.Called(ctx, limitName, projectID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectLimit")
	}

	var r0 map[string]interface{}
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]interface{}, *http.Response, error)); ok {
		return rf(ctx, limitName, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]interface{}); ok {
		r0 = rf(ctx, limitName, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *http.Response); ok {
		r1 = rf(ctx, limitName, projectID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, limitName, projectID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProjectSettings provides a mock function with given fields: ctx, groupID
func (_m *GroupProjectService) GetProjectSettings(ctx context.Context, groupID string) (*admin.GroupSettings, *http.Response, error) {
	ret := _m.Called(ctx, groupID)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectSettings")
	}

	var r0 *admin.GroupSettings
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*admin.GroupSettings, *http.Response, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *admin.GroupSettings); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GroupSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, groupID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListClusters provides a mock function with given fields: ctx, groupID
func (_m *GroupProjectService) ListClusters(ctx context.Context, groupID string) (*admin.PaginatedAdvancedClusterDescription, *http.Response, error) {
	ret := _m.Called(ctx, groupID)

	if len(ret) == 0 {
		panic("no return value specified for ListClusters")
	}

	var r0 *admin.PaginatedAdvancedClusterDescription
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*admin.PaginatedAdvancedClusterDescription, *http.Response, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *admin.PaginatedAdvancedClusterDescription); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedAdvancedClusterDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, groupID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListProjectLimits provides a mock function with given fields: ctx, groupID
func (_m *GroupProjectService) ListProjectLimits(ctx context.Context, groupID string) ([]admin.DataFederationLimit, *http.Response, error) {
	ret := _m.Called(ctx, groupID)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectLimits")
	}

	var r0 []admin.DataFederationLimit
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]admin.DataFederationLimit, *http.Response, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []admin.DataFederationLimit); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.DataFederationLimit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, groupID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListProjectTeams provides a mock function with given fields: ctx, groupID
func (_m *GroupProjectService) ListProjectTeams(ctx context.Context, groupID string) (*admin.PaginatedTeamRole, *http.Response, error) {
	ret := _m.Called(ctx, groupID)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectTeams")
	}

	var r0 *admin.PaginatedTeamRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*admin.PaginatedTeamRole, *http.Response, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *admin.PaginatedTeamRole); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedTeamRole)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, groupID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveProjectTeam provides a mock function with given fields: ctx, groupID, teamID
func (_m *GroupProjectService) RemoveProjectTeam(ctx context.Context, groupID string, teamID string) (*http.Response, error) {
	ret := _m.Called(ctx, groupID, teamID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProjectTeam")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*http.Response, error)); ok {
		return rf(ctx, groupID, teamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *http.Response); ok {
		r0 = rf(ctx, groupID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, groupID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnAllIPAddresses provides a mock function with given fields: ctx, groupID
func (_m *GroupProjectService) ReturnAllIPAddresses(ctx context.Context, groupID string) (*admin.GroupIPAddresses, *http.Response, error) {
	ret := _m.Called(ctx, groupID)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllIPAddresses")
	}

	var r0 *admin.GroupIPAddresses
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*admin.GroupIPAddresses, *http.Response, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *admin.GroupIPAddresses); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GroupIPAddresses)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, groupID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetProjectLimit provides a mock function with given fields: ctx, limitName, groupID, dataFederationLimit
func (_m *GroupProjectService) SetProjectLimit(ctx context.Context, limitName string, groupID string, dataFederationLimit *admin.DataFederationLimit) (*admin.DataFederationLimit, *http.Response, error) {
	ret := _m.Called(ctx, limitName, groupID, dataFederationLimit)

	if len(ret) == 0 {
		panic("no return value specified for SetProjectLimit")
	}

	var r0 *admin.DataFederationLimit
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.DataFederationLimit) (*admin.DataFederationLimit, *http.Response, error)); ok {
		return rf(ctx, limitName, groupID, dataFederationLimit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.DataFederationLimit) *admin.DataFederationLimit); ok {
		r0 = rf(ctx, limitName, groupID, dataFederationLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.DataFederationLimit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *admin.DataFederationLimit) *http.Response); ok {
		r1 = rf(ctx, limitName, groupID, dataFederationLimit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *admin.DataFederationLimit) error); ok {
		r2 = rf(ctx, limitName, groupID, dataFederationLimit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateProject provides a mock function with given fields: ctx, groupID, groupName
func (_m *GroupProjectService) UpdateProject(ctx context.Context, groupID string, groupName *admin.GroupName) (*admin.Group, *http.Response, error) {
	ret := _m.Called(ctx, groupID, groupName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProject")
	}

	var r0 *admin.Group
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.GroupName) (*admin.Group, *http.Response, error)); ok {
		return rf(ctx, groupID, groupName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.GroupName) *admin.Group); ok {
		r0 = rf(ctx, groupID, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *admin.GroupName) *http.Response); ok {
		r1 = rf(ctx, groupID, groupName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *admin.GroupName) error); ok {
		r2 = rf(ctx, groupID, groupName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTeamRoles provides a mock function with given fields: ctx, groupID, teamID, teamRole
func (_m *GroupProjectService) UpdateTeamRoles(ctx context.Context, groupID string, teamID string, teamRole *admin.TeamRole) (*admin.PaginatedTeamRole, *http.Response, error) {
	ret := _m.Called(ctx, groupID, teamID, teamRole)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTeamRoles")
	}

	var r0 *admin.PaginatedTeamRole
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.TeamRole) (*admin.PaginatedTeamRole, *http.Response, error)); ok {
		return rf(ctx, groupID, teamID, teamRole)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.TeamRole) *admin.PaginatedTeamRole); ok {
		r0 = rf(ctx, groupID, teamID, teamRole)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedTeamRole)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *admin.TeamRole) *http.Response); ok {
		r1 = rf(ctx, groupID, teamID, teamRole)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *admin.TeamRole) error); ok {
		r2 = rf(ctx, groupID, teamID, teamRole)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewGroupProjectService creates a new instance of GroupProjectService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupProjectService(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupProjectService {
	mock := &GroupProjectService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
