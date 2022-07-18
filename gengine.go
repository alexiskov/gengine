package gengine

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type MathExpression struct {
	Proection, Result string
	num1, num2        int
	Contraulation     bool
}

//Some Numbers & Expression Method-Generate...
func (gameExpression *MathExpression) gameExpressionGenerate() {
	rand.Seed(time.Now().UnixNano())
	gameExpression.num1 = rand.Intn(100)
	gameExpression.num2 = rand.Intn(100)
	randomizer := rand.Intn(24)

	if gameExpression.num1 > gameExpression.num2 {
		if randomizer*2 < gameExpression.num1 && gameExpression.num2 != 0 {
			gameExpression.Proection = "Чему равно выражение: " + strconv.Itoa(gameExpression.num1) + " / " + strconv.Itoa(gameExpression.num2) + " = ? (округлить до целых)"
			gameExpression.Result = strconv.Itoa(gameExpression.num1 / gameExpression.num2)
		} else {
			gameExpression.Proection = "Чему равно выражение: " + strconv.Itoa(gameExpression.num1) + " - " + strconv.Itoa(gameExpression.num2) + " = ? "
			gameExpression.Result = strconv.Itoa(gameExpression.num1 - gameExpression.num2)
		}
	} else {
		if randomizer*2 < gameExpression.num1 {
			gameExpression.Proection = "Чему равно выражение: " + strconv.Itoa(gameExpression.num1) + " + " + strconv.Itoa(gameExpression.num2) + " = ? "
			gameExpression.Result = strconv.Itoa(gameExpression.num1 + gameExpression.num2)
		} else {
			gameExpression.Proection = "Чему равно выражение: " + strconv.Itoa(gameExpression.num1) + " * " + strconv.Itoa(gameExpression.num2) + " = ? "
			gameExpression.Result = strconv.Itoa(gameExpression.num1 * gameExpression.num2)
		}
	}
}

//Game Engine...
func GameEngine(ch chan string) {
	randExpression := MathExpression{}
	for {
		answer := <-ch
		if !randExpression.Contraulation {
			randExpression.gameExpressionGenerate()
			randExpression.Contraulation = true
		}

		if strings.Trim(answer, " \n\r") == randExpression.Result {
			ch <- "Победа!"
			randExpression.Contraulation = false
		} else {
			ch <- randExpression.Proection
		}
	}
}
