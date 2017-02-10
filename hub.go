package main

// user is taken by os/user
type appUser struct {
	username string
}

//func addUser(userList User) chan User {
//return userList
//}

type appUserList []appUser

type serverContext struct {
	port   int
	users  appUserList
	logdir string
}
