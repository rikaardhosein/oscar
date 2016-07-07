package main

import ircevent "github.com/thoj/go-ircevent"

type IRCEngine struct {
	*ircevent.Connection
}

type Channel struct {
	IrcEngine   *IRCEngine
	ChannelName string
}

func (i *IRCEngine) Init(nick string, user string) {
	i.Connection = ircevent.IRC(nick, user)
}

func (i *IRCEngine) Connect(host string) error {
	hasReceivedMotd := make(chan bool, 1)
	i.AddCallback("376", func(event *ircevent.Event) {
		hasReceivedMotd <- true
	})
	err := i.Connection.Connect(host)
	if err != nil {
		return err
	}
	<-hasReceivedMotd
	return nil
}

//func (i *IRCEngine) Join(channel string) Channel {
//	c := Channel{IrcEngine: i, ChannelName: channel}
//	i.Join(channel)
//	return c
//}
