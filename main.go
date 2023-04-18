package main

import (
	"sesi_12/database"
	"sesi_12/routers"
)

func main() {
	// mux := http.NewServeMux()

	// endpoint := http.HandlerFunc(greet)

	// mux.Handle("/", middleware(middleware2(endpoint)))

	// fmt.Println("Linstening to port 3000")

	// err := http.ListenAndServe(":3000", mux)
	// log.Fatal(err)

	// beda
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World!!"))
// }

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware pertama")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func middleware2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware pertama")
// 		next.ServeHTTP(w, r)
// 	})
// }
