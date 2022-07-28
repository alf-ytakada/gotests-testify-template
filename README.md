# gotests-testify-template
custom templates for [gotests](https://github.com/cweill/gotests)

- using [stretchr/testify](https://github.com/stretchr/testify)
- checks error type

# Usage in Visual Studio Code

1. Add [vscode-go] extension
1. Clone this repository
1. Add config to setting.json
  ```
  "go.generateTestsFlags": [
    "-template_dir",
    "/path/to/gotests-testify-template/templates"
  ]
  ```

# Sample test code

See [samples/sample_test.go](samples/sample_test.go)
