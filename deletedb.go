package main

import (
	"fmt"
	"log"
)

func delete_userDB(userid string) (bool, bool) {

	err_responce := true
	result := true

	sqlStatement := `DELETE FROM users WHERE id = $1`
	user, err2 := DB.Exec(sqlStatement, userid)

	if err2 != nil {
		log.Fatal("ERror in DELETE: ", err2)
	}
	fmt.Println(user)
	return result, err_responce
}
