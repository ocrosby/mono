# Mono

A general purpose utility for building Go with multi-module workspaces.

[![CI](https://github.com/ocrosby/mono/actions/workflows/ci.yml/badge.svg)](https://github.com/ocrosby/mono/actions/workflows/ci.yml)

## Installation

To install Mono, you can use the following command:

```shell
go install github.com/ocrosby/mono@latest
```

Ensure that you have Go installed and configured on your machine.

## Usage

To use Mono, run the following command in your terminal:

```shell
mono <command> [options]
```

Replace <command> with the specific command you want to run.  For a list of available commands and options, use:

```shell
mono help
```

## Features

* **Multi-Module Workspaces**: Easily manage and build Go projects with multiple modules.
* **CI Integration**: Seamless itnegration with continuous integration systems.
* **Linting and Formatting**: Built-in support for golangci-lint to ensure code quality.
* **Configuration Management**: Leverage Cobra and Viper for command-line interface and configuraiton management.

## Contributing

We welcome contributions from the community.  To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes and commit them (git commit -am 'Add new feature').
4. Push the branch (git push origin feature-branch).
5. Create a new Pull Request.

Please ensure all contributions adhere to the project's coding standards and pass all tests.

## License

This project is licensed under the MIT License.  See the [LICENSE](LICENSE) file for details.

## Dependencies

* [Go](https://golang.org/) - The Go programming language.
* [golangci-lint](https://github.com/golangci/golangci-lint) - Fast linters runner for Go.
* [Cobra](https://github.com/spf13/cobra) - A library for creating powerful moderl CLI applications.
* [Viper](https://github.com/spf13/viper) - A complete configuration solution for Go applications.

## FAQs

**Q: How do I get started with Mono?**

A: Follow the installation instructions and run `mono help` to see available commands and options.

**Q: How can I contribute to the project?**

A: Follow the steps outlined in the [Contributing](docs/CONTRIBUTING.md) section.

## Acknowledgements

We would like to thank the following projects and contributors for their support and inspiration:

- The Go community for their continuous support and contributions.
- The maintainers of `golangci-lint`, `Cobra`, and `Viper` for their invaluable tools.

## References

* [How to Use The Cobra Package in Go](https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go)
