## Глава1
### Упражнение 1.1:
Измените программу `echo` так, чтобы она выводила также `os.Args[0]`, имя выполняемой команды.

### Упражнение 1.2:
Измените программу `echo` так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке.

### Упражнение 1.3:
Поэкспериментируйте с измерением разницы времени выполнения потенциально неэффективных версий и версии с применением `strings.Join`. (В разделе 1.6 демонстрируется часть пакета `time`, а в разделе 11.4 — как написать тест производительности для ее систематической оценки.)

### Упражнение 1.4:
Измените программу `dup2` так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.

#### Инструкции по запуску для Упражнения 1.4:

```bash
# Соберите исполняемый файл
go build -o exercise4.exe exercise4.go

# Запустите исполняемый файл без аргументов, чтобы увидеть стандартное поведение программы
.\exercise4.exe

# Запустите исполняемый файл с указанием конкретного входного файла (например, input_for_test.txt)
.\exercise4.exe input_for_test.txt

```

### Упражнение 1.5. 
Измените палитру программы lissajous так, чтобы изображение было зеленого цвета на черном фоне, чтобы быть более похожим на экран осциллографа. Чтобы создать веб-цвет #RRGGBB, воспользуйтесь инструкцией color.RGBA{0xRRj0xGG,0xBB,0xff}, в которой каждая пара шестнадцатеричных цифр представляет яркость красного, зеленого и синего компонентов пикселя.
### Упражнение 1.6. 
Измените программу lissajous так, чтобы она генерировала изображения разных цветов, добавляя в палитру palette больше значений, а затем выводя их путем изменения третьего аргумента функции SetColorIndex некоторым нетривиальным способом
### Упражнение 1.7
Вызов функции io.Copy(dst, src) выполняет чтение из src и запись в dst. Воспользуйтесь ею вместо ioutil.ReadAll для копирования тела ответа в поток os.Stdout без необходимости выделения достаточно большого буфера для хранения всего ответа. Не забудьте проверить, не произошла ли ошибка при вызове io.Copy.

Инструкция по запуску по запуску для Упражнения 1.7:

```bash
# Соберите исполняемый файл
go build -o fetch.exe fetch.go
# Запустите исполняемый файл с указанием конкретного входного url
.\fetch.exe http://google.com
```


### Упражнение 1.8.
Измените программу `fetch` так, чтобы к каждому аргументу URL автоматически добавлялся префикс `http://` в случае отсутствия в нем такового. Можете воспользоваться функцией `strings.HasPrefix`.
```bash
# запустите исполняемый файл, передавая url в качестве аргумента без указания протокола
go run fetch.go google.com google.com
```

### Упражнение 1.9.
Измените программу `fetch` так, чтобы она выводила код состояния HTTP, содержащийся в `resp.Status`.

```bash
go build -o fetch.exe .\fetch.go
.\fetch.exe  example.com 
```

### Упражнение 1.10. 
Найдите веб-сайт, который содержит большое количество данных. Исследуйте работу кеширования путем двукратного запуска fetchall и сравнения времени запросов. Получаете ли вы каждый раз одно и то же содержимое? Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить.
### Упражнение 1.11. 
Выполните fetchall с длинным списком аргументов, таким как образцы, доступные на сайте alexa.com. Как ведет себя программа, когда веб-сайт просто не отвечает? (В разделе 8.9 описан механизм отслеживания таких ситуаций.)

### Упражнение 1.12. 
Измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из URL. Например, URL вида `http://localhost:8000/?cycles=20` устанавливает количество циклов равным 20 вместо значения по умолчанию, равного 5. Используйте функцию `strconv.Atoi` для преобразования строкового параметра в целое число. Просмотреть документацию по данной функции можно с помощью команды `go doc strconv.Atoi`.
`http://localhost:8000/?cycles=5&delay=5&freq=0.2`
![Упражнение со схемой](1Tutorial/excercise12/file_for_readme.png)