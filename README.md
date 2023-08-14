# TODO-htmx-go

A proof-of-concept TODO app to explore the cababilities of [HTMX](https://htmx.org/). The backend server is written in [Go](https://go.dev/).

### Built With

* [![Htmx][htmx-shield]][htmx-url]
* [![Golang][golang-shield]][golang-url]

## Getting Started

### Prerequisites

This application only requires the Golang compiler, tooling, and libraries. These can be downloaded from the [Golang homepage](https://go.dev/) or installed via a package manager:
* apt
  ```sh
  > apt install golang-go
  ```
* Chocolatey
  ```pwsh
  > choco install -y golang
  ```
* Homebrew
  ```sh
  > brew install go
  ```
* winget
  ```pwsh
  > winget install GoLang.Go
  ```

### Installation & Execution

There is currently no additional installation necessary. Simply clone the repo and start the server.

1. Clone the repo
   ```sh
   > git clone git@github.com:Drakmyth/todo-htmx-go.git
   ```
1. Start the server
   ```sh
   > go run server.go
   ```

## Contributing

This is a simple proof of concept app, but if you find it useful and would like to contribute you are more than welcome to do so! If you have a suggestion, you can fork the repo and create a pull request. Alternatively you can open an issue with the `enhancement` tag.

Keep in mind that as this is a proof of concept app it is not under regular active development. Taking the initiative to provide a pull request will greatly increase the chances of a feature being added.

## License

Distributed under the MIT License. See [LICENSE.md](./LICENSE.md) for more information.


<!-- Reference Links -->
[htmx-url]: https://htmx.org
[htmx-shield]: https://img.shields.io/badge/htmx-4470d2?style=for-the-badge
[golang-url]: https://go.dev
[golang-shield]: https://img.shields.io/badge/golang-09657c?style=for-the-badge&logo=go&logoColor=79d2fa
