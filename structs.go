package main

type resError struct {
	Type        int    `json:"type"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type meta struct {
	Records int    `json:"records"`
	Type    string `json:"type"`
	Time    string `json:"time"`
}

type ocResponse struct {
	Body  interface{} `json:"body"`
	Error resError    `json:"error"`
	Meta  meta        `json:"meta"`
}
