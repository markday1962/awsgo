package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	Archive(src, dest string) error
}

type zipper struct{}

var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	//ensures the destination directory exists
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	//Used to create a new file that is in the dest path
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	// Create a new zip.Writer that will write to the file that has been created above
	w := zip.NewWriter(out)
	defer w.Close()

	//The filepath.Walk func, specifies a second argument of type ilepath.WalkFunc
	//as long as we adhere to the function specifics and inline function can be used
	//as we have done below.
	//The filepathWalk is recursive and will travel into the sub-folders
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
		return nil
	})
}
