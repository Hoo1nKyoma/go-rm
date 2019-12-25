package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "os"
  "path"
  "bufio"
  "strings"
)

var force = flag.Bool("force", false, "Attempt to remove the files without prompting for confirmation, regardless of the file's permissions.")
var filename = flag.String("filename", "", "You must enter the file or directory you want to delete.")
var depth = flag.Int("depth", 1, "This symbol represents the depth of the traversal.")

func main () {
  flag.Parse()
  if *filename == "" {
    fmt.Println("You must enter the file or directory you want to delete.")
    flag.Usage()
    os.Exit(1)
  }
  err := rm(".", *filename, *force, *depth)
  if err != nil {
    fmt.Println(err)
  }
}

func rm (pathname, filename string, force bool, depth int) error {
  depth--
  files, err := ioutil.ReadDir(pathname)
  if err != nil {
    return err
  }
  for _, file := range files {
    _filename := file.Name()
    pathname := path.Join(pathname, _filename)
    if _filename == filename {
      checked := force
      if !force {
        fmt.Print("Are you sure to remove ", pathname, " [Y/n] ")
        reader := bufio.NewReader(os.Stdin)
        input, _, _ := reader.ReadLine()
        answer := strings.ToLower(string(input))
        if answer == "y" || answer == "" {
          checked = true
        }
      }
      if checked {
        os.RemoveAll(pathname)
      }
      continue
    }
    if depth > 0 && file.IsDir() {
      // dont care error
      rm(pathname, filename, force, depth)
    }
  }
  return nil
}
