package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/storage"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/types"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create(storage storage.Storage) http.HandlerFunc {
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

		id, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
		}
		slog.Info("User Created successfully!", slog.String("StudentID", fmt.Sprint(id)))

		// w.Write([]byte("Welcome to Students Rest API"))
		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"status": "OK", "idCreated": id})
	}
}
