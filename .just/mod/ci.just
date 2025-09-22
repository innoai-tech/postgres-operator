export TTY := "0"
piper := 'piper' + if env("DEBUG", "0") == '1' { ' --log-level=debug' } else { '' }

build:
    {{ piper }} do go build

ship:
    {{ piper }} do ship
