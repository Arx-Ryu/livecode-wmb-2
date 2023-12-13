package main

import (	
	"livecode-wmb-rest-api/delivery"
	// "livecode-wmb-rest-api/model"
	// "livecode-wmb-rest-api/repository"
	// "livecode-wmb-rest-api/utils"
	//"fmt"
	// "log"
	// "strconv"
	// "strings"
	// "gorm.io/gorm"
	//generateid "github.com/jutionck/generate-id"
)

func main() {
	delivery.Server().Run()
	
	// setup(db)
	
	// customerRegis(db) //Livecode Nomor6

	// aktivasiMemberCustomer(db, 2) //Livecode Nomor 7

	// createBillAndDetails(db) //Livecode Nomor 8

	// payBill(db, 1) //Livecode Nomor 9

	// dailySales(db)	//Livecode Nomor 10
	//Catatan untuk dailySales:
	//Daily sales berbasis pada kolom created_at, dapat ubah secara manual tanggal dari kolom created_at untuk melihat sales dengan hari berbeda :)
}

// func customerRegis(db *gorm.DB) {
// 	newCust1 := model.Customer{
// 		Id: 1,
// 		CustomerName: "Kemal Roushdy Jenie",
// 		MobilePhoneNo: "0810242412",
// 	}
// 	newCust2 := model.Customer{
// 		Id: 2,
// 		CustomerName: "Samuel Maynard",
// 		MobilePhoneNo: "0412312334234",
// 	}
// 	newCust3 := model.Customer{
// 		Id: 3,
// 		CustomerName: "Wilbert Irwan",
// 		MobilePhoneNo: "0862312647378",
// 	}
// 	crepo := repository.NewCustomerRepository(db)
// 	err := crepo.Create(&newCust1)
// 	if utils.IsError(err){
// 		log.Println("Failed at Create Cust1")
// 		return
// 	}
// 	log.Println("Cust1 Success Create!")	
// 	err = crepo.Create(&newCust2)
// 	if utils.IsError(err){
// 		log.Println("Failed at Create Cust1")
// 		return
// 	}
// 	log.Println("Cust2 Success Create!")
// 	err = crepo.Create(&newCust3)
// 	if utils.IsError(err){
// 		log.Println("Failed at Create Cust1")
// 		return
// 	}
// 	log.Println("Cust3 Success Create!")
// }

// func aktivasiMemberCustomer(db *gorm.DB, id int) {
// 	crepo := repository.NewCustomerRepository(db)
// 	cust, err := crepo.FindById(id) //Ganti dengan ID Customer yang ingin diganti
// 	if utils.IsError(err) {
// 		log.Println("Failed at FindById Cust")
// 		return
// 	}
// 	disc, err := crepo.FindByIdDiscount(1)
// 	if utils.IsError(err) {
// 		log.Println("Failed at FindById Disc")
// 		return
// 	}
// 	newCust := cust
// 	newCust = model.Customer{
// 		IsMember: true,
// 	}
// 	newMember := model.CustomerDiscounts{
// 		CustomerID: cust.Id,
// 		DiscountID: disc.Id,
// 	}
// 	err  = crepo.Update(&cust, newCust)
// 	if utils.IsError(err) {
// 		log.Println("Failed at Update")
// 		return
// 	}
// 	err  = crepo.CreateCustDisc(&newMember)
// 	if utils.IsError(err) {
// 		log.Println("Failed at Create Assoc")
// 		return
// 	}	
// 	log.Println("Success Aktivasi Member")
// }

// func createBillAndDetails(db *gorm.DB) {
// 	//Test Case Succesful Bill	
// 	bill01 := model.Bill{
// 		Id: 1,
// 		CustomerID: 1,
// 		TableID: 99,
// 		TransTypeID: "TA",
// 		BillDetailID: []model.BillDetail{
// 			{
// 				MenuPriceID: 1,
// 				Quantity: 2,
// 			},
// 			{
// 				MenuPriceID: 2,
// 				Quantity: 4,
// 			},
// 		},
// 	}
// 	bill02 := model.Bill{
// 		Id: 2,
// 		CustomerID: 2,
// 		TableID: 3,
// 		TransTypeID: "EI",
// 		BillDetailID: []model.BillDetail{
// 			{
// 				MenuPriceID: 1,
// 				Quantity: 1,
// 			},
// 			{
// 				MenuPriceID: 2,
// 				Quantity: 1,
// 			},
// 		},
// 	}
// 	bill03 := model.Bill{
// 		Id: 3,
// 		CustomerID: 2,
// 		TableID: 2,
// 		TransTypeID: "EI",
// 		BillDetailID: []model.BillDetail{
// 			{
// 				MenuPriceID: 1,
// 				Quantity: 2,
// 			},
// 		},
// 	}
// 	bill04 := model.Bill{ //Test Case Table Not Available
// 		Id: 4,
// 		CustomerID: 2,
// 		TableID: 2,
// 		TransTypeID: "EI",
// 		BillDetailID: []model.BillDetail{
// 			{
// 				MenuPriceID: 1,
// 				Quantity: 2,
// 			},
// 		},
// 	}
// 	bill05 := model.Bill{ //Test Case Table Does Not Exist
// 		Id: 5,
// 		CustomerID: 2,
// 		TableID: 7,
// 		TransTypeID: "EI",
// 		BillDetailID: []model.BillDetail{
// 			{
// 				MenuPriceID: 1,
// 				Quantity: 2,
// 			},
// 		},
// 	}
// 	brepo := repository.NewBillRepository(db)
// 	if checkTable(db, int(bill01.TableID)) {//Test Case Succesful
// 		err := brepo.Create(&bill01)
// 		if utils.IsError(err) {
// 			log.Println("Failed at Create Bill01")
// 			return
// 		}
// 		tableAvailability(db, int(bill01.TableID))	
// 		log.Println("Bill01 Success Created!")	
// 	} 
// 	if checkTable(db, int(bill02.TableID)) {
// 		err := brepo.Create(&bill02)
// 		if utils.IsError(err) {
// 			log.Println("Failed at Create Bill02")
// 			return
// 		}
// 		tableAvailability(db, int(bill02.TableID))	
// 		log.Println("Bill02 Success Created!")	
// 	}
// 	if checkTable(db, int(bill03.TableID)) {
// 		err := brepo.Create(&bill03)
// 		if utils.IsError(err) {
// 			log.Println("Failed at Create Bill03")
// 			return
// 		}
// 		tableAvailability(db, int(bill03.TableID))	
// 		log.Println("Bill03 Success Created!")	
// 	}
// 	if checkTable(db, int(bill04.TableID)) { //Test Case Table Not Available
// 		err := brepo.Create(&bill04)
// 		if utils.IsError(err) {
// 			log.Println("Failed at Create Bill04")
// 			return
// 		}
// 		tableAvailability(db, int(bill04.TableID))	
// 		log.Println("Bill04 Success Created!")	
// 	}
// 	if checkTable(db, int(bill05.TableID)) { //Test Case Table Does Not Exist
// 		err := brepo.Create(&bill05)
// 		if utils.IsError(err) {
// 			log.Println("Failed at Create Bill05")
// 			return
// 		}
// 		tableAvailability(db, int(bill05.TableID))	
// 		log.Println("Bill05 Success Inserted!")	
// 	}	
// }

// func checkTable(db *gorm.DB, id int) bool {
// 	brepo := repository.NewBillRepository(db)
// 	table, err := brepo.FindByIdTable(id)
// 	if utils.IsError(err) {
// 		log.Println("Failed to create bill, table does not exist!")
// 		return false
// 	}
// 	if !table.IsAvailable {
// 		log.Println("Failed to create bill, table is not available!")
// 		return false
// 	}
// 	return true
// }

// func payBill(db *gorm.DB, id int) {
// 	brepo := repository.NewBillRepository(db)
// 	bill, err := brepo.FindById(id)
// 	if utils.IsError(err) {
// 		log.Println("Failed at FindById")
// 		return
// 	}
// 	billId := strconv.Itoa(bill.Id)
// 	details, err := brepo.FindSumBill(billId)
// 	if utils.IsError(err) {
// 		log.Println("Failed at FindAllBillDetailByBillId")
// 		return
// 	}
// 	total := 0
// 	for _, items := range details {
// 		total = total + (items.Quantity * items.Price)
// 	}
// 	fmt.Println("Total Bill:", total)
// 	tableAvailability(db, int(bill.TableID))	
// }

// func dailySales(db *gorm.DB) {
// 	brepo := repository.NewBillRepository(db)
// 	sales, err := brepo.DailySales()
// 	if utils.IsError(err) {
// 		log.Println("Failed at DailySales")
// 		return
// 	}
// 	for _, items := range sales {
// 		dateRev := strings.Split(items.Date, "T")
// 		sale := fmt.Sprintf("Sales Tanggal %s: %s", dateRev[0], items.Total)
// 		fmt.Println(sale)
// 	}
// }

// func tableAvailability(db *gorm.DB, id int) {
// 	brepo := repository.NewBillRepository(db)
// 	table, _ := brepo.FindByIdTable(id)
// 	brepo.UpdateTableAvailability(&table, !table.IsAvailable)
// }

// func setup(db *gorm.DB) {
// 	//Setup Table & Trans Type
// 	transType01 := model.TransType{
// 		Id: "TA",
// 		Description: "Take Away",
// 	}
// 	transType02 := model.TransType{
// 		Id: "EI",
// 		Description: "Eat In",
// 	}
// 	tableTA := model.Table{
// 		Id: 99,
// 		TableDescription: "Take Away",
// 	}
// 	table01 := model.Table{
// 		TableDescription: "Table 01",
// 	}
// 	table02 := model.Table{
// 		TableDescription: "Table 02",
// 		IsAvailable: false,
// 	}
// 	table03 := model.Table{
// 		TableDescription: "Table 03",
// 	}	
// 	brepo := repository.NewBillRepository(db)
// 	brepo.CreateTransType(&transType01)
// 	brepo.CreateTransType(&transType02)
// 	brepo.CreateTable(&tableTA)
// 	brepo.CreateTable(&table01)
// 	brepo.CreateTable(&table02)
// 	brepo.CreateTable(&table03)

// 	//Setup Menu & Price
// 	menu01 := model.Menu{
// 		MenuName: "Nasi Goreng Kampung",
// 	}
// 	menu02 := model.Menu{
// 		MenuName: "Nasi Goreng Teri",
// 	}
// 	menu03 := model.Menu{
// 		MenuName: "Es Teh",
// 	}
// 	mp01 := model.MenuPrice{
// 		Price: 15000,
// 		MenuID: []model.Menu{
// 			menu01,
// 			menu02,
// 		},
// 	}
// 	mp02 := model.MenuPrice{
// 		Price: 3000,
// 		MenuID: []model.Menu{
// 			menu03,
// 		},
// 	}
// 	mp03 := model.MenuPrice{
// 		Price: 10000,
// 	}	
// 	mrepo := repository.NewMenuRepository(db)
// 	mrepo.CreateMPrice(&mp01)
// 	mrepo.CreateMPrice(&mp02)
// 	mrepo.CreateMPrice(&mp03)

// 	//Setup Discount
// 	disc01 := model.Discount{
// 		Description: "Member Discount",
// 		Pct: 10,
// 	}
// 	crepo := repository.NewCustomerRepository(db)
// 	crepo.CreateDiscount(&disc01)
// }