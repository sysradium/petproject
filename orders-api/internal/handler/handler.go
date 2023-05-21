package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sysradium/petproject/orders-api/api/models"
	pbUsers "github.com/sysradium/petproject/users-api/proto/users/v1"
)

type Handler struct {
	client pbUsers.UsersServiceClient
}

// Returns a list of orders.
// (GET /orders)
func (s *Handler) GetOrders(ctx echo.Context) error {
	var rsp []*models.Order

	uRsp, err := s.client.List(
		ctx.Request().Context(),
		&pbUsers.ListRequest{},
	)
	if err != nil {
		return err
	}

	rsp = append(rsp, &models.Order{
		Id:     uuid.MustParse("9cb14230-b640-11ec-b909-0242ac120002"),
		UserId: uuid.MustParse(uRsp.Users[0].Id),
	})

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

func New(c pbUsers.UsersServiceClient) *Handler {
	return &Handler{
		client: c,
	}
}
