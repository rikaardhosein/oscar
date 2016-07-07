package main

import ircevent "github.com/thoj/go-ircevent"

func OneWayRelay(c1 *Channel, c2 *Channel) {
	c1.IrcEngine.AddCallback("PRIVMSG", func(e *ircevent.Event) {
		if e.Arguments[0] == c1.ChannelName {
			c2.IrcEngine.Privmsgf(c2.ChannelName, "<%s>: %s", e.Nick, e.Message())
		}
	})
}

func TwoWayRelay(c1 *Channel, c2 *Channel) {
	OneWayRelay(c1, c2)
	OneWayRelay(c2, c1)
}
