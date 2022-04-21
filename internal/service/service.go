package service

type Deps struct {
	GroupStorage   GroupStorage
	OfficeStorage  OfficeStorage
	SheduleStorage SheduleStorage
}

type Service struct {
	Group   *GroupService
	Office  *OfficeService
	Shedule *SheduleService
}

func New(deps Deps) *Service {
	return &Service{
		Group:   NewGroupService(deps.GroupStorage),
		Office:  NewOfficeService(deps.OfficeStorage),
		Shedule: NewSheduleService(deps.SheduleStorage),
	}
}
