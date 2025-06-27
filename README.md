
# Edge Gateway - Secure IoT Message Forwarder

A lightweight, reliable microservice written in Go that actus as an intermediary between IoT devices and cloud services. This project will be designed to ensure secure, fault-tolerant delivery of critical device events in real-time.

---

## Project Overview

The Edge Gateway simulates a production-style connectivity layer between smart devices and backend systems, supporting:

 - Secure event ingestion from IoT devices
 - Message buffering and retry logic for reliability
 - Forwarding events to a cloud service
 - Minimal, testable, containerized design

This project has been built to mirror real-world IoT communication patterns and and strives to highlight core backend engineering skills in reliability, observability, and secure communication.

---

## Architecture

`[IoT Devices]--->[Edge Gateway]--->[Cloud Service]`

 - Devices send telemetry events to the DEdge Gateway via HTTPS
 - Edge Gateway queues, validates, and forwards events to a configured cloud endpoint
 - Retries failed messages, ensuring eventual delivery

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| API Server | Go (`net/http`, `mux`) |
| Secure Communication | TLS, Token-based Auth |
| Message Queue | In-memory (extendable to Redis) |
| Deployment | Docker |
| Config | Environment Variables |

---

## Configuration

| Env Variable | Default | Description |
|--------------|---------|-------------|
| `LISTNE_ADDR` | `:8080` | Address and port to listen on |
| `CLOUD_URL` | `http://cloud-service:9000/events` | Cloud service endpoint |

---

## Planned Enhancements

- Protobuf message serialization
- Redis-backend persistent queue
- Prometheus metrics for observability
- Circuit breaker & exponential backoff for forwarding
- Unit tests for all components\
- Terrafrom ifrastructure examples

---

## Author

Kelsy Frank

M.S. in Computer Science | Learniung Tech Enthusiast

[LinkedIn](https://www.linkedin.com/in/kelsy-frank-36a20732a/)