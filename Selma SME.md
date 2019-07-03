# Selma SME

## Informationsinsamling

**Vx**
<!--reserverCaseId ska vid anrop första gången även anropa setCaseIdStatus

	Method: reserveCaseId/processId/{customerId}
	Method: setCaseIdStatus/{processId}/{caseId}/{caseIdStatus}
	caseIdStatus: Enum: STARTEDBYAPPLICANT
	-->
**Vx** LIME ska uppdateras att kund startat ansökan, befintlig kund, ny kund. LIME kundkort ska logga aktiviteten.

**CustomerId**: "19640120-3887": "d5744655-b71e-428a-98b9-2b6c66c8c95a"

**CustomerId**: "19650705-5579": "b2f86b36-7ff3-428e-ab45-8dad11952dae"

	
## Processes
**Nyckelbegrepp**
När en ny ansökan inkommer ska vi alltid skapa ett ProcessID, denna nyckel håller ihop hela ansökan. Under processId kan fler kunder finnas, representeras av customerId. Elementen processCreatedDate ska sättas när ansökan första gången skapas och elementet lastAccessed ska uppdateras när någon kunderna customerId har varit inne på ansökan. Dessa värden tillsammans med caseIdStatus används vid gallringen.

Genomgående i modellen är ProcessID, denna uuid håller ihop respektive ansökan, i REST-API så ingår det i http-header som "X-process-ID".

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"CustomerID": [
      	{
        	"customerId": "19640120-3887",
			"customerCreated": 2019-05-02T09:00:00"
      	},
      	{
        	"customerId": "19650705-5579",
			"customerCreated": 2019-05-02T19:00:00"
      	}
    	],
    	"processCreatedDate": "2019-06-01",
    	"lastAccessed": "2019-06-10T14:20:30"
  	}
	]

**caseIdStatus har följande värden**

	enum:
    - "CLOSEDBYAPPLICANT" # Sökande kund har valt att stänga ansökan
    - "STARTEDBYAPPLICANT" # Sökande kund har startat ett insamlingsflöde
    - "CLOSEDBYOFFICER"		# Handläggare stängt ärende, uppdatera insamingsflödet (Vx)
    - "CLOSEDBYTHINNING" # gallring (Vx)
    - "READYFOROFFICER" # inläst i LP (Vx)

    
	#Header parameters
	parameters:
        - $ref: '#/parameters/caseIdStatus'
    responses:
    	processIdCaseIdType:
    	description: return value after addProcess and caseid reservation i Loan Process
    	required:
      	- processId
      	- caseId
    	type: object
    	properties:
      		processId:
        		$ref: '#/definitions/processId'
      		caseId:
        		$ref: '#/definitions/caseId' # Ignoreras i V1
        		
    
Change OperationId på /addprocess, flytta in caseIdStatus som headerparameter. Ändra required. tabort caseId

**V2**
När en ny process startas ska ska ProcessId, CustomerId samt CaseID reservras i Loan Process.

### getprocessall

Hämta hela modellen för ett givet processid

### saveprocessall

Spara hela modellen för ett givet processid


## Applicant
**Sökande / (StakeholderType)**
Efter vi har fått CustomerID via BankID kommer vi anropa addApplicant med Personnummer, tillbaka kommer en uppdaterad ApplicantID (uuid), denna ska användas i kommande kommunikation med modellen och den givna sökande.

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19640120-3887",
    	"applicantId": "", 
    	"applicantName": "Anna Andersson",
    	"applicantAddress": "Stora vägen 1",
    	"applicantPostAddress": "420 20 Katrineholm",
    	"stakeholderType": "",
    	"ContactInformation": [
      	{
       	"applicantByeMail": "anna.andersson@gmail.com",
      		"applicantBySms": "07344455666"
      	}
    	],
    	"applicantEmployeed": false,
    	"applicantLPEmployment": "PERMANENT",
    	"applicantMember": false,
    	"applicantBySms": true,
    	"applicantByeMail": true
  	},
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19650705-5579",
    	"applicantId": "",
    	"applicantName": "Patrik Andersson",
    	"applicantAddress": "Stora vägen 1",
    	"applicantPostAddress": "420 20 Katrineholm",
    	"stakeholderType": "",
    	"ContactInformation": [
      	{
        	"applicantByeMail": "patrik.andersson@katrineholmrevision.se",
        	"applicantBySms": "07335533777"
      	}
    	],
    	"applicantEmployeed": false,
    	"applicantLPEmployment": "PERMANENT",
    	"applicantMember": false,
    	"applicantBySms": true,
    	"applicantByeMail": true
  	}
	]

## Companies
**Företag**

	[
	  {
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"companyId": "461460c2-3d14-11e9-b210-d663bd873d93",
    	"orgNr": "551010-8474",
    	"statusCode": "100",
    	"statusTextHigh": "Aktivt",
    	"companyName": "Anna Andersson Skog och djurhållning",
    	"created": "2012-01-01",
    	"businessFocuses": [
      	{
        	"businessId": "5c497a6e-3a5b-44a8-8302-21012cad32a2",
        	"businessCategory": "Jordbruksväxter",
        	"businessDirection": "Spanmål & oljeväxt.",
        	"businessPart": 100
      	}
    	],
    	"industriCode": "28300",
    	"industriText": "Manufacture of agricultural and forestry machinery",
    	"selectedCompany": true
  		}
	]

## CompanyEconomies
**Company Economies**

	[
	  {
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"companyId": "02d6a03e-5895-4077-98f7-7a5c192868b7",
    	"companyEconomyId": "4804f0c2-3d2d-11e9-b210-d663bd873d93",
    	"customerCategory": 1,
    	"Revenues": [
      	{
        	"revenueYear": 2016,
        	"revenue": 2000000
      	},
      	{
        	"revenueYear": 2017,
        	"revenue": 2100000
      	},
      	{
        	"revenueYear": 2018,
        	"revenue": 2300000
      	}
    	]
  	}
	]

## PersonalEconomies
**Personal Economies**

	[
 	 {
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19640120-3887",
    	"personalEconomyId": "52a50f80-3d28-11e9-b210-d663bd873d93",
    	"yearlyIncome": 340000,
    	"income": 0,
    	"typeOfEmployeement": "Tillsvidare",
    	"employeer": "Anna Andersson Skog och djurhållning",
    	"yearOfEmployment": "2012",
    	"mounthOfEmployment": "10"
  	},
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19650705-5579",
    	"personalEconomyId": "bf5ea49c-3d28-11e9-b210-d663bd873d93",
    	"yearlyIncome": 0,
    	"income": 460000,
    	"typeOfEmployeement": "Tillsvidare",
    	"employeer": "Katrineholm Revision AB",
    	"yearOfEmployment": "2009",
    	"mounthOfEmployment": "11"
  	}
	]
	
## Loans
**Lån**

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"loanId": "9b8e4822-3cb7-11e9-b210-d663bd873d93",
    	"loanNumber": "930101011212",
    	"loanTakeOver": false,
    	"loanAmount": 2300000,
    	"downPaymentType": "Övrigt",
	   	"downPaymentOther": "The mob",
    	"downPaymentAmount": 300000,
    	"purposeOfLoan": "Köp",
    	"Aims": [
      	{
        	"aimId": "fce3d0aa-4b04-11e9-8646-d663bd873d93",
        	"aimText": "Fastighetsköp - annan fastighet",
        	"loanAmountPart": 2000000
      	},
      	{
        	"aimId": "fce3d0aa-4b04-11e9-8646-d663bd873d94",
        	"aimText": "Renovering mjölkstall",
        	"loanAmountPart": 300000
      	}
    	]
  	}
	]

## ExtLoans
**Externa lån**

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"extloanId": "5aa735e8-3cbd-11e9-b210-d663bd873d93",
    	"ExtLoanOwners": [
      	{
        	"customerId": "19640120-3887"
      	}
    	],
    	"extCreditInstitut": "SEB",
    	"extLoanClearing": "5270",
    	"extloanNumber": "0028600",
    	"extLoanAmount": 100000,
    	"extRedeemLoan": true
  	}
	]

## Households
**Hushåll**

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"householdId": "4253017a-3d17-11e9-b210-d663bd873d93",
    	"HouseholdMembers": [
      	{
        	"householdMember": "19640120-3887"
      	},
      	{
        	"householdMember": "19650705-5579"
      	}
    	],
    	"numberOfChildsAtHome": 2,
    	"Childs": [
      	{
        	"childId": "248485ca-3d9d-11e9-b210-d663bd873d93",
        	"childsAge": 5,
        	"partInHousehold": true
      	},
      	{
        	"childId": "eb38da7c-3d9d-11e9-b210-d663bd873d93",
        	"childsAge": 8,
        	"partInHousehold": true
      	}
    	],
    	"numberOfCars": 1,
    	"childMaintenanceCost": 0
  	}
	]

## KycInformations
**Know your customer information**

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19640120-3887",
    	"kycId": "9bca3a55-2458-41d5-9420-f12bdcd0c809",
    	"kycAcceptUC": false,
    	"kycAcceptGDPR": false,
    	"kycUCAware": false,
    	"kycPublicFunction": false,
    	"kycRelatedPublicFunction": false,
    	"kycHighRiskIndustry": false
  	}
	]
	

## Collaterals
**Säkerheter (i form av fatighet)**

	[
  	{
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"customerId": "19640120-3887",
    	"typeOfCollateral": "FASTIGHET",
    	"collateralId": "b25961de-3cc3-11e9-b210-d663bd873d93",
    	"collateralCode": "110",
    	"collateralName": "ÅGERSTA 1:6",
    	"TaxedOwners": [
      	{
        	"taxedId": "c73119bc-3e71-11e9-b210-d663bd873d93",
        	"taxedOwner": "Anna Andersson"
      	}
    	],
    	"collateralMunicipality": "ENKÖPING",
    	"collateralStreet": "Bergsgatan 12",
    	"useAsCollateral": false,
    	"buyCollateral": true
  	}
	]
	
## Budgets
**Budget**

	[
	  {
    	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
    	"companyEconomyId": "4804f0c2-3d2d-11e9-b210-d663bd873d93",
    	"BudgetYears": [
      	{
        	"budgetId": "6737d7be-520f-47f5-a761-b24c1db2df5e",
        	"budgetYear": 2019,
        	"Values": [
          {
          	"value1": 210000,
          	"value2": 100000,
          	"value3": 200000,
            	"value4": 200000,
            	"value5": 200000,
            	"value6": 200000,
            	"value7": 200000,
            	"value8": 200000,
            	"value9": 200000,
            	"value10": 200000,
            	"value11": 10000,
            	"value12": 10000,
            	"value13": 300000,
            	"value14": 20000,
            	"value15": 30000,
            	"value16": 200000,
            	"value17": 20000,
            	"value18": 20000,
            	"value19": 20000,
            	"value20": 20000,
            	"value21": 20000,
            	"value22": 20000,
            	"value23": 20000,
            	"value24": 20000,
            	"value25": 20000
          	}
        	]
      	},
      	{
        	"budgetId": "e03f4a81-0c6a-49a7-a421-697a850f78cb",
        	"budgetYear": 2020,
        	"Values": [
          {
            	"value1": 220000,
            	"value2": 110000,
            	"value3": 200000,
            	"value4": 200000,
            	"value5": 200000,
            	"value6": 200000,
            	"value7": 200000,
            	"value8": 200000,
            	"value9": 200000,
            	"value10": 200000,
            	"value11": 10000,
            	"value12": 10000,
            	"value13": 300000,
            	"value14": 20000,
            	"value15": 30000,
            	"value16": 200000,
            	"value17": 20000,
            	"value18": 20000,
            	"value19": 20000,
            	"value20": 20000,
            	"value21": 20000,
            	"value22": 20000,
            	"value23": 20000,
            	"value24": 20000,
            	"value25": 20000
          }
        	]
      	}
    	] 
  	}
	]
	
## EUSupports
**EU-Stöd**

	{
  	"eusupports": [
    	{
      		"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
      		"euId": "20083dfd-b3c0-4cd4-ad26-47421124a8f6",
      		"euType": "EU-stöd och nationellt stöd",
      		"supportAmount": 850000.00,
      		"supportYear": "2019"
    	}
  	]
	}
	
## Guarantors
**Borgensmän**

	{
    "guarantors": [
      {
        "processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
        "gaurantorId": "",
        "gaurantorName": "Anna Andersson",
        "gaurantorPhone": "07012332144",
        "gaurantorCustomerId": "19640120-3887"
      }
    ]
  }
    	
## MaintananceCosts
**Driftkostnad övriga boende**

	{
  	"processId": "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
  	"maintenancecostid": "9e45fa0f-641a-4302-b236-c3b9df4de086",
  	"typeofhouses": [
   	 {
      "keepHouse": true,
      "loanInOtherInstitut": true,
      "redeemLoan": true,
      "loanOwner": "Anna Andersson",
      "houseLoanAmount": 850000,
      "maintenanceCost": 4500
   	}
  	]
	}