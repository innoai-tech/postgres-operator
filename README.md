# Postgres Operator

Self-management Postgres instance operator.

- [x] auto configuration by cpu/mem
- [x] readiness/liveness check
- [x] backup & restore
    - [x] cronjob
- [x] console dashboard
- [x] metrics
- [x] upgrade:
  - must mount old image host fs to /postgres-toolchain