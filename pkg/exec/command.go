package exec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"os"
	"os/exec"
	"slices"
	"syscall"

	"github.com/octohelm/x/logr"
)

type Command struct {
	Name    string
	Args    []string
	WorkDir string
	Flags   Flags
	Env     Env
	UID     int
	GID     int

	cmd *exec.Cmd
}

func (c *Command) Run(ctx context.Context) error {
	if err := c.Start(ctx); err != nil {
		return err
	}
	return c.Wait(ctx)
}

func (c *Command) Signal(ctx context.Context, signal os.Signal) error {
	if c.cmd == nil || c.cmd.Process == nil {
		return nil
	}
	return c.cmd.Process.Signal(signal)
}

func (c *Command) Wait(ctx context.Context) error {
	if c.cmd == nil {
		return nil
	}
	return c.cmd.Wait()
}

func (c *Command) Start(ctx context.Context) error {
	cmd := c.command(ctx)

	c.cmd = cmd

	logr.FromContext(ctx).WithValues(slog.Any("cmd", cmd.Args)).Info("exec")

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start %s failed: %w", c.Name, err)
	}
	return nil
}

func (c *Command) command(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.Name, slices.Concat(c.Args, c.Flags.ToArgs())...)

	cmd.Dir = c.WorkDir
	cmd.Env = c.Env.ToEnviron()

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(c.UID),
			Gid: uint32(c.GID),
		},
	}

	cmd.Stdout = prefix(os.Stdout, c.Name)
	cmd.Stderr = prefix(os.Stderr, c.Name)

	return cmd
}

func prefix(w io.Writer, name string) io.Writer {
	return &forward{
		w:    w,
		name: name,
	}
}

type forward struct {
	w    io.Writer
	name string
	buf  []byte
}

func (f *forward) Write(p []byte) (n int, err error) {
	for l := range bytes.Lines(append(f.buf, p...)) {
		if bytes.HasSuffix(l, []byte("\n")) {
			_, _ = fmt.Fprint(f.w, f.name, ": ", string(l))
			continue
		}
		f.buf = l
	}

	return len(p), nil
}

func (c *Command) Output(ctx context.Context) ([]byte, error) {
	cmd := c.command(ctx)

	stdout := new(bytes.Buffer)

	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return stdout.Bytes(), nil
}

type Flags map[string][]string

func (flags Flags) ToArgs() []string {
	if len(flags) == 0 {
		return nil
	}

	args := make([]string, 0, len(flags)*2)
	for _, flag := range slices.Sorted(maps.Keys(flags)) {
		for _, v := range flags[flag] {
			args = append(args, flag, v)
		}
	}
	return args
}

type Env map[string]string

func (env Env) ToEnviron() []string {
	if len(env) == 0 {
		return nil
	}

	envVars := make([]string, 0, len(env)*2)
	for _, e := range slices.Sorted(maps.Keys(env)) {
		envVars = append(envVars, fmt.Sprintf("%s=%s", e, env[e]))
	}

	return envVars
}
