package main

import "fmt"

func main() {
	name := "My name is Puspo"
	code := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body><h2>` + name + `
</body>
</html>`

	fmt.Println(code) //tocreate a html file command: go run main.go > index.html

}
