package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    tip := GetRandomTip()

    fmt.Println(tip)
}

func GetRandomTip() string {
    seed := rand.NewSource(time.Now().UnixNano())
    randomizer := rand.New(seed)
    tipArray := GetTips()

    tip := tipArray[randomizer.Intn(len(tipArray))]
    return tip
}

func GetTips() []string {
    var tipArray = []string{
        "To see all local users on your current machine, execute 'cut -d: -f1 /etc/passwd'",
        "Execute 'whoami' to see which user you are currently logged in as",
        "Install https://github.com/Russell91/sshrc to copy in an rc file of your own aliases and command anytime you log in to a new server through ssh.",
        "Learn to use the basic shell builtins thoroughly! For example, put `alias ll='ls -lha'` in your ~/.bash_profile or ~/.bashrc and instead of just the usual contents " +
        "of the passed in directory, you will see (-a) all files, including hidden (starting with .) files, (-l) long format, which is more easily readable and shows more information " +
        ", and (-h) which when used with (-l) will list the file size in human readable units (i.e. Byte, Gigabyte etc.)",
        "Execute `man ${command}` to see the manual entry for that command. For example, `man grep` will explain how to use the `grep` command",
    }

    return tipArray
}
