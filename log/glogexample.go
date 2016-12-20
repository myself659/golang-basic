/*
glog-example
------------
background
---
You probably want to read the source code comments at the top of the glog.go file in
the golang/glog repository on github.com. Located here: https://github.com/golang/glog/blob/master/glog.go
setup
---
	$ go get github.com/golang/glog
	$ mkdir log
run
---
	$ go run example.go -stderrthreshold=FATAL -log_dir=./log
or
	$ go run example.go -stderrthreshold=FATAL -log_dir=./log -v=2
or
	$ go run example.go -logtostderr=true
or
	$ go run example.go -logtostderr=true -v=2
*/

package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	flag.Lookup("log_dir").Value.Set("./logdir")
	number_of_lines := 100
	for i := 0; i < number_of_lines; i++ {
		glog.V(0).Infof("LINE: %d", i)
		message := fmt.Sprintf("TEST LINE: %d", i)
		glog.Error(message)
	}
}
