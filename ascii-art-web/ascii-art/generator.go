package asciiart

import (
	"ascii-art-web/ascii-art/utils"
	"errors"
	"fmt"
	"strings"
)

func Generate(inputStr, bannerName string) (string, error) {
	if bannerName == "" {
		bannerName = "standard"
	}
	// path to banner
	bannerPath := fmt.Sprintf("./ascii-art/banner/%s.txt", bannerName)
	// check the banner
	isValidBanner, err := utils.IsValidBannerFile(bannerPath)
	if err != nil {
		return "", err
	}
	if !isValidBanner {
		return "", errors.New("bannerFile is not valid")
	}
	// read each line from the banner
	contentBanner, err := utils.GetContentBanner(bannerPath)
	if err != nil {
		return "", err
	}
	// delete something that called "/r"
	for i, line := range contentBanner {
		contentBanner[i] = strings.ReplaceAll(line, "\r", "")
	}
	// replace all the new lines in the input real new line
	line := strings.Split(strings.ReplaceAll(inputStr, "\\n", "\n"), "\n")
	var result strings.Builder
	// check for empytiness
	isEmpty := true
	for _, each := range line {
		if each != "" {
			isEmpty = false
			break
		}
	}
	if isEmpty {
		result.WriteString(strings.Repeat("\n", len(line)-1))
		return result.String(), nil
	}
	// main
	for _, textLine := range line {
		if textLine == "" {
			result.WriteString("\n")
			continue
		}
		// 8 lines takes for one draw letter
		for i := 1; i <= 8; i++ {
			for _, char := range textLine {
				if char < 32 || char > 126 {
					continue
				}
				charIndex := (int(char) - 32) * 9
				if charIndex+i < len(contentBanner) {
					result.WriteString(contentBanner[charIndex+i])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run . \"text\" [banner]")
// 		fmt.Println("Example: go run . \"Hello\" standard")
// 		return
// 	}

// 	inputText := os.Args[1]

// 	// validate ASCII input
// 	if !utils.IsValidAsciiInput(inputText) {
// 		fmt.Println("invalid input: not asccii symbol")
// 		os.Exit(1)
// 	}

// 	bannerName := ""
// 	if len(os.Args) > 2 {
// 		bannerName = os.Args[2]
// 	}

// 	asciiArt, err := toAscii(inputText, bannerName)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Print(asciiArt)
// }
