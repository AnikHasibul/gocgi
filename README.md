# gocgi

usage:


	import (
	    "net/http"
	    "github.com/anikhasibul/gocgi"
	)

	func main() {    
	    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	    	gocgi.CGIfy(w, r, "php-cgi", "./server.php")
	    })
	    http.ListenAndServe(":8080", nil)
	}
	
server.php:

	<?php
	echo "Hello, {$_GET['name']}";
	
visit: http://localhost:8080/?name=World

	Hello, World
