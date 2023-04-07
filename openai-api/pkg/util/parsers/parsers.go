package parsers

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"strconv"
)

const (
	YAML = "yaml"
	JSON = "json"
)

func ParserConfigurationByFile(format, in string, out interface{}) error {
	data, err := fs.ReadFile(os.DirFS("."), in)

	if err != nil {
		return err
	}

	switch format {
	case YAML:
		return yaml.Unmarshal(data, out)
	case JSON:
		return json.Unmarshal(data, out)
	default:
		return errors.New("invalid file format")
	}
}

func ParserInt64(in string) (int64, error) {
	return strconv.ParseInt(in, 10, 64)
}
