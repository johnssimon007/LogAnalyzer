package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"regexp"
)

func LogAnalyze(fn string, content string) {

	var pattern = regexp.MustCompile(content)

	r, err := os.Open(fn)
	if err == nil {
		files, err := r.Readdir(0)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range files {
			reads, err := os.Open(fn + "/" + v.Name())
			if err != nil {
				fmt.Println(err)
			}
			sc := bufio.NewScanner(reads)

			for sc.Scan() {
				result := pattern.FindAllString(sc.Text(), -1)
				if len(result) != 0 {

					if result[0] != "" {

						fmt.Printf(string("\033[31m Found %v match in %s \033[34m %s \n"), len(result), fn+v.Name(), result)

					}

				}
			}

			reads.Close()
		}

		r.Close()

	}

}
func regex_file(fn string) {
	read, err := os.Open("regex.conf")
	if err == nil {
		file_read := bufio.NewScanner(read)

		for file_read.Scan() {

			content := file_read.Text()
			var comment = regexp.MustCompile(`^#.*`)
			var check bool = comment.MatchString(content)

			if !check {
				LogAnalyze(fn, content)

			}

		}

		read.Close()
	} else {
		panic(err)
	}

}
func main() {

	ascii := figure.NewColorFigure("Log analyzer", "", "green", true)
	ascii.Print()
	flag.Parse()
	fn := flag.Arg(0)
	if fn == "" {
		fmt.Println("Please specify the Log directory")
		os.Exit(1)
	} else {
		regex_file(fn)

	}
}
