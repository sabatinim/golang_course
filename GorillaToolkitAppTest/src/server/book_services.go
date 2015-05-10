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
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)



func listBooks(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	
	return books, nil
}

func getBook(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	// mux.Vars grabs variables from the path
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	b, index := getBookById(id)

	if index < 0 {
		return nil, &handlerError{nil, "Could not find book " + param, http.StatusNotFound}
	}
	return b, nil
}

func addBook(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	payload, e := parseBookRequest(r)
	if e != nil {
		return nil, e
	}

	// it's our job to assign IDs, ignore what (if anything) the client sent
	payload.Id = getNextId()
	books = append(books, payload)

	// we return the book we just made so the client can see the ID if they want
	return payload, nil
}

func updateBook(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	payload, e := parseBookRequest(r)
	if e != nil {
		return nil, e
	}

	_, index := getBookById(payload.Id)
	books[index] = payload
	return make(map[string]string), nil
}

func removeBook(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	// this is jsut to check to see if the book exists
	_, index := getBookById(id)

	if index < 0 {
		return nil, &handlerError{nil, "Could not find entry " + param, http.StatusNotFound}
	}

	// remove a book from the list
	books = append(books[:index], books[index+1:]...)
	return make(map[string]string), nil
}

// searches the books for the book with `id` and returns the book and it's index, or -1 for 404
func getBookById(id int) (book, int) {
	for i, b := range books {
		if b.Id == id {
			return b, i
		}
	}
	return book{}, -1
}

var id = 0

// increments id and returns the value
func getNextId() int {
	id += 1
	return id
}

