package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Tip struct {
    id int
    tags []string
    level string
    txt string
}

func main() {
    tip := GetRandomTip()

    fmt.Println(tip)
}

func GetRandomTip() string {
    seed := rand.NewSource(time.Now().UnixNano())
    randomizer := rand.New(seed)
    tipArray := GetTips()

    tip := tipArray[randomizer.Intn(len(tipArray))]

    return tip.txt
}

func GetTips() []Tip {
    tips := []Tip{
        Tip {
            id: 1,
            tags: []string{"bash", "cut", "builtins", "linux"},
            level: "i",
            txt: "To see all local users on your current machine, execute 'cut -d: -f1 /etc/passwd'",
        },
        Tip {
            id: 2,
            tags: []string{"bash","builtins"},
            level: "i",
            txt: "Execute 'whoami' to see which user you are currently logged in as",
        },
        Tip {
            id: 3,
            tags: []string{"utility"},
            level: "a",
            txt: "Install https://github.com/Russell91/sshrc to copy in an rc file of your own aliases and command anytime you log in to a new server through ssh.",
        },
        Tip {
            id: 4,
            tags: []string{"bash"},
            level: "b",
            txt: "Learn to use the basic shell builtins thoroughly! For example, put `alias ll='ls -lha'` in your ~/.bash_profile or ~/.bashrc and instead of just the usual contents " +
            "of the passed in directory, you will see (-a) all files, including hidden (starting with .) files, (-l) long format, which is more easily readable and shows more information " +
            ", and (-h) which when used with (-l) will list the file size in human readable units (i.e. Byte, Gigabyte etc.)",
        },
        Tip {
            id: 5,
            tags: []string{"bash"},
            level: "b",
            txt: "Execute `man ${command}` to see the manual entry for that command. For example, `man grep` will explain how to use the `grep` command",
        },
        Tip {
            id: 6,
            tags: []string{"bash", "autocompletion"},
            level: "b",
            txt: "Hit the `tab` key while typing the name of a file, path, or command to use autocompletion. Many third party tools such as git and the aws cli offer their own installable autocompletion tools as well.",
        },
        Tip {
            id: 7,
            tags: []string{"bash", "autocompletion"},
            level: "i",
            txt: "Surrounding a command with either backticks `` or $() will substitute the command with the output of the command. For example, `echo $(ls)` pipe the text listing the contents of your current directory through `echo`",
        },
        Tip {
            id: 8,
            tags: []string{"bash"},
            level: "i",
            txt: "There are many ways to loop through iterable collections of data in Bash. For example, to loop through every item in your current directory, enter: `for i in $( ls ); do echo item: ${i}; done`",
        },
    }

    return tips
}
