/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Part of the parser package, used to parse data based on the different
database in use. This file contains helper functions to download
locally a cache of the dictionaries in use.
At the moment we are using mainly a stripped down version of public
archives converted into csv format, so most of this functions revolve
around this. Every cached file will be save into $HOME/.cache/spg/

Usage
  func initDict(dict string) error
  | Given a dict path, will download it to the cache

  func cachePath() string {
  | Gives you the default cache path 

  func parsePathAndName(dict string) (string, string)
  | From a dict input (names/en.csv) gives you two string, one is the
  path (names/) the second one the file name (en.csv)

  func DeleteCache() error
  | Deletes every single file in the cache path 

Dependency
  "io"
  "os"
  "net/http"
  "errors"
  "strings"

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/
package parser

import (
  "io"
  "os"
  "net/http"
  "errors"
  "strings"
  "fmt"
)

func initDict(dict string) error {
  prefix := "https://raw.githubusercontent.com/otaleghani/spg/main/dict/"

  fmt.Println(dict)
  resp, err := http.Get(prefix + dict)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return errors.New("Server returned: " + resp.Status)
  }
  filePath, _ := parsePathAndName(dict)
  err = os.MkdirAll(cachePath() + filePath, os.ModePerm)
  if err != nil {
    return err
  }
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return err
  }
  err = os.WriteFile(cachePath() + dict, body, 0666)
  if err != nil {
    return err
  }
  return nil
}

func cachePath() string {
  homeDir := os.Getenv("HOME")
  if homeDir == "" {
    homeDir = os.Getenv("USERPROFILE")
  }
  path := "/.cache/spg/"
  return homeDir + path
}

func parsePathAndName(dict string) (string, string) {
  dictIndex := strings.Index(dict, "/")
  return dict[:dictIndex+1], dict[dictIndex:]
}

func DeleteCache() error {
  err := os.RemoveAll(cachePath())
  if err != nil {
    return err
  }
  return nil
}
