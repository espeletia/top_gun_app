package domain

import (
	"time"
)

const (
	WeaponEpee  = "EPEE"
	WeaponFoil  = "FOIL"
	WeaponSabre = "SABRE"

	GenderMen   = "MEN"
	GenderWomen = "WOMEN"
	GenderMixes = "MIXED"

	EventRoleReferee = "REFEREE"
	EventRoleAdmin   = "ADMIN"
	EventRoleAthlete = "ATHLETE"

	EventTypeIndividual   = "INDIVIDUAL"
	EventTypeTeam         = "TEAM"
	EventAgeU9            = "U9"
	EventAgeU10           = "U10"
	EventAgeU11           = "U11"
	EventAgeU12           = "U12"
	EventAgeU13           = "U13"
	EventAgeU14           = "U14"
	EventAgeU15           = "U15"
	EventAgeCADET16       = "CADET16"
	EventAgeCADET17       = "CADET17"
	EventAgeU18           = "U18"
	EventAgeJUNIOR        = "JUNIOR"
	EventAgeU23           = "U23"
	EventAgeSENIOR        = "SENIOR"
	EventAgeSENIOR13PLUS  = "SENIOR13PLUS"
	EventAgeVETERAN40PLUS = "VETERAN40PLUS"
	EventAgeVETERAN50     = "VETERAN50"
	EventAgeVETERAN60     = "VETERAN60"
	EventAgeVETERAN70     = "VETERAN70"

	AthleteFinished  = "FINISHED"
	AthleteCompeting = "COMPETING"

	EventNotFound = Error("Event does not exist")
)

type Event struct {
	ID           int64
	Status       string
	TournamentId int64
	EventData
}

type EventData struct {
	Name        string
	Description *string
	Start       time.Time
	End         time.Time
	Weapon      string
	Type        string
	Gender      string
	Category    string
	Athletes    []*Athlete
}

type Athlete struct {
	UserID         int64
	PooleSeeding   int64
	TableauSeeding *int64
	FinalRanking   *int64
	Status         string
}
