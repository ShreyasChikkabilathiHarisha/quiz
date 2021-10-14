package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

type Question struct {
	QuestionID     int
	QuestionValue  string
	QuestionAnswer string
	LatestAnswer   string
}

func TakeQuiz(db *sql.DB, reader *bufio.Reader) {
	result, err := db.Query("SELECT * FROM questions")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var question Question
		// for each row, scan the result into our question composite object
		err = result.Scan(&question.QuestionID, &question.QuestionValue, &question.QuestionAnswer, &question.LatestAnswer)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(question.QuestionValue)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimLeft(strings.TrimRight(text, " "), " ")
		if strings.TrimLeft(strings.TrimRight(question.QuestionAnswer, " "), " ") == text {
			fmt.Println("That's correct!")
		} else {
			fmt.Println("Sorry, that's not the correct answer. The correct answer is:", question.QuestionAnswer)
		}

		// save the latest answer
		fmt.Println("Saving your answer")
		query := fmt.Sprintf("UPDATE questions SET latestanswer = '%s' where id = %d;",
			text, question.QuestionID)
		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()
	fmt.Println("Thank you for takeing the quiz!")
}

func AddQuestions(db *sql.DB, reader *bufio.Reader) {
	question := Question{}
	fmt.Println("Please enter the question")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	question.QuestionValue = text

	fmt.Println("Please enter the answer")
	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	question.QuestionAnswer = text

	query := fmt.Sprintf("INSERT INTO questions(question, answer, latestanswer) VALUES ('%s', '%s', '')",
		question.QuestionValue, question.QuestionAnswer)
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The question and answer added:")
	fmt.Println("Question:", question.QuestionValue)
	fmt.Println("Answer:", question.QuestionAnswer)
}

func ShowPreviousAnswers(db *sql.DB) {
	result, err := db.Query("SELECT * FROM questions")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	questionNumber := 1

	for result.Next() {
		var question Question
		// for each row, scan the result into our question composite object
		err = result.Scan(&question.QuestionID, &question.QuestionValue, &question.QuestionAnswer, &question.LatestAnswer)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(questionNumber, ". Question: ", question.QuestionValue)
		fmt.Println(questionNumber, ". Latest answer: ", question.LatestAnswer)
		fmt.Println()
		questionNumber++
	}
}

//Go application entrypoint
func main() {
	//Instantiate a Welcome struct object and pass in some random information.
	//We shall get the name of the user as a query parameter from the URL
	welcome := Welcome{"to the quick Quiz app!", time.Now().Format(time.Stamp)}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/quiz")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter \"1\" to take the quiz, \"2\" to add questions, \"3\" to check the previous answers, \"4\" to start the localhost with a static page and \"quit\" to quit")
	fmt.Println("")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			TakeQuiz(db, reader)
			continue
		}
		if strings.Compare("2", text) == 0 {
			AddQuestions(db, reader)
			continue
		}
		if strings.Compare("3", text) == 0 {
			ShowPreviousAnswers(db)
			continue
		}
		if strings.Compare("4", text) == 0 {
			break
		}
		if strings.Compare("quit", text) == 0 {
			fmt.Println("Bye!")
			return
		}
		fmt.Println("Wrong option, please try again")
	}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	fmt.Println("Listening, Please go to \"localhost:8080\" on your browser")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
