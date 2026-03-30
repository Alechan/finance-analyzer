# finance-analyzer

Personal finance analysis tooling for credit-card statement extraction and a local-first web app for exploring spending, debt, categories, owners, and data quality.

This public repository contains only public-safe code, tests, and demo data. Real-data integration fixtures and sensitive overlays live in the companion private repo.

## Scope

Today:
1. extract structured data from supported credit-card statement PDFs,
2. provide a browser-based analysis surface using public/demo data by default,
3. keep the public repo runnable without private fixtures.

Later:
1. selectively reintroduce stable public docs,
2. broaden support for additional finance workflows,
3. continue tightening the public/private contract where it improves maintainability.

## Repository Layout

Top-level areas:
1. `pkg/` for Go code and CLI/WASM entrypoints
2. `web/` for the static web app, runtime, storage, and browser tests
3. `scripts/` for repo utility scripts
4. `.github/` and `.githooks/` for automation and local guardrails
5. `docs/` for a curated public documentation set

## Public / Private Split

Public repo:
1. source code
2. public/demo data
3. public test coverage

Private repo:
1. real integration fixtures
2. sensitive CSV overlays
3. private integration test flows

## Getting Started

Go:

```bash
go test ./...
go build -o finance-analyzer ./pkg/cmd/pdfcardssummarycli
```

Web:

```bash
cd web
npm install
npm run build:wasm
npm run test:unit
npm run test:smoke
```

## Current Status

This README is intentionally lightweight for the initial public cut. The immediate priority is a clean, runnable public-safe codebase with no private fixtures or sensitive artifacts.

## License

MIT. See [LICENSE](/Users/adanos/src/finance-analyzer/LICENSE).
