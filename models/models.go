package models

// Process documentation
type Process struct {
	ProcessID          string `json:"processId"`
	CustomerID         []CustomerID
	ProcessCreatedDate string `json:"processCreatedDate"`
}

// ProcessIDCaseID documentation
type ProcessIDCaseID struct {
	ProcessID string `json:"processId"`
	CaseID    string `json:"caseId"`
}

// CustomerID documentation
type CustomerID struct {
	CustomerID string `json:"customerId"`
}

// Household documentation
type Household struct {
	ProcessID            string `json:"processId"`
	HouseholdID          string `json:"householdId"`
	HouseholdMembers     []HouseholdMemberType
	NumberOfChildsAtHome int `json:"numberOfChildsAtHome"`
	Childs               []ChildType
	NumberOfCars         int     `json:"numberOfCars"`
	ChildMaintenanceCost float32 `json:"childMaintenanceCost"`
}

// HouseholdMemberType documentation
type HouseholdMemberType struct {
	CustomerIDs string `json:"householdMember"`
}

//
// ChildType documentation
type ChildType struct {
	ChildID         string `json:"childId"`
	ChildsAge       int    `json:"childsAge"`
	PartInHousehold bool   `json:"partInHousehold"`
}

// HouseholdID documentation
type HouseholdID struct {
	HouseholdID string `json:"householdId"`
}

// Applicant documentation
type Applicant struct {
	ProcessID             string `json:"processId"`
	CustomerID            string `json:"customerId"`
	ApplicantID           string `json:"applicantId"`
	ApplicantName         string `json:"applicantName"`
	ApplicantAddress      string `json:"applicantAddress"`
	ApplicantPostAddress  string `json:"applicantPostAddress"`
	StakeholderType       string `json:"stakeholderType"`
	ContactInformation    []ContactInformationType
	ApplicantEmployeed    bool   `json:"applicantEmployeed"`
	ApplicantLPEmployment string `json:"applicantLPEmployment"`
	ApplicantMember       bool   `json:"applicantMember"`
	ApplicantBySms        bool   `json:"applicantBySms"`
	ApplicantByeMail      bool   `json:"applicantByeMail"`
}

// UpdateApplicantType documentation
type UpdateApplicantType struct {
	ContactInformation []ContactInformationType
	StakeholderType    string `json:"stakeholderType"`
	ApplicantBySms     bool   `json:"applicantBySms"`
	ApplicantByeMail   bool   `json:"applicantByeMail"`
}

// ContactInformationType struct
type ContactInformationType struct {
	ApplicantEmail        string `json:"applicantByeMail"`
	ApplicantMobileNumber string `json:"applicantBySms"`
}

// Loan documentation
type Loan struct {
	ProcessID     string  `json:"processId"`
	LoanID        string  `json:"loanId"`
	LoanNumber    string  `json:"loanNumber"`
	LoanAmount    float32 `json:"loanAmount"`
	PurposeOfLoan string  `json:"purposeOfLoan"`
	Aims          []AimType
}

// LoanID documentation
type LoanID struct {
	LoanID string `json:"loanId"`
}

// AimType documentation
type AimType struct {
	AimID          string  `json:"aimId"`
	AimText        string  `json:"aimText"`
	LoanAmountPart float32 `json:"loanAmountPart"`
}

// ExtLoan documentation
type ExtLoan struct {
	ProcessID         string `json:"processId"`
	ExtLoanID         string `json:"extloanId"`
	ExtLoanOwners     []ExtLoanOwner
	ExtCreditInstitut string  `json:"extCreditInstitut"`
	ExtLoanClearing   string  `json:"extLoanClearing"`
	ExtLoanNumber     string  `json:"extloanNumber"`
	ExtLoanAmount     float32 `json:"extLoanAmount"`
	ExtRedeemLoan     bool    `json:"extRedeemLoan"`
}

//
// ExtLoanOwner documentation
type ExtLoanOwner struct {
	CustomerID string `json:"customerId"`
}

// ExtloanID documentation
type ExtLoanID struct {
	ExtLoanID string `json:"extloanId"`
}

// CompanyEconomy
type CompanyEconomy struct {
	ProcessID        string `json:"processId"`
	CompanyID        string `json:"companyId"`
	CompanyEconomyID string `json:"companyEconomyId"`
	Revenues         []Revenue
}

// CompanyEconomyID documentation
type CompanyEconomyID struct {
	CompanyEconomyID string `json:"companyEconomyId"`
}

// Revenue documentation
type Revenue struct {
	//RevenueID   string  `json:"revenueId"`
	RevenueYear int32   `json:"revenueYear"`
	Revenue     float32 `json:"revenue"`
}

// Company documentation
type Company struct {
	ProcessID       string `json:"processId"`
	CompanyID       string `json:"companyId"`
	OrgNumber       string `json:"orgNr"`
	CompanyName     string `json:"companyName"`
	Created         string `json:"created"`
	BusinessFocus   string `json:"businessFocus"`
	IndustriCode    string `json:"industriCode"` // SNI-kod
	IndustriText    string `json:"industriText"` // SNI-Text, kategori
	SelectedCompany bool   `json:"selectedCompany"`
}

// UpdateCompanyType documentation
type UpdateCompanyType struct {
	BusinessFocus string `json:"businessFocus"`
}

// CompanyID documentation
type CompanyID struct {
	CompanyID string `json:"companyId"`
}

// PersonalEconomy documentation
type PersonalEconomy struct {
	ProcessID          string  `json:"processId"`
	CustomerID         string  `json:"customerId"`
	PersonalEconomyID  string  `json:"personalEconomyId"`
	YearlyIncome       float32 `json:"yearlyIncome"`
	Income             float32 `json:"income"`
	TypeOfEmployeement string  `json:"typeOfEmployeement"`
	Employeer          string  `json:"employeer"`
	EmployeedFromYear  string  `json:"yearOfEmployment"`
}

// PersonalEconomyID documentation
type PersonalEconomyID struct {
	PersonalEconomyID string `json:"personalEconomyID"`
}

// Cases documentation
type Cases struct {
	CaseID string `json:"caseId"`
}

// Collateral documentation
type Collateral struct {
	ProcessID              string `json:"processId"`
	CustomerID             string `json:"customerId"`
	CollateralID           string `json:"collateralId"`
	CollateralCode         string `json:"collateralCode"`
	CollateralName         string `json:"collateralName"`
	TaxedOwners            []TaxedOwnerType
	CollateralMunicipality string `json:"collateralMunicipality"`
	CollateralStreet       string `json:"collateralStreet"`
	UseAsCollateral        bool   `json:"useAsCollateral"`
	BuyCollateral          bool   `json:"buyCollateral"`
}

// TaxedOwnerType documentation
type TaxedOwnerType struct {
	TaxedID    string `json:"taxedId"`
	TaxedOwner string `json:"taxedOwner"`
}

// CollateralID documentation
type CollateralID struct {
	CollateralID string `json:"collateralId"`
}

// KycInformation documentation
type KycInformation struct {
	ProcessID                string `json:"processId"`
	CustomerID               string `json:"customerId"`
	KycID                    string `json:"kycId"`
	KycAcceptUC              bool   `json:"kycAcceptUC"`
	KycAcceptGDPR            bool   `json:"kycAcceptGDPR"`
	KycUCAware               bool   `json:"kycUCAware"`
	KycPublicFunction        bool   `json:"kycPublicFunction"`
	KycRelatedPublicFunction bool   `json:"kycRelatedPublicFunction"`
}

// Budget documentation
type Budget struct {
	ProcessID        string `json:"processId"`
	CompanyEconomyID string `json:"companyEconomyId"`
	BudgetYears      []BudgetYear
}

// BudgetYear documentation
type BudgetYear struct {
	BudgetYear int     `json:"budgetYear"`
	Value1     float32 `json:"value1"`
	Value2     float32 `json:"value2"`
	Value3     float32 `json:"value3"`
	Value4     float32 `json:"value4"`
	Value5     float32 `json:"value5"`
	Value6     float32 `json:"value6"`
	Value7     float32 `json:"value7"`
	Value8     float32 `json:"value8"`
	Value9     float32 `json:"value9"`
	Value10    float32 `json:"value10"`
	Value11    float32 `json:"value11"`
	Value12    float32 `json:"value12"`
	Value13    float32 `json:"value13"`
	Value14    float32 `json:"value14"`
	Value15    float32 `json:"value15"`
	Value16    float32 `json:"value16"`
	Value17    float32 `json:"value17"`
	Value18    float32 `json:"value18"`
	Value19    float32 `json:"value19"`
	Value20    float32 `json:"value20"`
	Value21    float32 `json:"value21"`
	Value22    float32 `json:"value22"`
	Value23    float32 `json:"value23"`
	Value24    float32 `json:"value24"`
	Value25    float32 `json:"value25"`
}

//
// BudgetID documentation
//type BudgetID struct {
//	CompanyEconomyID string `json:"companyEconomyId"`
//}

// XAll documentation
type XAll struct {
	Persons []Person
	Animals []Animal
}

// Person documentation
type Person struct {
	Name   string
	Mobile string
}

// Animal documentation
type Animal struct {
	AnimalID string
	Animal   string
}
