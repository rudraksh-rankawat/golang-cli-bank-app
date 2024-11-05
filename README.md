# CLI Banking App

This is a simple command-line banking application written in Go. It allows users to check their balance, deposit money, and withdraw money. This project is intended for learning and practicing Go programming.

## Features

- **Check Balance**: View the current balance of your account.
- **Deposit Money**: Add money to your account balance.
- **Withdraw Money**: Remove money from your account balance.
- **Persistent Storage**: The balance is saved to a file (`balance.txt`) to maintain state between sessions.

## Getting Started

### Prerequisites

- Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/cli_banking_app.git
   cd cli_banking_app
   ```

2. Run the application:

   ```bash
   go run bank.go
   ```

## Usage

1. Upon running the application, you will be greeted with a welcome message.
2. You will be presented with a menu of options:
   - **1**: Check Balance
   - **2**: Deposit Money
   - **3**: Withdraw Money
   - **4**: Exit
3. Enter the number corresponding to your choice and follow the prompts.

## File Structure

- `bank.go`: The main application file containing the logic for the CLI banking operations.
- `balance.txt`: A text file used to store the account balance persistently.

## Contributing

This project is for personal learning and practice. However, if you have suggestions or improvements, feel free to fork the repository and submit a pull request.

## License

This project is open-source and available under the MIT License.

## Acknowledgments

- Thanks to the Go community for providing excellent resources and documentation.
