package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут
быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте
аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в
формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах
(linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*Шелл — это обычная консольная программа, которая будучи
запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный
сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/

func main() {
	printInvite()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		pipes := strings.Split(scanner.Text(), "|")
		output := strings.Builder{}
		var input string

		for _, pipe := range pipes {
			command := strings.Split(strings.Trim(pipe, " "), " ")
			input = output.String()
			output.Reset()

			switch command[0] {
			case "pwd":
				path, _ := filepath.Abs(".")
				output.WriteString(path)

			case "cd":
				err := os.Chdir(command[1])
				if err != nil {
					output.WriteString("Incorrect path")
				}

			case "echo":
				output.WriteString(strings.Join(command[1:], " "))

			case "ps":
				showProcesses(&output)

			case "kill":
				pid, err := strconv.Atoi(command[1])
				if err != nil {
					log.Println(err.Error())
				}
				killProcess(&output, pid)

			case "quit":
				return

			default:
				executeCommand(&output, strings.NewReader(input), command[0], command[1:]...)
			}
		}
		fmt.Println(output.String())
		printInvite()
	}
}

func executeCommand(output *strings.Builder, input *strings.Reader, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = input
	cmd.Stderr = os.Stderr
	cmd.Stdout = output
	err := cmd.Run()
	if err != nil {
		output.WriteString(err.Error())
	}
}

func printInvite() {
	path, _ := filepath.Abs(".")
	fmt.Print(path, " > ")
}

func showProcesses(output *strings.Builder) {
	processes, _ := process.Processes()
	for _, proc := range processes {
		name, _ := proc.Name()
		output.WriteString(fmt.Sprintf("%d %s\n", proc.Pid, name))
	}
}

func killProcess(output *strings.Builder, pid int) {
	prc, err := os.FindProcess(pid)
	if err != nil {
		output.WriteString(err.Error())
	}

	err = prc.Kill()
	if err != nil {
		output.WriteString(err.Error())
	}
}
