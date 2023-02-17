package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	evotorrestogo "github.com/softc24/evotor-resto-go"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalln("Token is empty")
	}

	ctx := context.TODO()

	// создаем клиент
	client := evotorrestogo.Client{
		BaseURL: evotorrestogo.DevURL,
		Token:   token,
	}

	// получаем список торговых точек
	stores, err := client.SelectStores(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", stores[0])

	storeId := stores[0].UUID
	// получаем меню торговой точки
	menu, err := client.SelectMenu(ctx, storeId)
	if err != nil {
		log.Fatalln(err)
	}

	product := menu[0]
	log.Printf("%+v\n", product)

	// создаем заказ
	order := evotorrestogo.MakeOrder(strconv.FormatInt(time.Now().UnixMilli(), 32), "Комментарий", evotorrestogo.Contacts{
		Phone: "79990001234",
	}, []evotorrestogo.OrderPosition{
		evotorrestogo.MakeOrderPosition(product.UUID, product.Name, product.Price+100, product.Price+50, 1000),
	})

	order, err = client.CreateOrder(ctx, storeId, order)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", order)

	// проверяем состояние заказа
	order, err = client.GetOrder(ctx, storeId, order.UUID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", order)
}
