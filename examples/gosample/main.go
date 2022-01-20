package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "os"
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/subosito/gotenv"
)

var maxLoop int

func init() {
    gotenv.Load()

    maxLoopStr := os.Getenv("MAX_LOOP")
    if maxLoopStr == "" {
        maxLoopStr = "1000"
    }

    maxLoop, _ = strconv.Atoi(maxLoopStr)

    log.Printf("MAX_LOOP = %d", maxLoop)
}

type SimpleResult struct {
    Message string `json:"message"`
}

func aHeavyTask() {
    var res uint64 = 1

    for i := 1; i <= maxLoop; i++ {
        res = (res % rand.Uint64()) * (uint64(i) % rand.Uint64())
    }
}

func main() {
    app := fiber.New()
    app.Use(logger.New(logger.Config{
        Format: "[${time}] ${status} ${latency} ${method} ${path}\n",
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Get("/test", func(c *fiber.Ctx) error {
        // a blocking / heavy task here for simulate heavy endpoints
        aHeavyTask()

        response, _ := json.Marshal(SimpleResult{Message: "Success!"})
        return c.Send(response)
    })

    log.Println("Application is running on port 3000")
    app.Listen(":3000")
}
