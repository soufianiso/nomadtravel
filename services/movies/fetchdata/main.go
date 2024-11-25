package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	// "io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type movie struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Original_language string `json:"original_language"`
	Overview string `json:"overview"`
	Poster_path string	`json:"poster_path"`
	Release_date string `json:"release_date"`
	Adult bool	`json:"adult"`
}

// Struct to represent the API response, which includes a list of movies
type movieResponse struct {
	Page    int     `json:"page"`
	Results []movie `json:"results"`
	Total_results int `json:"total_results"`
	Total_pages   int `json:"total_pages"`
}

func main() {
	godotenv.Load()
	conn := os.Getenv("postgres")
	

	for i:=1 ; i <= 100; i++ { 
    ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()

	rl := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=%v&sort_by=popularity.desc", i ) 
	req, err := http.NewRequestWithContext(ctx, "GET", rl, nil)


	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NjJlNWE2Y2VlOTZlZTU3ZDM1ZWI5N2Y1NTY2YmI1YiIsIm5iZiI6MTczMTY2ODk5Mi44NTk2Mjg0LCJzdWIiOiI2NmQxODk3YTVmNTk0NWUwNjg0NTU2Y2QiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.cEC2Mp8MnYu3vEZTMyvR9GyYrkk1IPd_sl5ckUy7ifw")

	    client := &http.Client{}

    // Perform the request
    res, err := client.Do(req)
    if err != nil {
        // Handle timeout or other errors
        return
    }

	movieResponse := new(movieResponse)

	if err := json.NewDecoder(res.Body).Decode(&movieResponse) ; err != nil{
		fmt.Println(err)
	}

	db ,err := sql.Open("postgres",conn)
	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	for _, m := range movieResponse.Results {
			fmt.Printf("ID: %d\n", m.Id)
			fmt.Printf("Title: %s\n", m.Title)
			fmt.Printf("Language: %s\n", m.Original_language)
			fmt.Printf("Overview: %s\n", m.Overview)
			fmt.Printf("Adult: %t\n", m.Adult)
			fmt.Printf("release_date: %s\n", m.Release_date)
			fmt.Printf("Poster Path: %s\n", m.Poster_path)
			fmt.Println("-------------------------------")
		_ , err := db.Exec(`INSERT INTO movies (
				original_title , original_language , overview, poster_path, release_date, adult) 
				VALUES ($1, $2 ,$3, $4, $5, $6)`,
			m.Title, m.Original_language, m.Overview, m.Poster_path, m.Release_date, m.Adult)
			if err != nil{
				fmt.Println(err)
			}
	}

	res.Body.Close()
	time.Sleep(4 * time.Second)
	}

}
