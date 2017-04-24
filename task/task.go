package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Задание на поиск по Url количества строк, попадающих
//под правило RegRule. Результат возвращается в канал Result
type Task struct {
	Url     string
	Result  *chan int
	RegRule *regexp.Regexp
}

//Процесс подсчета количества 'Go'
func (t *Task) Execute() {
	response, err := http.Get(t.Url)
	if err != nil {
		t.processError(err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.processError(err)
		return
	}
	count := len(t.RegRule.FindAll(body, -1))
	//fmt.Println(string(body))
	fmt.Printf("Count for %s: %d\n", t.Url, count)
	*(t.Result) <- count
}

//Обработка ошибки
func (t *Task) processError(err error) {
	fmt.Printf("Error for %s: %s\n", t.Url, err.Error())
	*(t.Result) <- 0
}
