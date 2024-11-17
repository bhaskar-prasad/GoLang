package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bhaskar-prasad/buildapi/models"
	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

var courses []models.Course

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Course API</h1>"))
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course models.Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.CourseName == "" {
		json.NewEncoder(w).Encode("Course name cannot be empty")
		return
	}

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course with this title already exists")
			return
		}
	}

	// Gen unq id
	rand.Seed(uint64(time.Now().Nanosecond()))
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse models.Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id")
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id")
}
