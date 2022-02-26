package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "time"
  "github.com/xuri/excelize/v2"
)

func main() {

  args := os.Args[1:]

  if len(args) < 4 { return }

  f, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  ct := ""

  if len(args) > 4 {
    ct = strings.ToLower(args[4])
  } else {
    ct = "s"
  }

  switch ct {
    case "b": {
      b, err := strconv.ParseBool(args[3])
      if err == nil {
        err = f.SetCellValue(args[1], args[2], b)
      }
    }
    case "n": {
      n, err := strconv.ParseFloat(args[3], 64)
      if err == nil {
        err = f.SetCellValue(args[1], args[2], n)
      }
    }
    case "s": {
      err = f.SetCellValue(args[1], args[2], args[3])
    }

    case "d": {
      d, err := time.ParseDuration(args[3])
      if err == nil {
        err = f.SetCellValue(args[1], args[2], d)
      }
    }

    case "t": {
      t, err := time.Parse("0000-00-00_00:00:00", args[3])
      if err == nil {
        err = f.SetCellValue(args[1], args[2], t)
      }
    }

  }

  if err != nil {
    fmt.Println(err)
    return
  }

  err = f.Save()
  if err != nil {
    fmt.Println(err)
  }

}