Employee Management System (CRUD) Application

Overview

This is an Employee Management System built using React for the frontend, Go-Gin for the backend, Redis for caching, and PostgreSQL with GORM for data storage. The application provides CRUD (Create, Read, Update, Delete) functionality for managing employee records. Additionally, OAuth authentication is implemented to secure user logins.

Features

CRUD Operations: Perform Create, Read, Update, and Delete operations on employee records.
React Frontend: User-friendly interface built with React for seamless interaction.
Go-Gin Backend: Powerful and efficient backend built using Go-Gin to handle API requests.
Redis Cache: Employ Redis for caching to enhance data storage and retrieval speed.
OAuth Authentication: Secure user authentication using OAuth for login functionality.
PostgreSQL with GORM: Utilize PostgreSQL as the relational database, managed by GORM for seamless data handling.
Setup

Prerequisites

Node.js and npm installed for React development.
Go installed for Go-Gin backend development.
Redis installed and running for caching.
PostgreSQL installed and configured.
GORM package installed for Go to interact with PostgreSQL.


Installation
Clone the Repository
bash
Copy code
git clone https://github.com/yourusername/employee-management-system.git
Frontend Setup
bash
Copy code
cd employee-management-system/frontend
npm install
npm start



Backend Setup
bash
Copy code
cd employee-management-system/backend
go get
go run main.go


Database Setup
Set up PostgreSQL and create a database for the application.
Update database configuration in the backend .env file.


Redis Setup
Ensure Redis is installed and running.
Update Redis configuration in the backend .env file.


OAuth Configuration
Set up OAuth credentials (e.g., from Google, Facebook, etc.).
Update OAuth configuration in the backend .env file.


Usage
Access the application by visiting http://localhost:3000 in your web browser.
Users can log in using OAuth authentication.
Perform CRUD operations on employee records.