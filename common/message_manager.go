package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Messages ...
type Messages struct {
	Path     string
	Language string
	Log      *Logger
}

// Supported Language
const (
	Korean  string = "KO"
	English string = "EN"
)

// New create message manager
func (Messages) New() *Messages {
	return &Messages{}
}

// Initialize 언어 설정에 따라 메시지 초기화
func (mm *Messages) Initialize() error {
	switch mm.Language {
	case Korean:
	case English:
	default:
		return fmt.Errorf("ERROR. Not supported language")
	}
	// mm.Log.Debugf("Initialze Message From Language[ %s ]. Path[ %s ]", mm.Language, mm.Path)
	// Path of File
	path := fmt.Sprintf("%s/%s.json", mm.Path, strings.ToLower(mm.Language))
	// Open jsonFile
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("ERROR. Can't Read Http Response Message File[ %s ]", err.Error())
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize our struct
	// var msg model.Messages

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	// json.Unmarshal(byteValue, &msg)

	mm.Log.Errorf("%+v", string(byteValue))

	// HttpErrMsg = make(map[int]string, 0)
	return nil
}
