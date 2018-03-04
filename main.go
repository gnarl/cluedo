package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "strconv"
)

//TODO do I need to use runes at all?

   const (
       NOT_ASKED = iota
       SEEN
       NOT_SEEN
   )

   type gameState struct {
       suspects map[rune]uint
       weapons map[rune]uint
       rooms map[rune]uint
   }

func main() {

    var state gameState
    state.init()

    // 3 cards placed in packet
    // 18 cards remaining to be dealt. up to 4 players. 2 players have extra cards


    // first input is n of suggestions
    // five cards you are dealt 'A' .. 'U'
    // remaining  lines contain one suggestion per line
    // 1
    // B I P C F
    // A G M - - -
    // 00000000000000
    // person, weapon, room, response1(player-to-right), response2, response3(upto)
    // '-' no evidence
    // 'S' you see 'S'
    // '*' some evidence was shown

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

func findNotSeen(m map[rune]uint) (key rune){
    for k, v := range m {
        if v == NOT_SEEN {
            key = k
            return
        }
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
            if v != "-"  {
                notSeenFlag = false
                break
            } else if v == "*" {
                
            }
        }
        if notSeenFlag {
            s.update(suspect, NOT_SEEN)
            s.update(weapon, NOT_SEEN)
            s.update(room, NOT_SEEN)
        }
    }



}

func (s* gameState)update(card rune, state uint) {
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

func initMap(start rune, finish rune) map[rune]uint {
    newMap := map[rune]uint{}
    for i := start; i < finish + 1; i++ {
        newMap[rune(i)] = NOT_ASKED
    }
    return newMap
}

func printMaps(m map[rune]uint) {
    for k, v := range m{
        fmt.Printf("key=%s   val=%d\n", k, v)
    }
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

