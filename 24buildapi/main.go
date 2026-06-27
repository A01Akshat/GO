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

// MODEL FOR COURSE

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB

var courses []Course

// Middleware, Checking conditions

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang",
		CoursePrice: 399, Author: &Author{Fullname: "Aksh", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Java",
		CoursePrice: 599, Author: &Author{Fullname: "Aksh", Website: "code.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("Delete")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL COURSES")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Sets the content type of the response to JSON, so the client knows what format to expect.
	// grab id from request
	params := mux.Vars(r)
	//loop through courses, find mathing id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) //Encodes it into JSON, and Sends it directly to the HTTP response (w), so the client receives it.
			return
		}
	}
	json.NewEncoder(w).Encode("NO COURSE FOUND") 
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE ONE COURSES")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("SEND SOME DATA!")
	}
	// what if data is sent but its empty, like : {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // dont care for return value, Also as we are refering to original value hence the decoded value is directly gone into course (i.e &course). Hence we dont care for return value
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("SEND SOME DATA!")
		return
	}

	// generate Unique Id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// We did above 2 steps coz initially in struct we have taken CourseId as string now to convert it to unique number and then to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE ONE COURSES")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from req

	params := mux.Vars(r) // grabs the id

	// loop,id,remove,add with same id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course                           // grabs updated data in course
			_ = json.NewDecoder(r.Body).Decode(&course) // decodes the resp sent by client
			course.CourseId = params["id"]              // Reassigns the original ID to the updated course — so the course keeps its identity even after the update.
			courses = append(courses, course)           // Appends new course back again
			json.NewEncoder(w).Encode(course)           //sends back the updated data as JSON
			return
		}
	}
	// send resp when id not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE ONE COURSES")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from req

	params := mux.Vars(r) // grabs the id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
