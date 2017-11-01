package main

import (
	"fmt"
	"github.com/gavruk/go-blog-example/models"
	"html/template"
	"net/http"
)

var posts map[string]*models.Post

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(posts)
	t.ExecuteTemplate(w, "index", posts)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "form", nil)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "form", post)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormaValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post = models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}
	http.Redirect(w, r, "/", 302)
}

func deleteHandler (w http.ResponseWriter, r *http.Request) {
	id := r.FormaValue("id")
	if id == "" {
		hhtp.NotFound(w, r)
	}
	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port :3000")

	posts = make(map[string]*models.Post, 0)

	//example.com/assets/css/app.css
	http.Handle("/assets/", httpStripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.ListenAndServe(":3000", nil)
	http.HandleFunc("/SavePost", savePostHandler)
}
