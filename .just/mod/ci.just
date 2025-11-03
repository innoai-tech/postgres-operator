export TTY := "0"
piper := 'piper' + if env("DEBUG", "0") == '1' { ' --log-level=debug' } else { '' }

build-webapp: (build-web "console")

build:
    {{ piper }} do go build

ship:
    {{ piper }} do ship postgres-16
    {{ piper }} do ship postgres-18
