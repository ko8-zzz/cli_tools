package main

import (
   "fmt"
   "ko8-zzz/cli_tools/internal/cli_selector"
)

func main() {

   optionFilePath := "./configs/option_list.json"
   val := cli_selector.Selector(optionFilePath)
   fmt.Printf(val)
}