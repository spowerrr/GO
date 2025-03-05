package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Exam struct {
	Dept        string `json:"Dept."`
	CourseCode  string `json:"Course Code"`
	CourseTitle string `json:"Course Title"`
	Section     string `json:"Section"`
	Teacher     string `json:"Teacher"`
	ExamDate    string `json:"Exam Date"`
	ExamTime    string `json:"Exam Time"`
	Room        string `json:"Room"`
}

// Function to extract ID ranges and check if the ID falls within the range
func idInRange(id string, room string) bool {
	re := regexp.MustCompile(`\(\d+-\d+\)`)
	matches := re.FindAllString(room, -1)

	for _, match := range matches {
		rangeParts := strings.Trim(match, "()")
		parts := strings.Split(rangeParts, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			userID, err3 := strconv.Atoi(id)

			if err1 == nil && err2 == nil && err3 == nil {
				if userID >= start && userID <= end {
					return true
				}
			}
		}
	}
	return false
}

type Course struct {
	Title   string `json:"title"`
	Section string `json:"section"`
}

type Request struct {
	UserID  string   `json:"userId"`
	Courses []Course `json:"courses"`
}

func main() {
	http.HandleFunc("/exam-schedule", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Open and parse JSON file
		file, err := ioutil.ReadFile("examSchedule.json")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Failed to read JSON file: %v", err)
			return
		}

		var data map[string][]Exam
		if err := json.Unmarshal(file, &data); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Failed to parse JSON: %v", err)
			return
		}

		if r.Method == "GET" {
			// For GET requests, return all exams
			var allExams []Exam
			for _, exams := range data {
				allExams = append(allExams, exams...)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(allExams)
			return
		}

		if r.Method == "POST" {
			// Parse request body
			var req Request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid request format", http.StatusBadRequest)
				return
			}

			// Validate User ID
			if !regexp.MustCompile(`^\d+$`).MatchString(req.UserID) {
				http.Error(w, "Invalid User ID format", http.StatusBadRequest)
				return
			}

			var results []Exam
			for _, course := range req.Courses {
				if course.Title == "" || course.Section == "" {
					http.Error(w, "Course title and section cannot be empty", http.StatusBadRequest)
					return
				}

				for _, exams := range data {
					for _, exam := range exams {
						if matchCourse(course.Title, exam.CourseTitle) && strings.EqualFold(course.Section, exam.Section) && idInRange(req.UserID, exam.Room) {
							results = append(results, exam)
						}
					}
				}
			}

			// Send response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(results)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Server starting on port 7070...")
	if err := http.ListenAndServe(":7070", nil); err != nil {
		log.Fatal(err)
	}
}

func matchCourse(input, examTitle string) bool {
	// Normalize strings
	normalizedInput := normalizeString(input)
	normalizedExam := normalizeString(examTitle)

	// Split into words
	inputWords := strings.Fields(normalizedInput)
	examWords := strings.Fields(normalizedExam)

	// Skip empty input
	if len(inputWords) == 0 {
		return false
	}

	// Check for exact match first
	if normalizedInput == normalizedExam {
		return true
	}

	// Count matches
	matchCount := 0
	totalWords := len(inputWords)

	// Create a map to track matched exam words
	matchedExamWords := make(map[string]bool)

	for _, inputWord := range inputWords {
		if len(inputWord) < 2 {
			totalWords--
			continue
		}

		for _, examWord := range examWords {
			// Skip if this exam word was already matched
			if matchedExamWords[examWord] {
				continue
			}

			if isRomanNumeral(inputWord) {
				if inputWord == examWord {
					matchCount++
					matchedExamWords[examWord] = true
					break
				}
			} else if examWord == inputWord || 
				(len(examWord) > 2 && strings.Contains(examWord, inputWord)) || 
				(len(inputWord) > 2 && strings.Contains(inputWord, examWord)) {
				matchCount++
				matchedExamWords[examWord] = true
				break
			}
		}
	}

	// Require a higher match threshold for better accuracy
	return totalWords > 0 && float64(matchCount)/float64(totalWords) >= 0.8
}

func normalizeString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' {
			return unicode.ToLower(r)
		}
		return ' '
	}, strings.TrimSpace(s))
}

func isRomanNumeral(s string) bool {
	return s == "i" || s == "ii" || s == "iii" || s == "iv" || s == "v"
}
