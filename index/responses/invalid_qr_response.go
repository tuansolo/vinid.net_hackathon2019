package responses

import (
	"fmt"
	"hackathon/index/resources"
	"net/http"
)

func RespondInvalidQR(w http.ResponseWriter) {
	_, _ = fmt.Fprint(w, resources.NewInvalidQRResource().ToJSON())
}
