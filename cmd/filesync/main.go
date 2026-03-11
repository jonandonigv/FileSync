/*
Copyright © 2026 Jon Andoni Galdos <jonandonigv@gmail.com>
*/
package main

import (
	_ "fmt"
	"log"

	_ "github.com/jonandonigv/FileSync/cmd/filesync/cmd"
	"github.com/jonandonigv/FileSync/internal/watcher"
)

func main() {
	err := watcher.FsWatch("./...")
	if err != nil {
		log.Fatal(err)
	}
	/* cmd.Execute() */
}
