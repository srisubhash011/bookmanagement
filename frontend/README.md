# Book Management Frontend

This is a React frontend for the Go book management API running in `cmd/bookapi`.

## Run backend
1. Open a terminal at the repository root.
2. Run `go run ./cmd/bookapi`.
3. The API listens on `http://localhost:8080`.

## Run frontend
1. Open a terminal in `frontend/`.
2. Run `npm install`.
3. Run `npm run dev`.
4. Open the URL shown by Vite (usually `http://localhost:5173`).

## Notes
- The frontend proxies requests from `/api` to the Go backend to avoid CORS issues.
- Available API endpoints:
  - `GET /api/book/{id}`
  - `GET /api/book/search/{query}`
  - `POST /api/book/add`