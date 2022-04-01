package mongo

import (
	"context"
	"strings"

	//"log"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceMongoStorage struct {
	storage *utils.MongoStorage
}

func NewServiceMongoStorage(ctx context.Context, appConfig *config.AppConfig) *ServiceMongoStorage {
	return &ServiceMongoStorage{
		utils.NewMongoStorage(
			ctx, appConfig,
		),
	}
}

func (m *ServiceMongoStorage) ListTokenPrice(optionDatas map[string]interface{}) ([]Token, error) {
	result := []Token{}
	bData := bson.M{}
	filter := []bson.M{}
	allowQuery := false
	limit := optionDatas["limit"].(int64)
	skip := optionDatas["offset"].(int64)
	pagination := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	if optionDatas["token_names"] != nil && optionDatas["token_names"].(string) != "" {
		token_names := strings.Split(optionDatas["token_names"].(string), ",")
		allowQuery = true
		for _, token := range token_names {
			filter = append(filter, bson.M{"name": token})
		}

	}

	if optionDatas["symbols"] != nil && optionDatas["symbols"].(string) != "" {
		symbols := strings.Split(optionDatas["symbols"].(string), ",")
		allowQuery = true
		for _, symbol := range symbols {
			filter = append(filter, bson.M{"symbol": strings.ToLower(symbol)})
		}
	}

	if optionDatas["contracts"] != nil && optionDatas["contracts"].(string) != "" {
		contracts := strings.Split(optionDatas["contracts"].(string), ",")
		allowQuery = true
		for _, contract := range contracts {
			filter = append(filter, bson.M{"contracts.contract_address": contract})
		}
	}

	if optionDatas["chains"] != nil && optionDatas["chains"].(string) != "" {
		chains := strings.Split(optionDatas["chains"].(string), ",")
		allowQuery = true
		for _, chain := range chains {
			filter = append(filter, bson.M{"contracts.chain": chain})
		}
	}
	if allowQuery {
		bData["$or"] = filter
	}

	if optionDatas["is_native_token"].(*bool) != nil {
		native := optionDatas["is_native_token"].(*bool)
		bData["is_native_token"] = &native
	}

	err := mgm.Coll(&Token{}).SimpleFind(&result, bData, pagination)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *ServiceMongoStorage) GetTokenPrice(optionDatas map[string]interface{}) (*Token, error) {
	//ctx := mgm.Ctx()
	token := &Token{}
	bData := bson.M{}
	allowQuery := false
	if optionDatas["is_native_token"].(*bool) != nil {
		native := optionDatas["is_native_token"].(*bool)
		bData["is_native_token"] = &native
		allowQuery = true
	}

	if optionDatas["token_name"] != nil && optionDatas["token_name"].(string) != "" {
		bData["token"] = optionDatas["token_name"].(string)
		allowQuery = true
	}

	if optionDatas["symbol"] != nil && optionDatas["symbol"].(string) != "" {
		bData["symbol"] = strings.ToLower(optionDatas["symbol"].(string))
		allowQuery = true
	}

	if optionDatas["contract"] != nil && optionDatas["contract"].(string) != "" {
		bData["contracts.contract_address"] = optionDatas["contract"].(string)
		allowQuery = true
	}

	if optionDatas["chain"] != nil && optionDatas["chain"].(string) != "" {
		bData["contracts.chain"] = optionDatas["chain"].(string)
		allowQuery = true
	}
	if !allowQuery {
		return token, nil
	}
	err := mgm.Coll(token).First(bData, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (m *ServiceMongoStorage) GetValidator(optionDatas map[string]interface{}) (*Validator, error) {
	//ctx := mgm.Ctx()
	validator := &Validator{}
	bData := bson.M{}
	allowQuery := false

	if optionDatas["name"] != nil && optionDatas["name"].(string) != "" {
		bData["name"] = optionDatas["name"].(string)
		allowQuery = true
	}

	if optionDatas["node_id"] != nil && optionDatas["node_id"].(string) != "" {
		bData["node_id"] = optionDatas["node_id"].(string)
		allowQuery = true
	}

	if !allowQuery {
		return validator, nil
	}

	err := mgm.Coll(validator).First(bData, validator)
	if err != nil {
		return nil, err
	}
	return validator, nil
}

func (m *ServiceMongoStorage) ListValidator(optionDatas map[string]interface{}) ([]Validator, error) {
	result := []Validator{}
	bData := bson.M{}

	limit := optionDatas["limit"].(int64)
	skip := optionDatas["offset"].(int64)
	pagination := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	if optionDatas["name"] != nil && optionDatas["name"].(string) != "" {
		bData["name"] = optionDatas["name"].(string)
	}

	if optionDatas["node_id"] != nil && optionDatas["node_id"].(string) != "" {
		bData["node_id"] = optionDatas["node_id"].(string)
	}

	err := mgm.Coll(&Validator{}).SimpleFind(&result, bData, pagination)
	if err != nil {
		return nil, err
	}
	return result, nil
}
