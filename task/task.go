package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Task struct {
	Url     string
	Result  *chan int
	RegRule *regexp.Regexp
}

func (t *Task) Execute() {
	response, err := http.Get(t.Url)
	if err != nil {
		t.ProcessError(err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.ProcessError(err)
		return
	}
	count := len(t.RegRule.FindAll(body, -1))
	//fmt.Println(string(body))
	fmt.Printf("Count for %s: %d\n", t.Url, count)
	*(t.Result) <- count
}

func (t *Task) ProcessError(err error) {
	fmt.Printf("Error for %s: %s\n", t.Url, err.Error())
	*(t.Result) <- 0
}
