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
	"gith