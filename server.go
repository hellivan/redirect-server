package main


import (
    "net/http"
    "os"
    "log"
)



func main() {

    var port string
    var dest string

    switch len(os.Args) {

    case 2:
    	 port = "80"
	 dest = os.Args[1]
    case 3:
    	 port = os.Args[1]
	 dest = os.Args[2]

    default:
         log.Printf("USAGE: %s [PORT] REDIRECT_URL\n", os.Args[0])
         os.Exit(1)
    }   

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, dest, 301)
    })

    log.Printf("Redirecting all incoming http requests on port %s to '%s'\n", port, dest)
    err := http.ListenAndServe(":" + port, nil)

    if err != nil {
        log.Fatal(err)
	os.Exit(1)
    }   

}
