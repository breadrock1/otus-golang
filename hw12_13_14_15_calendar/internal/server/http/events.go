package internalhttp

import (
	"encoding/json"
	"fmt"
	_ "github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/docs"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
	"github.com/labstack/echo/v4"
	"io"
	"strconv"
)

func (s *Server) CreateEventsGroup() {
	group := s.server.Group("/calendar")

	group.POST("/create", s.CreateEvent)
	group.PUT("/event/:id", s.UpdateEvent)
	group.DELETE("/event/:id", s.DeleteEvent)

	group.POST("/list/day", s.ListEventsPerDay)
	group.POST("/list/week", s.ListEventsPerWeek)
	group.POST("/list/month", s.ListEventsPerMonth)
}

// CreateEvent
// @Summary Create event
// @Description Create new event by form
// @ID create
// @Tags calendar
// @Produce json
// @Param jsonQuery body event.Event true "Event to create"
// @Success 200 {object} ResponseForm "Created event: 1345"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/create [post]
func (s *Server) CreateEvent(c echo.Context) error {
	eventForm := &event.Event{}
	form, err := s.extractForm(c.Request().Body, eventForm)
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	var eventID int
	cxt := c.Request().Context()
	ev := form.(*event.Event)
	if eventID, err = s.app.CreateEvent(cxt, *ev); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	msg := fmt.Sprintf("Created event: %d", eventID)
	return c.JSON(200, createStatusResponse(200, msg))
}

// UpdateEvent
// @Summary Update event
// @Description Updated existing event by form
// @ID update
// @Tags calendar
// @Produce json
// @Param id path int true "Event id"
// @Param jsonQuery body event.Event true "Event to update"
// @Success 200 {object} ResponseForm "Ok"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/event/{id} [put]
func (s *Server) UpdateEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	eventForm := &event.Event{}
	form, err := s.extractForm(c.Request().Body, eventForm)
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	cxt := c.Request().Context()
	ev := form.(*event.Event)
	if err = s.app.UpdateEvent(cxt, eventID, *ev); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	return c.JSON(200, createStatusResponse(200, "Updated!"))
}

// DeleteEvent
// @Summary Delete event
// @Description Delete existing event by form
// @ID delete
// @Tags calendar
// @Produce json
// @Param id path int true "Event id"
// @Success 200 {object} ResponseForm "Ok"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/event/{id} [delete]
func (s *Server) DeleteEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	cxt := c.Request().Context()
	if err = s.app.Delete(cxt, eventID); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	return c.JSON(200, createStatusResponse(200, "Deleted!"))
}

// ListEventsPerDay
// @Summary Get all events per day
// @Description Get all events per day by datetime form
// @ID list-per-day
// @Tags calendar
// @Produce json
// @Param jsonQuery body DatetimeForm true "Get events per day"
// @Success 200 {object} ResponseForm "Ok"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/list/day [post]
func (s *Server) ListEventsPerDay(c echo.Context) error {
	datetimeForm := &DatetimeForm{}
	form, err := s.extractForm(c.Request().Body, datetimeForm)
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	var perDay []event.Event
	cxt := c.Request().Context()
	datetime := form.(*DatetimeForm)
	if perDay, err = s.app.ListDay(cxt, *datetime.Datetime); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	return c.JSON(200, perDay)
}

// ListEventsPerWeek
// @Summary Get all events per week
// @Description Get all events per week by datetime form
// @ID list-per-week
// @Tags calendar
// @Produce json
// @Param jsonQuery body DatetimeForm true "Get events per week"
// @Success 200 {object} ResponseForm "Ok"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/list/week [post]
func (s *Server) ListEventsPerWeek(c echo.Context) error {
	datetimeForm := &DatetimeForm{}
	form, err := s.extractForm(c.Request().Body, datetimeForm)
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	var perWeek []event.Event
	cxt := c.Request().Context()
	datetime := form.(*DatetimeForm)
	if perWeek, err = s.app.ListDay(cxt, *datetime.Datetime); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	return c.JSON(200, perWeek)
}

// ListEventsPerMonth
// @Summary Get all events per month
// @Description Get all events per month by datetime form
// @ID list-per-month
// @Tags calendar
// @Produce json
// @Param jsonQuery body DatetimeForm true "Get events per month"
// @Success 200 {object} ResponseForm "Ok"
// @Failure 400 {object} BadRequestForm "Bad request message"
// @Failure	503 {object} ServerErrorForm "Server does not available"
// @Router /calendar/list/month [post]
func (s *Server) ListEventsPerMonth(c echo.Context) error {
	datetimeForm := &DatetimeForm{}
	form, err := s.extractForm(c.Request().Body, datetimeForm)
	if err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	var perMonth []event.Event
	cxt := c.Request().Context()
	datetime := form.(*DatetimeForm)
	if perMonth, err = s.app.ListDay(cxt, *datetime.Datetime); err != nil {
		respErr := createStatusResponse(400, err.Error())
		return c.JSON(400, respErr)
	}

	return c.JSON(200, perMonth)
}

func (s *Server) extractForm(body io.ReadCloser, form interface{}) (interface{}, error) {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(form); err != nil {
		return nil, err
	}
	return form, nil
}
