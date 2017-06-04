package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/rs/cors"

	"github.com/divergentdave/cidr-test-service/config"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		return
	}

	configuration := config.Config()
	cidrsText := configuration.GetStringSlice("cidrs")
	cidrs := make([]*net.IPNet, len(cidrsText))
	for i, s := range cidrsText {
		_, cidr, err := net.ParseCIDR(s)
		if err != nil {
			log.Fatal(err)
		}
		cidrs[i] = cidr
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		remoteAddr := req.RemoteAddr
		ipString := remoteAddr[:strings.LastIndex(remoteAddr, ":")]
		ipString = strings.Trim(ipString, "[]")
		ip := net.ParseIP(ipString)
		if ip == nil {
			log.Fatalf("Couldn't parse IP address from %s", remoteAddr)
		}
		for _, net := range cidrs {
			if net.Contains(ip) {
				io.WriteString(w, "true")
				return
			}
		}
		io.WriteString(w, "false")
	})

	listenAddress := configuration.GetString("listen_addr")
	corsMiddleware := cors.New(cors.Options{MaxAge: 86400})
	corsHandler := corsMiddleware.Handler(mux)
	log.Fatal(http.ListenAndServe(listenAddress, corsHandler))
}
