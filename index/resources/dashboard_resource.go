package resources

func NewDashboardResource() string {
	resource := NewBaseResource()
	elements := []Element{
		{
			Label:       "Chọn chức năng",
			Type:        "radio",
			DisplayType: "inline",
			Required:    true,
			Error:       "Vui lòng chọn một chức năng",
			Name:        "name",
			Options:     []Option{
				{
					Label: "Thông tin học sinh",
					Value: "information",
				},
				{
					Label: "Đăng ký đồng phục",
					Value: "uniform",
				},
				{
					Label: "Tra cứu thời khóa biểu",
					Value: "timetable",
				},
				{
					Label: "Thông tin học phí",
					Value: "tuition_fee",
				},
				{
					Label: "Đăng ký khóa học",
					Value: "course",
				},
			},
		},
	}
	resource.Data.Metadata.Elements = elements
	resource.Data.Metadata.SubmitButton = SubmitButton{
		Label:           "Tiếp theo",
		BackgroundColor: "#0275d8",
		CTA:             "request",
		URL:             "https://us-central1-index-hackathon.cloudfunctions.net/index",
	}
	return resource.ToJSON()
}
