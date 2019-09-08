package index

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"hackathon/index/resources"
	"hackathon/index/responses"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		show(w, r)
		return
	}

	if r.Method == http.MethodPost {
		redirect(w, r)
		return
	}
}

func show(w http.ResponseWriter, r *http.Request) {

	studentCode := r.URL.Query().Get("student_code")

	if len(studentCode) == 0 {
		responses.RespondInvalidQR(w)
		return
	}

	userID := r.Header.Get("user_id")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "index-hackathon")
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}
	defer client.Close()

	docsnap, err := client.Doc("students/" + studentCode).Get(ctx)
	if err != nil {
		responses.RespondInvalidQR(w)
		return
	}

	student := docsnap.Data()
	parentID := student["parent_id"].(string)

	if userID == parentID {
		respondToParent(w)
	} else {
		contactable := student["parent_contactable"].(bool)
		respondToStranger(w, contactable)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}

	var req request
	err = json.Unmarshal(b, &req)
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}
	switch req.Name {
	case "information":
		_, _ = fmt.Fprint(w, resources.NewInformationResource())
		break
	case "uniform":
		requestToUniformModule(w, r)
		break
	case "timetable":
		_, _ = fmt.Fprint(w, resources.NewTimetableResource())
		break
	case "tuition_fee":
		requestToTuitionFeeModule(w, r)
		break
	case "course":
		requestToCourseModule(w, r)
		break
	default:
		responses.RespondInternalServerError(w)
		break
	}
}

func respondToParent(w http.ResponseWriter) {
	_, _ = fmt.Fprint(w, resources.NewDashboardResource())
}

func respondToStranger(w http.ResponseWriter, contactable bool) {
	if contactable {
		_, _ = fmt.Fprint(w, resources.NewParentResource())
	} else {
		_, _ = fmt.Fprint(w, resources.NewStrangerResource())
	}
}

func requestToModule(w http.ResponseWriter, r *http.Request, url string) {
	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		responses.RespondInternalServerError(w)
		return
	}
	_, _ = fmt.Fprint(w, string(body))
}

func requestToInformationModule(w http.ResponseWriter, r *http.Request) {
	requestToModule(w, r, "https://us-central1-index-hackathon.cloudfunctions.net/getCourse")
}

func requestToUniformModule(w http.ResponseWriter, r *http.Request) {
	requestToModule(w, r, "https://us-central1-index-hackathon.cloudfunctions.net/getUniform")
}

func requestToTimetableModule(w http.ResponseWriter, r *http.Request) {
	requestToModule(w, r, "https://us-central1-index-hackathon.cloudfunctions.net/getCourse")
}

func requestToTuitionFeeModule(w http.ResponseWriter, r *http.Request) {
	requestToModule(w, r, "https://us-central1-gentle-epoch-234103.cloudfunctions.net/getStudentFee")
}

func requestToCourseModule(w http.ResponseWriter, r *http.Request) {
	requestToModule(w, r, "https://us-central1-index-hackathon.cloudfunctions.net/getCourse")
}
