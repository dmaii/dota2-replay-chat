package main

import (
  "os"
  "fmt"
  "github.com/dotabuff/manta"
  "github.com/dotabuff/manta/dota"
)

func main() {
  // Create a new parser instance from a file. Alternatively see NewParser([]byte)
  p, _ := manta.NewParserFromFile(os.Args[1])

  // Register a callback, this time for the OnCUserMessageSayText2 event.
//  p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
//    fmt.Printf("%#v \n", m)
//    fmt.Print(m.GetMessagename(), m.GetParam3())
//    //fmt.Printf("%s said: %s \n", m.GetParam1(), m.GetParam2())
//    return nil
//  })

  p.Callbacks.OnCUserMessageSayText(func(m *dota.CUserMessageSayText) error {
    return nil
  })

  p.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
    _ = "breakpoint"
    fmt.Println(m.GetGameInfo().GetDota().GetPlayerInfo())

    return nil
  })

  // Start parsing the replay!
  p.Start()
}
