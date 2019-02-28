# Selma SME

## Informationsinsamling

reserverCaseId ska vid anrop första gången även anropa setCaseIdStatus

	Method: reserveCaseId/processId/{customerId}
	Method: setCaseIdStatus/{processId}/{caseId}/{caseIdStatus}
	caseIdStatus: Enum: STARTEDBYAPPLICANT
	
LIME ska uppdateras att kund startat ansökan, befintlig kund, ny kund. LIME kundkort ska logga aktiviteten.

