# Employee Management Website

This is a simple employee management website built using Golang and MySQL. The application allows users to perform basic CRUD (Create, Read, Update, Delete) operations on employee data. It is designed to help manage employee information efficiently.

## Features

- Add new employees with relevant details.
- View a list of all employees.
- Update employee information.
- Delete employee records.

## Prerequisites

To run this project, ensure you have the following installed:

- [Golang](https://golang.org/) (version 1.16 or later)
- [MySQL](https://www.mysql.com/) (version 5.7 or later)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/MuhammadJundullah/Employee-data-management-website.git
    cd go-crud-employee
    ```

2. Set up the database:
    - Create a MySQL database named `crud_employee`.
    - Import the provided SQL script to create the necessary tables.

3. Run the application:
    ```bash
    go run main.go
    ```

4. Access the website at `http://localhost:8080`.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.