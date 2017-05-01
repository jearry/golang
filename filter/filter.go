package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	log.SetFlags(0)

	// 处理命令行参数
	algorithm, minSize, maxSize, suffixes, files := handleCommandLine()

	// 开始计算操作
	if algorithm == 1 {
		// 算法1是并行计算, 通过创建各个的goroutine

		// step1: 先通过source函数处理文件列表, 并把处理结果返回到管道里
		// step2: 将符合后缀的文件放到管道里
		// step3: 将符合大小的文件放到管道里
		// step4: 从管道获取结果数据
		sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
	} else {

		// 算法2是串行计算
		channel1 := source(files)
		channel2 := filterSuffixes(suffixes, channel1)
		channel3 := filterSize(minSize, maxSize, channel2)
		sink(channel3)
	}
}

// 命令行参数解析操作  
func handleCommandLine() (algorithm int, minSize, maxSize int64,
	suffixes, files []string) {

	// 将命令行参数绑定到对应的变量中
	// algorithm默认为1
	flag.IntVar(&algorithm, "algorithm", 1, "1 or 2")
	// minSize和maxSize默认为-1, 表示没有限制
	flag.Int64Var(&minSize, "min", -1,
		"minimum file size (-1 means no minimum)")
	flag.Int64Var(&maxSize, "max", -1,
		"maximum file size (-1 means no maximum)")
	// suffixes后缀列表默认为空
	var suffixesOpt *string = flag.String("suffixes", "",
		"comma-separated list of file suffixes")

	// 命令行预处理
	flag.Parse()

	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		// Fatalln is equivalent to Println() followed by a call to os.Exit(1)
		log.Fatalln("minimum size must be < maximum size")
	}

	// 将后缀列表用逗号分隔, 返回suffixes后缀切片
	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}

	// Args returns the non-flag command-line arguments
	// 认为非命令选项的参数全为文件参数
	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

// 创建管道, 处理文件列表并把结果返回到管道里  
func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

// 将符合后缀的文件放到管道里  
// 根据后缀切片处理管道里的文件, 同样再把结果返回到管道里  
// make the buffer the same size as for files to maximize throughput  
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {

			// 没有限制后缀的话, 则直接将文件塞到管道里
			if len(suffixes) == 0 {
				out <- filename
				continue
			}

			// 获取文件列表的后缀, 且全部转换为小写
			// Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot in the final element of path; it is empty if there is no dot
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

// 将符合文件大小的文件放到管道里  
// make the buffer the same size as for files to maximize throughput  
func filterSize(minimum, maximum int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {

			// 对文件大小没有限制, 直接将文件塞到管道里
			if minimum == -1 && maximum == -1 {
				out <- filename // don't do a stat call it not needed
				continue
			}

			// 使用操作系统的接口获取文件大小等信息
			// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError
			/*
					type FileInfo interface {
				    Name() string       // base name of the file
				    Size() int64        // length in bytes for regular files; system-dependent for others
				    Mode() FileMode     // file mode bits
				    ModTime() time.Time // modification time
				    IsDir() bool        // abbreviation for Mode().IsDir()
				    Sys() interface{}   // underlying data source (can return nil)
			    }
			*/
			finfo, err := os.Stat(filename)
			if err != nil {
				continue // ignore files we can't process
			}
			size := finfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) &&
				(maximum == -1 || maximum > -1 && maximum >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

// 从管道获取结果数据  
func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

/* 
output: 
mba:go gerryyang$ ./filter_t -min 1 -suffixes ".cpp" ../c++11/range_for.cpp ../c++11/test ../c++11/test.cpp routines.go 
../c++11/range_for.cpp 
../c++11/test.cpp 
mba:go gerryyang$ ./filter_t -min 1 -max -2 -suffixes ".cpp" ../c++11/range_for.cpp ../c++11/test ../c++11/test.cpp routines.go jjjj  
minimum size must be < maximum size 
*/  