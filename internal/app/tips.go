package tips

import (
    "fmt"
    "math/rand"
    "time"
    color "github.com/fatih/color"
)

type Tip struct {
    id int
    tags []string
    level string
    txt string
}

func GetRandomTip() string {
    seed := rand.NewSource(time.Now().UnixNano())
    randomizer := rand.New(seed)
    tipArray := GetTips()

    tip := tipArray[randomizer.Intn(len(tipArray))]

    return tip.txt
}

func code(snippet string) string {
    cyan := color.New(color.FgCyan).SprintFunc()
    return cyan(snippet)
}


func GetTips() []Tip {
    tips := []Tip{
        Tip {
            id: 1,
            tags: []string{"bash", "cut", "builtins", "linux"},
            level: "i",
            txt: fmt.Sprintf("To see all local users on your current machine, execute %s", code("cut -d: -f1 /etc/passwd")),
        },
        Tip {
            id: 2,
            tags: []string{"bash","builtins"},
            level: "i",
            txt: fmt.Sprintf("Execute %s to see which user you are currently logged in as", code("whoami")),
        },
        Tip {
            id: 3,
            tags: []string{"tag1", "tag2"},
            level: "a",
            txt: fmt.Sprintf("Install https://github.com/Russell91/sshrc to copy in an rc file of your own aliases and command anytime you log in to a new server through ssh."),
        },
        Tip {
            id: 4,
            tags: []string{"tag1", "tag2"},
            level: "b",
            txt: fmt.Sprintf("Learn to use the basic shell builtins thoroughly! For example, put %s in your ~/.bash_profile or ~/.bashrc and instead of just the usual contents " +
            "of the passed in directory, you will see (-a) all files, including hidden (starting with .) files, (-l) long format, which is more easily readable and shows more information " +
            ", and (-h) which when used with (-l) will list the file size in human readable units (i.e. Byte, Gigabyte etc.)", code("alias ll='ls -lha")),
        },
        Tip {
            id: 5,
            tags: []string{"tag1", "tag2"},
            level: "b",
            txt: fmt.Sprintf("Execute %s to see the manual entry for that command. For example, %s will explain how to use the %s command", code("man ${command}"),code("man grep"),code("grep")),
        },
    }

    return tips
}
