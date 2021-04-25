package main

import (
	"log"

	"github.com/envsh/go-purple/purple"
)

func init() {
	pi := purple.PluginInfo{
		Type: purple.PLUGIN_PROTOCOL,

		Id:          "purple-nostr",
		Name:        "Nostr",
		Version:     "1.0",
		Summary:     "Nostr Protocol Plugin",
		Description: "Nostr Protocol Plugin https://github.com/fiatjaf/nostr",
		Author:      "fiatajf",
		Homepage:    "https://github.com/fiatjaf/purple-nostr",

		Load: func(p *purple.Plugin) bool {
			log.Print("load called")
			return true
		},
		Unload: func(p *purple.Plugin) bool {
			log.Print("unload called")
			return true
		},
		Destroy: func(p *purple.Plugin) {
			log.Print("destroy called")
		},
	}
	ppi := purple.PluginProtocolInfo{
		BlistIcon: func() string { return "icon_name" },
		StatusTypes: func(acct *purple.Account) []*purple.StatusType {
			return []*purple.StatusType{
				purple.NewStatusType(purple.STATUS_AVAILABLE, "n_on", "Online", true),
				purple.NewStatusType(purple.STATUS_OFFLINE, "n_off", "Offline", true),
			}
		},
		Login: func(acct *purple.Account) {
			log.Print("login called ", *acct)
		},
		Close: func(conn *purple.Connection) {
			log.Print("close called ", *conn)
		},
		ChatInfo: func(conn *purple.Connection) []*purple.ProtoChatEntry {
			log.Print("proto chat entry called")
			return []*purple.ProtoChatEntry{
				purple.NewProtoChatEntry("nnn", "yyy", true),
			}
		},
		ChatInfoDefaults: func(conn *purple.Connection, chat_name string) map[string]string {
			return nil
		},
		SendIM: func(conn *purple.Connection, who string, message string) int {
			log.Print("sendim called ", who, " ", message)
			return 0
		},
	}

	purple.NewPlugin(&pi, &ppi, func(p *purple.Plugin) {
		log.Print("init")
	})
}

func main() {}
