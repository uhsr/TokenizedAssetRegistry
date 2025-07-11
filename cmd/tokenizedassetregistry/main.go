// cmd/tokenizedassetregistry/main.go
package main

import (
"flag"
"log"
"os"

"tokenizedassetregistry/internal/tokenizedassetregistry"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := tokenizedassetregistry.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
