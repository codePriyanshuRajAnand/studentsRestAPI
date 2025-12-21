package students

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/types"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/utils/response"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorWriter(err))
			return
		}
		slog.Info("Creating new student...")
		// w.Write([]byte("Welcome to Students Rest API"))
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
