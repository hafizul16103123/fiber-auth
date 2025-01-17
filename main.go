package main

import (
	"fiber-app/src/common"
	"fiber-app/src/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
    fmt.Println("Initializing environment...")
    err := common.LoadEnv()
    if err != nil {
        fmt.Println("Error loading environment:", err)
        return err
    }

    fmt.Println("Initializing database...")
    err = common.InitDB()
    if err != nil {
        fmt.Println("Error initializing database:", err)
        return err
    }

    defer func() {
        fmt.Println("Closing database...")
        common.CloseDB()
    }()

    fmt.Println("Creating Fiber app...")
    app := fiber.New()

    fmt.Println("Adding middleware...")
    app.Use(logger.New())
    app.Use(recover.New())
    app.Use(cors.New())

    fmt.Println("Adding routes...")
    router.AddBookGroup(app)

    fmt.Println("Starting server...")
    var port string
    if port = os.Getenv("PORT"); port == "" {
        port = "8080"
    }

    err = app.Listen(":" + port)
    if err != nil {
        fmt.Println("Error starting server:", err)
        return err
    }

    return nil
}



// package main

// import (
//     "github.com/gofiber/fiber/v2"
// )

// func main() {
//     app := fiber.New()

//     // Define a simple route
//     app.Get("/", func(c *fiber.Ctx) error {
//         return c.SendString("Hello, World!")
//     })

//     // Start the server on port 3000
//     app.Listen(":3000")
// }
