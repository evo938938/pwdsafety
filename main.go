/*
 *    Copyright (C) 2020 Edoardo Ottavianelli
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Author:
 *      Edoardo Ottavianelli <edoardott@gmail.com>
 */
package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"github.com/edoardottt/pwd-safety/utils"
)

func main() {
	password := utils.ReadSingleInput("Password")
	CheckPwd(password)
	words1 := utils.ReadWords("Known-pwds/known-pwd.txt")
	score := utils.Grader(words1, password)
	DisplayResult(score)
	if score <= 68 {
		crackTime := utils.CrackTime(password)
		println("A botnet can crack this pwd in " + utils.ShowCrackTime(crackTime))
		SuggestPwd(words1)
	}
}

//DisplayResult : Display the result for a password
func DisplayResult(score float64) {
	scoreRounded := fmt.Sprintf("%.2f", score)
	fmt.Printf("Score:%s/100", scoreRounded)
	print("\n")
	if score <= 35 {
		color.Red("VERY WEAK")
	}
	if score > 35 && score <= 59 {
		color.Red("WEAK")
	}
	if score > 59 && score <= 68 {
		color.Yellow("REASONABLE")
	}
	if score > 68 && score <= 80 {
		color.Green("STRONG")
	}
	if score > 80 {
		color.Green("VERY STRONG")
	}
}

//SuggestPwd : Suggest a new random password
func SuggestPwd(words1 []string) {
	engWords := utils.ReadWords("eng_words_list.txt")
	randomPwd := utils.GenerateRandom(engWords)
	scoreRandomPwd := utils.Grader(words1, randomPwd)
	println("You should use this instead...")
	println(randomPwd)
	DisplayResult(scoreRandomPwd)
}

//CheckPwd : Check if a password is unuseful
func CheckPwd(password string) {
	if len(password) <= 5 {
		println("Hey....Do you know what is password cracking?")
		os.Exit(1)
	}
}
