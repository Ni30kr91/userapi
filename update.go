package main

import (
	"fmt"
	"log"
)

func update_userDB(reqBody User, userid string) (bool, bool) {

	err_responce := true
	result := true
	sqlStatement := `
	UPDATE users SET name = $2, mobile=$3, mail=$4,city=$5 WHERE id = $1`
	user, err2 := DB.Exec(sqlStatement, reqBody.ID, reqBody.Name, reqBody.Mobile, reqBody.Mail, reqBody.City)
	fmt.Println(err2, user)
	if err2 != nil {
		log.Fatal("ERror in insert: ", err2)

	}

	return result, err_responce
}
