# delete-branches
`delete-branches` is a simple command-line tool written in Go that lists all local Git branches and deletes them except for a list of ignored branches provided as arguments.

## Features
Lists all local Git branches.
- Deletes local branches, excluding the branches specified as arguments.
- Utilizes git branch -D for forceful deletion.

## Prerequisites
Git: Ensure you have Git installed and accessible in your command line.
Go: To build the tool, you need Go installed on your system.

## Installation
1. Clone the Repository
```bash
git clone https://github.com/yourusername/delete-branches.git
cd delete-branches
```
2. Build the Executable
Run the following command to compile the Go program into a binary executable:

```bash
go build -o delete-branches delete_branches.go
```
3. Add to PATH (Optional)
To make delete-branches accessible from anywhere in your terminal, move the binary to a directory in your system's PATH.

```bash
sudo mv delete-branches /usr/local/bin/
```
Verify it works by running:

```bash
delete-branches --help
```

## Usage
```bash
delete-branches [branch1] [branch2] [...]
branch1, branch2, etc.: Branch names to be ignored during deletion.
```
## Example
```bash
delete-branches main develop feature-xyz
```
This command will:
- Skip deleting main, develop, and feature-xyz.
- Delete all other local branches.
## Error Handling
If the branch deletion fails (e.g., due to unmerged changes), the tool will output the error for each branch it cannot delete.

## Development
To modify the script, simply edit the delete_branches.go file, then recompile it using the go build command:

```bash
go build -o delete-branches delete_branches.go
```
## License
This project is licensed under the MIT License. See the LICENSE file for details.





