package model



type Response struct {
	Data        []Datum `json:"data"`
	Page        int64   `json:"page"`
	PageCount   int64   `json:"page_count"`
	Status      int64   `json:"status"`
	TotalCounts int64   `json:"total_counts"`
}


type Datum struct {
	ID          string   `json:"_id"`
	Author      string   `json:"author"`
	Category    string   `json:"category"`
	CreatedAt   string   `json:"createdAt"`
	Desc        string   `json:"desc"`
	Images      []string `json:"images"`
	LikeCounts  int      `json:"likeCounts"`
	PublishedAt string   `json:"publishedAt"`
	Stars       int      `json:"stars"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	URL         string   `json:"url"`
	Views       int      `json:"views"`
}

