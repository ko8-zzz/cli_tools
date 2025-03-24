package cli_selector

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	"ko8-zzz/cli_tools/internal/common"
)

type ActionNode struct {
    Name   string       `json:"name"`
    Action string       `json:"action,omitempty"`
    Next   []*ActionNode `json:"next,omitempty"`
}

func Selector(optionFilePath string) string {

    // デコードされたデータを格納するためのスライス
    var options []*ActionNode

    // JSONファイルの読み込みとデコード
    err := common.ReadJSONFile(optionFilePath, &options)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return "error"
    }

    selectedOptionName := navigateOptions(options)
    fmt.Printf("Selected Option: %s\n", selectedOptionName)
	
	return selectedOptionName
}

func navigateOptions(nodes []*ActionNode) string {
    reader := bufio.NewReader(os.Stdin)
    for {
        displayOptions(nodes)
        fmt.Printf("Select an option (0-%d): ", len(nodes)-1)

        choice, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        choice = choice[:len(choice)-1] // Remove newline character

        index, err := strconv.Atoi(choice)
        if err != nil || index < 0 || index >= len(nodes) {
            fmt.Println("Invalid option, please select a valid option.")
            continue
        }

        selectedNode := nodes[index]
        if len(selectedNode.Next) == 0 {
            return selectedNode.Name // Return the selected name if there are no next options
        }
        nodes = selectedNode.Next
    }
}

func displayOptions(nodes []*ActionNode) {
    for i, node := range nodes {
        fmt.Printf("%d: %s\n", i, node.Name)
    }
}