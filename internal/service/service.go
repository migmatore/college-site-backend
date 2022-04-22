package service

type Deps struct {
	GroupStorage   GroupStorage
	OfficeStorage  OfficeStorage
	SubjectStorage SubjectStorage
	SheduleStorage SheduleStorage
}

type Service struct {
	Group   *GroupService
	Office  *OfficeService
	Subject *SubjectService
	Shedule *SheduleService
}

func New(deps Deps) *Service {
	return &Service{
		Group:   NewGroupService(deps.GroupStorage),
		Office:  NewOfficeService(deps.OfficeStorage),
		Subject: NewSubjectService(deps.SubjectStorage),
		Shedule: NewSheduleService(deps.SheduleStorage),
	}
}
