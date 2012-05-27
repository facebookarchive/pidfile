// Package pidfile writes pid files.
package pidfile

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var pidfile = flag.String("pidfile", "", "If specified, write pid to file.")

// Write the pidfile based on the flag, if one is set.
func Write() {
	if *pidfile == "" {
		return
	}
	err := ioutil.WriteFile(
		*pidfile, []byte(strconv.Itoa(os.Getpid())), os.FileMode(0644))
	if err != nil {
		log.Fatalf("Failed to write pidfile %s: %s", *pidfile, err)
	}
}
