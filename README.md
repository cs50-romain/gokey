# GoPass - Simple Password Manager CLI Tool

GoPass is a straightforward command-line interface (CLI) tool written in Golang for managing passwords. It allows users to easily generate, store, and retrieve passwords securely.

## Installation

Make sure you have Golang installed on your machine. You can then install GoPass using the following command:

```bash
go get github.com/your-username/gopass
```

This will download and install GoPass in your `$GOPATH/bin` directory.

## Usage

### 1. Creating a new password

To generate a new password, use the following command:

```bash
gopass new
```

This will generate a random password and prompt you to name it.

To specify a custom name for the password, use:

```bash
gopass new <name>
```

### 2. Showing a password

To view a stored password, use the following command:

```bash
gopass show <name>
```

Replace `<name>` with the name of the password you want to retrieve. The password will be displayed on the terminal.

### 3. Copying a password to the clipboard

To copy a stored password to the clipboard, use the following command:

```bash
gopass copy <name>
```

Replace `<name>` with the name of the password you want to copy. The password will be copied to the clipboard, allowing you to paste it wherever needed.

## Security

GoPass takes security seriously. Passwords are stored using industry-standard encryption practices. Make sure to keep your master password secure and do not share it with others.

## Contributions

Contributions to GoPass are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
