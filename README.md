# Log File Analyzer

## Overview

The Log File Analyzer is a Go program designed to extract unique IP addresses from a log file based on a specified URL endpoint. It is optimized to handle large log files efficiently by processing them line by line rather than loading the entire file into memory.

## Features

- **Efficient File Handling:** Reads the log file line by line to handle large files without excessive memory usage.
- **IP Address Extraction:** Uses regular expressions to identify and extract IP addresses from log entries that match a specified URL endpoint.
- **Error Handling:** Logs and skips malformed lines to ensure that the program does not crash unexpectedly.
- **Unique IP Collection:** Stores IP addresses in a map to ensure uniqueness, avoiding duplicates.

## Log File Format

The program expects the log file to be in the following format:

[Date] [Time] [IP Address] [HTTP Method] [URL] [HTTP Status Code]

## Usage

1. **Prepare Your Log File:**
   - Ensure your log file follows the expected format and contains the relevant log entries.

2. **Run the Program:**
   - Use the command line to execute the program, specifying the path to your log file and the URL endpoint you want to search for.
   - Example command:
     ```sh
     go run main.go -file path/to/logfile.log -url /api/v1/users
     ```
   - Replace `path/to/logfile.log` with the actual path to your log file and `/api/v1/users` with the URL endpoint you want to filter by.

3. **View Results:**
   - The program will output a list of unique IP addresses that accessed the specified URL endpoint.

## Error Handling

- **Malformed Lines:** The program skips and logs lines that do not match the expected format.
- **File Access Issues:** Provides meaningful error messages for issues such as file not found or insufficient permissions.

## Conclusion

This Log File Analyzer is designed to be a robust tool for extracting and analyzing IP address data from large log files. Its efficient handling and error management make it suitable for various log analysis tasks.


