package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nedpals/supabase-go"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

	// Pull today's row from Supabase DB
	var results map[string]interface{}
	err := supabaseClient.DB.From("classes").Select("*").Single().Execute(&results)
  
	if err != nil {
		panic(err)
	}

	// Convert results to JSON
	jsonResults, _ := json.Marshal(results)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("User-Agent", "auto-repo")
	fmt.Fprintf(w, string(jsonResults))
}