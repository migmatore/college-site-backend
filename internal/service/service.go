package service

type Deps struct {
	GroupStorage   GroupStorage
	OfficeStorage  OfficeStorage
	SubjectStorage SubjectStorage
	TeacherStorage TeacherStorage
	SheduleStorage SheduleStorage
}

type Service struct {
	Group   *GroupService
	Office  *OfficeService
	Subject *SubjectService
	Teacher *TeacherService
	Shedule *SheduleService
}

func New(deps Deps) *Service {
	return &Service{
		Group:   NewGroupService(deps.GroupStorage),
		Office:  NewOfficeService(deps.OfficeStorage),
		Subject: NewSubjectService(deps.SubjectStorage),
		Teacher: NewTeacherService(deps.TeacherStorage),
		Shedule: NewSheduleService(deps.SheduleStorage),
	}
}
