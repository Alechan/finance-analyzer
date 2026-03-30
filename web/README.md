## Minimal WASM Demo Shell

> **Reminder:** Update `docs/index.md` if you add, rename, or remove documents.

This is a tiny static web shell to exercise the Go WASM API end-to-end:
- load demo CSV + mappings from embedded WASM functions
- compute all tables
- preview one selected table
- export the selected table as CSV

### Build the WASM binary

From repo root:

```sh
GOOS=js GOARCH=wasm go build -o web/finance.wasm ./pkg/cmd/financewasm
```

### Copy `wasm_exec.js`

From repo root:

```sh
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
```

### Run a static server

From repo root:

```sh
python3 -m http.server 8080 -d web
```

Then open:

```
http://localhost:8080
```

### OSS sensitive artifact guard

From `web/`:

```sh
npm run guard:oss-sensitive
```

Optional full-tree scan (expected to fail until private fixture migration is complete):

```sh
npm run guard:oss-sensitive:all
```

Enable local pre-commit enforcement from repo root:

```sh
./scripts/install-git-hooks.sh
```

### UX audit automation

From `web/`:

```sh
npm run audit:ux
```

Artifacts are written to:

1. `web/output/playwright/ux-audit/ux-audit-summary.json`
2. `web/output/playwright/ux-audit/ux-audit-*.png`
