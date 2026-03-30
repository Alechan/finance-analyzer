# Deployment

`finance-analyzer` publishes its static web app to GitHub Pages.

## How deployment works

1. A push to `main` triggers [/.github/workflows/pages.yml](/Users/adanos/src/finance-analyzer/.github/workflows/pages.yml#L1).
2. The workflow validates Go code, web unit tests, and the public sensitive-artifact guard.
3. The publishable site is assembled by [/scripts/build-pages-output.sh](/Users/adanos/src/finance-analyzer/scripts/build-pages-output.sh#L1) into `dist/pages/`.
4. GitHub Pages deploys that artifact.
5. A post-deploy smoke test checks the live Pages URL after the deployment finishes.

## Live site

- [https://alechan.github.io/finance-analyzer/](https://alechan.github.io/finance-analyzer/)

## What the post-deploy smoke checks

The deployed smoke test is meant to catch real publish and packaging regressions on the live GitHub Pages site.

It checks:
1. the deployed site becomes reachable,
2. the app boots from published same-origin assets,
3. the public demo profile loads without same-origin fetch failures.

It intentionally does not require full success from the external Highcharts CDN, because automated environments may receive `403` responses from `code.highcharts.com` even when the deployed site itself is packaged correctly.

## Local validation before pushing

From the repo root:

```bash
go test ./...
bash scripts/build-pages-output.sh
cd web
npm install
npm run test:unit
npm run guard:oss-sensitive
npm run test:smoke
```

To run the deployed-site smoke locally against a built Pages artifact:

```bash
bash scripts/build-pages-output.sh
node web/scripts/serve-no-store.mjs --root dist/pages --port 8787 --host 127.0.0.1
PLAYWRIGHT_BASE_URL=http://127.0.0.1:8787/ PLAYWRIGHT_SKIP_WEBSERVER=1 npm --prefix web run test:smoke:deployed
```
