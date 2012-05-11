package main

// TODO:
//  validate input
//  flag for format
//  could also loop over args and do out

import (
    "os"
    "flag"
    "time"
    "strconv"
)

func main() {
    os.Setenv("TZ", "GMT")
    flag.Parse()
    
    arg, _ := strconv.Atoi64( flag.Arg(0) )
    t := time.SecondsToLocalTime(arg)
    out := t.Format(time.RFC850)
    
    os.Stdout.WriteString(out + "\n")
}
