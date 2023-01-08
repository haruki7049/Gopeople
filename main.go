package main

import (
        "fmt"
        "os"
        "path/filepath"
)

const contactsPath = "/home/haruki/program_folder/Gopeople/people/"

func main() {
        if len(os.Args) <= 1{
                fmt.Println("no argments")
                os.Exit(1)
        }
        command := os.Args[1]
        switch command {
        case "ls":
                runListCommand()
                os.Exit(0)
        case "cat":
                runCatCommand()
                os.Exit(0)
        default:
                fmt.Printf("unknown command: %s\n", command)
                os.Exit(1)
        }
}

func runListCommand() {

}

func runCatCommand() {
        if len(os.Args) <= 2{
                fmt.Println("low commands, you should add a sub command after cat.")
                os.Exit(1)
        }
        subCommand := os.Args[2]
        switch subCommand{
        case "haruki":
                catPeople("haruki")
        default:
                fmt.Printf("unknown people: %s\n", subCommand)
                os.Exit(1)
        }
}

func checkIfPeopleExists(people string){
        //人物を引数に取り、その人物がデータに存在するかチェックする
        err := filepath.Walk(contactsPath, func(path string, info os.FileInfo, err error) error{
        })

        if err != nil{
                panic(err)
        }
}
