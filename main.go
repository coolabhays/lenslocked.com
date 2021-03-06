package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
	"github.com/gorilla/mux"
)

/* func handlerFunc(w http.ResponseWriter, req *http.Request) {
	} else {
		w.WriteHeader(http.StatusNotFound)
		// if WriteHeader is not called before Write(), the first call to Write
		// will trigger implicit WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>We couldn't find the page you are looking for :(</h1>" +
			"<p>Please email us if you keep being sent to an invalid page</p>")
	}
} */

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq_helper(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>FAQ!</h1>" +
		"<p><b>FAQ</b> will be added soon...</p>")
}

func status404(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Sorry :(</h1>" +
		"<p>The page you are looking for is not available</p>")
}

/**
 * To change a function to a http.Handler type, use http.HandlerFunc() wrapper
 * var h http.Handler = http.HandlerFunc(status404)
 * r := mux.NewRouter()
 * r.NotFoundHandler = h
 */

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()  // like we do with http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq_helper)
	r.NotFoundHandler = http.HandlerFunc(status404)  // gorilla/mux provides default page for 404 too
	http.ListenAndServe(":3000", r)
}

// must is helper wrapper, that panics upon an error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

/** Things to checkout:
 *  Dynamic reloading: https://github.com/pilu/fresh - Box3.2
 *  Another router: https://github.com/julienschmidt/httprouter - Ex3
 *	Closure: https://www.calhoun.io/what-is-a-closure
 *  Glob vs Regex: http://stackoverflow.com/questions/23702202/what-are-the-differences-between-glob-style-pattern-and-regular-expression
replicate the program we have written so far using this router instead */
