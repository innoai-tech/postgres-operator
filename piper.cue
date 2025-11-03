package main

import (
	"strings"

	"piper.octohelm.tech/wd"
	"piper.octohelm.tech/client"
	"piper.octohelm.tech/container"

	"github.com/octohelm/piper/cuepkg/golang"
	"github.com/octohelm/piper/cuepkg/containerutil"
)

hosts: local: wd.#Local & {}

pkg: {
	_ver: client.#RevInfo & {}

	version: _ver.version
}

actions: go: X = golang.#Project & {
	cwd:     hosts.local.dir
	main:    "./internal/cmd/postgres-operator"
	version: pkg.version
	goos: [
		"linux",
	]
	goarch: [
		"amd64",
		"arm64",
	]
	ldflags: [
		"-s", "-w",
		"-X", "\(X.module)/internal/version.version=\(X.version)",
	]
}

actions: ship: {
	for pgVersion, pgImageTag in {
		"16": "16.10"
		"18": "18.0"
	} {
		"postgres-\(pgVersion)": containerutil.#Ship & {
			name: "ghcr.io/innoai-tech/postgres-operator"
			tag:  "v\(pgImageTag).0-" + strings.Replace(pkg.version, "v0.0.0-", "", -1)

			from: "docker.io/library/postgres:\(pgImageTag)"

			steps: [
				{
					input: _

					_bin: container.#SourceFile & {
						file: actions.go.build[input.platform].file
					}

					_copy: container.#Copy & {
						"input":  input
						contents: _bin.output
						source:   "/"
						include: ["postgres-operator"]
						dest: "/usr/local/bin"
					}

					output: _copy.output
				},

				container.#Set & {
					config: {
						label: "org.opencontainers.image.source": "https://github.com/innoai-tech/postgres-operator"
						entrypoint: ["/usr/local/bin/postgres-operator"]
					}
				},
			]
		}
	}
}

settings: {
	_env: client.#Env & {
		GH_USERNAME!: string
		GH_PASSWORD!: client.#Secret
	}

	registry: container.#Config & {
		auths: "ghcr.io": {
			username: _env.GH_USERNAME
			secret:   _env.GH_PASSWORD
		}
	}
}
