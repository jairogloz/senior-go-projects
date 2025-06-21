package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	// Open the CSV file
	file, err := os.Open("./004_comment_picker/holabrackets_jrz_comment_list_20250223.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Map to store the number of unique tags for each user
	tagCount := make(map[string]map[string]struct{})

	// Process each record
	for i, record := range records {
		if i == 0 {
			continue
		}
		name := record[0]
		comment := record[2]

		// Extract tagged users from the comment
		tags := extractTags(comment)

		// Initialize the map for the user if not already done
		if _, exists := tagCount[name]; !exists {
			tagCount[name] = make(map[string]struct{})
		}

		// Add unique tags to the user's map
		for _, tag := range tags {
			tagCount[name][tag] = struct{}{}
		}
	}

	// Create a slice to store the sorted results
	type userTagCount struct {
		Name  string
		Count int
	}
	var sortedResults []userTagCount

	// Populate the slice with the tag counts
	for name, tags := range tagCount {
		sortedResults = append(sortedResults, userTagCount{Name: name, Count: len(tags)})
	}

	// Sort the results by the number of unique tags in descending order
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].Count > sortedResults[j].Count
	})

	for i, result := range sortedResults {
		fmt.Println(i+1, " ", result.Name, " ", result.Count)
	}

	// Define a slice of names to exclude from the raffle
	excludeNames := []string{"laurawho"}

	// Create a map for quick lookup of excluded names
	excludeMap := make(map[string]struct{})
	for _, name := range excludeNames {
		excludeMap[name] = struct{}{}
	}

	// Filter users with more than 3 separate tag counts and not in the exclusion list
	var eligibleUsers []userTagCount
	for _, result := range sortedResults {
		if result.Count > 3 {
			if _, excluded := excludeMap[result.Name]; !excluded {
				eligibleUsers = append(eligibleUsers, result)
			}
		}
	}

	// Create a weighted list of names
	var weightedList []string
	for _, user := range eligibleUsers {
		for i := 0; i < user.Count; i++ {
			weightedList = append(weightedList, user.Name)
		}
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select 5 names from the weighted list
	selectedNames := make(map[string]struct{})
	for len(selectedNames) < 3 && len(weightedList) > 0 {
		index := rand.Intn(len(weightedList))
		selectedNames[weightedList[index]] = struct{}{}
		// Remove the selected name from the weighted list to avoid duplicates
		weightedList = append(weightedList[:index], weightedList[index+1:]...)
	}

	// Simulate the raffle
	fmt.Print("Participants: ")
	for i := 0; i < 15; i++ {
		index := rand.Intn(len(weightedList))
		fmt.Printf("\r\033[KParticipants: %s", weightedList[index])
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Printf("\r\033[KGanadores:\n")

	// Print the selected names
	count := 1
	for name := range selectedNames {
		fmt.Println(count, " ", name, "count: ", len(tagCount[name]))
		count++
	}
}

// Function to extract tagged users from a comment
func extractTags(comment string) []string {
	words := strings.Fields(comment)
	var tags []string
	for _, word := range words {
		if strings.HasPrefix(word, "@") {
			tags = append(tags, word[1:])
		}
	}
	return tags
}
