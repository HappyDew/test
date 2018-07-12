// package main
//
// import (
// 	"fmt"
// 	"os"
// )
//
// func main() {
// 	dir, err := os.Open(".")
// 	if err != nil {
// 		return
// 	}
// 	defer dir.Close()
//
// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		return
// 	}
// 	for _, fi := range fileInfos {
// 		fmt.Println(fi.Name())
// 	}
// }

// package main
//
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )
//
// func main() {
// 	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
// 		fmt.Println(path)
// 		return nil
// 	})
// }

// package main
//
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
//
// func main() {
// 	m := new(sync.Mutex)
//
// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			m.Lock()
// 			fmt.Println(i, "start")
// 			time.Sleep(time.Second)
// 			fmt.Println(i, "end")
// 			m.Unlock()
// 		}(i)
// 	}
//
// 	var input string
// 	fmt.Scanln(&input)
// }

package main

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type Counter struct {
	dirs  int
	files int
}

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

// func (counter *Counter) output() string {
// 	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
// }

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)
		counter.index(subpath)

		if index == len(names)-1 {
			fmt.Println(prefix+"└──", name)
			tree(counter, subpath, prefix+"    ")
		} else {
			fmt.Println(prefix+"├──", name)
			tree(counter, subpath, prefix+"│   ")
		}
	}
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "."
	}

	counter := new(Counter)
	fmt.Println(directory)

	tree(counter, directory, "")
	//fmt.Println(counter.output())
}
