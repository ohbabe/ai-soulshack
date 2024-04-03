

// ____ _ ____ _ _
// / ___| ___ _ _ | | / ___| | __ _ ___ | | __
// \___ \ / _ \ | | | | | | \___ \ | '_ \ / _` | / __| | |/ /
// ___) | | (_) | | |_| | | | ___) | | | | | | (_| | | (__ | <
// |____/ \___/ \__,_| |_| |____/ |_| |_| \__,_| \___| |_||_\
// . . . because real people are overrated

AI-SoulShack is an AI-powered IRC chat bot that utilizes the OpenAI API to generate human-like responses.

## Features

- Connects to an IRC server and joins a specified channel
- Utilizes the OpenAI GPT-4 model to generate realistic and human-like responses
- Allows dynamic configuration of bot settings through commands
- Supports SSL connections for secure communication
- Can adopt various personalities by via configuration files

## Building

```bash
go build .
```

```
docker build . -t soulshack:dev
```

## Usage

```bash
soulshack --nick chatbot --server irc.freenode.net --port 6697 --channel '#soulshack' --ssl --openaikey ****************
```

## Configuration

SoulShack can be configured using command line flags, environment variables, or personality configuration files. It uses viper to manage configuration settings.

.........< Further Content in markdown>...........