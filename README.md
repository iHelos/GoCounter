# GoCounter
## Пример запуска:
```
'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run 1.go
```
## Флаги:
<b>k</b> - 	Максимальное количество горутин-воркеров (default - 5) <br>
<b>b</b> - 	Максимальный размер буфера тасков (default - 128)
```
'https://golang.org\nhttps://golang.org\nhttps://golang.org\nhttp://ihelos.ru' | go run 1.go -k=1 -b=1
```
