package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"math"
	"os"
	"strings"
)

func main() {
	var fileName string
	var nameInd, teamInd, start, speakerCounter int
	var found bool
	var speakerMap = make(map[string][]string)
	_, err := fmt.Fscan(os.Stdin, &fileName)
	if err != nil {
		panic("TUT CHTO-TO NAEBNULOS: " + err.Error())
	}

	xlFile, err := xlsx.OpenFile(fileName)

	if err != nil {
		panic("TUT CHTO-TO NAEBNULOS: " + err.Error())
	}

	for rowInd, row := range xlFile.Sheets[0].Rows {
		for cellInd, cell := range row.Cells {
			text := cell.String()
			if strings.TrimSpace(strings.ToLower(text)) == "команда" || strings.TrimSpace(strings.ToLower(text)) == "название команды" {
				nameInd = cellInd
				found = true
			} else if strings.TrimSpace(strings.ToLower(text)) == "имя" ||
				strings.TrimSpace(strings.ToLower(text)) == "полное имя" ||
				strings.TrimSpace(strings.ToLower(text)) == "ф.и.о." ||
				strings.TrimSpace(strings.ToLower(text)) == "фио" {
				teamInd = cellInd
				found = true
			}
			if found {
				start = rowInd
				break
			}
		}

	}

	for rowInd := start + 1; rowInd < len(xlFile.Sheets[0].Rows); rowInd++ {
		if len(xlFile.Sheets[0].Rows[rowInd].Cells) == 0 {
			break
		}

		name := strings.TrimSpace(xlFile.Sheets[0].Rows[rowInd].Cells[nameInd].String())
		team := strings.TrimSpace(xlFile.Sheets[0].Rows[rowInd].Cells[teamInd].String())

		if len(speakerMap[team]) >= 2 {
			team = ""
		}
		speakerMap[team] = append(speakerMap[team], name)
		speakerCounter++
	}

	if speakerCounter < 4 {
		panic("SLISHKOM MALO NARODU, IDI NAHOOY")
	}

	roomsFloat := float64(speakerCounter) / 8
	roomNum := int(math.Ceil(roomsFloat))

	var rooms = make([][]string, roomNum)

}

//asdasd
func fileWriter(teams []string) {

}
