package src

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"strconv"

	"github.com/go-redis/redis"
	"github.com/oyamo/telegram-compiler/config"
)

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	Text     string   `json:"text"`
	Chat     Chat     `json:"chat"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}


type Engine struct {
	RedisClient *redis.Client
}

func NewEngine() (*Engine, error) {
	redisClient,err := RedisClient()
	return &Engine{RedisClient: redisClient},err
}

func ParseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func sendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + config.AUTH_TOKEN + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func(e *Engine) HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	var codeResponse, errCompile = Compile(e, update)
	if errCompile != nil {
		codeResponse = errCompile.Error()
	}

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, codeResponse)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("punchline %s successfuly distributed to chat id %d", codeResponse, update.Message.Chat.Id)
	}
}



func Compile(e *Engine, u *Update) (string, error) {
	// Check request type
	if u.Message.Text == "" {
		return "", errors.New("no text in message")
	}

	if u.Message.Text == "/start" {
		return "Hi, I'm a Telegram bot that compiles code for you. Send me a code snippet and I'll compile it for you. \n\n" +
			"To start using me, send me any of these commands;\n" +
			"/kotlin\n" +
			"/java\n" +
			"/clang\n" +
			"/cpp\n" +
			"/python\n" +
			"/go\n" +
			"/javascript", nil
	}
			

	languageCommands := [...]string {
		"/start",
		"/kotlin",
		"/java",
		"/clang",
		"/cpp",
		"/python",
		"/go",
		"/javascript",

	}

	// check if text is a command


	for _, command := range languageCommands {
		if u.Message.Text == command {
			// update redis and set language
			err := e.RedisClient.Set(strconv.Itoa(u.Message.Chat.Id), command, 0).Err()
			if err != nil {
				return "", err
			}
			return "Language set to " + command[1:], nil
		}
	}

	// check if language is set else exit
	language, err := e.RedisClient.Get(strconv.Itoa(u.Message.Chat.Id)).Result()
	if err != nil {
		return "", errors.New("no language set. Please use one of the following commands: \n/kotlin, \n/java, \n/clang, \n/cpp, \n/python, \n/go, \n/javascript")
	}

	// compile code
	code := u.Message.Text

	var response *Response = nil

	switch language {
	case "/kotlin":
		response, err = Kotlin(code)
	case "/java":
		response, err = Java(code)
	case "/clang":
		response, err = Clang(code)
	case "/cpp":
		response, err = CPlus(code)
	case "/python":
		response, err = Python(code)
	case "/go":
		response, err = Golang(code)
	case "/javascript":
		response, err = Javascript(code)
	default:
		return "", errors.New("no language set. Please use one of the following commands: \n/kotlin, \n/java, \n/clang, \n/cpp, \n/python, \n/go, \n/javascript")
	}

	if err != nil {
		return "", err
	}

	if response == nil {
		return "", errors.New("no response from compiler")
	}

	return response.Body, nil

}	

// parseTelegramRequest handles incoming update from the Telegram web hook
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	if update.UpdateId == 0 {
		log.Printf("invalid update id, got update id = 0")
		return nil, errors.New("invalid update id of 0 indicates failure to parse incoming update")
	}
	return &update, nil
}

