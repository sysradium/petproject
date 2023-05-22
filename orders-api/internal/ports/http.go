package ports

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/api/models"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
	pbUsers "github.com/sysradium/petproject/users-api/proto/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ api.ServerInterface = (*HttpServer)(nil)

type HttpServer struct {
	client          pbUsers.UsersServiceClient
	orderRepository order.Repository
}

// Returns a list of orders.
// (GET /orders)
func (s *HttpServer) GetOrders(ctx echo.Context) error {
	rsp := []*models.Order{}

	for _, o := range []models.Order{
		{Name: "foo", UserId: uuid.New()},
	} {
		_, err := s.client.Get(ctx.Request().Context(), &pbUsers.GetRequest{
			Id: o.UserId.String(),
		})
		if err != nil {
			if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
				ctx.Logger().Warnj(log.JSON{"message": "skipping order", "id": o.Id, "user_id": o.UserId})

				continue
			}
			return err
		}

		rsp = append(rsp, &models.Order{
			Id:     o.Id,
			UserId: o.UserId,
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

	ctx.NoContent(http.StatusCreated)
	return nil
}

func NewHttpServer(c pbUsers.UsersServiceClient, r order.Repository) *HttpServer {
	return &HttpServer{
		client:          c,
		orderRepository: r,
	}
}
