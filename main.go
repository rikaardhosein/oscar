package main

import "crypto/tls"

func main() {
	done := make(chan bool, 1)

	zempirians := &IRCEngine{}
	zempirians.Init("Oscar", "Oscar")
	zempirians.UseTLS = true
	zempirians.Debug = true
	zempirians.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := zempirians.Connect("irc.zempirians.com:6697")
	if err != nil {
		panic(err)
	}

	rcode := &IRCEngine{}
	rcode.Init("Oscar", "Oscar")
	rcode.Debug = true
	err = rcode.Connect("irc.rcode.ca:6667")
	if err != nil {
		panic(err)
	}

	enigmagroup := &IRCEngine{}
	enigmagroup.Init("Oscar", "Oscar")
	enigmagroup.Debug = true
	err = enigmagroup.Connect("5.39.57.215:6667")
	if err != nil {
		panic(err)
	}

	enigmagroup.Join("#enigmagroup")
	zempirians.Join("#enigmagroup")
	zempirians.Join("#goatzzz")
	rcode.Join("#rcode")

	enigmagroupChannel := Channel{IrcEngine: enigmagroup, ChannelName: "#enigmagroup"}
	zempiriansChannel := Channel{IrcEngine: zempirians, ChannelName: "#enigmagroup"}
	TwoWayRelay(&enigmagroupChannel, &zempiriansChannel)

	<-done
}
