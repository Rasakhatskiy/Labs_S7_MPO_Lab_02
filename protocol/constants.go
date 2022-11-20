package protocol

const (
	CmdMessage   = "MESSAGE"
	CmdBroadcast = "BROADCAST"
	CmdName      = "NAME"

	DelimSP  = ' '
	DelimSPS = string(DelimSP)
	DelimNL  = '\n'
	DelimNLS = string(DelimNL)

	ServerName   = "@server"
	EveryoneName = "@everyone"

	ServerInboxID   = 0
	EveryoneInboxID = 1
)
