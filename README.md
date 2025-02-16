# 🚀🚀 Go Fiber + PostgreSQL API
- ✅ Fiber for fast and lightweight web framework
- ✅ PostgreSQL as the database
- ✅ Basic CRUD operations Create, Read

# 📂 Folder Structure

```sh 
db
|-db.go          # Handles DB connection
|-migrations.go  # Handles database migrations
models
|-models.go      # Defines the User model
routes
|-user.go        # Handles API routes for Users
go.mod           # Go module file
go.sum
main.go          # Entry point of the application
README.md        # Project documentation
```


# 🚀 Installation & Setup
### 1️⃣ Clone the repositry
``` sh
git clone https://github.com/vviveksharma/Gofiber-Server.git
cd Gofiber-Server
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Set Up PostgreSQL
Ensure you have PostgreSQL running and update the db.InitDB() function in db/db.go with your PostgreSQL credentials running on your local or your server: 

```sh
dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"

```
### 4️⃣ Run the Server

```sh
go run main.go
```

# 🛠 API Endpoints
## 👤 User APIs
| Method  | Endpoint | Description      | Request Body                                                  |
|---------|----------|------------------|---------------------------------------------------------------|         
| Get     | /users   | Get all the users| None                                                          |
| Post    | /user    | Add a new user   | ``` {"name": "vivek","email": "sharmavivek1709@gmail.com"}``` |




