# Zaxvat (Захват)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Zaxvat (Захват) is a subdomain takeover tool designed to check whether CNAME of the domain/subdomain can be used for takeover. As it checks for CNAMEs only, you will have to manually check domains to understand whether they are vulnerable or not. Check log includes documentation and discussion which will surely haelp you for the takeover. 

## Table of Contents
- [Features](https://github.com/grozdniyandy/zaxvat#features)
- [Usage](https://github.com/grozdniyandy/zaxvat#usage)
- [Installation](https://github.com/grozdniyandy/zaxvat#installation)
- [Dependencies](https://github.com/grozdniyandy/zaxvat#dependencies)
- [License](https://github.com/grozdniyandy/zaxvat#license)
- [Author](https://github.com/grozdniyandy/zaxvat#author)
- [Contributing](https://github.com/grozdniyandy/zaxvat#contributing)

## Features
- CNAME live update from https://github.com/EdOverflow/can-i-take-over-xyz

## Usage
1. **Clone or Download:** Clone this repository or download the code to your local machine.
2. **Install Golang:** Install the latest Golang from https://go.dev/dl/. For example:
    ```
    wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    go version
    ```
3. **Run the tool:** Run the tool using the following command:
   ```
   go run main.go example.com
   ```
   ![image](https://github.com/grozdniyandy/zaxvat/assets/143331225/7ee4abad-a71d-40b8-9d5a-bc874f01667d)

## Installation
You can either check the "Usage" or download already compiled code from "releases".

## Dependencies
This code uses the Go standard library, so there are no external dependencies to install.

## License
This code is released under the [MIT License](LICENSE).

## Author
Zaxvat is developed by GrozdniyAndy of [XSS.is](https://xss.is).

## Contributing
Feel free to contribute, report issues, or suggest improvements by creating pull requests or issues in the GitHub repository. Enjoy захватывать xD!
