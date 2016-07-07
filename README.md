# Oscar
IRC bot



## Architecture

### IRC Engine

#### Responsibilities
- Reply to PING messages with PONGs



The irc engine should be responsible for maintaining connections to all irc servers and receiving messages from all servers and routing them to the proper listeners.


Should be able to set a username, etc for the bot.


So there needs to be:
	A method for connecting to an IRC server with support for SSL[I guess this should only be relevant during the actual connection to the server, the rest
should be done transparently]

There also needs to be:
	A way to join channels, etc. However, we need an easy way to specify the order of operations. Maybe something along the lines of: WaitUntilResponseOrTimeout
	
		