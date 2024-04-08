# LRU Cache with Get/Set API and React App

This repository contains an LRU (Least Recently Used) cache implemented in Golang with Get/Set API endpoints exposed. Additionally, it includes a React application that consumes these API endpoints to interact with the cache.

## LRU Cache (Backend)

### Requirements

- Store Key/Value pairs with expiration time.
- Evict keys from the cache after expiration.
- Maximum of 1024 keys.
- Expose Get/Set methods as API endpoints.
- Backend built in Golang.

### Must Haves

- Golang backend.
- API endpoints for Get/Set methods.

### Good to Have

- Concurrency implementation in cache.

## React Application (Frontend)

### Requirements

- Consume Get API to retrieve keys from the cache.
- Consume Set API to add or update key/value pairs in the cache.

## Setup and Usage

### Backend (Golang)

1. **Clone Repository**: Clone this repository to your local machine.

```bash
git clone <repository-url>
```

2. **Navigate to Backend Directory**: Enter the backend directory.

```bash
cd backend
```

3. **Install Dependencies**: Make sure you have Golang installed on your machine. Then, install the dependencies.

```bash
go mod download
```

4. **Run the Server**: Start the Golang server.

```bash
go run main.go
```

### Frontend (React)

1. **Navigate to Frontend Directory**: Go to the frontend directory from the root of the repository.

```bash
cd frontend
```

2. **Install Dependencies**: Make sure you have Node.js and npm installed. Then, install the dependencies.

```bash
npm install
```

3. **Start the React App**: Run the React application.

```bash
npm start
```

### Using the Application

Once both the backend and frontend are running:

- Access the React application via `http://localhost:3000`.
- Use the provided UI to interact with the cache: set keys, retrieve keys, etc.
