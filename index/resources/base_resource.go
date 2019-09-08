package resources

import "encoding/json"

type BaseResource struct {
	Data Data `json:"data"`
}

type Data struct {
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	AppName      string       `json:"app_name"`
	AppID        int          `json:"app_id"`
	Title        string       `json:"title"`
	SubmitButton SubmitButton `json:"submit_button"`
	ResetButton  *ResetButton `json:"reset_button,omitempty"`
	Elements     []Element    `json:"elements"`
}

type SubmitButton struct {
	Label           string `json:"label"`
	BackgroundColor string `json:"background_color"`
	CTA             string `json:"cta"`
	URL             string `json:"url"`
}

type ResetButton struct {
	Label           string `json:"label"`
	BackgroundColor string `json:"background_color"`
}

type Element struct {
	Label       string   `json:"label"`
	Type        string   `json:"type"`
	DisplayType string   `json:"display_type"`
	Required    bool     `json:"required"`
	Error       string   `json:"error"`
	Name        string   `json:"name"`
	Placeholder string   `json:"placeholder"`
	Options     []Option `json:"options"`
	Content     string   `json:"content"`
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func NewBaseResource() *BaseResource {
	resource := new(BaseResource)
	resource.Data.Metadata.AppName = "IDSchool"
	resource.Data.Metadata.AppID = 123456
	resource.Data.Metadata.Title = "IDSchool"
	return resource
}

func (b *BaseResource) ToJSON() string {
	bytes, _ := json.Marshal(b)
	return string(bytes)
}
