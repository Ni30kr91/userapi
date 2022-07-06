package main

import (
	"fmt"
)

func delete_user(userID string) string {
	message := "User not found"
	user := User{}
	if user, ok := Data[userID]; ok {
		delete(Data, userID)
		message = "user deleted successfully"
		fmt.Println(user)
		return message
	}
	fmt.Println(user)
	return message
}
