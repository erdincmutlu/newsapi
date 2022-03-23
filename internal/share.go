package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erdincmutlu/newsapi/types"
)

func shareNews(w http.ResponseWriter, r *http.Request) {
	var request types.ShareRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Decode request body error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = share(request)
	if err != nil {
		log.Printf("share error %s", err.Error())
		if err == types.ErrInvalidInparams {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	resp := map[string]string{"message": "Status Ok"}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("json encode error: %s", err.Error())
	}
}

func share(request types.ShareRequest) error {
	log.Printf("Share request: %+v\n", request)
	var err error

	switch request.Action {
	case types.ActionEmail:
		err = sendEmail(request)
	case types.ActionTwit:
		// twit the news
	default:
		err = types.ErrInvalidInparams
	}

	return err
}
