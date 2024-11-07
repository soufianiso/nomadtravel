package types

import(
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	OriginalTitle string `json:"original_title"`
	OriginalLanguage string `json:"original_language"`
	Overview string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	CreatedAt time.Time `json:"createdAt"`
}

 







 







