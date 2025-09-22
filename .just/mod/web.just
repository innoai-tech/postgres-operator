bun := 'bun'
turbo := 'bunx --bun turbo'
vite := 'bunx --bun vite'
taze := 'bunx --bun taze'
export APP_ENV := env('APP_ENV', 'local')

dep-web:
    {{ bun }} install --no-cache

clean-web:
    rm -f bun.lock;
    find . -name '.swc' -type d -prune -print -exec rm -rf '{}' \;
    find . -name '.turbo' -type d -prune -print -exec rm -rf '{}' \;
    find . -name 'node_modules' -type d -prune -print -exec rm -rf '{}' \;

test-web:
    {{ bun }} test

lint-web:
    {{ turbo }} run lint --force
    {{ bun }} run lint

update-web:
    {{ taze }} -w -r latest

build-web $APP:
    {{ vite }} build --mode=production --outDir=./internal/webapp/public/{{ APP }}

dev-web $APP:
    {{ vite }} --host

preview-web $APP:
    {{ vite }} preview
