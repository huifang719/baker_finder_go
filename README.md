# Baker Finder Backend Service
## Description
Baker Finder is a backend service built using Go, Chi for routing, PostgreSQL as the database, and leveraging SQLC and Goose for efficient database management. The service allows users to create accounts, find bakers by postcode, create new baker profiles, and review/rate existing bakers.

## Features 
- User Management: Users can create accounts to access the service.
- Baker Creation: Users can create profiles for bakers including their information.
- Review and Rating: Users can review and rate bakers they have interacted with.
- Baker Search: Users can find bakers based on their location's postcode.

## :white_check_mark: Setup 
### Prerequisites
Before running the project, ensure you have the following installed:
vscode 
Go (1.15 or later)
PostgreSQL
Goose (goose command-line tool for database management)
SQLC (sqlc command-line tool for generating Go code from SQL)


## :rocket: Usage
- User Registration: Endpoint to register a new user.
- Baker Creation: Endpoint to create a new baker profile.
- Review and Rating: Endpoint to review and rate bakers.
- Baker Search: Endpoint to search for bakers by postcode.

## Contributing
Contributions are welcome! If you find any issues or want to add new features, please open an issue or submit a pull request.
