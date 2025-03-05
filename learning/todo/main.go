package main

import (
	"fmt"
	"net/http"
	// to build the output string
)

// func main() {
// 	// fmt.Println("###### Welcome to our Todolist app.######")
// 	short := "This is short"
// 	full := "This is full"
// 	reward := "This is reward"
// 	taskItems := []string{short, full, reward}
// 	// // print the tasks
// 	// printTasks(taskItems)
// 	// fmt.Println()
// 	// // adding new tasks
// 	taskItems = addTask(taskItems, "This is a new task.")
// 	taskItems = addTask(taskItems, "This is another task.")
// 	// printing newly added tasks with all the tasks
// 	// printTasks(taskItems)
// 	// fmt.Println()
//
// 	// http.HandleFunc("/", hellowUser) // home page or root page
// 	http.HandleFunc("/about", about) // about page
// 	// 🆕 Passing the `taskItems` list to the home route handler
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		hellowUser(w, taskItems) // Calls hellowUser with the task list
// 	})
//
// 	// start the server
// 	http.ListenAndServe(":8080", nil) // this basically starts a server on the port 8080 (localhost::8080)-->for the first parameter and the second parameter is handler if we cant reach that port
// }
//
// func about(writer http.ResponseWriter, request *http.Request) {
// 	titleAbout := "Welcome to about Page"
// 	fmt.Fprint(writer, titleAbout)
// }
//
// // 🆕 hellowUser now receives `taskItems` as a parameter and prints them on the page
// func hellowUser(writer http.ResponseWriter, taskItems []string) {
// 	greeting := "###Welcome to the home page of Our TODO List app.###"
// 	// 🆕 Gets the formatted TODO list as a string
// 	taskList := printTasks(taskItems)
// 	// 🆕 Writes the greeting + task list to the web page
// 	fmt.Fprint(writer, greeting+taskList)
// }
//
// // 🆕 printTasks now returns a formatted string instead of printing directly
// func printTasks(taskItems []string) string {
// 	// 🆕 Using a `strings.Builder` to efficiently create a multi-line string
// 	var result strings.Builder
// 	result.WriteString("List of my TODOS:\n")
//
// 	for _, items := range taskItems {
// 		result.WriteString("- " + items + "\n")
// 	}
// 	return result.String()
// }
//
// // func hellowUser(writer http.ResponseWriter, request *http.Request) {
// // 	greeting := "###Welcome to the home page of Our TODO List app.###"
// // 	// fmt.Println(greeting) // this will print it as terminal output
// // 	// if
// // 	fmt.Fprint(writer, greeting)
// // }
// //
// // func printTasks(taskItems []string) {
// // 	fmt.Println("List of my TODO:")
// // 	for _, items := range taskItems {
// // 		fmt.Println(items)
// // 	}
// // }
//
// func addTask(taskItems []string, newTask string) []string {
// 	updatedItems := append(taskItems, newTask)
// 	return updatedItems
// }

// another way without using strings package
var (
	short     = "This is short"
	full      = "This is full"
	reward    = "This is reward"
	taskItems = []string{short, full, reward}
)

func main() {
	taskItems = addTask(taskItems, "This is a new task.")
	taskItems = addTask(taskItems, "This is another task.")

	http.HandleFunc("/", hellowUser)       // home page or root page
	http.HandleFunc("/listview", listview) // about page

	// start the server
	http.ListenAndServe(":8080", nil) // this basically starts a server on the port 8080 (localhost::8080)-->for the first parameter and the second parameter is handler if we cant reach that port
}

func listview(writer http.ResponseWriter, request *http.Request) {
	titleAbout := "List:\n"
	fmt.Fprint(writer, titleAbout)
	for _, items := range taskItems {
		fmt.Fprintln(writer, items)
	}
}

func hellowUser(writer http.ResponseWriter, request *http.Request) {
	greeting := "###Welcome to the home page of Our TODO List app.###"
	fmt.Println(greeting) // this will print it as terminal output
	fmt.Fprint(writer, greeting)
}

//	func printTasks(taskItems []string) {
//		fmt.Println("List of my TODO:")
//		for _, items := range taskItems {
//			fmt.Println(items)
//		}
//	}
func addTask(taskItems []string, newTask string) []string {
	updatedItems := append(taskItems, newTask)
	return updatedItems
}
