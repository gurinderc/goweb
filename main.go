package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
			<html>
				<head>
					<title>Mario Landing Page</title>
				</head>
				<body>
					<h1>Welcome to the Mario Landing Page</h1>
					<img src="https://upload.wikimedia.org/wikipedia/en/thumb/a/a9/MarioNSMBUDeluxe.png/220px-MarioNSMBUDeluxe.png" alt="Mario" style="width:220px;">
					<p>Mario is a fictional character in the Mario video game franchise, created by Nintendo's Japanese video game designer, Shigeru Miyamoto. He is the eponymous protagonist of the series and has appeared in over 200 video games since his creation. Mario is one of the most recognizable and successful video game franchises of all time.</p>
				</body>
			</html>
		`)
	})

	http.ListenAndServe(":8080", nil)
}
