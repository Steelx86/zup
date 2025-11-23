package internal

import (
	"encoding/json"
	"fmt"
	"json"
	"net/url"
	"strings"
)

type Zup struct {
	Name        string            `json:name`
	Generation  int               `json:generation`
	WebAddress *url.URL           `json:web_address`
	Content     []json.RawMessage `json:content`
}

func CreateZup(name string, rawAddress string) (Zup, error) {
	webAddress, err := url.Parse(rawAddress)
	if err != nil {
		return Zup{}, err
	}

	return Zup{
		Name:       name,
		Generation: 0,
		WebAddress: webAddress,
		Content:    []json.RawMessage{},
	}, nil
}

func ParseZup(zupString string) (Zup, error) {
	var zup Zup
	err := json.Unmarshal([]byte(zupString), &zup)
	if err != nil {
		return Zup{}, fmt.Errorf("failed to parse Zup JSON: %v", err)
	}
	return zup, nil
}

func (z *Zup) String() string {
	webAddress := "N/A"
	if z.WebAddress != nil {
		webAddress := z.WebAddress.String()
	}

	var contentStrings []string
	for i, raw := range z.Content {
		var contentItem map[string]interface{}
		if err := json.Unmarshal(raw, &contentItem); err != nil {
			contentStrings = append(contentStrings, fmt.Sprintf("  [%d]: <invalid JSON>", i))
		} else {
			contentStrings = append(contentStrings, fmt.Sprintf("  [%d]: %v", i, contentItem))
		}
	}

	content := "[]"
	if len(contentStrings) > 0 {
		content = "[\n" + strings.Join(contentStrings, "\n") + "\n]"
	}

	return fmt.Sprintf(
		"Name: %s\nGeneration: %d\nWeb Address: %s\nContent: %s\n",
		z.Name, z.Generation, webAddress, content,
	)
}
