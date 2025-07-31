# Agent Guidelines for Go Projects

This document outlines the conventions and commands for Go projects within this repository.

## Build/Lint/Test Commands

*   **Build:** `go build ./...`
*   **Lint:** `go vet ./...`
*   **Run all tests:** `go test ./...`
*   **Run a single test file:** `go test <path_to_package>` (e.g., `go test ./lasagna`)
*   **Run a specific test function:** `go test -run <TestFunctionName> <path_to_package>` (e.g., `go test -run TestElapsedTime ./lasagna`)

## Code Style Guidelines (Go)

*   **Formatting:** Use `go fmt` to format code.
*   **Imports:** Group imports, standard library first, then third-party, then internal. Use blank lines to separate groups.
*   **Naming Conventions:**
    *   Package names: short, all lowercase, no underscores.
    *   Function/Variable names: `CamelCase` for exported, `camelCase` for unexported.
    *   Acronyms: `HTTP`, `URL` (not `Http`, `Url`).
*   **Error Handling:** Return errors explicitly as the last return value. Check errors immediately.
*   **Types:** Use specific types over `interface{}` where possible.
*   **Comments:** Explain *why* code does something, not *what* it does.
