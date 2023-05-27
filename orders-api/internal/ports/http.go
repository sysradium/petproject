package ports

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/api/models"
	"github.com/sysradium/petproject/orders-api/internal/app"
	"github.com/sysradium/petproject/orders-api/internal/app/commands"
	"github.com/sysradium/petproject/orders-api/internal/app/queries"
	pbUsers "github.com/sysradium/petproject/users-api/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ api.ServerInterface = (*HttpServer)(nil)

type HttpServer struct {
	client pbUsers.UsersServiceClient
	app    app.App
}

// Returns a list of orders.
// (GET /orders)
func (s *HttpServer) GetOrders(ctx echo.Context) error {

	orders, err := s.app.Queries.ListBookedOrders.Handle(
		ctx.Request().Context(),
		queries.BookedOrders{},
	)
	if err != nil {
		return err
	}

	rsp := []*models.Order{}
	for _, o := range orders {

		_, err := s.client.Get(ctx.Request().Context(), &pbUsers.GetRequest{
			Id: o.UserID.String(),
		})
		if err != nil {
			if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
				ctx.Logger().Warnj(log.JSON{"message": "skipping order", "id": o.ID, "user_id": o.UserID})
			}
			ctx.Logger().Error(err)
		}

		rsp = append(rsp, &models.Order{
			Id:     &o.ID,
			UserId: o.UserID,
			Name:   o.Name,
		})
	}

	ctx.JSON(http.StatusOK, rsp)
	return nil
}

// Creates a new order
// (POST /orders)
func (s *HttpServer) PostOrders(ctx echo.Context) error {
	u := &models.PostOrdersJSONRequestBody{}
	if err := ctx.Bind(u); err != nil {
		return err
	}

	newOrder, err := s.app.Commands.BookOrder.Handle(
		ctx.Request().Context(),
		commands.BookOrder{
			Name:   u.Name,
			UserID: u.UserId,
		})

	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, &models.Order{
		Id:     &newOrder.ID,
		UserId: newOrder.UserID,
		Name:   newOrder.Name,
	})
	return nil
}

func NewHttpServer(
	c pbUsers.UsersServiceClient,
	app app.App,
) *HttpServer {
	return &HttpServer{
		client: c,
		app:    app,
	}
}
