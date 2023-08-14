package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const dogsFile = "./data/dogs.json"

func main() {
	type dataDogs struct {
		Dogs []struct {
			ID          int64      `json:"id"`
			Name        string     `json:"name"`
			Birthday    *time.Time `json:"birthday,omitempty"`
			Deathday    *time.Time `json:"deathday,omitempty"`
			YearsOfLife int64      `json:"years_of_life,omitempty"`
		} `json:"dogs"`
	}

	http.HandleFunc("/api/v1/pets", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(dogsFile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("opening file: %v", err)))
			return
		}

		var res *dataDogs
		err = json.Unmarshal(data, &res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unmarshaling: %v", err)))
			return
		}

		for i, d := range res.Dogs {
			if d.Deathday == nil || d.Birthday == nil {
				continue
			}

			years := 0
			for year := d.Birthday.Year(); year < d.Deathday.Year(); year++ {
				years += 1
			}

			fmt.Printf("YearsOfLife: %v", years)
			d.Birthday = nil
			d.Deathday = nil
			d.YearsOfLife = int64(years)

			res.Dogs[i] = d
		}

		dogs, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("marshaling: %v", err)))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dogs)
	})

	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
		os.Exit(0)
	}
	if err != nil {
		fmt.Printf("error starting server: %s\n", err.Error())
		os.Exit(1)
	}
}
