package main

type GptResponse struct {
	Index        int            `json:"index"`
	Message      MessageContent `json:"message"`
	LogProbs     interface{}    `json:"logprobs"`
	FinishReason string         `json:"finish_reason"`
}

type MessageContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
