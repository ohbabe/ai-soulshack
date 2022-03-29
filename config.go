
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	ai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	vip "github.com/spf13/viper"
)

func init() {

	cobra.OnInitialize(initConfig)

	// irc client configuration
	root.PersistentFlags().StringP("nick", "n", "soulshack", "bot's nickname on the irc server")