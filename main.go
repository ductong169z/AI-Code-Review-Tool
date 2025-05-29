package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	genai "github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	diff, err := getGitDiff()
	if err != nil {
		fmt.Println("❌ Error getting git diff:", err)
		os.Exit(1)
	}

	if strings.TrimSpace(diff) == "" {
		fmt.Println("✅ No changes to review.")
		os.Exit(0)
	}
	fmt.Println("🔍 Reviewing staged changes...")
	review, err := reviewWithGemini(diff)
	if err != nil {
		fmt.Println("❌ Error calling Gemini:", err)
		os.Exit(1)
	}

	if isCritical(review) {
		fmt.Println("🚫 Critical issues found — please fix before pushing.")
		os.Exit(1)
	} else {
		fmt.Println("✅ No critical issues found. Please view review-<timestamp>.md for details.")
	}
	os.Exit(0)
}

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func reviewWithGemini(diff string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using OS env variables")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ GEMINI_API_KEY is not set")
		os.Exit(1)
	}

	promptBytes, err := os.ReadFile("prompt.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %w", err)
	}
	prompt := string(promptBytes)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt+"\n\n"+diff))
	if err != nil {
		return "", err
	}

	var output string
	for _, part := range resp.Candidates[0].Content.Parts {
		output += fmt.Sprint(part)
	}

	output = strings.TrimSpace(output)
	timeStamp := time.Now().Format("2006-01-02_15-04-05") // Use safe filename format
	filename := fmt.Sprintf("review_%s.md", timeStamp)

	err = os.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write %s: %w", filename, err)
	}

	return output, nil
}

func isCritical(review string) bool {
	review = strings.ToLower(review)
	return strings.Contains(review, "should") || strings.Contains(review, "must") || strings.Contains(review, "missing") || strings.Contains(review, "bug") || strings.Contains(review, "incorrect")
}
