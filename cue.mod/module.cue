module: "github.com/innoai-tech/postgres-operator"
language: {
	version: "v0.13.2"
}
deps: {
	"github.com/octohelm/kubepkgspec@v0": {
		v: "v0.0.0-20251028062053-35b39c4bda49"
	}
	"github.com/octohelm/piper@v0": {
		v: "v0.0.0-20250922044649-a61e5946afb1"
	}
	"piper.octohelm.tech@v0": {
		v:       "v0.0.0-builtin"
		default: true
	}
}
