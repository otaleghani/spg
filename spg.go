/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Here you will export all of the different functions to create a
placeholder. The idea is that from this functions you will call the
formater that in of itself calls the  parser. The flow is as follows.

Raw data ------> String ----------> Output data
Parser --------> Formatter -------> Export function

So based on this I could will give the export functions all of the
different data that it needs to actually call parser/formatter.

So it will need a language, then it will need 

func FirstName(lang string, format string)

// I could call the format here so that I could say like uppercase,
// lowercase etc. etc. directly without any trouble.

Usage


Dependency

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

package spg

import (
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"
  //"errors"
  "embed"
)

//go:embed dict/*
var csvFile embed.FS

func FirstName(lang string, format string) (string, error) {
  word, err := parser.ParseCvs("dict/names/" + lang + ".csv")
  if err != nil {
    return "", err
  }
  
  err = formatter.Switch(&word, format)
  if err != nil {
    return "", err
  }

  return word, nil
}

func LastName(lang string, format string) (string, error) {
  word, err := parser.ParseCvs("dict/surnames/" + lang + ".csv")
  if err != nil {
    return "", err
  }
  
  err = formatter.Switch(&word, format)
  if err != nil {
    return "", err
  }

  return word, nil
}

func FullName(lang string, format string, div string) (string, error) {
  var array []string
  firstName, err := FirstName(lang, format)
  if err != nil {
    return "", err
  }
  lastName, err := LastName(lang, format)
  if err != nil {
    return "", err
  }
  array = append(array, firstName, lastName)

  return formatter.Divider(array, div), nil
}

//func Words(lang string, )
