package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strconv"

	//"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"

	"log"
	"os"
	"time"
	//"github.com/fatih/structs"
	//"github.com/creasty/defaults"
	"github.com/urfave/cli"
	//"github.com/EZChain-core/price-service/logger"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/workers/repository/mongo"
	"github.com/EZChain-core/price-service/pkg/workers/usecase"
	//"github.com/EZChain-core/price-service/pkg/utils/abi"
	//"github.com/EZChain-core/price-service/pkg/utils/lbank/module"
	//"github.com/EZChain-core/price-service/pkg/utils/lbank/constant"
	gecko "github.com/enixdark/go-gecko/v3"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"
	//"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/common"

)

func ProvideConfig() *config.AppConfig {
	return config.InitConfig()
}

type CBalanceResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type SupplyResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Supply string `json:"supply"`
	} `json:"result"`
	ID int `json:"id"`
}

type ValidatorResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Validators []struct {
			TxID        string `json:"txID"`
			StartTime   string `json:"startTime"`
			EndTime     string `json:"endTime"`
			StakeAmount string `json:"stakeAmount"`
			NodeID      string `json:"nodeID"`
			RewardOwner struct {
				Locktime  string   `json:"locktime"`
				Threshold string   `json:"threshold"`
				Addresses []string `json:"addresses"`
			} `json:"rewardOwner"`
			PotentialReward string `json:"potentialReward"`
			DelegationFee   string `json:"delegationFee"`
			Uptime          string `json:"uptime"`
			Connected       bool   `json:"connected"`
			Delegators      []struct {
				TxID        string `json:"txID"`
				StartTime   string `json:"startTime"`
				EndTime     string `json:"endTime"`
				StakeAmount string `json:"stakeAmount"`
				NodeID      string `json:"nodeID"`
				RewardOwner struct {
					Locktime  string   `json:"locktime"`
					Threshold string   `json:"threshold"`
					Addresses []string `json:"addresses"`
				} `json:"rewardOwner"`
				PotentialReward string `json:"potentialReward"`
			} `json:"delegators"`
		} `json:"validators"`
	} `json:"result"`
	ID int `json:"id"`
}

type CoinmarketCapReponse struct {
	Data   Data   `json:"data,omitempty"`
	Status Status `json:"status,omitempty"`
}

type PointObject struct {
	V []float64 `json:"v,omitempty"`
}

type Data struct {
	Points map[string]PointObject `json:"points,omitempty"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp,omitempty"`
	ErrorCode    string    `json:"error_code,omitempty"`
	ErrorMessage string    `json:"error_message,omitempty"`
	Elapsed      string    `json:"elapsed,omitempty"`
	CreditCount  int       `json:"credit_count,omitempty"`
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
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
			//{
			//	Name:  "vesting",
			//	Usage: "run vesting contract address token in background",
			//	Action: func(c *cli.Context) error {
			//		fmt.Println("Processing  for vesting ezc tokens")
			//
			//		conn, err := ethclient.Dial("https://api.ezchain.com/ext/bc/C/rpc")
			//		if err != nil {
			//			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
			//		}
			//
			//		// Instantiate the contract and display its name
			//		address := common.HexToAddress("0x05E4dfbB6f26E568D846C95C0C716C4338fd1C0A")
			//		SM, err := abi.NewAbi(address, conn)
			//		if err != nil {
			//			log.Fatalf("Failed to instantiate a smart contract contract: %v", err)
			//		}
			//
			//		totalVesting, err := SM.GetVestingSchedulesCount(nil)
			//		if err != nil {
			//			log.Fatalf("Failed to call smart contract : %v", err)
			//		}
			//		fmt.Println(totalVesting)
			//
			//		var VestingList struct {
			//			Index int
			//			VestingResult interface{}
			//		}
			//
			//		for i := 0; i < totalVesting; i++ {
			//			data, err := SM.GetVestingIdAtIndex(nil, i)
			//			if err != nil {
			//				log.Fatalf("Failed to call smart contract : %v", err)
			//			}
			//
			//			.then(result => { resolve({ index: i, vestingID: result }) });
			//		}
			//
			//		for (const id of vestingIDs) {
			//			contract.getVestingIdAtIndex(i).then(result => { resolve({ index: i, vestingID: result }) });
			//		}
			//
			//		return nil
			//	},
			//},
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
							for i := 1; i < 5; i++ {
								task := NewTask(func(data interface{}) error {
									taskID := data.(int)
									cids := []string{}
									pcp := geckoTypes.PriceChangePercentageObject
									priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
									order := geckoTypes.OrderTypeObject.MarketCapDesc
									market, _ := cg.CoinsMarket(vsCurrency, cids, order, perPage, page, sparkline, priceChangePercentage)
									serviceUseCase.Upsert(context, market)
									if page >= 53 {
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
			{
				Name:  "lbank",
				Usage: "run fetch price in background for lbank",
				Action: func(c *cli.Context) error {
					fmt.Println("Processing fetch lbank")
					//client := module.NewLbankClientWithKey(appConfig.LBankApiKey, appConfig.LBankSecretKey)
					for {
						cg := gecko.NewClient(nil)
						// find specific coins
						vsCurrency := "usd"
						ids := []string{"ezchain"}
						perPage := 1
						page := 1
						sparkline := true
						pcp := geckoTypes.PriceChangePercentageObject
						priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
						order := geckoTypes.OrderTypeObject.MarketCapDesc
						market, _ := cg.CoinsMarket(vsCurrency, ids, order, perPage, page, sparkline, priceChangePercentage)
						coinmarketCapReponse := CoinmarketCapReponse{}
						url := "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/detail/chart?id=19455&range=1D"
						method := "GET"

						payload := strings.NewReader(``)

						client := &http.Client {
						}
						req, err := http.NewRequest(method, url, payload)

						if err != nil {
							log.Fatal(err)
						}
						res, err := client.Do(req)
						if err != nil {
							log.Fatal(err)
						}
						defer res.Body.Close()

						body, err := ioutil.ReadAll(res.Body)
						if err != nil {
							log.Fatal(err)
						}
						//fmt.Println(string(body))

						if err := json.Unmarshal(body, &coinmarketCapReponse); err != nil {
							log.Fatal(err)
						}
						keys := reflect.ValueOf(coinmarketCapReponse.Data.Points).MapKeys()
						//sort.SliceStable(keys, func(i, j int) bool {
						//	return keys[i] > keys[j]
						//})
						max := 0
						key := ""
						for _, i := range keys {
							m, _ := strconv.Atoi(i.String())
							if max < m {
								max = m
								key = i.String()
							}
						}
						price := coinmarketCapReponse.Data.Points[key].V[0]

						serviceUseCase.UpsertEZC(context, market, price)
						fmt.Printf("Task %d processed %s\n", market, price)
						time.Sleep(time.Duration(appConfig.LBankIntervalTime) * time.Second)
					}
					return nil
				},
			},
			{
				Name:  "supply",
				Usage: "run update supply for ezchain",
				Action: func(c *cli.Context) error {
					fmt.Println("Processing update supply")
					for {
						// get balance for circult
						payload := map[string]interface{}{
							"jsonrpc": "2.0",
							"method":  "eth_getBalance",
							"params": []string{
								"0x8d38762C9B2Ea11BFE198972ED0e173E89cE149b",
								"latest",
							},
							"id": 2612,
						}

						bytesRepresentation, err := json.Marshal(payload)
						if err != nil {
							log.Fatalln(err)
						}

						resp, err := http.Post(appConfig.RPCEZChainCURI, "application/json", bytes.NewBuffer(bytesRepresentation))

						if err != nil {
							log.Fatalln(err)
						}

						dataCBalanceResponse := &CBalanceResponse{}

						bodyBalance, err := ioutil.ReadAll(resp.Body)

						if err != nil {
							log.Fatalln(err)
						}

						err = json.Unmarshal(bodyBalance, &dataCBalanceResponse)

						if err != nil {
							log.Fatalln(err)
						}

						defer resp.Body.Close()

						// get balance for 0x01
						payloadC := strings.NewReader(`{
								"jsonrpc": "2.0",
								"method": "eth_getBalance",
								"params": [
									"0x0100000000000000000000000000000000000000",
									"latest"
								],
								"id": 2612
							}`)

						respCBalance, err := http.Post(appConfig.RPCEZChainCURI, "application/json", payloadC)
						if err != nil {
							log.Fatalln(err)
						}

						dataCBalance := &CBalanceResponse{}

						bodyCBalance, err := ioutil.ReadAll(respCBalance.Body)

						if err != nil {
							log.Fatalln(err)
						}

						err = json.Unmarshal(bodyCBalance, &dataCBalance)

						if err != nil {
							log.Fatalln(err)
						}

						defer respCBalance.Body.Close()

						// load for p chain supply

						payloadP := strings.NewReader(`{
								"jsonrpc":"2.0",
								"id"     :2612,
								"method" :"platform.getCurrentSupply",
								"params": {}
							}`)

						respSupply, err := http.Post(appConfig.RPCEZChainPURI, "application/json", payloadP)
						if err != nil {
							log.Fatalln(err)
						}

						dataCSupply := &SupplyResponse{}

						body, err := ioutil.ReadAll(respSupply.Body)

						if err != nil {
							log.Fatalln(err)
						}

						err = json.Unmarshal(body, &dataCSupply)

						if err != nil {
							log.Fatalln(err)
						}

						defer respSupply.Body.Close()

						// get payload validator
						payloadValidator := strings.NewReader(`{
								"jsonrpc": "2.0",
								"method": "platform.getCurrentValidators",
								"id": 2612
							}`)

						respValidator, err := http.Post(appConfig.RPCEZChainPURI, "application/json", payloadValidator)
						if err != nil {
							log.Fatalln(err)
						}

						dataCValidator := &ValidatorResponse{}

						bodyValidator, err := ioutil.ReadAll(respValidator.Body)

						if err != nil {
							log.Fatalln(err)
						}

						err = json.Unmarshal(bodyValidator, &dataCValidator)

						if err != nil {
							log.Fatalln(err)
						}

						defer respValidator.Body.Close()

						currentSupply, _ := new(big.Int).SetString(dataCSupply.Result.Supply, 10)
						currentSupply = currentSupply.Mul(currentSupply, big.NewInt(1000000000))
						totalValidatorBalance := big.NewInt(0)
						for _, validator := range dataCValidator.Result.Validators {
							temp, _ := new(big.Int).SetString(validator.PotentialReward, 10)
							temp = temp.Mul(temp, big.NewInt(1000000000))
							if len(validator.Delegators) > 0 {
								for _, delegator := range validator.Delegators {
									tempDelagator, _ := new(big.Int).SetString(delegator.PotentialReward, 10)
									tempDelagator = tempDelagator.Mul(tempDelagator, big.NewInt(1000000000))
									temp = temp.Add(temp, tempDelagator)
								}
							}
							totalValidatorBalance = totalValidatorBalance.Add(totalValidatorBalance, temp)
						}
						circulatingSupply, _ := new(big.Int).SetString(hexaNumberToInteger(dataCBalanceResponse.Result), 16)
						CSupply, _ := new(big.Int).SetString(hexaNumberToInteger(dataCBalance.Result), 16)
						totalSupply := currentSupply.Sub(currentSupply, totalValidatorBalance)
						totalSupply = totalSupply.Add(totalSupply, CSupply)
						if totalSupply.Cmp(big.NewInt(1).Mul(big.NewInt(500000000), big.NewInt(1000000000000000000))) > 0 {
							totalSupply = big.NewInt(1).Mul(big.NewInt(500000000), big.NewInt(1000000000000000000))
						}
						maxSupply := new(big.Int).SetInt64(0)
						circulatingCaculate := new(big.Int).Set(totalSupply)
						circulatingCaculate = circulatingCaculate.Sub(circulatingCaculate, circulatingSupply)
						//totalSupply = totalSupply.Add(totalSupply, CSupply)

						fmt.Println(circulatingSupply)
						fmt.Println(CSupply)
						fmt.Println(totalSupply)
						//fmt.Println("Current Supply: ", totalSupply.Sub(totalSupply, circulatingSupply))

						fmt.Println(totalValidatorBalance)
						fmt.Println("-----------------------------------------------------")
						serviceUseCase.UpdateEZCSupply(context, maxSupply, totalSupply.Div(totalSupply, big.NewInt(1000000000000000000)), circulatingCaculate.Div(circulatingCaculate, big.NewInt(1000000000000000000)))
						time.Sleep(time.Duration(appConfig.LBankIntervalTime) * time.Second)
					}
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
