/*
Copyright © 2026 Jon Andoni Galdos <jonandonigv@gmail.com>
*/
package main

import (
	_ "fmt"
	"log"

	_ "github.com/jonandonigv/FileSync/cmd/filesync/cmd"
	"github.com/rjeczalik/notify"
)

func fsWatch() {
	c := make(chan notify.EventInfo, 1)
	err := notify.Watch("./...", c, notify.All)
	if err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)
	for ei := range c {
		log.Println("Got:", ei.Event(), ei.Path())
	}
}

func main() {
	fsWatch()
	/* cmd.Execute() */
}
