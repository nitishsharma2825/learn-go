package main

import (
	jsonparse "encoding/json"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// Args hold arguments passed to JSON RPC service
type Args struct {
	Name string
}

// TwitterProfile struct holds JSON struct
type TwitterProfile struct {
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Followers string `json:"followers,omitempty"`
	Following string `json:"following,omitempty"`
}

type JSONServer struct{}

// TwitterProfileDetail
func (t *JSONServer) TwitterProfileDetail(r *http.Request, args *Args, reply *TwitterProfile) error {
	var twitterprofiles []TwitterProfile
	// Read JSON file and load data
	raw, readerr := os.ReadFile("../twitterProfile.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	// Unmarshal JSON raw data into twitterprofiles array
	marshalerr := jsonparse.Unmarshal(raw, &twitterprofiles)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each twitterprofile to find the given twitterprofile
	for _, twitterprofile := range twitterprofiles {
		if twitterprofile.Name == args.Name {
			// if twitterprofile found, fill reply with it
			*reply = twitterprofile
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC server
	jsonServer := new(JSONServer)
	rpc.Register(jsonServer)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)
}
