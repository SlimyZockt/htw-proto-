#!/usr/bin/env bash
set -e

# go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/a-h/templ/cmd/templ@latest
# goose -dir=assets/migrations/ sqlite3 app.db up
pnpm install tailwindcss @tailwindcss/cli
npx @tailwindcss/cli -o ./include_dir/output.css
templ generate
