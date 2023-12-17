# Employee Management System (CRUD) Application

## Overview

This Employee Management System is a full-stack application designed to manage employee records efficiently. The system provides CRUD (Create, Read, Update, Delete) functionality, utilizes React for the frontend, Go-Gin for the backend, Redis for caching, and PostgreSQL with GORM for data storage. OAuth authentication is implemented to secure user logins.

## Features

- **CRUD Operations**: Perform Create, Read, Update, and Delete operations on employee records.
- **React Frontend**: User-friendly interface built with React for seamless interaction.
- **Go-Gin Backend**: Powerful and efficient backend built using Go-Gin to handle API requests.
- **Redis Cache**: Employ Redis for caching to enhance data storage and retrieval speed.
- **OAuth Authentication**: Secure user authentication using OAuth for login functionality.
- **PostgreSQL with GORM**: Utilize PostgreSQL as the relational database, managed by GORM for seamless data handling.

## Setup

### Prerequisites

Before you begin, make sure you have the following installed:

- Node.js and npm for React development.
- Go for Go-Gin backend development.
- Redis for caching.
- PostgreSQL and GORM package for Go.
- OAuth credentials for authentication.

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/employee-management-system.git


2. Frontend Setup

#### cd your_project/frontend
#### npm install
#### npm start

3. Backend Setup

#### go run main.go

4. Database Setup

    Set up PostgreSQL and create a database for the application.
    Update database configuration in the backend .env file.


5. Redis Setup
Ensure Redis is installed and running.
Update Redis configuration in the backend .env file.


6. OAuth Configuration
Set up OAuth credentials (e.g., from Google, Facebook, etc.).
Update OAuth configuration in the backend .env file.


7. Usage
Access the application by visiting http://localhost:3000 in your web browser.
Users can log in using OAuth authentication.
Perform CRUD operations on employee records.
