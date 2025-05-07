# 🔐 Encryption SDK
A pluggable and testable AES-GCM encryption SDK for Go that supports key rotation, caching, and secure encryption/decryption using derived or static keys.

## ✨ Features
- AES-GCM (AES-256) encryption
- Pluggable KeyProvider interface (e.g., HMAC-based, static, from Redis, etc.)
- In-memory key caching
- Key rotation support
- Designed for secure, versioned encryption in production or testing environments

