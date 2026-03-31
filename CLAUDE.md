# CLAUDE.md

This file provides guidance to AI assistants (Claude and others) working in this repository.

## Repository Overview

**Project**: taco  
**Owner**: maddenmanel  
**Status**: Initial setup — no application code yet.

This file should be updated as the project evolves to reflect the actual codebase structure, conventions, and workflows.

---

## Development Branch

All Claude-initiated development should occur on branches prefixed with `claude/`. The current documentation branch is `claude/add-claude-documentation-huUHz`.

Never push directly to `main` without explicit user approval.

---

## Git Conventions

### Commit Messages
- Use the imperative mood: "Add feature" not "Added feature"
- Keep the subject line under 72 characters
- Include a blank line before the body if more context is needed
- Reference issue numbers when applicable: `Fix login bug (#42)`

### Branch Naming
- Features: `feature/<short-description>`
- Bug fixes: `fix/<short-description>`
- Claude-initiated work: `claude/<short-description>-<id>`
- Documentation: `docs/<short-description>`

### Commit Signing
This repository is configured with SSH commit signing (`gpgformat = ssh`). All commits must be signed. Do not bypass signing with `--no-gpg-sign`.

### Pull Requests
- Do not create a pull request unless explicitly asked by the user
- PR titles should be short (under 70 characters)
- Include a summary and test plan in the PR body

---

## AI Assistant Behavior

### General Rules
- Read files before modifying them
- Do not create files unless strictly necessary
- Do not add features or refactoring beyond what is asked
- Do not add comments, docstrings, or type annotations to unchanged code
- Prefer editing existing files over creating new ones
- Match the scope of changes to what was actually requested

### Security
- Never introduce command injection, XSS, SQL injection, or other OWASP Top 10 vulnerabilities
- Validate only at system boundaries (user input, external APIs)
- Do not commit secrets, credentials, or `.env` files

### Reversibility
Confirm with the user before taking irreversible or high-blast-radius actions:
- Deleting files or branches
- Force-pushing
- Dropping database tables or data
- Modifying CI/CD pipelines
- Pushing to shared branches

---

## Codebase Structure

> To be populated as the project develops.

```
taco/
├── CLAUDE.md          # This file
└── ...                # Application code to be added
```

---

## Technology Stack

> To be documented once the project stack is established.

Typical sections to fill in:
- **Language(s)**: e.g., TypeScript, Python, Go
- **Framework(s)**: e.g., Next.js, FastAPI, Gin
- **Package manager**: e.g., npm, pnpm, poetry, go modules
- **Database**: e.g., PostgreSQL, SQLite
- **Infrastructure**: e.g., Docker, Kubernetes, AWS

---

## Development Workflow

> To be documented once tooling is in place.

Typical sections to fill in:
- How to install dependencies
- How to run the application locally
- How to run tests
- How to build for production
- Environment variable setup (`.env.example`)

---

## Testing

> To be documented once a test framework is chosen.

Typical sections to fill in:
- Test framework and runner
- How to run unit tests
- How to run integration/e2e tests
- Coverage requirements

---

## Code Conventions

> To be documented as conventions emerge.

Typical sections to fill in:
- Linter/formatter configuration
- Naming conventions (files, functions, variables)
- Module/package organization patterns
- Error handling patterns
- Logging conventions

---

## Key Files to Know

> Update this section as the project grows.

| File | Purpose |
|------|---------|
| `CLAUDE.md` | AI assistant guidance (this file) |

---

## Updating This File

When the project grows, update this file to reflect:
1. Actual directory structure
2. Real technology stack and version requirements
3. Concrete setup and run instructions
4. Observed code conventions and patterns
5. Test commands and coverage expectations
