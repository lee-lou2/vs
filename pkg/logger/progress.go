package logger

import "fmt"

func ProgressBar(current, total int, rps *int64) {
	const barLength = 50
	progress := float64(current) / float64(total) * float64(barLength)
	// 프로그레스 바를 그림
	fmt.Print("[")
	for i := 0; i < barLength; i++ {
		if float64(i) < progress {
			fmt.Print("■")
		} else {
			fmt.Print("□")
		}
	}
	if (current + 1) == total {
		fmt.Println("]")
	} else {
		rpsString := fmt.Sprintf("%6d", *rps)
		fmt.Print(fmt.Sprint("] ", current+1, "/", total, " | RPS :", rpsString, "/s\r"))
	}
}
