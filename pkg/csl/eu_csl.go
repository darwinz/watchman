package csl

// CLS - Consolidated List Sanctions from European Union
// TODO: does this need to be in csl (ask Adam from moov)

// TODO: get this from env
// download uri
// https://webgate.ec.europa.eu/fsd/fsf/public/files/csvFullSanctionsList_1_1/content?token=dG9rZW4tMjAxNw
// protocol: https://
// hostname: webgate.ec.europa.eu
// path: /fsd/fsf/public/files/csvFullSanctionsList_1_1/content
// query: ?token=dG9rZW4tMjAxNw

// struct to hold the rows from the csv data
type EUCSL map[int]*EUCSLRow

// type EUCSLSheet []*EUCSLRow

type EUCSLRow struct {
	FileGenerationDate string            `json:"fileGenerationDate"`
	Entity             *Entity           `json:"entity"`
	NameAliases        []*NameAlias      `json:"nameAliases"`
	Addresses          []*Address        `json:"addresses"`
	BirthDates         []*BirthDate      `json:"birthDates"`
	Identifications    []*Identification `json:"identifications"`
}

type Entity struct {
	LogicalID       int    `json:"logicalId"`
	ReferenceNumber string `json:"referenceNumber"`
	// UnitiedNationsID   string
	// DesignationDate    string
	// DesignationDetails string
	Remark      string       `json:"remark"`
	SubjectType *SubjectType `json:"subjectType"`
	Regulation  *Regulation  `json:"regulation"`
}
type SubjectType struct {
	// SingleLetter       string
	ClassificationCode string `json:"classificationCode"`
}
type Regulation struct {
	// Type               string
	// OrganizationType   string
	// PublicationDate    string
	// EntryInfoForceDate string
	// NumberTitle        string
	// Programme          string
	PublicationURL string `json:"publicationURL"`
}

type NameAlias struct { // AltNames
	// LastName           string
	// FirstName          string
	// MiddleName         string
	WholeName string `json:"wholeName"`
	// NameLanguage       string
	// Gender             string
	Title string `json:"title"`
	// Function           string
	// LogicalID          int64
	// RegulationLanguage string
	// Remark             string
	// Regulation         *Regulation
}
type Address struct { // addresses
	City    string `json:"city"`
	Street  string `json:"street"`
	PoBox   string `json:"poBox"`
	ZipCode string `json:"zipCode"`
	// Region             string
	// Place              string
	// AsAtListingTime    string
	// ContactInfo        string
	// CountryIso2code    string
	CountryDescription string `json:"countryDescription"`
	// LogicalID          int64
	// RegulationLanguage string
	// Remark             string
	// Regulation         *Regulation
}
type BirthDate struct {
	BirthDate string // keep
	// Day                int64
	// Month              int64
	// Year               int64
	// YearRangeFrom      string
	// YearRangeTo        string
	// Circa              string
	// CaldendarType      string // TODO: this could be an enum
	// ZipCode            string
	// Region             string
	// Place              string
	City string `json:"city"`
	// CountryIso2code    string
	CountryDescription string `json:"countryDescription"`
	// LogicalID          int64
	// Regulation         *Regulation
}

type Identification struct {
	// Regulation         *Regulation
	// Number             int64
	// KnownExpired       bool
	// KnownFalse         bool
	// ReportedLost       bool
	// RevokedByIssuer    bool
	// LogicalID          int64
	// Diplomatic         string // TODO: not sure about this field
	// IssuedBy           string
	// IssuedDate         string
	ValidFrom string `json:"validFrom"`
	ValidTo   string `json:"validTo"`
	// NameOnDocument     string
	// TypeCode           string
	// TypeDescription    string
	// Region             string
	// CountryIso2code    string
	// CountryDescription string
	// RegulationLanguage string
	// Remark             string
}

// header indicies
const (
	FileGenerationDateIdx             = 0
	EntityLogicalIdx                  = 1
	ReferenceNumberIdx                = 2
	EntityRemarkIdx                   = 6
	EntitySubjectTypeIdx              = 8
	EntityRegulationPublicationURLIdx = 15

	NameAliasWholeNameIdx = 19
	NameAliasTitleIdx     = 22

	AddressCityIdx               = 34
	AddressStreetIdx             = 35
	AddressPoBoxIdx              = 36
	AddressZipCodeIdx            = 37
	AddressCountryDescriptionIdx = 43

	BirthDateIdx        = 54
	BirthDateCityIdx    = 65
	BirthDateCountryIdx = 67

	IdentificationValidFromIdx = 86
	IdentificationValidToIdx   = 87
)

func NewEUCSLRow() *EUCSLRow {
	row := new(EUCSLRow)
	row.Entity = new(Entity)
	row.Entity.SubjectType = new(SubjectType)
	row.Entity.Regulation = new(Regulation)
	row.NameAliases = make([]*NameAlias, 0)
	row.Addresses = make([]*Address, 0)
	row.BirthDates = make([]*BirthDate, 0)
	row.Identifications = make([]*Identification, 0)

	return row
}
