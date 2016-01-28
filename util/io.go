package util

import (
  "os"
  "bufio"
)

func ReadFullyFromStdin() string {
  input := ""
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      line := scanner.Text()
      input = input + line
  }

  return input
}
