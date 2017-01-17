/**
 * Telegram notifier
 * Verri Andriawan <verri@tiduronline.com>
 *
 * http://tiduronline.com
 * 17th, January '17
 ---------------------------------- */
package main

import (
    "fmt"
    "net/http"
    "bytes"
    "net/url"
    "os"
)

var TOKEN     = ""
var BASE_URL  = fmt.Sprintf("https://api.telegram.org/bot%s", TOKEN)


// To get the chat ID, you have to send a chat to your bot. bot will record the chat with the id chat.
// then go to  following link to get recorded some last chats. 
// https://api.telegram.org/bot<BOT_TOKEN>/getUpdates
// --------------------------------------------------------------------
var CHAT_ID   = ""

func notify(msg string) {
    base_url  := fmt.Sprintf("%s/sendMessage", BASE_URL)
    params    := url.Values{}

    params.Set("chat_id", CHAT_ID)
    params.Set("text", msg)

    cl    := http.Client{}
    datas := bytes.NewBufferString(params.Encode())

    request, _ := http.NewRequest("POST", base_url, datas)
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    cl.Do(request)
}

func main() {
    args := os.Args

    if len(args[1:]) > 0 {
        notify(args[1:][0])
        return
    }
}
