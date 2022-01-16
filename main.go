package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
	"regexp"
)

func LogAnalyze(fn string, content string , keys reflect.Value) {

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

						fmt.Printf(string("\033[31m Found %v match(%s) in %s \033[34m %s \n"), len(result),keys, fn+v.Name(), result)

					}

				}
			}

			reads.Close()
		}

		r.Close()

	}

}

func regex_file(fn string) {
	m:=make(map[string]bool)
	read, err := os.Open("regex.conf")
	if err == nil {
		file_read := bufio.NewScanner(read)

		for file_read.Scan() {

			content := file_read.Text()
			var comment = regexp.MustCompile(`^#.*`)
			var check bool = comment.MatchString(content)
			if check==true{
        m[content]=true
			}

			if !check {
				keys := reflect.ValueOf(m).MapKeys()
				if(len(keys)==0){
				 continue
				}else
				{
					LogAnalyze(fn, content,keys[len(keys)-1])

				}

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
