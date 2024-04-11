package database

import (
	"github.com/go-openapi/strfmt"
	"github.com/lib/pq"
)

const (
	DigitalTraceRole                 = "digital_trace.role"
	DigitalTraceUser                 = "digital_trace.user"
	DigitalTraceSession              = "digital_trace.session"
	DigitalTraceMailUserVerefication = "digital_trace.mail_user_verification"
	DigitalTraceTest                 = "digital_trace.test"
	DigitalTraceStatus               = "digital_trace.status"
	DigitalTraceAttempt              = "digital_trace.attempt"
	DigitalTraceTestKlimov           = "digital_trace.test_klimov"
	DigitalTraceTestAzbel            = "digital_trace.test_azbel"
	DigitalTraceButtonType           = "digital_trace.button_type"
	DigitalTraceQuestion             = "digital_trace.question"
	DigitalTraceUserAnswer           = "digital_trace.user_answer"
	DigitalTraceAchievementType      = "digital_trace.achievement_type"
	DigitalTraceAchievement          = "digital_trace.achievement"
	DigitalTraceUserAvatar           = "digital_trace.user_avatar"
)

type Role struct {
	Id   *uint64 `psql:"id"`
	Name *string `psql:"name"`
}

func roleStatic() Role {
	id := uint64(0)
	name := ""

	return Role{
		Id:   &id,
		Name: &name,
	}
}

type User struct {
	Id        *uint64          `psql:"id"`
	CreatedAt *strfmt.DateTime `psql:"created_at"`
	UpdatedAt *strfmt.DateTime `psql:"updated_at"`
	RoleId    *uint64          `psql:"role_id"`
	Email     *string          `psql:"email"`
	Login     *string          `psql:"login"`
	Name      *string          `psql:"name"`
	Surname   *string          `psql:"surname"`
	Password  *string          `psql:"password"`
}

func userStatic() User {
	id := uint64(0)
	createdAt := strfmt.DateTime{}
	updatedAt := strfmt.DateTime{}
	roleId := uint64(0)
	email := ""
	login := ""
	name := ""
	surname := ""
	password := ""

	return User{
		Id:        &id,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		RoleId:    &roleId,
		Email:     &email,
		Login:     &login,
		Name:      &name,
		Surname:   &surname,
		Password:  &password,
	}
}

type Session struct {
	Id     *uint64 `psql:"id"`
	UserId *uint64 `psql:"user_id"`
	Token  *[]byte `psql:"token"`
}

func sessionStatic() Session {
	id := uint64(0)
	UserId := uint64(0)
	token := []byte{}

	return Session{
		Id:     &id,
		UserId: &UserId,
		Token:  &token,
	}
}

type MailUserVerefication struct {
	Id        *uint64          `psql:"id"`
	CreatedAt *strfmt.DateTime `psql:"created_at"`
	Token     *string          `psql:"token"`
	Email     *string          `psql:"email"`
	Login     *string          `psql:"login"`
	Name      *string          `psql:"name"`
	Surname   *string          `psql:"surname"`
	Password  *string          `psql:"password"`
}

func mailUserVereficationStatic() MailUserVerefication {
	id := uint64(0)
	createdAt := strfmt.DateTime{}
	token := ""
	email := ""
	login := ""
	name := ""
	surname := ""
	password := ""

	return MailUserVerefication{
		Id:        &id,
		CreatedAt: &createdAt,
		Token:     &token,
		Email:     &email,
		Login:     &login,
		Name:      &name,
		Surname:   &surname,
		Password:  &password,
	}
}

type Test struct {
	Id          *uint64 `psql:"id"`
	Name        *string `psql:"name"`
	PgName      *string `psql:"pg_name"`
	Content     *string `psql:"content"`
	Description *string `psql:"description"`
}

func testStatic() Test {
	id := uint64(0)
	name := ""
	pgName := ""
	content := ""
	description := ""

	return Test{
		Id:          &id,
		Name:        &name,
		PgName:      &pgName,
		Content:     &content,
		Description: &description,
	}
}

type Status struct {
	Id   *uint64 `psql:"id"`
	Name *string `psql:"name"`
}

func statusStatic() Status {
	id := uint64(0)
	name := ""

	return Status{
		Id:   &id,
		Name: &name,
	}
}

type Attempt struct {
	Id       *uint64 `psql:"id"`
	UserId   *uint64 `psql:"user_id"`
	TestId   *uint64 `psql:"test_id"`
	StatusId *uint64 `psql:"status_id"`
}

func attemptStatic() Attempt {
	id := uint64(0)
	UserId := uint64(0)
	TestId := uint64(0)
	StatusId := uint64(0)

	return Attempt{
		Id:       &id,
		UserId:   &UserId,
		TestId:   &TestId,
		StatusId: &StatusId,
	}
}

type TestKlimov struct {
	Id              *uint64 `psql:"id"`
	UserId          *uint64 `psql:"user_id"`
	AttemptId       *uint64 `psql:"attempt_id"`
	HumanSign       *uint64 `psql:"human_sign"`
	HumanHuman      *uint64 `psql:"human_human"`
	HumanNature     *uint64 `psql:"human_nature"`
	HumanTechnic    *uint64 `psql:"human_technic"`
	HumanSignSystem *uint64 `psql:"human_sign_system"`
}

func testKlimovStatic() TestKlimov {
	id := uint64(0)
	UserId := uint64(0)
	AttemptId := uint64(0)
	HumanSign := uint64(0)
	HumanHuman := uint64(0)
	HumanNature := uint64(0)
	HumanTechnic := uint64(0)
	HumanSignSystem := uint64(0)

	return TestKlimov{
		Id:              &id,
		UserId:          &UserId,
		AttemptId:       &AttemptId,
		HumanSign:       &HumanSign,
		HumanHuman:      &HumanHuman,
		HumanNature:     &HumanNature,
		HumanTechnic:    &HumanTechnic,
		HumanSignSystem: &HumanSignSystem,
	}
}

type TestAzbel struct {
	Id                            *uint64 `psql:"id"`
	UserId                        *uint64 `psql:"user_id"`
	AttemptId                     *uint64 `psql:"attempt_id"`
	PhysicsMaths                  *uint64 `psql:"physics_maths"`
	BiologyChemistry              *uint64 `psql:"biology_chemistry"`
	Tourism                       *uint64 `psql:"tourism"`
	Medicine                      *uint64 `psql:"medicine"`
	InformationTechnology         *uint64 `psql:"information_technology"`
	Construction                  *uint64 `psql:"construction"`
	EngineeringAndTechnicalSphere *uint64 `psql:"engineering_and_technical_sphere"`
	EconomicsFinance              *uint64 `psql:"economics_finance"`
	BusinessManagement            *uint64 `psql:"business_management"`
	ForeignLanguages              *uint64 `psql:"foreign_languages"`
	TransportLogistics            *uint64 `psql:"transport_logistics"`
	StrongStructure               *uint64 `psql:"strong_structure"`
	SocioPoliticalSphere          *uint64 `psql:"socio_political_sphere"`
	Journalism                    *uint64 `psql:"journalism"`
	Jurisprudence                 *uint64 `psql:"jurisprudence"`
	Education                     *uint64 `psql:"education"`
	ServiceSectorTrade            *uint64 `psql:"service_sector_trade"`
	PhysicalEducationAndSports    *uint64 `psql:"physical_education_and_sports"`
	MusicalAndPerformingArts      *uint64 `psql:"musical_and_performing_arts"`
	FineArtDesign                 *uint64 `psql:"fine_arts_design"`
}

func testAzbelStatic() TestAzbel {
	id := uint64(0)
	UserId := uint64(0)
	AttemptId := uint64(0)
	PhysicsMaths := uint64(0)
	BiologyChemistry := uint64(0)
	Tourism := uint64(0)
	Medicine := uint64(0)
	InformationTechnology := uint64(0)
	Construction := uint64(0)
	EngineeringAndTechnicalSphere := uint64(0)
	EconomicsFinance := uint64(0)
	BusinessManagement := uint64(0)
	ForeignLanguages := uint64(0)
	TransportLogistics := uint64(0)
	StrongStructure := uint64(0)
	SocioPoliticalSphere := uint64(0)
	Journalism := uint64(0)
	Jurisprudence := uint64(0)
	Education := uint64(0)
	ServiceSectorTrade := uint64(0)
	PhysicalEducationAndSports := uint64(0)
	MusicalAndPerformingArts := uint64(0)
	FineArtDesign := uint64(0)

	return TestAzbel{
		Id:                            &id,
		UserId:                        &UserId,
		AttemptId:                     &AttemptId,
		PhysicsMaths:                  &PhysicsMaths,
		BiologyChemistry:              &BiologyChemistry,
		Tourism:                       &Tourism,
		Medicine:                      &Medicine,
		InformationTechnology:         &InformationTechnology,
		Construction:                  &Construction,
		EngineeringAndTechnicalSphere: &EngineeringAndTechnicalSphere,
		EconomicsFinance:              &EconomicsFinance,
		BusinessManagement:            &BusinessManagement,
		ForeignLanguages:              &ForeignLanguages,
		TransportLogistics:            &TransportLogistics,
		StrongStructure:               &StrongStructure,
		SocioPoliticalSphere:          &SocioPoliticalSphere,
		Journalism:                    &Journalism,
		Jurisprudence:                 &Jurisprudence,
		Education:                     &Education,
		ServiceSectorTrade:            &ServiceSectorTrade,
		PhysicalEducationAndSports:    &PhysicalEducationAndSports,
		MusicalAndPerformingArts:      &MusicalAndPerformingArts,
		FineArtDesign:                 &FineArtDesign,
	}
}

type ButtonType struct {
	Id   *uint64 `psql:"id"`
	Name *string `psql:"name"`
}

func buttonTypeStatic() ButtonType {
	id := uint64(0)
	name := ""

	return ButtonType{
		Id:   &id,
		Name: &name,
	}
}

type Question struct {
	Id           *uint64         `psql:"id"`
	TestId       *uint64         `psql:"test_id"`
	ButtonTypeId *uint64         `psql:"button_type_id"`
	Number       *uint64         `psql:"number"`
	Content      *string         `psql:"content"`
	Answer       *pq.StringArray `psql:"answer"`
}

func questionStatic() Question {
	id := uint64(0)
	TestId := uint64(0)
	ButtonTypeId := uint64(0)
	Number := uint64(0)
	Content := ""
	Answer := pq.StringArray{}

	return Question{
		Id:           &id,
		TestId:       &TestId,
		ButtonTypeId: &ButtonTypeId,
		Number:       &Number,
		Content:      &Content,
		Answer:       &Answer,
	}
}

type UserAnswer struct {
	Id         *uint64        `psql:"id"`
	AttemptId  *uint64        `psql:"attempt_id"`
	QuestionId *uint64        `psql:"question_id"`
	Answer     *pq.Int64Array `psql:"answer"`
}

func userAnswerStatic() UserAnswer {
	id := uint64(0)
	AttemptId := uint64(0)
	QuestionId := uint64(0)
	Answer := pq.Int64Array{}

	return UserAnswer{
		Id:         &id,
		AttemptId:  &AttemptId,
		QuestionId: &QuestionId,
		Answer:     &Answer,
	}
}

type AchievementType struct {
	Id   *uint64 `psql:"id"`
	Name *string `psql:"name"`
}

func achievementTypeStatic() AchievementType {
	id := uint64(0)
	name := ""

	return AchievementType{
		Id:   &id,
		Name: &name,
	}
}

type Achievement struct {
	Id               *uint64        `psql:"id"`
	UserId           *uint64        `psql:"user_id"`
	AchievementID    *uint64        `psql:"achievement_id"`
	AchievementTypes *pq.Int64Array `psql:"achievement_types"`
}

func achievementStatic() Achievement {
	id := uint64(0)
	UserId := uint64(0)
	AchievementID := uint64(0)
	AchievementTypes := pq.Int64Array{}

	return Achievement{
		Id:               &id,
		UserId:           &UserId,
		AchievementID:    &AchievementID,
		AchievementTypes: &AchievementTypes,
	}
}

type UserAvatar struct {
	Id     *uint64 `psql:"id"`
	UserId *uint64 `psql:"user_id"`
	Prefix *string `psql:"prefix"`
	Path   *string `psql:"path"`
}

func userAvatarStatic() UserAvatar {
	id := uint64(0)
	UserId := uint64(0)
	Prefix := ""
	Path := ""

	return UserAvatar{
		Id:     &id,
		UserId: &UserId,
		Prefix: &Prefix,
		Path:   &Path,
	}
}
