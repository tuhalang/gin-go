package schema

import "time"

// Role
type Role struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Sequence  int       `json:"sequence"`
	Memo      string    `json:"memo"`
	Status    int       `json:"status" binding:"required,max=2,min=1"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleMenus RoleMenus `json:"role_menus" binding:"required,gt=0"`
}

// RoleQueryParam 查询条件
type RoleQueryParam struct {
	PaginationParam
	IDs        []string `form:"-"`
	Name       string   `form:"-"`
	QueryValue string   `form:"queryValue"`
	UserID     string   `form:"-"`
	Status     int      `form:"status"`
}

// RoleQueryOptions
type RoleQueryOptions struct {
	OrderFields []*OrderField
}

// RoleQueryResult
type RoleQueryResult struct {
	Data       Roles
	PageResult *PaginationResult
}

// Roles 角色对象列表
type Roles []*Role

// ToNames 获取角色名称列表
func (a Roles) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Roles) ToMap() map[string]*Role {
	m := make(map[string]*Role)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// ----------------------------------------RoleMenu--------------------------------------

// RoleMenu
type RoleMenu struct {
	ID       string `json:"id"`
	RoleID   string `json:"role_id" binding:"required"`
	MenuID   string `json:"menu_id" binding:"required"`
	ActionID string `json:"action_id" binding:"required"`
}
// RoleMenu Query Param
type RoleMenuQueryParam struct {
	PaginationParam
	RoleID  string
	RoleIDs []string
}

// RoleMenuQueryOptions
type RoleMenuQueryOptions struct {
	OrderFields []*OrderField
}

// RoleMenuQueryResult
type RoleMenuQueryResult struct {
	Data       RoleMenus
	PageResult *PaginationResult
}

// RoleMenus
type RoleMenus []*RoleMenu

// ToMap
func (a RoleMenus) ToMap() map[string]*RoleMenu {
	m := make(map[string]*RoleMenu)
	for _, item := range a {
		m[item.MenuID+"-"+item.ActionID] = item
	}
	return m
}

// ToRoleIDMap
func (a RoleMenus) ToRoleIDMap() map[string]RoleMenus {
	m := make(map[string]RoleMenus)
	for _, item := range a {
		m[item.RoleID] = append(m[item.RoleID], item)
	}
	return m
}

// ToMenuIDs
func (a RoleMenus) ToMenuIDs() []string {
	var idList []string
	m := make(map[string]struct{})

	for _, item := range a {
		if _, ok := m[item.MenuID]; ok {
			continue
		}
		idList = append(idList, item.MenuID)
		m[item.MenuID] = struct{}{}
	}

	return idList
}

// ToActionIDs
func (a RoleMenus) ToActionIDs() []string {
	idList := make([]string, len(a))
	m := make(map[string]struct{})
	for i, item := range a {
		if _, ok := m[item.ActionID]; ok {
			continue
		}
		idList[i] = item.ActionID
		m[item.ActionID] = struct{}{}
	}
	return idList
}
