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

package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Beautify : If the terminal size is enough, print the label PWD-SAFETY.
func Beautify() {
	minSize := 55
	firstLine := "                  _                  __      _         \n"
	secondLine := " _ ____      ____| |      ___  __ _ / _| ___| |_ _   _ \n"
	thirdLine := "| '_ \\ \\ /\\ / / _` |_____/ __|/ _` | |_ / _ \\ __| | | |\n"
	fourthLine := "| |_) \\ V  V / (_| |_____\\__ \\ (_| |  _|  __/ |_| |_| |\n"
	fifthLine := "| .__/ \\_/\\_/ \\__,_|     |___/\\__,_|_|  \\___|\\__|\\__, |\n"
	sixthLine := "|_|                                              |___/ \n"
	beauty := firstLine + secondLine + thirdLine + fourthLine + fifthLine + sixthLine
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(out)) // words[1] == size
	size, _ := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	if size >= minSize {
		println(beauty)
	}
}

//WrapTable : It wraps some key-value entries inside a CLI table
func WrapTable(title string, input map[string]string) {
	maxValue := 0
	result := ""
	for _, key := range input {
		currentLength := len(key) + len(input[key])
		if currentLength > maxValue {
			maxValue = currentLength
		}
	}
	basicDeco := "-"
	titBar1 := strings.Repeat(basicDeco, maxValue) + "\n"
	titBar := "|"
	if !(int(maxValue/2)-len(title)/2-1 >= 0) {
		maxValue = len(title)
	}
	titBar += strings.Repeat(" ", int(maxValue/2)-int(len(title)/2)-1)
	titBar += title
	titBar += strings.Repeat(" ", int(maxValue/2)-int(len(title)/2)-1)
	titBar += "|\n"
	titBar2 := strings.Repeat(basicDeco, maxValue) + "\n"
	result += result + titBar1 + titBar + titBar2
	for _, elem := range input {
		row := "|"

	}
	fmt.Println(result)
}
