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
## Database Setup

### Install and Configure PostgreSQL

1. Ensure PostgreSQL is installed and running.

2. Create a database for the application.

### Update Database Configuration

1. Open the backend `.env` file.

2. Update the database connection details.

## Redis Setup

### Install and Start Redis

1. Ensure Redis is installed and running.

### Update Redis Configuration

1. Open the backend `.env` file.

2. Update the Redis connection details.

## OAuth Configuration

### Set Up OAuth Credentials

1. Obtain OAuth credentials from your preferred provider (e.g., Google, Facebook).

2. Configure OAuth in the backend `.env` file.

## Usage

1. **Access the Application**

   - Open your web browser and go to [http://localhost:3000](http://localhost:3000).

2. **OAuth Authentication**

   - Log in using OAuth authentication.

3. **Perform CRUD Operations**

   - Use the application to perform Create, Read, Update, and Delete operations on employee records.
