# Утилита sort

Утилита сортирует строки в файле. На вход программе подается файл из несортированными
строками, на выходе - файл с отсортированными строками.

Утилита поддерживает следующие ключи:

* ```-k ``` - указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
* ```-n``` - сортировка по числовому значению 
* ```-r``` - сортировка в обратном порядке
* ```-u``` - не выводить повторяющиеся строки

# Установка:

```bash
git clone https://github.com/bersen66/sort.git
cd sort
go install
```
