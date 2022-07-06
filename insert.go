package main

import (
	"fmt"
	"log"
)

func insert(reqBody User) (bool, bool) {

	err_responce := true
	result := true

	sqlStatement := `
    INSERT INTO users(ID, Name, UserID, Mobile, Mail, City, Password)
    VALUES ($1, $2, $3, $4, $5, $6, $7)`
	user, err2 := DB.Exec(sqlStatement, reqBody.ID, reqBody.Name, reqBody.UserID, reqBody.Mobile, reqBody.Mail, reqBody.City, reqBody.Password)

	if err2 != nil {
		log.Fatal("ERror in insert: ", err2)
	}
	fmt.Println(user)
	return result, err_responce
}
