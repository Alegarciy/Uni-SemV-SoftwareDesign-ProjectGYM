package Requests

import (
	"API/Database"
	"API/Models"
	"fmt"
)

func CancelSession(pYear int, pMonth int, pWeekDay int, pRoomId  int, pHour string) Models.VoidOperationResult {
	query := fmt.Sprintf(`EXEC SP_DeletePreliminary %d, %d, %d, %d, %q `, pYear, pMonth, pWeekDay, pRoomId, pHour)

	return VoidRequest(query)
}

func GetPreliminarySchedule(pMonth int, pYear int) Models.PreliminarySchedule {

	query := fmt.Sprintf(`EXEC SP_GetPreliminarySchedule %d, %d;`, pMonth, pYear)

	resultSet, err := Database.ReadTransaction(query)

	if err != nil {
		return Models.PreliminarySchedule {}
	}

	schedule := ParsePreliminarySchedule(resultSet)

	return schedule

}

func DeletePreliminarySession(pSessionID int) Models.VoidOperationResult {
	// TODO: real db request

	dummyResult := Models.VoidOperationResult{Success: true}

	return dummyResult
}

func InsertPreliminarySession(	pName string,
								pWeekDay int,
								pMonth int,
								pYear int,
								pStartTime string,
								pDurationMins int,
								pService string,
								pInstructorIdentification string,
								pRoomId int,
							 ) 	Models.VoidOperationResult {

	query := fmt.Sprintf(`EXEC SP_InsertPreliminarySession %q, %d, %d, %d,%q, %d, %q, %q,%d;`,
		pName,
		pWeekDay,
		pMonth,
		pYear,
		pStartTime,
		pDurationMins,
		pService,
		pInstructorIdentification,
		pRoomId,
	)

	return VoidRequest(query)
}

func ConfirmPreliminarySchedule() Models.VoidOperationResult {
	// TODO: real db request

	dummyResult := Models.VoidOperationResult{Success: true}
	return dummyResult
}
