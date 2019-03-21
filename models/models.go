package models

// Process struct
type Process struct {
	ProcessID          string `json:"processId"`
	CustomerID         []CustomerID
	ProcessCreatedDate string `json:"processCreatedDate"`
}

// CustomerID
type CustomerID struct {
	CustomerID string `json:"customerId"`
}

// Household
type Household struct {
	ProcessID            string `json:"processId"`
	HouseholdMembers     []HouseholdMemberType
	HouseholdID          string `json:"householdId"`
	NumberOfChildsAtHome int    `json:"numberOfChildsAtHome"`
	Childs               []ChildType
	NumberOfCars         int     `json:"numberOfCars"`
	ChildMaintenanceCost float32 `json:"childMaintenanceCost"`
}

// HouseholdMemberType
type HouseholdMemberType struct {
	CustomerIDs string `json:"householdMember"`
}

//
// ChildType
type ChildType struct {
	ChildID         string `json:"childId"`
	ChildsAge       int    `json:"childsAge"`
	PartInHousehold bool   `json:"partInHousehold"`
}

type HouseholdID struct {
	HosueholdID string `json:"householdId"`
}

// Applicant struct
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

// UpdateApplicantType
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

// Loan
type Loan struct {
	ProcessID     string  `json:"processId"`
	LoanID        string  `json:"loanId"`
	LoanNumber    string  `json:"loanNumber"`
	LoanAmount    float32 `json:"loanAmount"`
	PurposeOfLoan string  `json:"purposeOfLoan"`
	Aims          []AimType
}

// LoanId
type LoanID struct {
	LoanID string `json:"loanId"`
}

// aimType
type AimType struct {
	AimID          string  `json:"aimId"`
	AimText        string  `json:"aimText"`
	LoanAmountPart float32 `json:"loanAmountPart"`
}

// ExtLoan
type ExtLoan struct {
	ProcessID         string `json:"processId"`
	ExtLoanOwners     []ExtLoanOwner
	ExtLoanID         string  `json:"extloanId"`
	ExtCreditInstitut string  `json:"extCreditInstitut"`
	ExtLoanClearing   string  `json:"extLoanClearing"`
	ExtLoanNumber     string  `json:"extloanNumber"`
	ExtLoanAmount     float32 `json:"extLoanAmount"`
	ExtRedeemLoan     bool    `json:"extRedeemLoan"`
}

//
// ExtLoanOwner
type ExtLoanOwner struct {
	CustomerID string `json:"customerId"`
}

// ExtloanID
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

// CompanyEconomyID
type CompanyEconomyID struct {
	CompanyEconomyID string `json:"companyEconomyId"`
}

// Revenue
type Revenue struct {
	RevenueID   string  `json:"revenueId"`
	RevenueYear int32   `json:"revenueYear"`
	Revenue     float32 `json:"revenue"`
}

// Company
type Company struct {
	ProcessID       string `json:"processId"`
	CompanyID       string `json:"companyId"`
	OrgNumber       string `json:"orgNr"`
	CompanyName     string `json:"companyName"`
	NumberOfLoans   string `json:"numberOfLoans"`
	Created         string `json:"created"`
	BusinessFocus   string `json:"businessFocus"`
	SelectedCompany bool   `json:"selectedCompany"`
}

// UpdateCopmpanyType
type UpdateCompanyType struct {
	BusinessFocus string `json:"businessFocus"`
}

// CompanyID
type CompanyID struct {
	CompanyID string `json:"companyId"`
}

// PersonalEconomy
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

//
type PersonalEconomyID struct {
	PersonalEconomyID string `json:"personalEconomyID"`
}

// Cases struct
type Cases struct {
	CaseID string `json:"caseId"`
}

// Collateral
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

// TaxedOwnerType
type TaxedOwnerType struct {
	TaxedID    string `json:"taxedId"`
	TaxedOwner string `json:"taxedOwner"`
}

// CollateralID
type CollateralID struct {
	CollateralID string `json:"collateralId"`
}

// KycInformation
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

// Xall
type XAll struct {
	Persons []Person
	Animals []Animal
}

// Person
type Person struct {
	Name   string
	Mobile string
}

// Animal
type Animal struct {
	AnimalID string
	Animal   string
}
