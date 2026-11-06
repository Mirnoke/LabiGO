# LabiGo

A simple console maze game I created to teach myself Go. / Простенькая консольная игра-лабиринт созданная мною для самостоятельного изучения golang.

## Table of Contents / Оглавление

- [About the Game](#about-the-game) / [Об игре](#об-игре)
- [Features](#features) / [Возможности](#возможности)
- [Requirements](#requirements) / [Требования](#требования)
- [Installation](#installation) / [Установка](#установка)
- [Menu Commands](#menu-commands) / [Команды меню](#команды-меню)
- [Controls](#controls) / [Управление](#управление)
- [Field Legend](#field-legend) / [Обозначения на поле](#обозначения-на-поле)

## About the Game

**LabiGo** is a console game where you control a player (`✖`) on a field littered with asteroids (`■`). The goal is to reach the finish flag (`⚑`) without stepping on an asteroid. The field is randomly generated for each new game.

Written in Go, using only the standard library.

## Features

- Procedurally generated field
- Random placement of the player, finish, and asteroids
- Turn-based movement with wrap-around edges
- Multiple commands per turn (e.g., `wasd`)

## Requirements

- Go 1.22 or newer

## Installation

Run directly:

    go run main.go

Or build a binary:

    go build -o labigo main.go
    ./labigo

## Menu Commands

| Command | Description |
|---------|-------------|
| `game` | New game |
| `help` | Help |
| `menu` | Return to menu |
| `diff` | Change difficulty (easy / normal / hard) |
| `lang` | Change language (ru / en) |
| `exit` | Exit |

> The `diff` and `lang` commands are not implemented yet.

## Controls

| Key | Action |
|-----|--------|
| `w` | Up |
| `s` | Down |
| `a` | Left |
| `d` | Right |

> You can enter several commands in a row

## Field Legend

| Symbol | Meaning |
|--------|---------|
| `.` | Empty space |
| `■` | Asteroid (death) |
| `✖` | Player |
| `⚑` | Finish |

## Об игре

**LabiGo** - консольная игра, в которой вы управляете игроком (`✖`) на поле, усеянном астероидами (`■`). Цель — добраться до финишного флага (`⚑`), не наступив на астероид. Поле генерируется случайным образом при каждой новой игре.

Написана на Go, использует только стандартную библиотеку.

## Возможности

- Процедурно генерируемое поле
- Случайное размещение игрока, финиша и астероидов
- Пошаговое перемещение с переходом через края поля
- Несколько команд за один ход (например, `wasd`)

## Требования

- Go 1.22 или новее

## Установка

Запуск напрямую:

    go run main.go

Или сборка бинарника:

    go build -o labigo main.go
    ./labigo

## Команды меню

| Команда | Описание |
|---------|----------|
| `game` | Новая игра |
| `help` | Справка |
| `menu` | Вернуться в меню |
| `diff` | Сменить сложность (easy / normal / hard) |
| `lang` | Сменить язык (ru / en) |
| `exit` | Выйти |

> Команды `diff` и `lang` пока не реализованны

## Управление

| Клавиша | Действие |
|---------|----------|
| `w` | Вверх |
| `s` | Вниз |
| `a` | Влево |
| `d` | Вправо |

> Можно вводить сразу несколько комманд подряд

## Обозначения на поле

| Символ | Значение |
|--------|----------|
| `.` | Пустота |
| `■` | Астероид (смерть) |
| `✖` | Игрок |
| `⚑` | Финиш |
