package main

import (
	"context"
	"fmt"
	"math/rand"

	//"fmt"
	//"fmt"
	"log"
	"time"

	//"net/http"
	"os"
	//"github.com/fatih/structs"
	//"github.com/creasty/defaults"
	"github.com/urfave/cli"
	//"github.com/EZChain-core/price-service/logger"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/workers/usecase"
	"github.com/EZChain-core/price-service/pkg/workers/repository/mongo"
	gecko "github.com/enixdark/go-gecko/v3"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"

)

func ProvideConfig() *config.AppConfig {
	return config.InitConfig()
}

func main() {

	context := context.Background()

	// Create usecase via dependency injection
	appConfig := ProvideConfig()
	serviceMongoStorage := mongo.NewServiceMongoStorage(context, appConfig)
	serviceUseCase := usecase.NewServiceUseCase(serviceMongoStorage)
	// Prepare the data

	var allTask []*Task

	pool := NewPool(allTask, 5)
	cg := gecko.NewClient(nil)

	app := &cli.App{
		Name:  "go worker pool fetch price",
		Usage: "check different work loads with worker pool",
		Action: func(c *cli.Context) error {
			fmt.Println("You need more parameters")
			return nil
		},
		Commands: []cli.Command{
			{
				Name:  "wpool",
				Usage: "run robust worker pool in background",
				Action: func(c *cli.Context) error {
					go func() {
						taskID := rand.Intn(1)
						page := 1
						perPage := 250
						vsCurrency := "usd"
						sparkline := false
						for {
							time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
							for i:= 1; i < 5; i++ {
								task := NewTask(func(data interface{}) error {
									taskID := data.(int)
									cids := []string{}
									pcp := geckoTypes.PriceChangePercentageObject
									priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
									order := geckoTypes.OrderTypeObject.MarketCapDesc
									market, _ := cg.CoinsMarket(vsCurrency, cids, order, perPage, page, sparkline, priceChangePercentage)
									serviceUseCase.Upsert(context, market)
									if (page >= 53) {
										page = 1
									} else {
										page += 1
									}
									fmt.Printf("Task %d processed\n", taskID)
									return nil
								}, taskID)
								pool.AddTask(task)
							}
							time.Sleep(10000 * time.Millisecond)
						}
					}()
					pool.RunBackground()
					return nil
				},
			},
			{
				Name:  "import",
				Usage: "run import contract address token in background",
				Action: func(c *cli.Context) error {
					fmt.Println("Processing import for tokens")
					cg := gecko.NewClient(nil)
					list, err := cg.CoinsList(true)
					if err != nil {
						log.Fatal(err)
					}
					serviceUseCase.Import(context, list)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
