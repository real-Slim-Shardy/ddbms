package main

import (
	"ddbms/lib/alarm"
	hb "ddbms/lib/heartbeat"
)

func main() {
	alarm.Testme()
	hb.Pulse()
}