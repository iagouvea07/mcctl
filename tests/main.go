package main

import (
	"time"
	"github.com/schollz/progressbar/v3"
)


func main() {

	bar := progressbar.DefaultBytes(1000)
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
}
}