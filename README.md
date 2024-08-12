# Twitter Archive Extractor

This Go program extracts and processes JavaScript files from a ZIP archive, specifically targeting files in the /data directory. It replaces certain window. assignments with var data = and outputs the JSON representation of the processed data.

## How It Works

1. The program opens a ZIP file specified as a command-line argument.
2. It scans for JavaScript files within the /data directory.
3. For each JavaScript file, it replaces any window. assignments (e.g., window.__THAR_CONFIG = {) with var data =.
4. It then uses the goja JavaScript interpreter to execute the modified script and extract the data variable.
5. The extracted data is marshaled into JSON format and output to the console.

## Todo

- Select a destination for export
  - separate files? via an ORM to a database?
- Release to go.dev

## Prerequisites

Ensure you have Go installed on your system.

## Setup

1. Clone this repository or download the source files.
2. Navigate to the project directory.
3. Run the following command to install dependencies:

```bash
go mod tidy 
```

## Usage

1. Place your target ZIP file in the project directory. A sample `twitter.zip` file is included for testing purposes.
2. Run the program with the following command:

```bash
go run main.go /path/to/your/zip.zip 
```

Replace `/path/to/your/zip.zip` with the actual path to your own ZIP file.

### Example

To run the program with the provided sample twitter.zip:

```bash
go run main.go twitter.zip 
```

The program will output the JSON representation of the data extracted from the JavaScript files inside the /data directory of the ZIP file.

## Dependencies

The dependencies for this project are managed with Go modules. All necessary dependencies are listed in the go.mod and go.sum files.

Key dependencies include:

- github.com/dop251/goja: A JavaScript interpreter written in Go.

## Notes

- The program processes all .js files found in the /data directory of the ZIP archive.
- The replacement of window. assignments is based on regular expressions, so the code may need adjustments if you encounter different patterns.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.