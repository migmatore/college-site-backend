package service

type Deps struct {
	GroupStorage   GroupStorage
	SheduleStorage SheduleStorage
}

type Service struct {
	Group   *GroupService
	Shedule *SheduleService
}

func New(deps Deps) *Service {
	return &Service{
		Group:   NewGroupService(deps.GroupStorage),
		Shedule: NewSheduleService(deps.SheduleStorage),
	}
}
