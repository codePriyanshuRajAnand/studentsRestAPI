package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/types"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		slog.Info("Creating new student...")
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorWriter(fmt.Errorf("request body is empty")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorWriter(err))
			return
		}
		//request validation
		if err := validator.New().Struct(student); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		// w.Write([]byte("Welcome to Students Rest API"))
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
