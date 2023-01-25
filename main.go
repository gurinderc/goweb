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
					<style>
						body {
							background-color: #f0f0f0;
							font-family: Arial, sans-serif;
							text-align: center;
						}
						h1 {
							color: #333333;
							margin-top: 50px;
						}
						img {
							border-radius: 20%;
							margin: 20px 0;
							box-shadow: 10px 10px 5px #888888;
						}
						p {
							color: #555555;
							font-size: 22px;
							line-height: 1.6;
							margin: 20px 0;
							text-align: center;
						}
				    </style>
				</head>
				<body>
					<h1>Welcome to the Mario Landing Page</h1>
					<img src="https://upload.wikimedia.org/wikipedia/en/thumb/a/a9/MarioNSMBUDeluxe.png/220px-MarioNSMBUDeluxe.png" alt="Mario" style="width:220px;">
					<p>Mario is a fictional character in the Mario video game franchise, created by Nintendo's Japanese video game designer, Shigeru Miyamoto. He is the eponymous protagonist of the series and has appeared in over 200 video games since his creation. Mario is one of the most recognizable and successful video game franchises of all time.</p>
				</body>
			</html>
		`)
	})

	http.ListenAndServe(":8081", nil)
}
