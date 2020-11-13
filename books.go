package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Book Management !")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", getAllBooks).Methods("GET")
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", returnSingleBook)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Software Project Management", Desc: "a category of computer software designed to help streamline the complexity of large projects and tasks as well as facilitate team collaboration and project reporting", Content: "Introduction to Go"},
		Book{Id: "2", Title: "Golang", Desc: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software", Content: "Introduction to Software Project Management"},
	}
	handleRequests()
}

/*

package main

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// Library Details
type LibDetails struct {
	TotalBooks     int `json:"TotalBooks"`
	TotalAvaiBooks int `json:"TotalAvaiBooks"`
	TotalIssued    int `json:"TotalIssued"`
}

// Book Data Structure
type Book struct {
	Name      string `json:"Name"`
	Issues    int    `json:"Issues"`
	Available bool   `json:"Available"`
	Issuedto  string `json:"Issuedto"`
}

// Data for Books
var AllBooks = []Book{
	Book{Name: "Lean Startup", Issues: 20, Available: true, Issuedto: ""},
	Book{Name: "The secret", Issues: 40, Available: false, Issuedto: "sanket"},
	Book{Name: "Steve Jobs", Issues: 200, Available: true, Issuedto: ""},
	Book{Name: "Moby Dick", Issues: 10, Available: true, Issuedto: ""},
	Book{Name: "Intro to Algorithms", Issues: 5, Available: false, Issuedto: "sayam"},
	Book{Name: "Game of Thrones", Issues: 90, Available: true, Issuedto: ""},
	Book{Name: "The Snape", Issues: 40, Available: true, Issuedto: ""},
	Book{Name: "Hamlet", Issues: 14, Available: false, Issuedto: "sayam"},
	Book{Name: "The Wit and Wisdom", Issues: 10, Available: true, Issuedto: ""},
	Book{Name: "Infinite Game", Issues: 1, Available: false, Issuedto: "sayam"},
	Book{Name: "War and Peace", Issues: 0, Available: true, Issuedto: ""},
	Book{Name: "Zero to One", Issues: 100, Available: false, Issuedto: "akshay"},
	Book{Name: "Madame Bovary", Issues: 5, Available: false, Issuedto: "sayam"},
	Book{Name: "Artificial Intelligence", Issues: 34, Available: true, Issuedto: ""},
	Book{Name: "Why i Killed Gandhi", Issues: 22, Available: true, Issuedto: ""},
}

// Home page
func homepage(w http.ResponseWriter, r *http.Request) {

	var totalIssued, totalAvailable int
	var totalBooks = len(AllBooks)

	for i, s := range AllBooks {
		if s.Available == true {
			totalAvailable++
		}
		totalIssued += s.Issues
		i++
	}

	details := LibDetails{
		TotalBooks:     totalBooks,
		TotalAvaiBooks: totalAvailable,
		TotalIssued:    totalIssued,
	}

	fmt.Fprintf(w, "<h1>Data For Sayam's Library </h1> <br></br>")
	json.NewEncoder(w).Encode(details)
}

// Return All the Available Books
func AllBooksAvailable(w http.ResponseWriter, r *http.Request) {

	var books []Book

	for i, s := range AllBooks {
		if s.Available == true {
			books = append(books, s)
			fmt.Println(i, s.Name)
		}
	}
	fmt.Println("Endpoint hit: Available Books")
	json.NewEncoder(w).Encode(books)
}

// Return True if book is available false if not
func BookAvailable(w http.ResponseWriter, r *http.Request) {

	bookName := r.URL.Query().Get("book")
	for i, s := range AllBooks {
		if s.Name == bookName {
			if s.Available == false {
				json.NewEncoder(w).Encode(false)
			} else {
				json.NewEncoder(w).Encode(true)
			}
			break
		}
		i++
	}
}

// Return the user Book is Issued To
func IssuedUser(w http.ResponseWriter, r *http.Request) {

	bookName := r.URL.Query().Get("book")

	// Check If Book has Issued User return false if not
	for i, s := range AllBooks {
		if s.Name == bookName {

			if s.Issuedto == "" {
				json.NewEncoder(w).Encode(false)
			} else {
				json.NewEncoder(w).Encode(s.Issuedto)
			}
			break
		}
		i++
	}
}

// Return Most Issued
func MostIssued(w http.ResponseWriter, r *http.Request) {

	sort.Slice(AllBooks, func(i, j int) bool {
		return AllBooks[i].Issues < AllBooks[j].Issues
	})

	fmt.Println("Endpoint hit: Most Issued Book")
	json.NewEncoder(w).Encode(AllBooks[len(AllBooks)-1])
}

// Return Top Trending
func TopTrending(w http.ResponseWriter, r *http.Request) {

	var trendingBook = AllBooks[0]
	fmt.Println("Endpoint hit: Top Trending Book")
	json.NewEncoder(w).Encode(trendingBook)

}

func handleRequest() {

	// Handlers
	http.HandleFunc("/api/", homepage)
	http.HandleFunc("/api/booksAvailable", AllBooksAvailable)
	http.HandleFunc("/api/bookAvailable", BookAvailable)
	http.HandleFunc("/api/MostIssued", MostIssued)
	http.HandleFunc("/api/IssuedTo", IssuedUser)
	http.HandleFunc("/api/TopTrending", TopTrending)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}


*/
