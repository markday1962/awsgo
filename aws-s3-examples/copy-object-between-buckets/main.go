package main

import (
	"fmt"
	"os"
)

//Create a function we use to display errors and exit.
func exitErrorf(msg string, args ...interface{}){
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

//The following example copies an item from one bucket to another with the names specified
//as command line arguments.
func main() {
	if len(os.Args) != 4 {
		exitErrorf("Source bucket, Object, Target bucket are required" +
			"\nUsage: %s source_bucket object_name target_bucket", os.Args[0])
	}

}
