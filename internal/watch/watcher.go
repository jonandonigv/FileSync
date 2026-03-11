package watcher

import (
	"github.com/rjeczalik/notify"
	"log"
)

func FsWatch(dir string) error {
	c := make(chan notify.EventInfo, 1)
	err := notify.Watch(dir, c, notify.All)
	if err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)
	for ei := range c {
		log.Println("Got:", ei.Event(), ei.Path())
	}
	return nil
}
