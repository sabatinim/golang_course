package main

import (
	"fmt"
	//"time"
	"strconv"
	"encoding/json"
)

func generate_data() {

	db := GetDb()
	// Associations will be inserted automatically when save the record
	for i := 0; i < 10; i++ {
		billingAddr := Address{Address1: "B.A. - Addrd 1-" + strconv.Itoa(i),
			Address2: "B.A. - Addrd 2-" + strconv.Itoa(i)}
		shoppingAddr := Address{Address1: "S.A. - Addrd 1-" + strconv.Itoa(i),
			Address2: "S.A. - Addrd 2-" + strconv.Itoa(i)}
		user2 := User{
			Name:            "Value inserted " + strconv.Itoa(i),
			BillingAddress:  billingAddr,
			ShippingAddress: shoppingAddr,
			Emails:          []Email{{Email: "jinzhu@example.com"}, {Email: "jinzhu-2@example@example.com"}},
			Languages:       []Language{{Name: "IT"}, {Name: "EN"}},
		}
		db.NewRecord(user2)
		db.Create(&user2)
	}
}

func some_query() {
	db := GetDb()
	var users []User
	
	db.Where("name LIKE ?", "%alue%").Find(&users)
	 for _, user := range users {
        //fmt.Printf("%+v\n", user)
    		var languages []Language
		//map model sub structure
		db.Model(&user).Association("Languages").Find(&languages)
        fmt.Printf("%+v\n", languages)
    		user.Languages = languages
      
        js,_:= json.Marshal(user)
        fmt.Printf("%s\n", js )
    }
    fmt.Println()
}

func main() {
	fmt.Print("DB ", myDb)
	some_query()

}
