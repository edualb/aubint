package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type Dog struct {
	ID           int64        `json:"id"`
	WalkMetadata WalkMetadata `json:"walk_metadata"`
	Name         string       `json:"name"`
	Birthday     time.Time    `json:"birthday"`
	Deathday     time.Time    `json:"deathday"`
}

type WalkMetadata struct {
	WalksPerMonth  int64 `json:"walks_per_month"`
	MinutesPerWalk int64 `json:"minutes_per_walk"`
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	dogs := []Dog{}
	for i := 0; i < 1000; i++ {
		dog := Dog{
			ID:       int64(i + 1),
			Name:     generateRandomPetName(r),
			Birthday: generateRandomDate(r, time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)),
			Deathday: generateRandomDate(r, time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC)),
			WalkMetadata: WalkMetadata{
				WalksPerMonth:  int64(generateRandomNumberBetween(r, 8, 64)),
				MinutesPerWalk: int64(generateRandomNumberBetween(r, 10, 40)),
			},
		}
		dogs = append(dogs, dog)
	}

	output := struct {
		Dogs []Dog `json:"dogs"`
	}{
		Dogs: dogs,
	}
	data, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/dogs.json", data, 0o444)
	if err != nil {
		panic(err)
	}
}

func generateRandomPetName(r *rand.Rand) string {
	prefixes := []string{"Fluffy", "Whisker", "Fuzzy", "Spike", "Cuddly", "Paws", "Mittens", "Buddy", "Mocha", "Sunny"}
	suffixes := []string{"kins", "paws", "tail", "face", "buddy", "whiskers", "nose", "paws", "fur", "spot"}

	prefixIndex := r.Intn(len(prefixes))
	suffixIndex := r.Intn(len(suffixes))

	name := prefixes[prefixIndex] + suffixes[suffixIndex]
	return name
}

func generateRandomDate(r *rand.Rand, start, end time.Time) time.Time {
	diff := end.Sub(start)
	randomDiff := time.Duration(r.Int63n(int64(diff)))
	randomDate := start.Add(randomDiff)
	return randomDate
}

func generateRandomNumberBetween(r *rand.Rand, min, max int) int {
	return min + r.Intn(max-min+1)
}
