package system

import (
    "github.com/codeskyblue/go-sh"
    "bytes"
    "bufio"
    "strconv"
    "fmt"
)

func GetAllPids() []int {
    out, err := sh.Command("ps", "ax").Command("awk", "{print $1}").Output()
    if err != nil {
        fmt.Print("Failed running command", err.Error())
    }

    reader := bytes.NewReader(out)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)

    pids := make([]int, len(out)+1)
    scanner.Scan() // Skip 'PID'
    for scanner.Scan() {
        pid, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Error parsing integer", err.Error())
        }
        pids = append(pids, pid)
    }
    return pids
}
