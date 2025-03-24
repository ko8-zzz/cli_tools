package common

import (
	"encoding/json"
    "io/ioutil"
    "os"
)

func Sample () {
	
}

// JSONファイルの内容を読み込む関数
func ReadJSONFile(optionFilePath string, v interface{}) error {
    // ファイルを開く
    file, err := os.Open(optionFilePath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // ファイルの内容を読み取る
    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }

    // JSONデコード
    err = json.Unmarshal(byteValue, v)
    if err != nil {
        return err
    }

    return nil
}