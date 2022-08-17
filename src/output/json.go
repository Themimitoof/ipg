package output

import (
	"encoding/json"
	"os"
)

type IpgJsonOutput struct {
	Config *OutputInformationSettings
	Data   IpgOutputData
}

func (c IpgJsonOutput) Render() []byte {
	b, err := json.Marshal(c.Data)

	if err != nil {
		os.Stderr.WriteString("Unable to output the JSON.")
		os.Exit(2)
	}

	return append(b, '\n')
}
