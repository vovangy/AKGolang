package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// printTree рекурсивно печатает содержимое директории в виде дерева.
func printTree(path string, prefix string, isLast bool, depth int) {
	if depth < 0 {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, entry := range entries {
		if i == len(entries)-1 {
			// Последний элемент
			fmt.Printf("%s└── %s\n", prefix, entry.Name())
			newPrefix := prefix + "    "
			if entry.IsDir() {
				printTree(filepath.Join(path, entry.Name()), newPrefix, true, depth-1)
			}
		} else {
			// Не последний элемент
			fmt.Printf("%s├── %s\n", prefix, entry.Name())
			newPrefix := prefix + "│   "
			if entry.IsDir() {
				printTree(filepath.Join(path, entry.Name()), newPrefix, false, depth-1)
			}
		}
	}
}

func main() {
	// Определение флагов
	var depth int
	flag.IntVar(&depth, "n", -1, "Глубина дерева")
	flag.Parse()

	// Получение пути к директории из позиционных аргументов
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Использование: tree [-n глубина] <путь_к_директории>")
		os.Exit(1)
	}

	path := args[0]

	// Проверка, если путь не абсолютный, то сделать его абсолютным
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}

	// Печать дерева
	printTree(path, "", true, depth)
}
