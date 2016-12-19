package main

import (
	"fmt"
	"log"
	"os"
  "strings"
	"strconv"
	"net/http"
)
import _ "github.com/joho/godotenv/autoload"

 func chatbotProcess(session Session, message string ) (string, error) {
	 
	 gx:= Team{}
	 
	 session["x"] = append(session["x"],"0")

	 session["team"] = append(session["team"],"")
	 team1 := session["team"]
	 
	 
	 for _, v := range team1 {
	 if strings.EqualFold(v,message) {
 		return "", fmt.Errorf("I am sorry you wrote the team name again %s!", message)
 	}
		 
	 }
	 
	  x1,_ := strconv.ParseInt(session["x"] [0],0,64)
	 
	 if(x1==0){
	 var testlg string = strings.ToLower(message)
	 
	 	
	 
	 switch testlg {
	case "premier league","spanish league","italian league","bundesliga":
		x1++ 
		 i := int(x1)
 		 session["x"][0]= strconv.Itoa(i)
			return fmt.Sprintf(" your favourite League is %s . What's your favourite team?" , message),nil
		 
	 default : return "", fmt.Errorf("This is not a league!")
	 }
	 }
	 
	 	client := NewClient(http.DefaultClient)

	// Tell it to use our API token
	client.AuthToken = "c11f9c4454e542d9806f2b2242083b0c"

	// Get team info ...
	
	var fc Team
	 
	 
	 var test string = strings.ToLower(message)
	 switch test {
		 
  case "manchester united": fc.Id= 66
  case "as roma": fc.Id  =100
	case "borussia dortmund","dortmund": fc.Id  =4 
	case "bayern munchen","bayern": fc.Id  =5
	case "schalke": fc.Id  =6
	case "arsenal": fc.Id  =57 
	case "chelsea": fc.Id  =61
	case "liverpool": fc.Id  =64
	case "manchester city": fc.Id  =65
	case "tottenham spurs","tottenham": fc.Id  =73
	case "atletico madrid": fc.Id  =78
	case "barcelona": fc.Id  =81
	case "real madrid": fc.Id  =86	 
	case "villareal": fc.Id  =94
	case "valencia": fc.Id  =95	 
	case "ac milan": fc.Id  =98
	case "fiorentina": fc.Id  =99	 
	case "inter milan": fc.Id  =108
	case "juventus": fc.Id  =109 
	case "lazio": fc.Id  =110
	case "napoli": fc.Id  =113
		 
  default: fmt.Errorf("No Team with name %s", message)
}
	
	
 	bb := TeamRequest{client.req("teams/%d", fc.Id), fc.Id}
	 
	 cec := PlayerList{}
	var fc1 PlayerList
	fc1.Count = fc.Id

	
 	tet2 := TeamPlayersRequest{client.req("teams/%d/players", fc1.Count)}
	
	cec ,_ = tet2.Do() 
	 
	 
	 jrs := strconv.Itoa(int((cec.Players[0].JerseyNumber)))

	 vplay := cec.Players[0].Name + "  His Position : "+ cec.Players[0].Position + "  And his Jersy is " + jrs + "  ..."
	
	

	gx, _= bb.Do()
	 
	 
	 for _, val:= range session["hany"] {
		if strings.EqualFold(val, gx.Name) {
			return "", fmt.Errorf("You've already entered %s before!", gx.Name)
		}
	}

	 session["hany"] = append(session["hany"], gx.Name)
	// Form a sentence out of the history in the form Message 1, Message 2, and Message 3
	words := session["hany"]
	l := len(words)
	wordsForSentence := make([]string, l)
	copy(wordsForSentence, words)
	if l > 1 {
		wordsForSentence[l-1] = " " + wordsForSentence[l-1]
	}
	sentence := strings.Join(wordsForSentence, " , ")
	 
	 	 message = "Team Name : "+gx.Name + "  , " +"ShortName : "+ gx.ShortName + " , One of its famous players is " + vplay + "You have searched : "+ sentence
	return fmt.Sprintf(" The Details for the team is ==> %s." , message), nil 
	 
	 
	 for _, val:= range session["playSlice"] {
		if strings.EqualFold(val, message) {
			return "", fmt.Errorf("You've already entered %s before!", message)
		}
	}
	 return "",nil
 }



func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}


func main() {
		
	// Uncomment the following lines to customize the chatbot
	WelcomeMessage = "What's your Favorite League?"
	ProcessFunc(chatbotProcess)
	
	
	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(Engage(":" + port))
}
