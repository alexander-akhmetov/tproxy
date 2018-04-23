package main

import (
	"log"
	"os"
	"strings"

	"github.com/armon/go-socks5"
)

type config struct {
	listenTo  string
	usernames []string
	passwords []string
}

func getConfig(conf config) *socks5.Config {
	credentials := readCredentials(conf)
	authenticator := socks5.UserPassAuthenticator{
		Credentials: credentials,
	}

	return &socks5.Config{
		AuthMethods: []socks5.Authenticator{authenticator},
		Rules:       filterTelegram(),
	}
}

func readCredentials(conf config) socks5.StaticCredentials {
	credentials := socks5.StaticCredentials{}
	for i := range conf.usernames {
		credentials[conf.usernames[i]] = conf.passwords[i]
	}
	return credentials
}

func createServer(conf config) *socks5.Server {
	server, err := socks5.New(getConfig(conf))
	if err != nil {
		panic(err)
	}
	return server
}

func runServer(conf config) {
	server := createServer(conf)
	err := server.ListenAndServe("tcp", conf.listenTo)
	if err != nil {
		panic(err)
	}
}

func main() {
	listenTo := os.Getenv("SOCKS5_LISTEN")
	if listenTo == "" {
		listenTo = "127.0.0.1:8000"
	}

	usernames := strings.Split(os.Getenv("SOCKS5_USERNAME"), ",")
	passwords := strings.Split(os.Getenv("SOCKS5_PASSWORD"), ",")

	if len(usernames) == 0 || len(passwords) == 0 {
		log.Panic("you have to define SOCKS5_USERNAME and SOCKS5_PASSWORD!")
	}
	if len(usernames) != len(passwords) {
		log.Panic("The SOCKS5_USERNAME and SOCKS5_PASSWORD have different length! (splitted by ',')")
	}

	conf := config{
		listenTo:  listenTo,
		usernames: usernames,
		passwords: passwords,
	}
	log.Println(
		"Starting SOCKS5 server on",
		conf.listenTo,
		"usernames: ",
		conf.usernames,
	)
	runServer(conf)
}
