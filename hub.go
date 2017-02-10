package main

// user is taken by os/user
type appUser struct {
	c        *conn
	channels []string
}

type appUserList map[string]*appUser

type serverContext struct {
	port   int
	users  appUserList
	logdir string
}
