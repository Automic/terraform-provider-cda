package cda

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func SetUpTest() error {
	file, err := os.Open("test.config")
	if err != nil {
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var mapConfig map[string]string
	mapConfig = make(map[string]string)

	for {
		line, err := reader.ReadString('\n')

		// check if the line has = sign
		// and process the line. Ignore the rest.
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the config map
				mapConfig[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	os.Setenv("TF_ACC", mapConfig["TF_ACC"])
	os.Setenv(environmentVariableMaps[cdaServer], mapConfig["CDA_SERVER"])
	os.Setenv(environmentVariableMaps[cdaUser], mapConfig["CDA_USER"])
	os.Setenv(environmentVariableMaps[password], mapConfig["CDA_PASSWORD"])

	return nil
}
