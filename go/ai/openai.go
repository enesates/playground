package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/openai"
)

func main() {
    prompt := os.Args[1]

    if prompt == "" {
        fmt.Println("No prompt given")
        return
    }

    ctx := context.Background()
    llm, err := openai.New()
    if err != nil {
        log.Fatal(err)
    }

    llms.WithTemperature(0.0)

    //// FIRST ANSWER ////
    completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(completion)

    //// SECOND ANSWER ////
    fmt.Println("")

    llms.WithTemperature(1.0)
    newPrompt := fmt.Sprintf("%s %s", prompt, " and should happen in the future but in German with Sheakespeare style")

    completion, err = llms.GenerateFromSinglePrompt(ctx, llm, newPrompt)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(completion)
}
