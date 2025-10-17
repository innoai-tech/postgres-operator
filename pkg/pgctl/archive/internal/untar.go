package internal

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Untar(r io.Reader, dir string) (err error) {
	tr := tar.NewReader(r)

	madeDir := map[string]bool{}

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("tar error: %v", err)
		}
		if !validRelPath(f.Name) {
			return fmt.Errorf("tar contained invalid name error %q", f.Name)
		}

		rel := filepath.FromSlash(f.Name)
		abs := filepath.Join(dir, rel)

		mode := f.FileInfo().Mode()

		switch f.Typeflag {
		case tar.TypeReg:
			dir := filepath.Dir(abs)
			if !madeDir[dir] {
				if err := os.MkdirAll(filepath.Dir(abs), 0o750); err != nil {
					return err
				}
				madeDir[dir] = true
			}
			wf, err := os.OpenFile(abs, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode.Perm())
			if err != nil {
				return err
			}
			n, err := io.Copy(wf, tr)
			if closeErr := wf.Close(); closeErr != nil && err == nil {
				err = closeErr
			}

			if err != nil {
				return fmt.Errorf("error writing to %s: %v", abs, err)
			}

			if n != f.Size {
				return fmt.Errorf("only wrote %d bytes to %s; expected %d", n, abs, f.Size)
			}

			if !f.ModTime.IsZero() {
				if err := os.Chtimes(abs, f.AccessTime, f.ModTime); err != nil {
					return err
				}
			}
		case tar.TypeDir:
			if err := os.MkdirAll(abs, 0o750); err != nil {
				return err
			}
			madeDir[abs] = true
		case tar.TypeXGlobalHeader:
			// ignore
		default:
			return fmt.Errorf("tar file entry %s contained unsupported file type %v", f.Name, mode)
		}
	}
	return nil
}

func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}
