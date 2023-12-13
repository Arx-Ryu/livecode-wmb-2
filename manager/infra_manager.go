package manager

import (
	"fmt"
	"livecode-wmb-rest-api/config"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/service"
	"livecode-wmb-rest-api/utils"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
	LopeiClientConn() service.LopeiPaymentClient
	Token() utils.Token
}

type infra struct {
	db *gorm.DB
	lopeiClient service.LopeiPaymentClient
	cfg         config.Config
	token		utils.Token
}

func (i *infra) Token() utils.Token {
	return i.token
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func (i *infra) LopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	infra := infra{
		cfg: config,
		db: resource,
	}
	infra.initGrpcClient()
	infra.newJwtToken()
	return &infra
}

func (i *infra) newJwtToken() {
	tokenService := utils.NewTokenService(i.cfg.TokenConfig)	
	i.token = tokenService
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	enigmaDb, _ := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to DB")
	}
	env := os.Getenv("ENV")
	if env == "dev" {
		db.Debug()
	} else if env == "migration" {
		db.Debug()
		err := db.AutoMigrate(&model.CustomerDiscounts{}, &model.Customer{}, &model.Discount{}, &model.MenuPrice{}, &model.Menu{}, &model.Table{}, &model.TransType{}, &model.Bill{}, &model.BillDetail{})
		if err != nil {
			log.Println(err.Error())
			panic("Error at AutoMigrate")
		}
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (i *infra) initGrpcClient() {
	dial, err := grpc.Dial(i.cfg.GrpcConfig.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not conenct...", err)
	}

	client := service.NewLopeiPaymentClient(dial)
	i.lopeiClient = client
	fmt.Println("GRPC client connected...")
}