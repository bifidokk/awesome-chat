### Chat with Authentication Microservice

This Go-based microservice offers chat and authentication functionalities using gRPC. Key features include:

- **User Management:** CRUD operations for users with built-in request validation.
- **Chat System:** Create and manage chats, and send messages.
- **Rate Limiting:** Token bucket algorithm to manage request rates.
- **Circuit Breaking:** Handles service failures gracefully.
- **Metrics Collection:** Prometheus integration for performance monitoring.
- **Logging:** Structured logging with Zap for detailed traceability.

The service includes middleware for request validation, rate limiting, and metrics collection, ensuring reliable and scalable operation.
