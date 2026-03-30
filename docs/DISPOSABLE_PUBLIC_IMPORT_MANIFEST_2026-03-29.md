# Disposable Public Import Manifest (2026-03-29)

Status: Draft
Purpose: define the first safe copy from `/Users/adanos/src/pdf_tmp` into `/Users/adanos/src/finance-analyzer` without bringing private fixtures, stale planning docs, or local-generated residue.

This document is intentionally disposable:
1. it is only for the initial split/import,
2. it should be deleted or archived after the public repo is populated and validated,
3. it does not define the long-term public documentation set.

## Import Goal

Populate `finance-analyzer` with:
1. the public Go code,
2. the public web app and its tests,
3. public/demo data only,
4. no sensitive fixtures,
5. no transition planning residue by default.

## Move Into Public Repo

### Repository-Level
Move as-is:
1. `.github/`
2. `.githooks/`
3. `.gitignore`
4. `go.mod`
5. `go.sum`
6. `scripts/install-git-hooks.sh`

Do not copy:
1. `.git/`
2. `.idea/`
3. `.cursor/`
4. `.playwright-cli/`
5. `pdfcardssummarycli`

### Go Tree
Move as-is:
1. `pkg/cmd/`
2. `pkg/demo_dataset/`
3. `pkg/internal/`

Do not move:
1. `pkg/integration_tests/`

Reason:
1. `pkg/integration_tests/` is part of the private test/data workflow and already contains sensitive and private-oriented material.

### Web Tree
Move as-is:
1. `web/README.md`
2. `web/package.json`
3. `web/package-lock.json`
4. `web/playwright.config.js`
5. `web/index.html`
6. `web/csvCombine.js`
7. `web/csvCombine.test.js`
8. `web/format.js`
9. `web/format.test.js`
10. `web/mappingsCsvParse.js`
11. `web/mappingsCsvParse.test.js`
12. `web/tableRenderer.js`
13. `web/tableRenderer.test.js`
14. `web/tableRenderer.fixtures.js`
15. `web/wasm_exec.js`
16. `web/__snapshots__/`
17. `web/e2e/`
18. `web/runtime/`
19. `web/scripts/`
20. `web/storage/`
21. `web/mockups_lab/app/`
22. `web/mockups_lab/shared/`
23. `web/mockups_lab/reference_pack/`
24. `web/mockups_lab/tmp_public_data/`
25. `web/mockups_lab/vendor/`

Move, but regenerate after copy:
1. `web/finance.wasm`

Do not move:
1. `web/mockups_lab/tmp_sensitive_data/`
2. `web/node_modules/`
3. `web/output/`
4. `web/test-results/`

### Docs
Create fresh in `finance-analyzer` now:
1. `README.md`
2. `LICENSE`
3. this disposable import manifest

Do not copy in the first import:
1. `docs/`
2. `web/mockups_lab/docs/`

Reason:
1. the current docs trees still contain planning, audit, migration, and private-split residue,
2. the first public import should stay conservative,
3. stable public docs can be reintroduced later by explicit allowlist.

## Explicit Exclusion List

These paths must stay out of the first public import:
1. `pkg/integration_tests/`
2. `web/mockups_lab/tmp_sensitive_data/`
3. `.idea/`
4. `.cursor/`
5. `.playwright-cli/`
6. `web/node_modules/`
7. `web/output/`
8. `web/test-results/`
9. `pdfcardssummarycli`
10. `docs/OSS_SENSITIVE_DATA_INVENTORY.md`
11. `docs/PRIVATE_REPO_GO_BOOTSTRAP_PLAN_2026-03-22.md`
12. `docs/DISPOSABLE_PUBLIC_SPLIT_CLEANUP_2026-03-29.md`
13. `docs/FINANCE_DASHBOARD_LIVE_TASKS.md`
14. `docs/FINANCE_DASHBOARD_MVP_PLAN.md`
15. `docs/FRONTEND_TESTS_PLAN.md`
16. `docs/INDEXEDDB_PLAN.md`
17. `docs/T13_FILE_IMPORT_PLAN.md`
18. `docs/T21_INDEXEDDB_PERSISTENCE_PLAN.md`
19. `docs/T22_MAPPINGS_PLAN.md`
20. `docs/T25_MAPPINGS_CSV_PLAN.md`
21. `docs/T26_DISABLE_DRAG_DROP_PLAN.md`
22. `docs/T2_EXPLORE_UX_DATA_MODEL_PLAN.md`
23. `docs/WEB_APP_TRANSITION_TASKS.md`
24. `docs/WEB_APP_UX_UI_AGENT_RESEARCH.md`
25. `docs/WEB_APP_UX_UI_AGENT_WORKFLOW.md`
26. `docs/WEB_APP_UX_UI_AUDIT_2026-03-14.md`
27. `docs/WEB_APP_UX_UI_AUDIT_2026-03-15_WTT-072.md`
28. `docs/WEB_APP_UX_UI_PARTS_PLAN_2026-03-15.md`
29. `docs/WEB_APP_VIEW_QUESTIONS_WORKBOOK.md`
30. `docs/WEB_APP_VIEW_RESET_PLAN.md`

## Post-Copy Follow-Up

After the public-safe folders are copied:
1. confirm the imported tree uses `github.com/Alechan/finance-analyzer` consistently,
2. update any remaining repo-specific references accordingly,
3. regenerate `web/finance.wasm`,
4. rerun public validation.

## Third-Party Licensing Caveat

Before publishing the public repo:
1. confirm redistribution and licensing posture for vendored third-party assets under `web/mockups_lab/vendor/`,
2. in particular, review `highcharts.js` and `highcharts-dark-unica.js` before treating the repo as publication-ready,
3. if redistribution is not acceptable, replace vendored assets with an allowed alternative or a documented install step before the first public release.

## Public Validation Checklist

After code is copied into `finance-analyzer`:
1. `go test ./...`
2. `npm run test:unit` from `web/`
3. `npm run test:smoke` from `web/`
4. `npm run audit:ux` from `web/`
5. `npm run guard:oss-sensitive` from `web/`

## Public Safety Checks

Before declaring the repo import-ready:
1. `rg -n "tmp_sensitive_data|integration_tests|PRIVATE_REPO|OSS_SENSITIVE|LIVE_TASKS|_PLAN|WORKBOOK|ux-audit|ux-parts-analysis" /Users/adanos/src/finance-analyzer`
2. confirm only the intentionally created disposable import manifest appears among planning-oriented docs,
3. confirm no private CSVs, PDFs, expected Go fixtures, or extracted private text artifacts exist in the public repo.
