// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Athlete struct {
	UserID         string `json:"userId"`
	User           *User  `json:"User"`
	PooleSeeding   int64  `json:"PooleSeeding"`
	TableauSeeding *int64 `json:"TableauSeeding"`
	FinalRanking   *int64 `json:"FinalRanking"`
}

type AthleteSeedingInput struct {
	UserID string `json:"UserId"`
	Seed   int64  `json:"Seed"`
}

type Club struct {
	ID         string    `json:"Id"`
	Name       string    `json:"Name"`
	Location   *Location `json:"Location"`
	OwnerID    string    `json:"OwnerId"`
	Owner      *User     `json:"Owner"`
	MembersIds []string  `json:"MembersIds"`
	Members    []*User   `json:"Members"`
	Country    string    `json:"Country"`
}

type CreateEventInput struct {
	Name        string                 `json:"Name"`
	Description *string                `json:"Description"`
	RefereeIds  []string               `json:"RefereeIds"`
	AthleteIds  []*AthleteSeedingInput `json:"AthleteIds"`
	Start       int64                  `json:"start"`
	End         int64                  `json:"end"`
	Details     *DetailsInput          `json:"Details"`
}

type CreateMatchInput struct {
	LeftAthlete string `json:"LeftAthlete"`
}

type CreateTournamentInput struct {
	Start       int64               `json:"start"`
	End         int64               `json:"end"`
	Name        string              `json:"name"`
	Location    *LocationInput      `json:"Location"`
	City        string              `json:"City"`
	Country     string              `json:"Country"`
	OwnerID     string              `json:"OwnerId"`
	Events      []*CreateEventInput `json:"Events"`
	Description *string             `json:"Description"`
}

type CreateUserInput struct {
	Email          string `json:"Email"`
	BornIn         string `json:"BornIn"`
	HashedPassword string `json:"HashedPassword"`
	UserName       string `json:"UserName"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Nationality    string `json:"Nationality"`
}

type DetailsInput struct {
	Weapon   EventWeapon      `json:"Weapon"`
	Type     EventType        `json:"Type"`
	Gender   EventGenderMix   `json:"Gender"`
	Category EventAgeCategory `json:"Category"`
}

type Event struct {
	ID           string           `json:"Id"`
	Name         string           `json:"Name"`
	Description  *string          `json:"Description"`
	Tournament   *Tournament      `json:"Tournament"`
	TournamentID string           `json:"TournamentId"`
	RefereeIds   []string         `json:"RefereeIds"`
	Referees     []*User          `json:"Referees"`
	AthleteIds   []string         `json:"AthleteIds"`
	Athletes     []*Athlete       `json:"Athletes"`
	Start        int64            `json:"start"`
	Pooles       []*Poole         `json:"Pooles"`
	Tableaus     []*Tableau       `json:"Tableaus"`
	End          int64            `json:"end"`
	Status       TournamentStatus `json:"Status"`
	Details      *EventDetails    `json:"Details"`
}

type EventDetails struct {
	Weapon   EventWeapon      `json:"Weapon"`
	Type     EventType        `json:"Type"`
	Gender   EventGenderMix   `json:"Gender"`
	Category EventAgeCategory `json:"Category"`
}

type Location struct {
	Lat     float64 `json:"Lat"`
	Lon     float64 `json:"Lon"`
	Address string  `json:"Address"`
}

type LocationInput struct {
	Lat     float64 `json:"Lat"`
	Lon     float64 `json:"Lon"`
	Address string  `json:"Address"`
}

type Match struct {
	ID             string      `json:"Id"`
	LeftAthleteID  string      `json:"LeftAthleteId"`
	LeftAthlete    *User       `json:"LeftAthlete"`
	RightAthleteID string      `json:"RightAthleteId"`
	RightAthlete   *User       `json:"RightAthlete"`
	RefereeID      string      `json:"RefereeId"`
	Referee        *User       `json:"Referee"`
	RightScore     int64       `json:"RightScore"`
	LeftScore      int64       `json:"LeftScore"`
	Status         MatchStatus `json:"Status"`
}

type Poole struct {
	ID         string           `json:"Id"`
	EventID    string           `json:"EventId"`
	RefereeID  string           `json:"RefereeId"`
	Referee    []*User          `json:"Referee"`
	AthleteIds string           `json:"AthleteIds"`
	Athletes   []*User          `json:"Athletes"`
	MatchIds   []string         `json:"MatchIds"`
	Matches    []*Match         `json:"Matches"`
	Status     TournamentStatus `json:"Status"`
}

type Tableau struct {
	ID       string           `json:"Id"`
	EventID  string           `json:"EventId"`
	Matches  []*Match         `json:"Matches"`
	MatchIds []string         `json:"MatchIds"`
	Name     string           `json:"Name"`
	Status   TournamentStatus `json:"Status"`
}

type Tournament struct {
	ID          string           `json:"Id"`
	Start       int64            `json:"start"`
	End         int64            `json:"end"`
	Name        string           `json:"name"`
	Location    *Location        `json:"Location"`
	City        string           `json:"City"`
	Country     string           `json:"Country"`
	OwnerID     string           `json:"OwnerId"`
	Owner       *User            `json:"Owner"`
	Events      []*Event         `json:"Events"`
	Status      TournamentStatus `json:"Status"`
	Description *string          `json:"Description"`
}

type User struct {
	ID                          string        `json:"Id"`
	Email                       string        `json:"Email"`
	BornIn                      string        `json:"BornIn"`
	UserName                    string        `json:"UserName"`
	FirstName                   *string       `json:"FirstName"`
	LastName                    *string       `json:"LastName"`
	ParticipatingTournamentsIds []string      `json:"ParticipatingTournamentsIds"`
	ParticipatingTournaments    []*Tournament `json:"ParticipatingTournaments"`
	LikedTournamentsIds         []string      `json:"LikedTournamentsIds"`
	LikedTournaments            []*Tournament `json:"LikedTournaments"`
	FollowingUserIds            []string      `json:"FollowingUserIds"`
	Following                   []*User       `json:"Following"`
	FollowersUserIds            []string      `json:"FollowersUserIds"`
	Followers                   []*User       `json:"Followers"`
	Nationality                 string        `json:"Nationality"`
}

type AthleteStatus string

const (
	AthleteStatusFinished  AthleteStatus = "FINISHED"
	AthleteStatusCompeting AthleteStatus = "COMPETING"
)

var AllAthleteStatus = []AthleteStatus{
	AthleteStatusFinished,
	AthleteStatusCompeting,
}

func (e AthleteStatus) IsValid() bool {
	switch e {
	case AthleteStatusFinished, AthleteStatusCompeting:
		return true
	}
	return false
}

func (e AthleteStatus) String() string {
	return string(e)
}

func (e *AthleteStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AthleteStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AthleteStatus", str)
	}
	return nil
}

func (e AthleteStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventAgeCategory string

const (
	EventAgeCategoryU9            EventAgeCategory = "U9"
	EventAgeCategoryU10           EventAgeCategory = "U10"
	EventAgeCategoryU11           EventAgeCategory = "U11"
	EventAgeCategoryU12           EventAgeCategory = "U12"
	EventAgeCategoryU13           EventAgeCategory = "U13"
	EventAgeCategoryU14           EventAgeCategory = "U14"
	EventAgeCategoryU15           EventAgeCategory = "U15"
	EventAgeCategoryCadet16       EventAgeCategory = "CADET16"
	EventAgeCategoryCadet17       EventAgeCategory = "CADET17"
	EventAgeCategoryU18           EventAgeCategory = "U18"
	EventAgeCategoryJunior        EventAgeCategory = "JUNIOR"
	EventAgeCategoryU23           EventAgeCategory = "U23"
	EventAgeCategorySenior        EventAgeCategory = "SENIOR"
	EventAgeCategorySenior13plus  EventAgeCategory = "SENIOR13PLUS"
	EventAgeCategoryVeteran40plus EventAgeCategory = "VETERAN40PLUS"
	EventAgeCategoryVeteran50     EventAgeCategory = "VETERAN50"
	EventAgeCategoryVeteran60     EventAgeCategory = "VETERAN60"
	EventAgeCategoryVeteran70     EventAgeCategory = "VETERAN70"
)

var AllEventAgeCategory = []EventAgeCategory{
	EventAgeCategoryU9,
	EventAgeCategoryU10,
	EventAgeCategoryU11,
	EventAgeCategoryU12,
	EventAgeCategoryU13,
	EventAgeCategoryU14,
	EventAgeCategoryU15,
	EventAgeCategoryCadet16,
	EventAgeCategoryCadet17,
	EventAgeCategoryU18,
	EventAgeCategoryJunior,
	EventAgeCategoryU23,
	EventAgeCategorySenior,
	EventAgeCategorySenior13plus,
	EventAgeCategoryVeteran40plus,
	EventAgeCategoryVeteran50,
	EventAgeCategoryVeteran60,
	EventAgeCategoryVeteran70,
}

func (e EventAgeCategory) IsValid() bool {
	switch e {
	case EventAgeCategoryU9, EventAgeCategoryU10, EventAgeCategoryU11, EventAgeCategoryU12, EventAgeCategoryU13, EventAgeCategoryU14, EventAgeCategoryU15, EventAgeCategoryCadet16, EventAgeCategoryCadet17, EventAgeCategoryU18, EventAgeCategoryJunior, EventAgeCategoryU23, EventAgeCategorySenior, EventAgeCategorySenior13plus, EventAgeCategoryVeteran40plus, EventAgeCategoryVeteran50, EventAgeCategoryVeteran60, EventAgeCategoryVeteran70:
		return true
	}
	return false
}

func (e EventAgeCategory) String() string {
	return string(e)
}

func (e *EventAgeCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventAgeCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventAgeCategory", str)
	}
	return nil
}

func (e EventAgeCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventGenderMix string

const (
	EventGenderMixMixed EventGenderMix = "MIXED"
	EventGenderMixMen   EventGenderMix = "MEN"
	EventGenderMixWomen EventGenderMix = "WOMEN"
)

var AllEventGenderMix = []EventGenderMix{
	EventGenderMixMixed,
	EventGenderMixMen,
	EventGenderMixWomen,
}

func (e EventGenderMix) IsValid() bool {
	switch e {
	case EventGenderMixMixed, EventGenderMixMen, EventGenderMixWomen:
		return true
	}
	return false
}

func (e EventGenderMix) String() string {
	return string(e)
}

func (e *EventGenderMix) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventGenderMix(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventGenderMix", str)
	}
	return nil
}

func (e EventGenderMix) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventRoles string

const (
	EventRolesReferee  EventRoles = "REFEREE"
	EventRolesAdmin    EventRoles = "ADMIN"
	EventRolesAthelete EventRoles = "ATHELETE"
)

var AllEventRoles = []EventRoles{
	EventRolesReferee,
	EventRolesAdmin,
	EventRolesAthelete,
}

func (e EventRoles) IsValid() bool {
	switch e {
	case EventRolesReferee, EventRolesAdmin, EventRolesAthelete:
		return true
	}
	return false
}

func (e EventRoles) String() string {
	return string(e)
}

func (e *EventRoles) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventRoles(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventRoles", str)
	}
	return nil
}

func (e EventRoles) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventType string

const (
	EventTypeIndividual EventType = "INDIVIDUAL"
	EventTypeTeam       EventType = "TEAM"
)

var AllEventType = []EventType{
	EventTypeIndividual,
	EventTypeTeam,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeIndividual, EventTypeTeam:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventWeapon string

const (
	EventWeaponEpee  EventWeapon = "EPEE"
	EventWeaponFoil  EventWeapon = "FOIL"
	EventWeaponSabre EventWeapon = "SABRE"
)

var AllEventWeapon = []EventWeapon{
	EventWeaponEpee,
	EventWeaponFoil,
	EventWeaponSabre,
}

func (e EventWeapon) IsValid() bool {
	switch e {
	case EventWeaponEpee, EventWeaponFoil, EventWeaponSabre:
		return true
	}
	return false
}

func (e EventWeapon) String() string {
	return string(e)
}

func (e *EventWeapon) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventWeapon(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventWeapon", str)
	}
	return nil
}

func (e EventWeapon) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GlobalRoles string

const (
	GlobalRolesAdmin GlobalRoles = "ADMIN"
	GlobalRolesUser  GlobalRoles = "USER"
)

var AllGlobalRoles = []GlobalRoles{
	GlobalRolesAdmin,
	GlobalRolesUser,
}

func (e GlobalRoles) IsValid() bool {
	switch e {
	case GlobalRolesAdmin, GlobalRolesUser:
		return true
	}
	return false
}

func (e GlobalRoles) String() string {
	return string(e)
}

func (e *GlobalRoles) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GlobalRoles(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GlobalRoles", str)
	}
	return nil
}

func (e GlobalRoles) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MatchStatus string

const (
	MatchStatusPending  MatchStatus = "PENDING"
	MatchStatusFinished MatchStatus = "FINISHED"
)

var AllMatchStatus = []MatchStatus{
	MatchStatusPending,
	MatchStatusFinished,
}

func (e MatchStatus) IsValid() bool {
	switch e {
	case MatchStatusPending, MatchStatusFinished:
		return true
	}
	return false
}

func (e MatchStatus) String() string {
	return string(e)
}

func (e *MatchStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MatchStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MatchStatus", str)
	}
	return nil
}

func (e MatchStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Side string

const (
	SideLeft  Side = "LEFT"
	SideRight Side = "RIGHT"
	SideNone  Side = "NONE"
)

var AllSide = []Side{
	SideLeft,
	SideRight,
	SideNone,
}

func (e Side) IsValid() bool {
	switch e {
	case SideLeft, SideRight, SideNone:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TournamentStatus string

const (
	TournamentStatusCreated  TournamentStatus = "CREATED"
	TournamentStatusStarted  TournamentStatus = "STARTED"
	TournamentStatusFinished TournamentStatus = "FINISHED"
)

var AllTournamentStatus = []TournamentStatus{
	TournamentStatusCreated,
	TournamentStatusStarted,
	TournamentStatusFinished,
}

func (e TournamentStatus) IsValid() bool {
	switch e {
	case TournamentStatusCreated, TournamentStatusStarted, TournamentStatusFinished:
		return true
	}
	return false
}

func (e TournamentStatus) String() string {
	return string(e)
}

func (e *TournamentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TournamentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TournamentStatus", str)
	}
	return nil
}

func (e TournamentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRoles string

const (
	UserRolesFencer  UserRoles = "FENCER"
	UserRolesTrainer UserRoles = "TRAINER"
	UserRolesReferee UserRoles = "REFEREE"
	UserRolesNone    UserRoles = "NONE"
)

var AllUserRoles = []UserRoles{
	UserRolesFencer,
	UserRolesTrainer,
	UserRolesReferee,
	UserRolesNone,
}

func (e UserRoles) IsValid() bool {
	switch e {
	case UserRolesFencer, UserRolesTrainer, UserRolesReferee, UserRolesNone:
		return true
	}
	return false
}

func (e UserRoles) String() string {
	return string(e)
}

func (e *UserRoles) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRoles(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRoles", str)
	}
	return nil
}

func (e UserRoles) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
