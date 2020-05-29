package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//命令行工具
	//不加参数的话就默认problesm.csv
	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'Q&A'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	//打开文件
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open CSV file : %s\n", *csvFilename))
	}

	//调用csv的读方法
	//返回一个阅读器
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file\n")
	}
	//把文件中的值写成一个切片返回
	problems := parseLines(lines)
	//定义一个定时器,把定时器的时间输入到channel中
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	for i, p := range problems {
		fmt.Printf("problem #%d:%s=\n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		//把管道中的时间输出来
		case <-timer.C:
			fmt.Printf("you scored %d out of %d", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("you scored %d out of %d", correct, len(problems))
}

//解析文件每一行，返回一个问题切片，把每一行的第一个字符串作为问题，第二个作为答案
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			//去除空格
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
