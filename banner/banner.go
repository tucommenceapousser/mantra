package banner

import (
  "fmt"
  "time"
)

func colorfulBanner(text string) {
  colors := []string{
    "\033[31m", // Red
    "\033[32m", // Green
    "\033[33m", // Yellow
    "\033[34m", // Blue
    "\033[35m", // Magenta
    "\033[36m", // Cyan
  }

  for _, char := range text {
    color := colors[time.Now().Nanosecond()%len(colors)]
    fmt.Printf("%s%c\033[0m", color, char)
    time.Sleep(10 * time.Millisecond)
  }
  fmt.Println()
}

func Banner() {
  bannerText := `
   ██████╗  ██████╗       ████████╗ ██████╗  ██████╗ ██╗     
  ██╔════╝ ██╔═══██╗      ╚══██╔══╝██╔═══██╗██╔═══██╗██║     
  ██║  ███╗██║   ██║█████╗   ██║   ██║   ██║██║   ██║██║     
  ██║   ██║██║   ██║╚════╝   ██║   ██║   ██║██║   ██║██║     
  ╚██████╔╝╚██████╔╝         ██║   ╚██████╔╝╚██████╔╝███████╗
   ╚═════╝  ╚═════╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝

    Go tool for find leaked JS by TRHACKNON
  `
  colorfulBanner(bannerText)
}
