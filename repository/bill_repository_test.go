package repository

import (
	"livecode-wmb-rest-api/model"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repoMock struct {
	mock.Mock
}

type CustomerRepositoryTestSuite struct {
	suite.Suite
	mockDb  *gorm.DB
	mockSql sqlmock.Sqlmock
}

func (suite *CustomerRepositoryTestSuite) SetupTest() {
	db, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("An error when opening a stub database connection", err)
	}
	mockDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *CustomerRepositoryTestSuite) TestTableCreate_Failed() {
	dummyData := model.Table{
		Id: 99,
		TableDescription: "Take Away",
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO 'm_table' (id, table_description) VALUES (?,?) RETURNING 'id','id'").
	WithArgs(dummyData.Id, dummyData.TableDescription).	
	WillReturnResult(sqlmock.NewResult(1,1))
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.CreateTable(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.NotNil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestTransTypeCreate_Failed() {
	dummyData := model.TransType{
		Id: "TA",
		Description: "Take Away",
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO 'trans_types' (id, description) VALUES (?,?) RETURNING 'id','id'").
		WithArgs(dummyData.Id, dummyData.Description).
		WillReturnResult(sqlmock.NewResult(1,1))
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.CreateTransType(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.NotNil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestBillDetailCreate_Failed() {
	dummyData := model.BillDetail{
		Id: 1,
		BillID: 1,
		MenuPriceID: 1,
		Quantity: 1,
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO 'bill_details' (id, bill_id, menu_price_id, quantity) VALUES (?,?,?,?) RETURNING 'id','id'").
		WithArgs(dummyData.Id, dummyData.BillID, dummyData.MenuPriceID, dummyData.Quantity).	
		WillReturnResult(sqlmock.NewResult(1,1))
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.CreateBillDetail(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.NotNil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestBillCreate_Failed() {
	dummyData := model.Bill{
		Id: 1,
		CustomerID: 1,
		TableID: 1,
		TransTypeID: "TA",
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO 't_bill' (id, customer_id, table_id, trans_type_id) VALUES ($1,$2,$3,$4) RETURNING 'id','id'").
		WithArgs(dummyData.Id, dummyData.CustomerID, dummyData.TableID, dummyData.TransTypeID).	
		WillReturnResult(sqlmock.NewResult(1,1))
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.Create(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.NotNil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestBillReadAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1)
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("SELECT * FROM t_bill").
		WillReturnRows(rows)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	actual, err := repo.FindAll()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),1,len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestBillDetailsReadAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1)
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("SELECT * FROM bill_details").
		WillReturnRows(rows)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	actual, err := repo.FindAllBillDetail()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),1,len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestTableReadAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1)
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("SELECT * FROM table").
		WillReturnRows(rows)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	actual, err := repo.FindAllTable()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),1,len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestTransTypeReadAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow("1")
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("SELECT * FROM trans_type").
		WillReturnRows(rows)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	actual, err := repo.FindAllTransType()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),1,len(actual))
}

func (suite *CustomerRepositoryTestSuite) TestBillDelete_Failed() {
	dummyData := model.Bill{
		Id: 1,
		CustomerID: 1,
		TableID: 1,
		TransTypeID: "TA",
	}
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("DELETE FROM t_bill WHERE id=?").WithArgs(dummyData.Id)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.Delete(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestBillDetailsDelete_Failed() {
	dummyData := model.BillDetail{
		Id: 1,
		BillID: 1,
		MenuPriceID: 1,
		Quantity: 1,
	}
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("DELETE FROM bill_details WHERE id=?").WithArgs(dummyData.Id)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.DeleteBillDetail(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestTableDelete_Failed() {
	dummyData := model.Table{
		Id: 1,
		TableDescription: "Dummy",
	}
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("DELETE FROM table WHERE id=?").WithArgs(dummyData.Id)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.DeleteTable(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestTransTypeDelete_Failed() {
	dummyData := model.TransType{
		Id: "TA",
		Description: "Dummy",
	}
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery("DELETE FROM trans_type WHERE id=?").WithArgs(dummyData.Id)
	suite.mockSql.ExpectRollback()
	repo := NewBillRepository(suite.mockDb)
	err := repo.DeleteTransType(&dummyData)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(suite.T(), err)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}