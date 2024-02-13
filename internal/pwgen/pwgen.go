package pwgen

import (
	"math/rand"
	"strings"
)

func randRuneListFromRuneList(runeList []rune, length int) []rune {
	output := []rune{}

	for i := 0; i < length; i++ {
		randIndex := rand.Intn(len(runeList))

		output = append(output, runeList[randIndex])
	}

	return output
}

func RandPW() string {
	nums := []rune("0123456789")
	alpha := []rune("abcdefghijklmnopqrstuvwxyz")
	alphaUpper := []rune(strings.ToUpper(string(alpha)))
	special := []rune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	segmentCount := 32

	randList := []rune{}
	randList = append(randList, randRuneListFromRuneList(nums, segmentCount)...)
	randList = append(randList, randRuneListFromRuneList(alpha, segmentCount)...)
	randList = append(randList, randRuneListFromRuneList(alphaUpper, segmentCount)...)
	randList = append(randList, randRuneListFromRuneList(special, segmentCount)...)

	rand.Shuffle(len(randList), func(i int, j int) {
		randList[i], randList[j] = randList[j], randList[i]
	})

	return string(randList)
}
