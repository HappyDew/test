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

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
