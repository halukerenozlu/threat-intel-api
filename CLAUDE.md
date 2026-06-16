# CLAUDE.md

Behavioral guidelines to reduce common LLM coding mistakes.Merge with
project-specific instructions as needed.

**Tradeoff note:** These guidelines bias toward caution over speed. For
trivial tasks (typo fixes, obvious one-liners), use judgment — not every
change needs the full rigor.

## 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

Before implementing:

- State your assumptions explicitly. If uncertain, ask rather than guess.
- If multiple interpretations exist, present them — don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, stop. Name what's confusing. Ask.

## 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If you write 200 lines and it could be 50, rewrite it.

Ask yourself: "Would a senior engineer say this is overcomplicated?"
If yes, simplify.

## 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:

- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, mention it — don't delete it.

When your changes create orphans:

- Remove imports/variables/functions that YOUR changes made unused.
- Don't remove pre-existing dead code unless asked.

The test: every changed line should trace directly to the user's request.

## 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:

- "Add validation" → "Write tests for invalid inputs, then make them pass"
- "Fix the bug" → "Write a test that reproduces it, then make it pass"
- "Refactor X" → "Ensure tests pass before and after"

For multi-step tasks, state a brief plan:

1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]

Strong success criteria let you loop independently. Weak criteria
("just make it work") require constant clarification.

---

**These guidelines are working if:** fewer unnecessary changes appear in
diffs, fewer rewrites happen due to overcomplication, and clarifying
questions come before implementation rather than after mistakes.

---

## Project-Specific Guidelines — threat-intel-api

- Language: Go 1.26. Prefer the standard library; avoid unnecessary dependencies.
- Strict layer separation: `handler` only deals with HTTP request/response,
  business logic lives in `service`, database access only goes through
  `repository`. Don't suggest code that violates these boundaries.
- When adding a new external API client, follow the structure and naming
  pattern of existing files under `internal/service/threatclient/`
  (e.g. `abuseipdb.go`).
- NEVER touch the `.env` file or real API keys, and never suggest committing them.
- When adding a new SQL migration, follow the existing numbering sequence
  (pattern like `000001_create_users.up.sql`).
- When giving code examples for learning-critical parts (auth, concurrency,
  SQL), don't hand over the complete solution directly — explain the
  approach first and let the user write it themselves.
- Whenever a piece of code is completed, the following steps are mandatory
  before it can be marked "done":
  1. Verify the project builds with `go build ./...`.
  2. Check for suspicious patterns with `go vet ./...`.
  3. If tests exist for the relevant package, run `go test ./...` and
     confirm they pass.
  4. If a new function/handler/endpoint was added, write at least one
     happy-path test and one error-case test — if no test was added,
     say so explicitly rather than silently skipping it.
  5. For code that depends on an external API, use an interface + mock so
     it can be tested without making real network calls; never make tests
     depend on network connectivity.
- If any of these verification steps fails, do not move to the next step —
  fix that failure first, then proceed. Don't ignore a failing step with a
  justification like "it works in the big picture."
