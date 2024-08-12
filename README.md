## Project structure
### `cmd/app/main.go`
Configuration and logger initialization. Then the main function "continues" in
`internal/app/app.go`.

### `config`
Configuration. First, `config.yml` is read, then environment variables overwrite the YAML config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in YAML or in environment variables).

For configuration, we chose the [viper](https://github.com/spf13/viper) library.