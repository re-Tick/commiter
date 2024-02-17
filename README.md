# Commiter CLI
This CLI tool is designed to automate the process of committing and pushing changes to a Git repository on a daily basis. It creates a new branch for each day, commits any changes in the working directory, and pushes the branch to the remote repository.

## Installation
Clone this repository to your local machine:
```sh
git clone https://github.com/re-Tick/commiter.git
cd commiter
go build -o commiter . && sudo mv ./commiter /usr/local/bin/
```
Make the script executable:

## Usage
To use the CLI, simply run the script:
```sh
commiter
```
