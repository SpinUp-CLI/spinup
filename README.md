![SpinUp Banner](assets/SpinUpBanner.svg)

# 🌀 SpinUp CLI

> Build web-apps, with the frameworks of your choice, with a single command. 100% customisable.

SpinUp is a powerful CLI tool that enables developers to quickly scaffold full-stack web applications using customizable templates. It supports various frontend frameworks and backend technologies, all configurable through a simple YAML configuration file.

## ✨ Features

- **🚀 Rapid Project Setup**: Create full-stack projects with a single command
- **🔧 Fully Customizable**: Configure frontend, and backend services
- **📦 Template System**: Support for remote Git repositories as templates
- **🔐 Private Repository Support**: Authentication for private template repositories
- **🌍 Environment Variables**: Dynamic configuration using environment files
- **🎨 Beautiful CLI Output**: Colored and formatted terminal output
- **⚡ Concurrent Deployment**: Parallel service creation for faster setup

## 📦 Installation

### Prerequisites
- Go 1.25.2 or later
- Git

### Build from Source
```bash
git clone https://github.com/SpinUp-CLI/spinup
cd spinup
make all
```

This will create a `spin` executable in the project directory.

## 🚀 Quick Start

### 1. Initialize Configuration
```bash
./spin init
```

This creates a default configuration file at `~/.config/spinup/defaults.yaml`.

### 2. Create a Project
```bash
# Using default project name
./spin create

# Using custom project name
./spin create my-awesome-app
```

### 3. Reset Configuration
```bash
./spin reset
```

## ⚙️ Configuration

SpinUp uses a YAML configuration file located at `~/.config/spinup/defaults.yaml`.

### Default Configuration Structure

```yaml
env_file: none
templates:
  remote_templates:
    - name: none
      url: https://github.com/SpinUp-CLI/
      secret: none
defaults:
  frontend: vue-template
  backend: go
  project_name: web-app
```

### Configuration Options

#### Environment File
```yaml
env_file: path/to/.env  # Path to environment variables file
```

#### Template Remotes
```yaml
templates:
  remote_templates:
    - name: github
      url: https://github.com/your-org/
      secret: your_github_token
    - name: gitlab
      url: https://gitlab.com/your-org/
      secret: your_gitlab_token
```

#### Default Stacks
```yaml
defaults:
  frontend: vue-template     # Frontend framework template
  backend: go               # Backend framework template
  project_name: web-app     # Default project name
```

### Environment Variables

You can use environment variables in your configuration:

```yaml
templates:
  remote_templates:
    - name: github
      url: https://github.com/my-org/
      secret: ${GITHUB_TOKEN:default_value}
```

## 🏗️ Project Structure

When you create a project, SpinUp generates the following structure:

```
my-project/
├── my-project-front/    # Frontend service
└── my-project-back/     # Backend service
```

Each service is cloned from the corresponding template repository based on your configuration.

## 🛠️ Available Commands

### `spin init`
Initialize SpinUp configuration. Creates a default configuration file if it doesn't exist.

**Usage:**
```bash
./spin init
```

### `spin create [project-name]`
Create a new full-stack project.

**Usage:**
```bash
./spin create                    # Uses default project name
./spin create my-awesome-app     # Uses custom project name
```

**What it does:**
1. Checks/creates configuration
2. Sets up project directory
3. Clones frontend template (if required)
4. Clones backend template (if required)
5. Configures services concurrently

### `spin reset`
Reset the configuration to defaults. Recreates the configuration file.

**Usage:**
```bash
./spin reset
```

## 🎯 Template System

SpinUp uses Git repositories as templates. Templates are organized by framework name and fetched from configured remote URLs.

### Template URL Structure
```
{remote_url}{framework_name}
```

Example:
- Remote URL: `https://github.com/SpinUp-CLI/`
- Framework: `vue-template`
- Final URL: `https://github.com/SpinUp-CLI/vue-template`

### Supported Authentication
- **Public repositories**: No authentication required
- **Private repositories**: Basic authentication using tokens

## 🔧 Development

### Project Architecture

```
spinup/
├── cmd/                    # Cobra CLI commands
│   ├── root.go            # Root command and help
│   ├── init.go            # Init command
│   ├── create.go          # Create command
│   └── reset.go           # Reset command
├── internal/
│   ├── commands/          # Command implementations
│   │   ├── init_cmd.go
│   │   ├── create_cmd.go
│   │   ├── reset_cmd.go
│   │   └── utils/         # Command utilities
│   ├── config/            # Configuration management
│   │   ├── config.go      # Config structures
│   │   ├── yaml.go        # YAML parsing
│   │   └── env_manager.go # Environment variable handling
│   └── project/           # Project creation logic
│       ├── project.go     # Project structures
│       ├── create_project.go # Project creation
│       └── cloner.go      # Git repository cloning
├── pkg/                   # Reusable packages
│   ├── fileutils/         # File system utilities
│   ├── iostream/          # Formatted output
│   └── utils/             # ANSI colors and formatting
└── main.go               # Application entry point
```

### Key Components

#### Configuration System ([internal/config](internal/config))
- **[`config.Config`](internal/config/config.go)**: Main configuration structure
- **[`config.ParseConfig`](internal/config/yaml.go)**: YAML parsing with environment variable support
- **[`config.GetConfigPath`](internal/config/config.go)**: Configuration file location management

#### Project Management ([internal/project](internal/project))
- **[`project.Project`](internal/project/project.go)**: Project representation
- **[`project.CreateProject`](internal/project/create_project.go)**: Project creation orchestration
- **[`project.TryRemote`](internal/project/cloner.go)**: Git repository cloning with authentication

#### Output System ([pkg/iostream](pkg/iostream))
- **[`iostream.Error`](pkg/iostream/output.go)**: Error messages
- **[`iostream.Success`](pkg/iostream/output.go)**: Success messages
- **[`iostream.Info`](pkg/iostream/output.go)**: Information messages
- **[`iostream.Log`](pkg/iostream/output.go)**: General logging

### Building and Testing

```bash
# Build the project
make all

# Run the application
make run

# Format code
make fmt

# Clean build artifacts
make clean
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Format code: `make fmt`
6. Submit a pull request

## 📝 License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## 👤 Author

**Elouann Hosta**
- Email: ehosta@student.42lyon.fr

---

*Built with ❤️ using Go and Cobra CLI*