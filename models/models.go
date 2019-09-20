package models

// ProcessType documentation
type ProcessType struct {
	ProcessID          string           `json:"processId"`
	CustomerID         []CustomerIDType `json:"Customers"`
	CaseID             string           `json:"caseId"`
	CaseIDStatus       string           `json:"caseIdStatus"`
	ProcessCreatedDate string           `json:"processCreatedDate"`
	LastAccessed       string           `json:"lastAccessed"`
	ReferenceID        int              `json:"referenceId"`
	ProcessVersion     string           `json:"processVersion"`
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

// CustomerIDType documentation
type CustomerIDType struct {
	CustomerID    string `json:"customerId"`
	CustomerAdded string `json:"customerAdded"`
}

// HouseholdType documentation
type HouseholdType struct {
	ProcessID                    string `json:"processId"`
	HouseholdID                  string `json:"householdId"`
	HouseholdMembers             []HouseholdMemberType
	NumberOfChildsAtHome         int `json:"numberOfChildsAtHome"`
	Childs                       []ChildType
	NumberOfCars                 int     `json:"numberOfCars"`
	ChildMaintenanceCost         float32 `json:"childMaintenanceCost"`
	ReceivedChildMaintenanceCost float32 `json:"receivedchildMaintenanceCost"`
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

// ApplicantType documentation
type ApplicantType struct {
	ProcessID             string `json:"processId"`
	CustomerID            string `json:"customerId"`
	ApplicantID           string `json:"applicantId"`
	ApplicantName         string `json:"applicantName"`
	ApplicantAddress      string `json:"applicantAddress"`
	ApplicantPostalCode   string `json:"applicantPostalCode"`
	ApplicantPostAddress  string `json:"applicantPostAddress"`
	StakeholderType       string `json:"stakeholderType"`
	ContactInformation    []ContactInformationType
	ApplicantEmployed     bool   `json:"applicantEmployed"`
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

// TakeoverLoanType struct
type TakeoverLoanType struct {
	ProcessID       string  `json:"processId"`
	TakeoverLoanID  string  `json:"takeoverLoanId"`
	CreditInstitute string  `json:"creditInstitute"`
	LoanNumber      string  `json:"loanNumber"`
	DebtAmount      float32 `json:"debtAmount"`
	PurposeText     string  `json:"purposeText"`
	AimText         string  `json:"aimText"`
}

// TakeoverLoanID documentation
type TakeoverLoanID struct {
	TakeoverLoanID string `json:"takeoverLoanId"`
}

// LoanType documentation
type LoanType struct {
	ProcessID         string  `json:"processId"`
	LoanID            string  `json:"loanId"`
	LoanNumber        string  `json:"loanNumber"`
	LoanTakeOver      bool    `json:"loanTakeOver"`
	LoanAmount        float32 `json:"loanAmount"`
	DownPaymentType   string  `json:"downPaymentType"`
	DownPaymentOther  string  `json:"downPaymentOther"`
	DownPaymentAmount float32 `json:"downPaymentAmount"`
	//PurposeOfLoan     string  `json:"purposeOfLoan"`
	Aims []AimType
}

// LoanID documentation
type LoanID struct {
	LoanID string `json:"loanId"`
}

// AimType documentation
type AimType struct {
	AimID          string  `json:"aimId"`
	PurposeText    string  `json:"purposeText"`
	AimText        string  `json:"aimText"`
	LoanAmountPart float32 `json:"loanAmountPart"`
}

// ExtLoanType documentation
type ExtLoanType struct {
	ProcessID          string `json:"processId"`
	ExtLoanID          string `json:"extloanId"`
	ExtLoanOwners      []ExtLoanOwner
	TypeOfLoan         string  `json:"typeOfLoan"`
	ExtCreditInstitute string  `json:"extCreditInstitute"`
	ExtLoanClearing    string  `json:"extLoanClearing"`
	ExtLoanNumber      string  `json:"extloanNumber"`
	ExtLoanAmount      float32 `json:"extLoanAmount"`
	ExtRedeemLoan      bool    `json:"extRedeemLoan"`
	ExtMonthlyCost     float32 `json:"extMonthlyCost"`
}

//
// ExtLoanOwner documentation
type ExtLoanOwner struct {
	CustomerID string `json:"customerId"`
}

// ExtLoanID documentation
type ExtLoanID struct {
	ExtLoanID string `json:"extloanId"`
}

// CompanyEconomyType documentation
type CompanyEconomyType struct {
	ProcessID        string `json:"processId"`
	CompanyID        string `json:"companyId"`
	CompanyEconomyID string `json:"companyEconomyId"`
	CustomerCategory int16  `json:"customerCategory"` // Kundklass från SAS
	Revenues         []RevenueType
}

// CompanyEconomyIdBudgetID documentation
type CompanyEconomyIdBudgetID struct {
	CompanyEconomyID string `json:"companyEconomyId"`
	BudgetID         string `json:"budgetId"`
}

// RevenueType documentation
type RevenueType struct {
	//RevenueID   string  `json:"revenueId"`
	RevenueYear int32   `json:"revenueYear"`
	Revenue     float32 `json:"revenue"`
}

// CompanyEconomyID documentation
type CompanyEconomyID struct {
	CompanyEconomyID string `json:"companyEconomyId"`
}

// CompanyType documentation
type CompanyType struct {
	ProcessID       string `json:"processId"`
	CompanyID       string `json:"companyId"`
	OrgNumber       string `json:"orgNr"`
	CompanyName     string `json:"companyName"`
	StatusCode      string `json:"statusCode"`
	StatusTextHigh  string `json:"statusTextHigh"`
	Created         string `json:"created"`
	BusinessFocuses []BusinessFocusType
	//BusinessFocus   string `json:"businessFocus"`
	IndustriCode      string `json:"industriCode"` // SNI-kod
	IndustriText      string `json:"industriText"` // SNI-Text, kategori
	LegalGroupCode    string `json:"legalGroupCode"`
	LegalGroupText    string `json:"legalGroupText"`
	SelectedCompany   bool   `json:"selectedCompany"`
	Principals        []PrincipalType
	BrokenFiscalYear  bool   `json:"brokenFiscalYear"`
	FiscalYearEndDate string `json:"fiscalYearEndDate"`
}

// PrincipalType documentation
type PrincipalType struct {
	CustomerID    string `json:"customerId"`
	PrincipalName string `json:"principalName"`
}

// BusinessFocusType documentation
type BusinessFocusType struct {
	BusinessID        string `json:"businessId"`
	BusinessCategory  string `json:"businessCategory"`
	BusinessDirection string `json:"businessDirection"`
	BusinessPart      int    `json:"businessPart"`
	MainBusiness      bool   `json:"mainBusiness"`
}

// UpdateCompanyType documentation
type UpdateCompanyType struct {
	BusinessFocus string `json:"businessFocus"`
}

// CompanyID documentation
type CompanyID struct {
	CompanyID string `json:"companyId"`
}

// PersonalEconomyType documentation
type PersonalEconomyType struct {
	ProcessID          string  `json:"processId"`
	CustomerID         string  `json:"customerId"`
	PersonalEconomyID  string  `json:"personalEconomyId"`
	YearlyIncome       float32 `json:"yearlyIncome"`
	Income             float32 `json:"income"`
	TypeOfEmployeement string  `json:"typeOfEmployeement"`
	Employeer          string  `json:"employeer"`
	EmployeedFromYear  int     `json:"yearOfEmployment"`
	EmployeedFromMonth int     `json:"monthOfEmployment"`
}

// PersonalEconomyID documentation
type PersonalEconomyID struct {
	PersonalEconomyID string `json:"personalEconomyID"`
}

// Cases documentation
type Cases struct {
	CaseID string `json:"caseId"`
}

// CollateralType documentation
type CollateralType struct {
	ProcessID              string `json:"processId"`
	CustomerID             string `json:"customerId"`
	TypeOfCollateral       string `json:"typeOfCollateral"`
	CollateralID           string `json:"collateralId"`
	CollateralCode         string `json:"collateralCode"`
	CollateralName         string `json:"collateralName"`
	TaxedOwners            []TaxedOwnerType
	CollateralMunicipality string `json:"collateralMunicipality"`
	CollateralStreet       string `json:"collateralStreet"`
	CollateralAreal        int    `json:"collateralAreal"`
	CollateralAge          string `json:"collateralAge"`
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

// KycInformationType documentation
type KycInformationType struct {
	ProcessID                string `json:"processId"`
	CustomerID               string `json:"customerId"`
	KycID                    string `json:"kycId"`
	KycAcceptUC              bool   `json:"kycAcceptUC"`
	KycAcceptGDPR            bool   `json:"kycAcceptGDPR"`
	KycUCAware               bool   `json:"kycUCAware"`
	KycPublicFunction        bool   `json:"kycPublicFunction"`
	KycRelatedPublicFunction bool   `json:"kycRelatedPublicFunction"`
	KycHighRiskIndustry      bool   `json:"kycHighRiskIndustry"`
	KycHighRiskIndustryText  string `json:"kycHighRiskIndustryText"`
}

// KycID documentation
type KycID struct {
	KycID string `json:"kycId"`
}

// BudgetType documentation
type BudgetType struct {
	ProcessID        string          `json:"processId"`
	CompanyEconomyID string          `json:"companyEconomyId"`
	BudgetYears      []ValueYearType `json:"budgetYears"`
}

// BudgetID documentation
type BudgetID struct {
	BudgetID string `json:"budgetId"`
}

// ValueYearType documentation
type ValueYearType struct {
	BudgetID   string    `json:"budgetId"`
	BudgetYear int       `json:"budgetYear"`
	Value      ValueType `json:"value"`
}

// ValueType documentation
type ValueType struct {
	Value1  float32 `json:"value1"`
	Value2  float32 `json:"value2"`
	Value3  float32 `json:"value3"`
	Value4  float32 `json:"value4"`
	Value5  float32 `json:"value5"`
	Value6  float32 `json:"value6"`
	Value7  float32 `json:"value7"`
	Value8  float32 `json:"value8"`
	Value9  float32 `json:"value9"`
	Value10 float32 `json:"value10"`
	Value11 float32 `json:"value11"`
	Value12 float32 `json:"value12"`
	Value13 float32 `json:"value13"`
	Value14 float32 `json:"value14"`
	Value15 float32 `json:"value15"`
	Value16 float32 `json:"value16"`
	Value17 float32 `json:"value17"`
	Value18 float32 `json:"value18"`
	Value19 float32 `json:"value19"`
	Value20 float32 `json:"value20"`
	Value21 float32 `json:"value21"`
	Value22 float32 `json:"value22"`
	Value23 float32 `json:"value23"`
	Value24 float32 `json:"value24"`
	Value25 float32 `json:"value25"`
}

// SubmitApplication documentation
type SubmitApplication struct {
	SubmitApplication bool `json:"submitApplication"`
}

//
// Error struct
type Error struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// MessageBody documentation
type MessageBody struct {
	Status      bool
	MessageText string
	Filename    string
}

// ProcessAllType documentation
type ProcessAllType struct {
	Processes         []ProcessType
	Applicants        []ApplicantType
	Loans             []LoanType
	ExtLoans          []ExtLoanType
	TakeoverLoans     []TakeoverLoanType
	Households        []HouseholdType
	Companies         []CompanyType
	CompanyEconomies  []CompanyEconomyType
	PersonalEconomies []PersonalEconomyType
	Collaterals       []CollateralType
	KycInformations   []KycInformationType
	Budgets           []BudgetType
	EUSupports        []EUSupportType
	Guarantors        []GuarantorType
	MaintenanceCosts  []MaintenanceCostType
}

// TextValue documentation
type TextValue struct {
	Text string
}

// EUID documentation
type EUID struct {
	EUID string `json:"euId"`
}

// EUSupportType documentation
type EUSupportType struct {
	ProcessID     string  `json:"processId"`
	EUID          string  `json:"euId"`
	EUType        string  `json:"EUType"`
	SupportAmount float32 `json:"supportAmount"`
	SupportYear   string  `json:"supportYear"`
}

// GuarantorID documentation
type GuarantorID struct {
	GuarantorID string `json:"guarantorId"`
}

// GuarantorType documentation
type GuarantorType struct {
	ProcessID           string `json:"processId"`
	GuarantorID         string `json:"guarantorId"`
	GuarantorName       string `json:"guarantorName"`
	GuarantorPhone      string `json:"guarantorPhone"`
	GuarantorCustomerID string `json:"guarantorCustomerId"`
}

// MaintenanceCostID documentation
type MaintenanceCostID struct {
	MaintenanceCostID string `json:"maintenanceCostId"`
}

// MaintenanceCostType documentation
type MaintenanceCostType struct {
	ProcessID         string `json:"processId"`
	MaintenanceCostID string `json:"maintenanceCostId"`
	CustomerID        string `json:"customerId"`
	TypeOfHouses      []TypeOfHouse
}

// TypeOfHouse (driftkostnad på andra boenden)
type TypeOfHouse struct {
	HouseID              string `json:"houseId"`
	TypeOfHouse          string `json:"typeOfHouse"`
	KeepHouse            bool   `json:"keepHouse"`
	LoanInOtherInstitute bool   `json:"loanInOtherInstitute"`
	RedeemLoan           bool   `json:"redeemLoan"`
	CreditInstitute      string `json:"creditInstitute"`
	LoanClearing         string `json:"loanClearing"`
	InstituteLoanNumber  string `json:"instituteLoanNumber"`
	LoanOwners           []LoanOwnerType
	MaintenanceCost      float32 `json:"maintenanceCost"`
	LoanAmount           float32 `json:"loanAmount"`
}

// LoanOwnerType documentation
type LoanOwnerType struct {
	CustomerID    string `json:"customerId"`
	OwnershipPart int    `json:"ownershipPart"`
}

// Header struct documentation
type Header struct {
	AccessToken   string
	TokenType     string
	Scope         string
	ExpiresIn     string
	Authorization string
}

// PurposeTextType documentation
type PurposeTextType struct {
	ID              string        `json:"id"`
	MainText        string        `json:"mainText"`
	MainDisplayText string        `json:"mainDisplayText"`
	Aims            []AimTextType `json:"aims"`
}

// AimTextType documentation
type AimTextType struct {
	AimText        string `json:"aimText"`
	AimDisplayText string `json:"aimDisplayText"`
}
