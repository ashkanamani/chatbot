/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ashkanamani/chatbot/cmd"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
