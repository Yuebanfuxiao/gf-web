package repository

type PermissionPolicy struct {
	Abstract
}

func NewPermissionPolicy() *PermissionPolicy {
	repository := &PermissionPolicy{}

	return repository
}
