package service

type Deps struct {
	GroupStorage   GroupStorage
	OfficeStorage  OfficeStorage
	SubjectStorage SubjectStorage
	TeacherStorage TeacherStorage
	WeekdayStorage WeekdayStorage
	SheduleStorage SheduleStorage
}

type Service struct {
	Group   *GroupService
	Office  *OfficeService
	Subject *SubjectService
	Teacher *TeacherService
	Weekday *WeekdayService
	Shedule *SheduleService
}

func New(deps Deps) *Service {
	return &Service{
		Group:   NewGroupService(deps.GroupStorage),
		Office:  NewOfficeService(deps.OfficeStorage),
		Subject: NewSubjectService(deps.SubjectStorage),
		Teacher: NewTeacherService(deps.TeacherStorage),
		Weekday: NewWeekdayService(deps.WeekdayStorage),
		Shedule: NewSheduleService(deps.SheduleStorage),
	}
}
