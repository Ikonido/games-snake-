package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Константи для розміру ігрового поля
const (
	width  = 20
	height = 20
)

// Структура для змійки
type Snake struct {
	body  []Position
	grow  int
	alive bool
}

// Структура для позиції
type Position struct {
	x, y int
}

// Функція для очищення консолі
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Функція для виведення ігрового поля
func drawField(s *Snake, food Position) {
	clearScreen()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == food.x && y == food.y {
				fmt.Print("@")
			} else {
				isBody := false
				for _, p := range s.body {
					if p.x == x && p.y == y {
						fmt.Print("o")
						isBody = true
						break
					}
				}
				if !isBody {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

// Функція для генерації випадкової їжі
func generateFood(s *Snake) Position {
	var food Position
	for {
		food.x = rand.Intn(width)
		food.y = rand.Intn(height)
		unique := true
		for _, p := range s.body {
			if p.x == food.x && p.y == food.y {
				unique = false
				break
			}
		}
		if unique {
			break
		}
	}
	return food
}

// Функція для керування змійкою
func moveSnake(s *Snake, dir Position) {
	head := s.body[0]
	head.x += dir.x
	head.y += dir.y

	if head.x < 0 || head.x >= width || head.y < 0 || head.y >= height {
		s.alive = false
		return
	}

	for _, p := range s.body[1:] {
		if p.x == head.x && p.y == head.y {
			s.alive = false
			return
		}
	}

	s.body = append([]Position{head}, s.body...)
	if s.grow == 0 {
		s.body = s.body[:len(s.body)-1]
	} else {
		s.grow--
	}
}

func main() {
	// Ініціалізація змійки
	snake := &Snake{
		body:  []Position{{2, 2}, {2, 1}, {2, 0}},
		grow:  0,
		alive: true,
	}

	// Ініціалізація їжі
	food := generateFood(snake)

	// Ініціалізація напрямку руху
	dir := Position{0, 1}

	// Головний цикл гри
	for snake.alive {
		drawField(snake, food)

		// Перевірка, чи з'їла змійка їжу
		head := snake.body[0]
		if head.x == food.x && head.y == food.y {
			snake.grow++
			food = generateFood(snake)
		}

		moveSnake(snake, dir)

		// Отримання вводу від користувача
		fmt.Print("Напрямок руху (w/a/s/d): ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "w":
			dir = Position{0, -1}
		case "a":
			dir = Position{-1, 0}
		case "s":
			dir = Position{0, 1}
		case "d":
			dir = Position{1, 0}
		}

		// Затримка для швидкості гри
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Гра завершена.")
}
