Task API - Talent Growth

This is a simple Task Management REST API built with Go (Golang), Gin framework, and MongoDB. It allows users to perform CRUD operations on task resources. This project is part of the Talent Growth technical assignment.

Features

- Create new tasks
- Get all tasks or a specific task by ID
- Update task details
- Delete a task
- Organized using MVC structure (Models, Controllers, Routes)
- MongoDB integration for persistent storage


Tech Stack

- Language: Go (Golang)
- Framework: Gin Gonic
- Database: MongoDB
- Tools: Postman (for API testing), Git (for version control)


Folder Structure
.
├── controllers/ # Handles logic for each endpoint
├── models/ # Defines data structures (MongoDB models)
├── routes/ # Maps endpoints to controllers
├── main.go # Entry point of the app
├── go.mod / go.sum # Dependency tracking
└── README.md # Project documentation


Setup Instructions

1. Clone the Repository

```bash
git clone https://github.com/your-username/task-api.git
cd task-api
2. Install Dependencies
go mod tidy
3. Set Up MongoDB
Make sure MongoDB is running locally or use a cloud service (MongoDB Atlas). Set the URI in your main.go file:
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
4. Run the Project
go run main.go
Server will run on http://localhost:8080

API Endpoints
Method	Endpoint	Description
GET	/tasks	Get all tasks
GET	/tasks/:id	Get a task by ID
POST	/tasks	Create a new task
PUT	/tasks/:id	Update a task by ID
DELETE	/tasks/:id	Delete a task by ID

Sample JSON Body (POST / PUT)
json
{
  "title": "Finish project report",
  "description": "Complete and submit the weekly progress report",
  "category": "Work",
  "priority": "High"
}

🌐 API Documentation with Postman
Import the Postman collection.

Use raw > JSON format in body when creating/updating.

Update base_url variable with your running server URL.

Use scripts in "Tests" tab for checking response.