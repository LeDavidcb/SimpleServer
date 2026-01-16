# SimpleServer

A simple HTTP server written in Go to demonstrate secure request handling using JWT (JSON Web Tokens). The project keeps dependencies minimal, using only `.env` and `jwt` libraries.

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/ledavid/SimpleServer.git
   ```
2. Add a `.env` file with:
   ```env
   SECRET_KEY=your_jwt_secret_key
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

## Available Routes
- `GET /home` - Home route (protected with JWT authorization).
- `POST /login` - Login and get a token. The body must provide the next json: {"email": <email>, "password" <password>}
- `GET /refrehs` - Refresh your token. The body must provide the last refreshToken in the body
