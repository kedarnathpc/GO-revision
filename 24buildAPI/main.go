package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building the API !!!")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "ReackJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh", Website: "lco.dev"}})
	courses = append(courses, Course{CourseID: "4", CourseName: "MERN stack", CoursePrice: 299, Author: &Author{Fullname: "Hitesh", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOncCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

//serve home - route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Kedarnath Chavan !!!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}
	var cour Course
	_ = json.NewDecoder(r.Body).Decode(&cour)
	if cour.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//todo : check only if title is duplicate
	//loop, title matches course.Coursename, JSON

	//genearte unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	cour.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, cour)
	json.NewEncoder(w).Encode(cour)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	//loop, id, remove, add with the same id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// to do : send a response when id is not found
}

func deleteOncCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
