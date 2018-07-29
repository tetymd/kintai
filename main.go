package main

import (
    "github.com/sclevine/agouti"
    "log"
)

func main() {
    driver := agouti.ChromeDriver()
    if err := driver.Start(); err != nil {
        log.Fatalf("Failed to start driver:%v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage(agouti.Browser("chrome"))
    if err != nil {
        log.Fatalf("Failed to open page:%v", err)
    }

    if err := page.Navigate("https://accounts.secure.freee.co.jp/login/hr"); err != nil {
        log.Fatalf("Failed to navigate:%v", err)
    }
    page.Screenshot("/tmp/chrome_qiita.jpg")
}
