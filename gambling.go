package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	gaming(30, 60, 7)
}
func gaming(stake int, goal int, trial int){
	totalWin := 0
	totalBets := 0
	for i := 0; i < trial; i++ {
		betMade := 0
		currentStake := stake
		for currentStake > 0 && currentStake < goal {
			betMade++
			rand.Seed(time.Now().UnixNano())
			outcome := rand.Intn(2)
			if outcome == 1{
				currentStake++
			} else {
				currentStake--
			}
		}
		if currentStake == goal{
			totalWin++
		}
		totalBets += betMade
	}
	winPercentage := float64(totalWin) / float64(trial) * 100
	winAvg := float64(totalWin) / float64(trial)
	fmt.Println("Number of wins :", totalWin)
	fmt.Println("Number of bets:", totalBets)
	fmt.Println("Number of win percentage :", winPercentage)
	fmt.Println("Nverage win average : ", winAvg)
}
