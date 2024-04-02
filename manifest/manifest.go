package manifest

import (
	"encoding/json"
	"io"
	"os"

	"github.com/emilwidlund/init/utils"
)

type Manifest struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

func LoadManifest(FilePath string) *Manifest {
	data, err := os.Open(FilePath)
	utils.Check(err)

	defer data.Close()

	bytes, _ := io.ReadAll(data)

	manifest := Manifest{}

	json.Unmarshal(bytes, &manifest)

	return &manifest
}
