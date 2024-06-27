# Gossip Forum (Backend)

Backend web service for Gossip Forum, a web forum application.

### Project Structure: 
```
go-backend/
├── controllers/
│   ├── commentController.go
│   ├── threadController.go
│   ├── usersController.go
├── initialisers/
│   ├── connectToDB.go
│   ├── loadEnvVariables.go
│   ├── syncDB.go
├── middleware/
│   ├── requireAuth.go
├── models/
│   ├── commentModel.go
│   ├── threadModel.go
│   ├── tagModel.go
│   ├── userModel.go
├── routes/
│   ├── comments.go
│   ├── threads.go
│   ├── users.go
├── utils/
│   ├── hash.go
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
```

