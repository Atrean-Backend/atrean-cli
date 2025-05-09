# Atrean CLI

A command-line tool for setting up and managing Atrean environments.

## Installation

Install the Atrean CLI with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/Atrean-Backend/atrean-cli/main/install.sh | bash
```

This will download and install the latest version of the CLI to your system.

## Usage

### Setting up your environment

To set up your environment with the required configuration files:

```bash
atrean setup
```

This will create:
- `docker-compose.yml` - Configuration for Docker containers
- `.env` - Environment variables for your deployment

### Starting the containers

After setup, you can start the containers in two ways:

Option 1: Using the setup start command:
```bash
atrean setup start
```

Option 2: Using the up command:
```bash
atrean up
```

Both commands will start your Atrean containers in detached mode.

## Commands

- `atrean setup` - Create Docker Compose and environment files
- `atrean setup start` - Start the Docker containers
- `atrean up` - Alternative command to start the containers

## Development

This CLI is built using [Cobra](https://github.com/spf13/cobra), a powerful library for creating modern CLI applications.

To build from source:

```bash
git clone https://github.com/Atrean-Backend/atrean-cli.git
cd atrean-cli
go build -o atrean
```

## License

[License details here] 