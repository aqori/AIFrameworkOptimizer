// cmd/aiframeworkoptimizer/main.go
package main

import (
"flag"
"log"
"os"

"aiframeworkoptimizer/internal/aiframeworkoptimizer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := aiframeworkoptimizer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
