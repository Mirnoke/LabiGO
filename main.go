package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

// шанс выпадения препядствий
const chance float32 = 0.25

// размер поля
const (
	rows    int = 10
	columns int = 20
)

// создание псевдослучайная последовательность 0 и 1 для создания поля
func getline(n int, chance float32) []byte {
	result := make([]byte, n)
	for i := range result {
		if rand.Float32() > chance {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	// добавление финиша и старта
	perm := rand.Perm(n)
	result[perm[0]] = '2'
	result[perm[1]] = '3'

	return result
}

// создание нового поля на основе констант
func genNewField(rawFieldMask []byte) [rows][columns]rune {
	var field [rows][columns]rune
	fieldMask := string(rawFieldMask)

	// генерируем поле на основе маски
	for i := range field {
		for j := range field[i] {
			switch fieldMask[i*columns+j] {
			case '0':
				field[i][j] = '.' // пустота
			case '1':
				field[i][j] = '■' // астероид
			case '2':
				field[i][j] = '✖' // игрок
			case '3':
				field[i][j] = '⚑' // финиш
			}
		}
	}

	return field
}

/*
* пустота : . : 0
* астероид : ■ : 1
* игрок : ✖ : 2
* финиш : ⚑ : 3
 */

// выводит поле превращая его из списка в красивый построчный вывод
func printOutField(field [rows][columns]rune) {
	for i := range field {
		for j := range field[i] {
			fmt.Printf("%c ", field[i][j])
		}
		fmt.Println()
	}
}

// функция для получения команды от игрока, любой
func getCommand() string {
	var command string

	fmt.Print("Ваши действия: ")
	fmt.Scanln(&command)

	return command
}

// получает позицию игрока или финиша
func getElementPosition(mask []byte, element byte) int {
	pos := -1

	for i, b := range mask {
		if b == element {
			pos = i
			break
		}
	}
	return pos
}

// получает позицию всех указанных элементов
func getManyPositions(mask []byte, element byte) []int {
	var pos []int

	for i, b := range mask {
		if b == element {
			pos = append(pos, i)
		}
	}
	return pos
}

// перемещение игрока
func movePlayer(mask []byte, cmd string) []byte {
	pos := getElementPosition(mask, '2')

	if pos <= -1 {
		return mask
	}

	// вычсиляем новую позицию в зависимости от команды
	cmd = strings.ToLower(cmd)
	newPos := pos
	switch cmd {
	case "d":
		if pos%columns == columns-1 {
			newPos = pos - columns + 1
		} else {
			newPos = pos + 1
		}
	case "a":
		if pos%columns == 0 {
			newPos = pos + columns - 1
		} else {
			newPos = pos - 1
		}
	case "w":
		if pos < columns {
			newPos = pos + (rows-1)*columns
		} else {
			newPos = pos - columns
		}
	case "s":
		if pos >= (rows-1)*columns {
			newPos = pos - (rows-1)*columns
		} else {
			newPos = pos + columns
		}
	}

	result := make([]byte, len(mask))
	copy(result, mask)

	result[pos] = '0'
	result[newPos] = '2'

	return result
}

// рассматривает каждый ход игрока, возвращает какрту как результат ходу и требуется ли окночить игру
func makeTurn(fieldMask []byte, input string, finishPos int, asteroids []int) (bool, []byte) {
	end := false

	fieldMask = movePlayer(fieldMask, input)

	playerPos := getElementPosition(fieldMask, '2')

	clearScreen()

	if playerPos == finishPos {
		fmt.Println("Победа!")
		end = true
	} else if slices.Contains(asteroids, playerPos) {
		fmt.Println("Смерть!")
		condition = "menu"
		end = true
	}
	return end, fieldMask
}

// Функции вывода:

// экран меню
func printMenu() {
	clearScreen()
	fmt.Println("=== МЕНЮ ===")
	fmt.Println()
	fmt.Printf("Рекорд: %d\n", highScore)
	fmt.Println()
	fmt.Println("Новая игра:  game")
	fmt.Printf("Сложность:   %s\n", difficulty)
	fmt.Printf("Язык:        %s\n", language)
	fmt.Println()
	fmt.Println("* Используйте help для помощи")
	fmt.Println()
}

// экран помощи
func printHelp() {
	clearScreen()
	fmt.Println("=== ПОМОЩЬ ===")
	fmt.Println()
	fmt.Println("Управление:")
	fmt.Println("  w - вверх")
	fmt.Println("  s - вниз")
	fmt.Println("  a - влево")
	fmt.Println("  d - вправо")
	fmt.Println("  Можно вводить несколько команд подряд, например: ddss")
	fmt.Println()
	fmt.Println("Обозначения на поле:")
	fmt.Println("  .  - пустота")
	fmt.Println("  ■  - астероид (смерть при наступании)")
	fmt.Println("  ✖  - игрок")
	fmt.Println("  ⚑  - финиш")
	fmt.Println()
	fmt.Println("Команды меню:")
	fmt.Println("  game   - начать новую игру")
	fmt.Println("  help   - эта справка")
	fmt.Println("  menu   - вернуться в меню")
	fmt.Println("  diff   - сменить сложность (easy/normal/hard)")
	fmt.Println("  lang   - сменить язык (ru/en)")
	fmt.Println("  exit   - выйти из игры")
	fmt.Println()
}

// функция чистящая экран
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

/*
* game - состояние запущенной игры
* menu - состояние открытого меню
* help - состояние открытого меню помощи в игре
* exit - состояние окончание работы
 */

var condition string = "menu"

// переменные состояния
var (
	highScore  int    = 0
	difficulty string = "normal"
	language   string = "ru"
)

func main() {
	fmt.Println("Успешный запуск!")

	// переменные для состояния game
	var iteration int // подсчёт итераци
	var input string  // ввод ходов
	var end bool

	var finishPos int   // позиция финиша
	var asteroids []int // позиция астероидов

	// основной цикл
	for condition != "exit" {

		// цикл игры
		if condition == "game" {

			iteration++

			end = false

			fmt.Printf("Итерация №%d", iteration)
			fmt.Println()

			fieldMask := getline(rows*columns, chance)
			field := genNewField(fieldMask)

			asteroids = getManyPositions(fieldMask, '1')
			finishPos = getElementPosition(fieldMask, '3')

			printOutField(field)

			for !end {

				input = getCommand()

				if len(input) == 1 {
					end, fieldMask = makeTurn(fieldMask, input, finishPos, asteroids)
				} else {
					for _, iterInp := range input {
						if end {
							break
						}
						end, fieldMask = makeTurn(fieldMask, string(iterInp), finishPos, asteroids)
					}
				}

				if end {
					highScore = iteration
					continue
				}

				fmt.Printf("Итерация №%d", iteration)
				fmt.Println()
				field := genNewField(fieldMask)
				printOutField(field)
			}
		}

		if condition == "menu" {

			// переменные которые должны вернуться к изначальному значению
			iteration = 0

			printMenu()
			input = getCommand()
			clearScreen()

			switch input {
			case "exit":
				condition = "exit"
			case "help":
				condition = "help"
			case "game":
				condition = "game"
			default:
				fmt.Println("Неизвестная команда\nhelp - для списка команд")
			}

		}

		if condition == "help" {

			printHelp()
			input = getCommand()
			clearScreen()

			switch input {
			case "exit":
				condition = "exit"
			case "menu":
				condition = "menu"
			default:
				fmt.Println("Неизвестная команда\nhelp - для списка команд")
			}
		}
	}
}
