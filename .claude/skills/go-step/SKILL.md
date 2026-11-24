---
name: go-step
description: Teach the next step of the Go REST API (Todo API) tutorial in the established step-by-step format. Use when the user says "next step", asks to learn the next Go/REST concept, or wants a tutorial step explained.
---

# Go REST API tutorial step

Teach ONE step at a time. The user is learning Go while building a Todo API with the standard library only (`net/http`, `encoding/json`, Go 1.22+ `ServeMux` patterns). Do not add frameworks.

## Before writing the step

1. Read `/home/sarthak/projects/learn-go/main.go` (and any other `.go` files) to see what the user actually has. Never assume they applied the previous step.
2. If their code is behind or differs from the previous step, say so briefly and tell them what to fix before continuing.
3. Pick the next step from this roadmap, or the step the user names:
   1. Hello server (`go mod init`, `package main`, `net/http`)
   2. JSON and structs (struct tags, `json.NewEncoder`)
   3. Routing, path params, reading JSON (`ServeMux`, `PathValue`, `Decode`, errors)
   4. In-memory storage (maps, mutexes, concurrency)
   5. Full CRUD (PUT/PATCH, DELETE)
   6. Validation and errors (status codes, helper for JSON errors)
   7. Project structure (packages, handlers, separation)
   8. Database (SQLite)
   9. Middleware, tests, config

## Output format (follow exactly)

Use this structure, in this order:

```
## Step N: <Title>

<1-2 sentences: what this step adds and why it matters.>

### Concepts first
<Short explanations of the new Go/HTTP ideas, each with a tiny snippet.
Bold the term, then explain it plainly. Explain WHY, not just what.>

### Update `main.go`   (or "Create `file.go`" when adding files)
<One complete Go code block the user can paste. Replace the whole file
unless the change is tiny. Keep the code small and runnable.>

### What's new
<Bullet list: each new construct from the code, in bold code format,
followed by a one-line explanation. Call out gotchas such as
ordering (headers before WriteHeader before body) and the need for `return`
after http.Error.>

### Run it
<Exact commands: restart the server (`Ctrl+C`, then `go run .`), then
`curl` commands in a second terminal. State the expected output. Include
the error cases (404, 400, 405) when relevant.>

### Try this
<2-3 small experiments the user does themselves. Include a "what happens
and why" question. Make the last one point at the next step's topic.>

<Closing line: mention any deliberate shortcut taken in this step
(e.g. ignored errors) and which later step fixes it. Then tell the user
to reply when it works and name the next step with a short preview.>
```

## Style rules

- Teach, don't dump: introduce at most 3-5 new ideas per step.
- Go concepts are explained as they first appear (pointers, slices, `:=`, `_`, exported names, error handling, etc.). Do not re-explain ones already covered.
- Compare to concepts from other languages only if the user mentions one.
- Keep headings, bold terms, and short bullets; no long paragraphs.
- Do NOT write or edit the user's files unless they ask. They type the code themselves. Offer to create the files at the end if that seems helpful.
- Use `http.StatusXxx` constants rather than raw numbers in code, and mention the number once in prose.
- When the user reports a problem ("I don't see X"), first inspect their files and check the running server with `curl` yourself, then list the most likely causes before moving on.
- Keep `CLAUDE.md` accurate if the project structure changes (new packages, test commands).
