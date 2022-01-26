package foodHelper

import (
	"foodHelper/session"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Then, initialize the session manager
	globalSessionManager, err := session.NewManager("memory", "gosessionid", 3600)
	if err != nil {
		panic(err)
	}

	// Start deleting expired session periodically
	go globalSessionManager.GC()

	//connect to postgres
	//start https server
}
