# Contributing to docker-credential-acr

Thanks for your interest in contributing! Here's how to get started.

## Prerequisites

- [Go 1.23+](https://go.dev/dl/)
- A working Docker installation (for end-to-end testing)

## Getting started

1. Fork the repository and clone your fork.
2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build:

   ```bash
   make build
   ```

4. Run tests:

   ```bash
   make test
   ```

## Making changes

1. Create a branch from `main`:

   ```bash
   git checkout -b my-change
   ```

2. Make your changes and add tests where appropriate.
3. Run the linter:

   ```bash
   make lint
   ```

4. Commit with a clear, descriptive message.
5. Push your branch and open a pull request against `main`.

## Pull request guidelines

- Keep PRs focused — one logical change per PR.
- Include tests for new functionality or bug fixes.
- Make sure CI passes before requesting review.
- Update documentation if your change affects public APIs or usage.

## Reporting issues

Use the [GitHub issue tracker](https://github.com/gaganhr94/docker-credential-acr/issues).
Please include:

- What you expected to happen
- What actually happened
- Steps to reproduce
- Go version and OS

## Code of conduct

Be respectful and constructive. We're all here to build something useful.
