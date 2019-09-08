package resources

func NewStrangerResource() string {
	resource := NewBaseResource()
	elements := []Element{
		{
			Type:    "web",
			Content: `<html><head><style>body div {display: block;width: 100%;margin: 0px auto;text-align: center;}body div h2 {color: #d9534f}</style></head><body><div><img src="data:image/png;base64,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" /><h2>TRƯỜNG TIỂU HỌC IDSCHOOL</h2><a href="tel:0912345678"><h4>0912345678</h4></a></div></body></html>`,
		},
	}
	resource.Data.Metadata.Elements = elements
	resource.Data.Metadata.SubmitButton = SubmitButton{
		Label:           "Đóng",
		BackgroundColor: "#0275d8",
		CTA:             "close",
		URL:             "https://us-central1-index-hackathon.cloudfunctions.net/index",
	}
	return resource.ToJSON()
}