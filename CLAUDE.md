# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

This is a learning project: the user is learning Go while building a Todo REST API step by step, guided by Claude. The planned progression is hello server → JSON/structs → routing (GET/POST/path params) → in-memory storage (maps, mutexes) → full CRUD → validation/errors → package structure → SQLite → middleware/tests/config. Teach in small steps, explain Go concepts as they appear, and let the user write the code unless they ask otherwise.

## Commands

- Run: `go run .` (serves on `:8080`; blocks the terminal, so test with `curl` from a second one)
- Build: `go build -o todo-api .`
- Vet: `go vet ./...`
- Format: `gofmt -w .`
- Tests (none yet): `go test ./...`; single test: `go test -run TestName ./...`

## Architecture

- Module name is `todo-api` (`go.mod`, Go 1.27). Currently a single `main.go` in `package main`.
- Standard library only (`net/http`, `encoding/json`). The goal is to use Go 1.22+ `http.ServeMux` method/path patterns (e.g. `mux.HandleFunc("GET /todos/{id}", ...)`, `r.PathValue("id")`) rather than a third-party router; avoid introducing frameworks.
- The `Todo` struct (`ID`, `Title`, `Done`) carries `json` struct tags; fields must stay capitalized (exported) or `encoding/json` skips them.
