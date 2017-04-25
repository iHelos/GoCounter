# GoCounter
## Пример запуска:
Запуск с заранее заданным потоком входных данных:
```
echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run 1.go
```
Интерактивный запуск:
```
go run 1.go
```
Входные URL'ы разделяются переходом на новую строку. <br> Для выхода следует закрыть входной поток<i> (Ctrl + D) </i>.
## Флаги:
<b>k</b> - 	Максимальное количество горутин-воркеров (default - 5) <br>
<b>b</b> - 	Максимальный размер буфера тасков (default - 128)
```
echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org\nhttp://ihelos.ru' | go run 1.go -k=1 -b=1
```
