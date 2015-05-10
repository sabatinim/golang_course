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
// new dependency: ngResource is included just above
var myapp = new angular.module("myapp", [ "ngResource" ]);

// inject the $resource dependency here
myapp.controller("MainCtl", [ "$scope", "$resource",
		function($scope, $resource) {
			// I designed the backend to play nicely with angularjs so this is
			// all the
			// setup we need to do all of the ususal operations.
			var Book = $resource("/books/:id", {
				id : '@id'
			}, {});

			$scope.selected = null;

			$scope.list = function(idx) {
				// Notice calls to Book are often given callbacks.
				Book.query(function(data) {
					$scope.books = data;
					if (idx != undefined) {
						$scope.selected = $scope.books[idx];
						$scope.selected.idx = idx;
					}
				}, function(error) {
					alert(error.data);
				});
			};

			$scope.list();

			$scope.get = function(idx) {
				// Passing parameters to Book calls will become arguments if
				// we haven't defined it as part of the path (we did with id)
				Book.get({
					id : $scope.books[idx].id
				}, function(data) {
					$scope.selected = data;
					$scope.selected.idx = idx;
				});
			};

			$scope.add = function() {
				// I was lazy with the user input.
				var title = prompt("Enter the book's title.");
				if (title == null) {
					return;
				}
				var author = prompt("Enter the book's author.");
				if (author == null) {
					return;
				}
				// Creating a blank book object means you can still $save
				var newBook = new Book();
				newBook.title = title;
				newBook.author = author;
				newBook.$save();

				$scope.list();
			};

			$scope.update = function(idx) {
				var book = $scope.books[idx];
				var title = prompt("Enter a new title", book.title);
				if (title == null) {
					return;
				}
				var author = prompt("Enter a new author", book.author);
				if (author == null) {
					return;
				}
				book.title = title;
				book.author = author;
				// Noticed I never created a new Book()?
				book.$save();

				$scope.list(idx);
			};

			$scope.remove = function(idx) {
				$scope.books[idx].$delete();
				$scope.selected = null;
				$scope.list();
			};
		} ]);
