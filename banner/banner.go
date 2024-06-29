package banner

import (
  "fmt"
  "time"
)

func colorfulBanner(text string) {
  colors := []string{
    "\033[91m", // Red
    "\033[92m", // Green
    "\033[93m", // Yellow
    "\033[94m", // Blue
    "\033[95m", // Magenta
    "\033[96m", // Cyan
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
