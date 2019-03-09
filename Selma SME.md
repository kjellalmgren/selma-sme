# Selma SME

## Informationsinsamling

reserverCaseId ska vid anrop första gången även anropa setCaseIdStatus

	Method: reserveCaseId/processId/{customerId}
	Method: setCaseIdStatus/{processId}/{caseId}/{caseIdStatus}
	caseIdStatus: Enum: STARTEDBYAPPLICANT
	
LIME ska uppdateras att kund startat ansökan, befintlig kund, ny kund. LIME kundkort ska logga aktiviteten.



## To Do
#### General
	Se över alla omitempty i koden (go-code)
	I alla Type definitions, se över required
#### LoanType
	TBD: loanNumber has to be changed to a uuid
	TBD: purposeOfLoan has to be changed
	
#### ExtLoanType
	DONE: add extloanAmount to the Type
TBD: can we remove customerId from ExtLoanType or do we need it for KALP. ExtLoan till hushållet kanske?
	
#### HosueholdMemberType
	DONE: We have changed the naming convention regarding household members	
#### PersonalEconomies
	TBD: Årlig inkomst från UC måste nog läggas till
	DONE: yearOfEmployment has changed name

#### CompanyEconomies
	DONE: Added companyId to point to company

#### Budgets mybe should be under CompanyEconomies as a array list

#### CollateralType should remove customerId
	We probaly don´t need customerId in type of collateral
	
# Processes
GetProcess
När en ny process startas ska 