package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	ai "github.com/sashabaranov/go-openai"
	vip "github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestChatSession(t *testing.T) {
	vip.Set("session", 1*time.Hour)
	vip.Set("history", 10)

	//log.SetOutput(io.Discard)

	ctx := &ChatContext{
		Personality: &Personality{
			Prompt: "You are a helpful 