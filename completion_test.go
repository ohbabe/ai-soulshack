package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChunker_Chunk(t *testing.T) {
	timeout := 1000 * time.Millisecond

	tests := []struct {
		name     string
		input