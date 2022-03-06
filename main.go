// fileSystemWalk/main.go

package main

import (
  "flag"
  "fmt"
  "io"
  "os"
  "path/filepath"
)

type config struct {
  ext string  // extension to filter out
  size int64  // min file size
  list bool
}

func run( root string, out io.Writer, cfg config) error {
  return filepath.Walk(root,
    func(path string, info os.FileInfo, err error) error {
      if err != nil {
        return err
      }
      if filterOut(path, cfg.ext, cfg.size, info) {
        return nil
      }
      if cfg.list {
        return listFile(path, out)
      }
      return listFile(path, out)
    })
}

func main() {
  root := flag.String("root", ".", "Root directory to start")
  list := flag.Bool("list", false, "Only list files")
  ext  := flag.String("ext", "", "File extension to filter out")
  size := flag.Int64("size", 0, "Minimum file size")
  flag.Parse()

  c := config{
    ext:  *ext,
    size: *size,
    list: *list,
  }

  if err := run(*root, os.Stdout, c); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

}
