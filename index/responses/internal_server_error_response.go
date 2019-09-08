package responses

import (
	"fmt"
	"hackathon/index/resources"
	"net/http"
)

func RespondInternalServerError(w http.ResponseWriter) {
	_, _ = fmt.Fprint(w, resources.NewInternalServerErrorResource().ToJSON())
}
