# Communications

A lightweight backend service for handling "Contact Us" form submissions from static web apps. This service receives lead data via HTTP, then sends notifications to a pre-registered client via email and SMS using Azure Communication Services. Designed to keep your frontend static and secure, while offloading sensitive operations (like email/SMS sending and credential management) to a robust backend.

---

## Features

- **REST API**: Exposes a simple HTTP API for submitting leads and checking health.
- **Email & SMS Notifications**: Sends both email and SMS to the client when a new lead is submitted.
- **Azure Communication Services Integration**: Uses Azure APIs for reliable delivery.
- **PostgreSQL Database**: Stores clients and leads, with migrations managed automatically.
- **Rate Limiting**: Protects against abuse with configurable per-IP throttling.
- **CORS & Security Headers**: Safe for use with static web apps.
- **Validation**: Strict input validation for phone, email, and message fields.
- **Dockerized**: Easy to run locally or deploy anywhere with Docker.
- **SSH Access**: Optional SSH server for container debugging (port 2222).

---

## API Endpoints

### `GET /api/v1/health`

- Checks database connectivity.
- Returns:  
  - `200 OK` if healthy  
  - `503 Service Unavailable Error` if DB is unreachable

### `POST /api/v1/leads/:id`

- Submits a new lead for the client with the given UUID.
- Request body (JSON):

  ```json
  {
    "name": "John Doe",
    "phone": "+12345678901",
    "email": "john@example.com",
    "message": "Optional message"
  }
  ```

- Returns:
  - `200 OK` on success (email and SMS sent)
  - `400 Bad Request` for invalid input
  - `404 Not Found` if client does not exist
  - `429 Too Many Requests` if rate limit exceeded
  - `500 Internal Server Error` if notification fails

---

## Database Schema

- **clients**: Stores client info (id, name, email, phone, website, timestamps)
- **leads**: Stores each lead submission (id, datetime, name, email, phone, client_id)
- See [`migrations/0001_init.up.sql`](migrations/0001_init.up.sql) for full schema.

---

## Installation & Running Locally

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go 1.24+](https://go.dev/) if running natively

### 1. Clone the Repository

```sh
git clone https://github.com/anicmarko9/communications.git
cd communications
```

### 2. Configure Environment Variables

Copy `.env.example` to `.env` and fill in all required values:

```sh
cp .env.example .env
```

- **Azure Communication Services**:  
  - `AZURE_URL` should be in the format:  
    `endpoint=https://<resource-name>.communication.azure.com;accesskey=<access-key>`
  - `EMAIL_FROM` and `SMS_FROM` must match your Azure sender identities.

- **Database**:  
  - Set `POSTGRES_*` variables as needed.

- **CORS**:  
  - Set `ALLOWED_ORIGINS` to your frontend's URL (e.g., `http://localhost:3000`).

### 3. Start the Stack

```sh
docker-compose up -d
```

- This will start both PostgreSQL and the Go backend.
- The backend will be available at `http://localhost:<PORT>` (as set in `.env`).

### 4. Add Clients

Insert at least one client into the `clients` table (use a DB tool or psql):

```sql
INSERT INTO clients (id, name, email, phone) VALUES (
  '<uuid>', 'Client Name', 'client@example.com', '+12345678901'
);
```

- Use a UUID for `id` (e.g., generate with `uuidgen`).

### 5. Test the API

Send a POST request to `/api/v1/leads/<client-uuid>` with the required JSON body.

---

## Development

- Main entry: [`cmd/main.go`](cmd/main.go)
- API logic: [`internal/handlers/`](internal/handlers/)
- Business logic: [`internal/services/`](internal/services/)
- DB models & DTOs: [`internal/database/models/`](internal/database/models/), [`internal/database/dto/`](internal/database/dto/)
- Config: [`internal/config/config.go`](internal/config/config.go)
- Utilities: [`internal/utils/`](internal/utils/)

---

## Security Notes

- SSH server runs on port 2222 (for debugging).  
  - Default key: [`ssh/authorized_keys`](ssh/authorized_keys)
  - Private key: [`ssh/id_rsa`](ssh/id_rsa) (excluded from git)
- Never expose your `.env` or private keys publicly.
- Rate limiting and CORS are enforced by default.

---

## License

MIT
