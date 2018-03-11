package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "strconv"
)

   const (
       NOT_ASKED = iota
       SEEN
       NOT_SEEN
   )

   type gameState struct {
       suspects map[rune]int
       weapons map[rune]int
       rooms map[rune]int
   }

func main() {

    var state gameState
    state.init()

    inputPath := os.Args[1]
    fmt.Printf("reading %s \n", inputPath)
    inputs, _ := readLines(inputPath)

    for _, v := range inputs {
        fmt.Printf("line = %s\n", v)
    }

    //process inputs
    numSuggestions, _ := strconv.Atoi(inputs[0])
    cards := strings.Fields(string(inputs[1]))
    var suggestions []string
    for i := 2; i < (2 + numSuggestions); i++ {
        suggestions = append(suggestions, inputs[i])
    }

    for _, v := range cards {
        state.update(rune(v[0]), SEEN)
    }

    state.process(suggestions)
    state.deduction()

}

func (s* gameState)deduction() {
  fmt.Printf("%s%s%s\n",
      string(findNotSeen(s.suspects)),
      string(findNotSeen(s.weapons)),
      string(findNotSeen(s.rooms)))
}

func findNotSeen(m map[rune]int) (key rune){
    for k, v := range m {
        if v == NOT_SEEN {
            return k
        }
    }

    count := 0
    for k, v := range m {
        if v == NOT_ASKED {
                key = k
                count++
        }
    }
    if count == 1 {
        return key
    }

    return '?'
}

func (s* gameState)process(suggestions []string) {
    var suspect rune
    var weapon rune
    var room rune

    for i := 0; i < len(suggestions); i++ {
        suspect = rune(suggestions[i][0])
        weapon = rune(suggestions[i][2])
        room = rune(suggestions[i][4])


        suggestion := strings.Fields(suggestions[i][6:])

        notSeenFlag := true
        for _, v := range suggestion {
            if v == "*" {
                notSeenFlag = false
                if s.suspects[suspect] == NOT_ASKED  {
                    if (s.weapons[weapon] != NOT_ASKED) &&
                        (s.rooms[room] != NOT_ASKED) {
                            s.update(suspect, SEEN)
                } else if s.weapons[weapon] == NOT_ASKED {
                    if s.rooms[room] != NOT_ASKED {
                            s.update(weapon, SEEN)
                        }
                    }
                } else if s.rooms[room] == NOT_ASKED {
                    s.update(room, SEEN)
                }
            } else if v != "-"{
                notSeenFlag = false
                s.update(rune(v[0]), SEEN)
            }
        }
        if notSeenFlag {
            s.update(suspect, NOT_SEEN)
            s.update(weapon, NOT_SEEN)
            s.update(room, NOT_SEEN)
        }
    }



}

func (s* gameState)update(card rune, state int) {
    switch {
    case card < 'G':
        s.suspects[card] = state
    case card <= 'L':
        s.weapons[card] = state
    case card > 'L':
        s.rooms[card] = state
    }
}

func (s* gameState)init() {
    s.suspects = initMap('A', 'F')
    s.weapons = initMap('G', 'L')
    s.rooms = initMap('M', 'U')
}

func initMap(start rune, finish rune) map[rune]int {
    newMap := map[rune]int{}
    for i := start; i < finish + 1; i++ {
        newMap[rune(i)] = NOT_ASKED
    }
    return newMap
}

func readLines(path string) (lines []string, err error) {
    buffer, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    lines = strings.Split(string(buffer), "\n")
    lines = lines[0 : len(lines)-1]

    return lines, err
}
