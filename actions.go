// fileSystemWalk/actions.go

package main

import (
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
)

// func filterOut returns true if info is a dir, below minSize, if 
//   ext is not "" and if filepath.Ext(path) is not ext.
// Otherwise returns false
func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {
  if info.IsDir() || info.Size() < minSize {
    return true
  }
  if ext != "" && filepath.Ext(path) != ext {
    return true
  }
  return false
}

// listFile lists files
func listFile(path string, out io.Writer) error {
  _,err := fmt.Fprintln(out, path)
  return err
}

// delFile removes the given path and returns any error.
func delFile(path string, delLogger *log.Logger) error {
  if err := os.Remove(path); err != nil {
    return err
  }
  delLogger.Println(path)
  return nil
}
