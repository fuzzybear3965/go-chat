package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	// Retrieves the user's session if it exists. If it doesn't exist then we'll
	// make one.
	session, err := cookieStore.Get(r, "channelname")
	if err != nil {
		fmt.Println("Could not get user's session.")
		http.Error(w, err.Error(), 500)
		return nil
	}
	_, userIDExists := session.Values["userID"]
	fmt.Println(session.IsNew, userIDExists, session.Values["userID"])
	if session.IsNew || !userIDExists || session.Values["userID"] == "" {
		fmt.Println("Someone is trying to fly under the radar!")
		session.Values["userID"] = randomString(10)
		session.Save(r, w)
	}
	return session
}

// Below is an optimized random string generator devised by user icza found at:
// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
