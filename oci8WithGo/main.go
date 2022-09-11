package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mattn/go-oci8"
)

var openString string = "ebill/ebill[*]@202.79.32.157:1521/arcdb"

type Document struct {
	Username      string         `json:"username"`
	Type          sql.NullString `json:"type"`
	Remarks       sql.NullString `json:"remarks"`
	Uploaded_By   string         `json:"uploaded_by"`
	Updated_On    time.Time      `json:"updated_on"`
	Uploaded_From string         `json:"uploaded_from"`
}

type DocumentSingle struct {
	Data []Document
}

func databaseQuery(queryType string, param []string, w http.ResponseWriter) {
	oci8.Driver.Logger = log.New(os.Stderr, "oci8", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	db, err := sql.Open("oci8", openString)
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("db is nil")
			return
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	err = db.PingContext(ctx)
	cancel()
	if err != nil {
		fmt.Println("PingContext error is not nil:", err)
		return
	}

	var queryString string
	var rows *sql.Rows
	ctx, cancel = context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()

	fmt.Println(param)

	switch {
	case queryType == "index":
		queryString = "select username, type, remarks, uploaded_by, updated_on, uploaded_from from newaccount_imageupload_arc where username = '" + param[0] + "' group by username, type, remarks, uploaded_by, updated_on, uploaded_from order by updated_on desc"
	case queryType == "show":
		queryString = "select username, type,remarks,uploaded_by,updated_on,uploaded_from from newaccount_imageupload_arc where username = '" + param[0] + "' and type = '" + param[1] + "'"
	default:
		queryString = ""
	}
	rows, err = db.QueryContext(ctx, queryString)
	if err != nil {
		fmt.Println("QueryContext error is not nil:", err)
		return
	}

	// defer close rows
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
			return
		}
	}()

	// countRows, _ := rows.Columns()
	// dest := make([]interface{}, len(countRows))
	// destPointer := make([]interface{}, len(countRows))

	// Scan copies the columns in the current row into the values pointed at by destPointer in this scenario which is the value of dest. So that's why we are assigning &dest[i] to destPointer[i] above
	for rows.Next() {
		var document Document
		err := rows.Scan(&document.Username, &document.Type, &document.Remarks, &document.Uploaded_By, &document.Updated_On, &document.Uploaded_From)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(document)

		// for i := 0; i < len(countRows); i++ {
		// 	destPointer[i] = &dest[i]
		// }
		// err = rows.Scan(destPointer...)
		// if err != nil {
		// 	fmt.Println("Scan error is not nil:", err)
		// 	return
		// }

		// fmt.Println(destPointer...)

		// sqlResult := make(map[string]interface{})

		// for i := 0; i < len(countRows); i++ {
		// sqlResult[countRows[i]] = dest[i]
		// fmt.Println(dest[i])
		// }

		// var document Document
		// var arrayOfDocs []interface{}

		// err = mapstructure.Decode(sqlResult, &document)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(sqlResult)

		// for _, value := range sqlResult {
		// fmt.Println(key)
		// fmt.Println(value)
		// arrayOfDocs = append(arrayOfDocs, value)
		// }
		// fmt.Println(arrayOfDocs)
		// fmt.Println(document)

		json.NewEncoder(w).Encode(document)

	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Err error is not nil:", err)
		return
	}

	// sqlResult2 := []Document{
	// 	{"something", "something", "something", "something", time.Now(), "something"},
	// 	{"something", "something", "something", "something", time.Now(), "something"},
	// }
	// json.NewEncoder(w).Encode(sqlResult2)

}

func index(w http.ResponseWriter, r *http.Request) {
	var param []string
	param = append(param, chi.URLParam(r, "username"))
	databaseQuery("index", param, w)
}

func show(w http.ResponseWriter, r *http.Request) {
	var param []string
	param = append(param, chi.URLParam(r, "username"))
	param = append(param, chi.URLParam(r, "type"))
	databaseQuery("show", param, w)
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is document View"))
}

func store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is document Store"))
}

func destroy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is document destroy"))
}

func main() {
	fmt.Println("Working with oci8 for go")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/documents/{username}", func(r chi.Router) {
		r.Get("/", index)
		r.Get("/{type}", show)
		r.Get("/{type}/view", view)
		r.Post("/{type}", store)
		r.Delete("/{type}", destroy)
	})

	fmt.Println("Listening at port 4500")
	http.ListenAndServe(":4500", r)

}
