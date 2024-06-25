# Greptile Integration Plugin

This plugin provides a WebAssembly-based integration with the Greptile API, allowing for repository indexing, querying, and searching operations. It's designed to be used with the Extism runtime.

## Features

- Repository indexing
- Repository querying
- Code searching
- Seamless integration with Greptile API
- WebAssembly-based for portability and security

## Prerequisites

- Go 1.16 or later
- TinyGo
- Extism CLI
- Greptile API key and GitHub token

## Installation

1. Clone this repository:

   ```
   git clone https://github.com/your-username/greptile-integration-plugin.git
   cd greptile-integration-plugin
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Build the WebAssembly module:
   ```
   make build
   ```

## Configuration

Create a `.env` file in the project root with your Greptile API key and GitHub token:

```
GREPTILE_API_KEY=your_greptile_api_key_here
GITHUB_TOKEN=your_github_token_here
```

## Usage

The plugin supports three main operations: index, query, and search. You can test each operation using the provided Makefile targets.

### Indexing a Repository

```
make test-index
```

This will index the specified repository in Greptile.

### Querying a Repository

```
make test-query
```

This will send a query to Greptile about the specified repository.

### Searching a Repository

```
make test-search
```

This will perform a code search in the specified repository.

## Customizing Operations

You can customize the input for each operation by modifying the JSON in the Makefile. Here's an example of the structure for each operation:

### Index

```json
{
  "operation": "index",
  "repository": "username/repo",
  "remote": "github",
  "branch": "main",
  "api_key": "your_api_key",
  "github_token": "your_github_token"
}
```

### Query

```json
{
  "operation": "query",
  "repository": "username/repo",
  "remote": "github",
  "branch": "main",
  "api_key": "your_api_key",
  "github_token": "your_github_token",
  "messages": [
    {
      "id": "1",
      "content": "What is this repository about?",
      "role": "user"
    }
  ],
  "session_id": "test-session",
  "stream": false,
  "genius": true
}
```

### Search

```json
{
  "operation": "search",
  "repository": "username/repo",
  "remote": "github",
  "branch": "main",
  "api_key": "your_api_key",
  "github_token": "your_github_token",
  "query": "Functions that use recursion",
  "session_id": "test-session",
  "stream": false
}
```

## Integration

To integrate this plugin into your own projects, you can use the Extism runtime to call the `run` function with the appropriate JSON input for the desired operation.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Public domain.
