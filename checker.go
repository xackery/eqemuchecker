package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	//	"os/exec"
)

var (
	perlFileCount = 0
	luaFileCount  = 0
)

func main() {
	log.Println("checking perl and lua scripts...")
	start := time.Now()
	output, err := run()
	fmt.Println(output)
	log.Println("finished in", time.Since(start).Seconds(), "seconds procesed", perlFileCount, "perl and", luaFileCount, "lua files")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if len(output) > 0 {
		os.Exit(1)
	}

}

func run() (output string, err error) {

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".pl" {
			perlFileCount++
			cmd := exec.Command("perl", "-c", path, "2>/dev/null")
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				output += stderr.String()
			}
		}
		if filepath.Ext(path) == ".lua" {
			luaFileCount++
			cmd := exec.Command("luac", "-p", path)
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				output += stderr.String()
			}
		}
		return nil
	})
	return

}
