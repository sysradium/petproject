package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sysradium/petproject/orders-api/api/models"
)

type Handler struct{}

// Returns a list of orders.
// (GET /orders)
func (s *Handler) GetOrders(ctx echo.Context) error {
	var rsp []*models.Order

	rsp = append(rsp, &models.Order{Id: uuid.MustParse("9cb14230-b640-11ec-b909-0242ac120002")})

	ctx.JSON(http.StatusOK, rsp)
	return nil
}

// Creates a new order
// (POST /orders)
func (s *Handler) PostOrders(ctx echo.Context) error {
	u := &models.PostOrdersJSONRequestBody{}
	if err := ctx.Bind(u); err != nil {
		return err
	}

	ctx.NoContent(http.StatusCreated)
	return nil
}

func New() *Handler {
	return &Handler{}
}
