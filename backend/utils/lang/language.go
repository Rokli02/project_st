package lang

import (
	"fmt"
	"os"
	"path"
	"st/backend/utils/logger"

	"gopkg.in/yaml.v3"
)

var Text *Language = &Language{}

func LoadLanguage(languageId string) {
	logger.InfoF("Loading texts for the choosen language (%s) ...", languageId)

	workingDirectory, _ := os.Getwd()
	dbPath := path.Join(workingDirectory, "data", "lang", languageId+".yaml")

	file, err := os.OpenFile(dbPath, os.O_RDONLY, os.ModeDevice)
	if err != nil || file == nil {
		logger.ErrorF("Couldn't open language file, (%s)", err)

		panic(-1)
	}

	// Determine the length of the choosen language file
	fileLength, err := file.Seek(0, 2)
	if err != nil {
		logger.Error("Error:", err)
	}

	// Settings back the cursor to the begining of language file
	file.Seek(0, 0)

	fileBytes := make([]byte, fileLength)
	file.Read(fileBytes)

	yaml.Unmarshal(fileBytes, Text)

	logger.Info("Texts for the choosen language is loaded")
}

func (m *TextMap) Get(key string) string {
	if value, has := (*m)[key]; has {
		return value
	}

	return fmt.Sprintf(Text.Unknown, key)
}
