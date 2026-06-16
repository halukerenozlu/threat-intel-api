# threat-intel-api

A Go project that collects and aggregates threat intelligence data from multiple sources
and serves it via an authenticated REST API.

## What does it do?

When an IP address, URL, or file hash is queried, the system
simultaneously sends requests to the following sources and combines the results to return a single threat score:

- AbuseIPDB
- VirusTotal
- URLhaus
- AlienVault OTX

## Technology stack

- **Language**: Go 1.26
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)
- **Internal communication** (planned): gRPC
- **AI Feature** (planned): LLM-based report summarization

## Project Status

🚧 Under development — no working version available yet.
