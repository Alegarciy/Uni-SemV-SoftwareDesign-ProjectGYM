package Controllers

import (
	"API/Database/Requests"
	"API/WebServer/Common"
	"API/WebServer/Controllers/AssistanceVisitor"
	"API/WebServer/Controllers/FilteredScheduleStrategy"
	"API/WebServer/Controllers/Observer/FreeSpaceObserver"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var sessionFilter = &FilteredScheduleStrategy.ScheduleFilter{}
var assistanceVisitor = &AssistanceVisitor.AssistanceVisitor{}

var freeSpaceNotifier = &FreeSpaceObserver.FreeSpaceNotifier{}
var sessionsQueues = make(map[int]FreeSpaceObserver.SessionQueue)

func GetActiveSchedule(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}
	schedule := Requests.GetCurrentSessionSchedule()

	return Common.GiveJSONResponse(context, schedule, fiber.StatusOK)
}

func GetFilteredSchedule(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	filterType, _ := strconv.Atoi(data["filterType"])
	filterTerm := data["filterTerm"]

	strategies := []FilteredScheduleStrategy.FilterStrategy{
		&FilteredScheduleStrategy.FilterByInstructor{},
		&FilteredScheduleStrategy.FilterByServiceType{},
		&FilteredScheduleStrategy.FilterByDate{},
		&FilteredScheduleStrategy.FilterByTime{},
	}

	sessionFilter.SetFilterStrategy(strategies[filterType])

	schedule := sessionFilter.Filter(filterTerm)

	return Common.GiveJSONResponse(context, schedule, fiber.StatusOK)
}

func GetReservedSessions(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	membershipNumber, _ := strconv.Atoi(data["membershipNumber"])

	sessions := Requests.GetReservedSessions(membershipNumber)

	return Common.GiveJSONResponse(context, sessions, fiber.StatusOK)
}

func BookSession(context *fiber.Ctx) error {

	// token := Common.AnalyzeToken(context)
	// if token == nil {
	// 	return nil
	// }

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	clientnumber, _ := strconv.Atoi(data["clientIdentification"])
	date := data["date"]
	roomId, _ := strconv.Atoi(data["roomId"])
	startTime := data["startTime"]

	result := Requests.BookSession(clientnumber, date, roomId, startTime)

	// ? Se le puede poner a result el codigo de error?
	// Para saber si es por no haber espacio?
	if !result.Success {
		//queue, exist := sessionsQueues[]
		// Se puede revisar en un hash? Si existe la lista ya?
		timeDate := date + "T" + startTime + ":00"
		println(timeDate)
		t, _ := time.Parse("2006-01-02T15:04:05", timeDate)
		println(t.String())
	}

	return Common.GiveVoidOperationResponse(context, result)

}

func CancelBooking(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)
	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	clientnumber, _ := strconv.Atoi(data["clientIdentification"])
	date := data["date"]
	roomId, _ := strconv.Atoi(data["roomId"])
	startTime := data["startTime"]

	result := Requests.CancelBooking(clientnumber, date, roomId, startTime)

	if result.Success {
		freeSpaceNotifier.NotifyAll()
	}

	return Common.GiveVoidOperationResponse(context, result)

}

func CancelSession(context *fiber.Ctx) error {
	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	date := data["date"]
	roomId, _ := strconv.Atoi(data["roomId"])
	startTime := data["startTime"]

	result := Requests.CancelSession(date, roomId, startTime)

	return Common.GiveVoidOperationResponse(context, result)
}

func ChangeSessionInstructor(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	newInstructorNumber, _ := strconv.Atoi(data["instructorNumber"])
	date := data["date"]
	roomId, _ := strconv.Atoi(data["roomId"])
	startTime := data["startTime"]

	result := Requests.ChangeSessionInstructor(date, roomId, startTime, newInstructorNumber)

	return Common.GiveVoidOperationResponse(context, result)
}

func SetSessionAttendance(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}

	var data map[string]string
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	date := data["date"]
	roomId, _ := strconv.Atoi(data["roomId"])
	startTime := data["startTime"]

	attendants := data["attendants"]

	assistanceVisitor.Date = date
	assistanceVisitor.RoomId = roomId
	assistanceVisitor.StartTime = startTime

	attendantsSlice := strings.Split(attendants, ",")

	for _, attendantNumber := range attendantsSlice {
		intAttendantNumber, _ := strconv.Atoi(attendantNumber)
		bookingAttendance := &AssistanceVisitor.BookingAttendant{ClientMembershipNumber: intAttendantNumber}
		assistanceVisitor.VisitAttendant(bookingAttendance)
	}

	result := Requests.MarkSessionAttendanceTaken(date, roomId, startTime)

	return Common.GiveVoidOperationResponse(context, result)
}

func GetAttendancePendingSession(context *fiber.Ctx) error {

	token := Common.AnalyzeToken(context)

	if token == nil {
		return nil
	}
	schedule := Requests.GetAttendancePendingSessions()

	return Common.GiveJSONResponse(context, schedule, fiber.StatusOK)
}
