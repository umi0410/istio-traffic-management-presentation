package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"server/orderHistory"
	"server/product"
)

func main() {
	if len(os.Args) != 2 {
		log.Errorf("실행 인자의 길이가 2가 아닙니다: %+v\n", os.Args)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "order-history":
		srv := orderHistory.NewHistoryServer()
		if err := srv.App.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
			err = errors.WithStack(err)
			log.Errorf("%+v", err)
			os.Exit(1)
		}
	case "product":
		srv := product.NewProductServer(os.Getenv("ORDER_HISTORY_URL"), os.Getenv("VERSION"))
		if err := srv.App.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
			err = errors.WithStack(err)
			log.Errorf("%+v", err)
			os.Exit(1)
		}
	default:
		log.Errorf("Unknown server type: %+v", os.Args[1])
		os.Exit(1)
	}
}
