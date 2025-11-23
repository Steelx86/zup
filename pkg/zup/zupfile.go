package zup

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Zup struct {
	Name       string              `json:"name"`
	CreatedOn  string              `json:"created_on"`
	UpdatedOn  string              `json:"updated_on"`
	Generation int                 `json:"generation"`
	WebAddress *url.URL            `json:"web_address"`
	Formats    map[string][]string `json:"formats"`
	Content    []json.RawMessage   `json:"content"`
}

func CreateZup(name string, rawAddress string) (*Zup, error) {
	webAddress, err := url.Parse(rawAddress)
	if err != nil {
		return &Zup{}, err
	}

	currentTime := time.Now().Format(time.RFC3339)

	return &Zup{
		Name:       name,
		CreatedOn:  currentTime,
		UpdatedOn:  currentTime,
		Generation: 0,
		WebAddress: webAddress,
		Formats:    make(map[string][]string),
		Content:    []json.RawMessage{},
	}, nil
}

func (z *Zup) String() string {
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
		z.Name, z.Generation, z.WebAddress.String(), content,
	)
}

func (z *Zup) AddFormat(formatName string, fields []string) {
	z.Formats[formatName] = fields
}

func (z *Zup) AddContentItem(item json.RawMessage) {
	z.Content = append(z.Content, item)
}
