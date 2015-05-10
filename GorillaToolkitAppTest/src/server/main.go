/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main 

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


// list of all of the books
var books = make([]book, 0)

func main() {
	// command line flags
	port := flag.Int("port", 8080, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/web/", 302))
	router.Handle("/books", handler(listBooks)).Methods("GET")
	router.Handle("/books", handler(addBook)).Methods("POST")
	router.Handle("/books/{id}", handler(getBook)).Methods("GET")
	router.Handle("/books/{id}", handler(updateBook)).Methods("POST")
	router.Handle("/books/{id}", handler(removeBook)).Methods("DELETE")
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web", fileHandler))
	http.Handle("/", router)

	// bootstrap some data
	books = append(books, book{"Ender's Game", "Orson Scott Card", getNextId()})
	books = append(books, book{"Code Complete", "Steve McConnell", getNextId()})
	books = append(books, book{"World War Z", "Max Brooks", getNextId()})
	books = append(books, book{"Pragmatic Programmer", "David Thomas", getNextId()})

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
