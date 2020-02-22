# PEMS-API
People Management System is a platform for businesses to store and manage user profiles.

This is the API for the system.

## Features
- in development

## API Endpoints
<table>
<tr><th>HTTP VERB</th><th>ENDPOINTS</th><th>DESCRIPTION</th><th>BODY PARAMS</th></tr>
<tr><td>GET</td><td>/api/v1/</td><td>Returns the welcome message</td><td></td></tr>

</table>

# Getting Started
**Via Cloning The Repository**
```
# Clone the app
git clone https://github.com/stylll/pems-api.git

# Switch to directory
cd pems-api

# Create .env file in the root directory
touch .env
update env file with required information
use the .sample.env file as a guideline

# Create application docker image
make build

# Start the application
make run

# Start the application without Docker (Development)
go mod download
go run ./src/main.go

# View the application
navigate to localhost:{PORT} to view the application
```


## Copyright
&copy; Stephen Aribaba

