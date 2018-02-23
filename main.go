package main


import (
	"fmt"
	"os"
	"log"
	"time"
	"net/http"
	"github.com/shomali11/slacker"
	"./createIssue"
)

func heartBeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main () {
	token := os.Getenv("SLACK_API_TOKEN")

	bot := slacker.NewClient(token)

	bot.Command("Hello", "I will say hello hello!", func(r *slacker.Request, res slacker.ResponseWriter){
		res.Reply("Hello guy")
	})
	bot.Command("What time is it?", "I will tell you the time", func(r *slacker.Request, res slacker.ResponseWriter){
		now := time.Now()
		res.Reply(now.String())
	})
	bot.Command("Create an issue", "Creates an issue in jira", func(r *slacker.Request, res slacker.ResponseWriter){
		res.Reply("Have you tried restarting the server?")
		res.Reply("Haha only joking")
		res.Reply("Give a short description, the issue type, the accociated project key and the summary of the issue")
		bot.Command("<description> <type> <project> <summary>", "Give a short description, the issue type, the accociated project key and the summary of the issue", func(request *slacker.Request, response slacker.ResponseWriter){
			d := request.Param("description")
			t := request.Param("type")
			p := request.Param("project")
			s := request.Param("summary")
			if d != "" && t != "" && p != "" && s != "" {
				err := createIssue.Create(d, t, p, s)
				if err != nil {
					response.Reply(fmt.Sprintf("ERROR! %s", err))
				} else {
					response.Reply("issue has been created")
				}
			} else {
				response.Reply("You need to give me complete information")
			}
		})
	})
	bot.Command("Who is the coolest?", "I will give my opinion", func(r *slacker.Request, res slacker.ResponseWriter){
		res.Reply("ME!")
		res.Reply("haha")
	})
	bot.Command("haha", "I will tell you how it is", func(r *slacker.Request, res slacker.ResponseWriter){
		res.Reply("You're not funny! :C")
	})
	fmt.Println("Bot is now up and running, your wish is my command master!")
	err := bot.Listen()
	if err != nil {
		log.Fatal(err)
	} 

}