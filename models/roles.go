package models

type Roles struct {
	RoleID uint `json:"role_id" gorm:"primary_key"`
	RoleName string `json:"role_name"`
}