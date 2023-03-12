#!/bin/sh -e

# Install modd
which modd || go install github.com/cortesi/modd/cmd/modd@latest

# Install golang-migrate CLI
which migrate || go install github.com/golang-migrate/migrate/v4

# Install dlv
which dlv || go install github.com/go-delve/delve/cmd/dlv@latest

