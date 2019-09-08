package resources

func NewInvalidQRResource() *BaseResource {
	resource := NewBaseResource()
	elements := []Element{
		{
			Type:    "web",
			Content: `<html><head><style>body div {display: block;width: 100%;margin: 0px auto;text-align: center;}body div h2 {color: #d9534f}</style></head><body><div><img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAMAAACahl6sAAAAbFBMVEX///8AAACrq6vHx8fz8/Ovr69jY2NHR0crKyvDw8MnJycfHx9LS0uPj4/39/enp6ff398TExN3d3d/f39ra2tbW1uXl5cPDw+Dg4M7OzuLi4tTU1Ofn59vb2/X19c/Pz/l5eXb29u5ubkyMjK85//bAAAEa0lEQVR4nO2d2WKqMBRFRdSqOOGsVYr6//94yVFPGgwhKqHBu9dTMx2yHgpkILZaAAAAAAAAAADAJ9OxY+K2FxPLbhhCBHbM3YrMLbsBEYhABCIfJzI0s3crsi+5vL1I7Laj7xJDxDMg4hsQ8Q2I+AZEfAMivgER33hPJG3bYZpt+kXHMlxauUg7sCO0ExlYhutCBCKNEhn2TEQs0j5kaKa6OyK/yyKRMdzQoUjP0LbV+mIR6kP6WGMm8vssMjKG60EEIjeaJhKOHzm9LnLShEtqERkHj7QLRMb7G+tCkakm3No7kc29wgYiEPkIkcUV3frHayLbhRKyGpH9LaZBxMALIhPmHqQakfd4QeSHi2cQgQhENCKj+EYQplfCZopwW7UuRCACkbwIb76IGi7CNPT267vIqWjSPC+y+L5zd9KJ6Kb5T7WIFPLak70QiJiBiJci0ZeJWBX54QH7+T5EzInExnCRQxELpIiEg+RELIAIRBoj0gntoHX2RM3jIBORSp4OBwD4Ew7qlnsaihxL9uXf+FEjDbiAJqQPfYXqb1g5cgOrg8j7truZRmqkLhcsRXKvCQwRiDRfZBNl0Bv7VEzvXkpFkuhOzAU6EQp8rk9EQhfdlYtoCnQihMtbF0QgUindzZ0jf4e95OtN+bNsEglFMuLS3Pfjmi+61yLwnJM9hyJyYHXkvCPntTmPRGjIIUXKo69FNfmd5hIi5UDEc5EFJ+WbxBfnaUQ6gcKuSEQXGCIQaZbIUjyy6ON5ej7K633zM/NULqKebEIiRzXw1LWIZGxsYRQh5LLCWi1Yu3GACERqFKFxKP2fd2YF0ECWhrQipXt514lQ4DEHqf68pZyI3OLZC0zQjamoUCeyFcl6XlEgApFKOaszzPK+suibmLCISMkluc0lQw4I2mrgLSddzms9D/Vc/DFjkV1ZGy+BiG94JdIZ2EG3qZCTUkSk6Na32WUsOXDKdWecN1PDVcqLGwY0t9++GlhuBax78gEiEKleJBqZyG0804hsRLWlGnjA7WlUkHQzRqLyReRZfh7/rMhTWwE1IpGxPSFnUVxOPkAEItXza0v59JFUFTmuMmjJ+iBKuW+xpmmoxqQV+4toP3AtYrHJn4g09TSsRN26Z1EgAhHvReJlxvUBvs5YNVaEBlbdolKIQOS/ENnOf3NUG5BIIgpGMo8rb+/HqexPHoioFC6G/snACiJ1i1R1nIhbEYvjRG6t3z7gRbne9YqTrQoXLkSK/uN3IjUoF6nxpBrCuDtIftGTe9eyWB+BCEQ8F0nWj+SObsuJLB+Zc9Mzx6Th+k6U9msRKeTpd60ctvt+IQIRf0WGKxNyNp7o8qQbd4sm6Ki/sWiQexVKxJGto0DBr2NyuTC3yV/3Yf5CDQcRiDRGJO3aoVsY50J6inc4qZNO1HBp5SIeARHfgIhvQMQ3IOIbEPENiPgGRHzDXiSIzDj+OdptyeUDe5ESmvMDwRCBCEQg0gyRiR1uPbzpBgAAAAAA+DD+ASLirBB9tfN/AAAAAElFTkSuQmCC" /><h2>MÃ QR KHÔNG ĐÚNG</h2></div></body></html>`,
		},
	}
	resource.Data.Metadata.Elements = elements
	resource.Data.Metadata.SubmitButton = SubmitButton{
		Label:           "Đóng",
		BackgroundColor: "#0275d8",
		CTA:             "close",
		URL:             "https://us-central1-index-hackathon.cloudfunctions.net/index",
	}
	return resource
}
