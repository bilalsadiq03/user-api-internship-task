# Architectural Reasoning and Design Decisions

This document explains the reasoning behind key technical and architectural choices made while building this project.

---

## 1. Why Layered Architecture?

The project follows a layered architecture:
- Handler
- Service
- Repository

This separation ensures:
- Clean responsibility boundaries
- Easier testing
- Maintainability and scalability

---

## 2. Why SQLC?

SQLC was chosen to:
- Generate type-safe database queries
- Eliminate runtime SQL errors
- Improve developer productivity
- Keep SQL readable and explicit

---

## 3. Why Not Store Age in Database?

Age is derived data and changes every year.

Storing age would:
- Cause data inconsistency
- Require periodic updates

Instead:
- `dob` is stored permanently
- `age` is calculated dynamically using Go’s `time` package

This ensures correctness at all times.

---

## 4. Why DTOs for Requests?

DTOs (Data Transfer Objects) are used to:
- Separate API input from domain models
- Apply strict validation rules
- Prevent accidental exposure of internal structures

---

## 5. Validation Strategy

The `go-playground/validator` library is used to:
- Validate required fields
- Enforce correct date format (`YYYY-MM-DD`)
- Fail fast on invalid input

Validation happens at the API boundary (handler layer).

---

## 6. Logging Strategy

Uber Zap is used because:
- It is fast and production-ready
- Supports structured logging
- Makes debugging and monitoring easier

Key events such as request handling and errors are logged.

---

## 7. Middleware Usage

Middleware is used to:
- Inject a unique request ID into each request
- Log request duration for performance monitoring

This improves observability without polluting business logic.

---

## 8. Error Handling Approach

- Proper HTTP status codes are returned
- Internal errors are logged
- Client-facing error messages remain clean

---

## 9. Scalability Considerations

- Repository pattern allows DB changes with minimal impact
- Service layer can grow with business logic
- Pagination can be easily added to list endpoints

---

## Conclusion

This project was designed to follow real-world backend best practices, emphasizing:
- Clean architecture
- Type safety
- Correctness
- Maintainability

The design choices reflect how production-grade Go services are built.
