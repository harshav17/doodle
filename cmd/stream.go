package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(streamCommand)
}

var streamCommand = &cobra.Command{
	Use:   "stream",
	Short: "play with the openAI stream API",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		stream()
	},
}

type RequestBody struct {
	Model            string      `json:"model"`
	Messages         []Message   `json:"messages"`
	Functions        []Function  `json:"functions,omitempty"`
	FunctionCall     interface{} `json:"function_call,omitempty"`
	Temperature      *float64    `json:"temperature,omitempty"`
	TopP             *float64    `json:"top_p,omitempty"`
	N                *int        `json:"n,omitempty"`
	Stream           *bool       `json:"stream,omitempty"`
	Stop             interface{} `json:"stop,omitempty"`
	MaxTokens        *int        `json:"max_tokens,omitempty"`
	PresencePenalty  *float64    `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64    `json:"frequency_penalty,omitempty"`
	LogitBias        interface{} `json:"logit_bias,omitempty"`
	User             *string     `json:"user,omitempty"`
}

type Message struct {
	Role         string        `json:"role"`
	Content      string        `json:"content"`
	Name         string        `json:"name,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

type Function struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ResponseData struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	Delta        Delta   `json:"delta"`
	FinishReason *string `json:"finish_reason"`
}

type Delta struct {
	Content string `json:"content"`
}

func stream() {
	// load OPENAI_API_KEY from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	client := &http.Client{}
	doStream := true

	reqBody := &RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: "Hello, I'm an user",
			},
		},
		Stream: &doStream,
		// ... set other fields as needed
	}

	reqBodyJson, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBodyJson))
	// Set any required headers, e.g., for authentication
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Exit the loop on error
		}
		if strings.TrimSpace(line) == "data: [DONE]" {
			break // Exit the loop when the DONE message is received
		}
		jsonData := strings.TrimPrefix(line, "data: ")
		jsonData = strings.TrimSpace(jsonData)
		if jsonData == "" {
			continue
		}

		var data ResponseData
		err = json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err, jsonData)
			continue
		}
		// Now 'data' contains the parsed data
		fmt.Print(data.Choices[0].Delta.Content)
	}
}
