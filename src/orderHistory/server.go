package orderHistory

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type OrderHistoryServer struct {
	App *fiber.App
}

func NewHistoryServer() *OrderHistoryServer {
	orderHistoryServer := &OrderHistoryServer{
		App: fiber.New(),
	}
	orderHistoryServer.App.Use(logger.New())
	orderHistoryServer.App.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	orderHistoryServer.App.Get("/order-histories", orderHistoryServer.getOrderHistories)

	return orderHistoryServer
}

func (s *OrderHistoryServer) getOrderHistories(ctx *fiber.Ctx) error {
	var histories []*OrderHistory
	orderer := ctx.Query("orderer")
	product, err := strconv.Atoi(ctx.Query("product"))
	if err != nil {
		if len(ctx.Query("product")) == 0 {
			product = 0
		} else {
			err = errors.WithStack(err)
			log.Errorf("%+v", err)
			return err
		}
	}

	for _, fh := range fakeOrderHistories {
		if (orderer == "" && product == 0) ||
			(orderer == "" && product == fh.Product) ||
			(orderer == fh.Orderer && product == 0) ||
			(orderer == fh.Orderer && product == fh.Product) {
			h := *fh
			histories = append(histories, &h)
		}
	}

	return ctx.JSON(histories)
}

var (
	fakeOrderHistories = []*OrderHistory{
		{
			ID:        1,
			Orderer:   "jinsu",
			Product:   1,
			OrderedAt: "2022-10-18 19:30",
		},
		{
			ID:        2,
			Orderer:   "jinsu",
			Product:   1,
			OrderedAt: "2022-10-21 18:03",
		},
		{
			ID:        3,
			Orderer:   "jinsu",
			Product:   1,
			OrderedAt: "2022-10-23 12:25",
		},
		{
			ID:        4,
			Orderer:   "jinsu",
			Product:   1,
			OrderedAt: "2022-10-29 13:16",
		},
		{
			ID:        5,
			Orderer:   "jinsu",
			Product:   2,
			OrderedAt: "2022-11-07 07:03",
		},
		{
			ID:        6,
			Orderer:   "jinsu",
			Product:   3,
			OrderedAt: "2022-11-08 11:33",
		},
		{
			ID:        7,
			Orderer:   "foo",
			Product:   1,
			OrderedAt: "2022-11-08 12:53",
		},
		{
			ID:        8,
			Orderer:   "bar",
			Product:   1,
			OrderedAt: "2022-11-11 12:16",
		},
	}
)
