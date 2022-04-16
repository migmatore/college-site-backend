package service

type Deps struct {
	GroupStorage GroupStorage
}

type Service struct {
	Group *GroupService
}

func New(deps Deps) *Service {
	return &Service{Group: NewGroupService(deps.GroupStorage)}
}
