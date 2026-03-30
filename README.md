# finance-analyzer

Personal finance analysis tooling for extracting structured data from supported credit-card statement PDFs and exploring the results in a local-first web app. It focuses on spending, debt, categories, owners, and data quality.

## What can I do with it?

You can:
1. extract transactions from supported PDF statements,
2. inspect the results in the browser without sending financial data to a backend,
3. keep working with your own CSVs and workspace state locally.

## Supported banks

The CLI currently supports only:
1. Santander
2. Visa Prisma

## Install The CLI

From the repo root:

```bash
go install ./pkg/cmd/pdfcardssummarycli
```

## Extract PDFs

The CLI reads one or more statement PDFs from the current folder and writes a CSV next to each PDF.

Example: one PDF.

```bash
ls -1
# march.pdf

pdfcardssummarycli --bank santander march.pdf

ls -1
# march.pdf
# march.pdf.csv
```

Example: multiple PDFs and one combined CSV.

```bash
ls -1
# jan.pdf
# feb.pdf

pdfcardssummarycli --bank visa-prisma --join-csvs all.csv jan.pdf feb.pdf

ls -1
# all.csv
# feb.pdf
# feb.pdf.csv
# jan.pdf
# jan.pdf.csv
```

## Use The Web App

The web app is a static site that keeps your finance data in browser storage on your machine. It does not send CSVs, mappings, or workspace state to a backend because there is no backend in this repo.

Why it works this way:
1. your financial data stays on your device,
2. the site can be hosted as plain static files,
3. it stays simple to run locally and on GitHub Pages.

From the repo root:

```bash
cd web
npm install
npm run build:wasm
npm run test:unit
npm run test:smoke
```

Then open the static site locally with a simple server such as:

```bash
python3 -m http.server 8080 -d web
```

## Repository Layout

The main folders are:
1. `pkg/` for Go code and CLI/WASM entrypoints
2. `web/` for the static web app, runtime, storage, and browser tests
3. `scripts/` for repo utility scripts
4. `.github/` and `.githooks/` for automation and local guardrails
5. `docs/` for public documentation and deployment notes

## More Info

The one network dependency in the public web app is the pinned Highcharts CDN for charting code. It fetches JavaScript assets, not your financial data.

## License

MIT. See [LICENSE](./LICENSE).
