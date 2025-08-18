## Step 1: REST Basics - TODO APP

**Goal:**
Understand the fundamentals of RESTful API design, HTTP methods, status codes, and basic web development.

**Tasks:**

-   Implement the following endpoints:
    -   Create a todo item ( `POST /todos` ) 
    -   Toggle/untoggle a todo item ( `PUT /todos/{id}` )
    -   Delete a todo item ( `DELETE /todos` )
    -   List all todo items ( `GET /todos` ) 
-   Use an in-memory or simple local database (e.g., SQLite or MySql)
-   Use GORM as ORM library

**Bonus:**

-   Containerize the app with Docker    
-   Run the application using a `docker-compose.yml` file