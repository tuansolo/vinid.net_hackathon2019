package resources

func NewInformationResource() string {
	resource := NewBaseResource()
	elements := []Element{
		{
			Type:    "web",
			Content: `<html><head><style>.image {display: block;width: 100%;text-align: center;}.image img {max-width: 200px;}div.information {margin-top: 1em;}div.information div {margin-top: 1em;}div.information div label:first-child {padding-bottom: .5em;display: block;font-weight: bold;}div.information div label:last-child {display: block;}</style></head><body><div class="image"><img src="https://img.thedailybeast.com/image/upload/dpr_2.0/c_crop,h_2147,w_2147,x_0,y_0/c_limit,w_128/d_placeholder_euli9k,fl_lossy,q_auto/v1567102268/190829-swin-stein-trump-puerto-rico-recovery-hero_eobp5s"></div><div class="information"><div><label>Mã học sinh</label><label>000001</label></div><div><label>Họ và tên học sinh</label><label>Đỗ Nam Trung</label></div><div><label>Trường</label><label>Trường tiểu học IDSchool</label></div><div><label>Lớp</label><label>3E</label></div></div></body></html>`,
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
