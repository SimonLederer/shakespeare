package shakespeare

import (
	"awesomeProject/file"
	"regexp"
)
func DidShakespeareSayThat(input string) []string {
	numResults := 10
	charPadding := 10
	ans := make([]string, numResults)

	sp, err := file.ReadFile("shakespeare.txt")
	if err!=nil {
		return []string{"A serious error has occurred"}
	}else{
		re,err := regexp.Compile(input)
		if err != nil{
			return []string{"Invalid input"}
		}

		results:=re.FindAllIndex(sp, 10)
		for i,result := range results {
			if result[0] < charPadding{ // Too close to the start
				ans[i] = wordRoundOff("dummy " + string(sp[:result[1] + charPadding]))
			}else if result[1] > len(sp) - charPadding - 1{ //Too close to the end
				ans[i] = wordRoundOff(string(sp[result[0]-charPadding:]) + " dummy") // Dummy to be removed in wordRoundOff
			}else{ //Somewhere in the middle
				ans[i] = wordRoundOff(string(sp[result[0]-charPadding:result[1] + charPadding]))
			}
		}
	}
	return ans
}


func wordRoundOff(input string)(rounded string){
	re:=regexp.MustCompile(" .")
	spaces := re.FindAllStringIndex(input, -1)
	rounded = input[spaces[0][0]+1:spaces[len(spaces)-1][0]]
	return
}

