package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ProductServer struct {
	App             *fiber.App
	orderHistoryURL string
	version         string
}

func NewProductServer(orderHistoryURL string, version string) *ProductServer {
	productServer := &ProductServer{
		App:             fiber.New(),
		orderHistoryURL: orderHistoryURL,
		version:         version,
	}
	productServer.App.Use(logger.New())
	productServer.App.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	productServer.App.Get("/products", productServer.getProducts)

	return productServer
}

func (s *ProductServer) getProducts(ctx *fiber.Ctx) error {
	var products []*Product
	for _, fp := range fakeProducts {
		product := *fp

		req, err := http.Get(fmt.Sprintf("%s/order-histories?orderer=%s&product=%d", s.orderHistoryURL, fakeUser, fp.ID))
		if err != nil {
			err = errors.WithStack(err)
			log.Errorf("%+v", err)
			return err
		}

		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			err = errors.WithStack(err)
			log.Errorf("%+v", err)
			return err
		}
		log.Infof("GetOrderHistoriesResponse: %s", string(data))

		if err := json.Unmarshal(data, &product.MyOrderHistories); err != nil {
			log.Errorf("%+v", err)
			return err
		}
		products = append(products, &product)
	}

	return ctx.JSON(products)
}

var (
	fakeUser     = "jinsu"
	fakeProducts = []*Product{
		{
			ID:    1,
			Name:  "라즈베리파이 4 8GB + 개발 키트",
			Price: 250000,
		},
		{
			ID:    2,
			Name:  "USB 32GB",
			Price: 12000,
		}, {
			ID:    3,
			Name:  "맥북 프로 M1X",
			Price: 3200000,
		},
	}
)
