package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"log"
)

//struct
type human struct {
  name string
  age int
  isAdult bool
}

func main() {
	sarawut := human{name: "guy", age: 103}
  	sarawut.printInfo()

	homie := "sarawut"
	bakes := "kodkaew"
	printFullName(homie, bakes)

    // ใช้ & แทนการอ้างถึง reference
    setAdult(&sarawut)
    fmt.Println(sarawut)

	// var message string
	message := "Hello, Go"
	fmt.Println(message)

	names := []string{"guys", "sarawut", "kodkaew"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	// ฟังก์ชันคืนค่าสองค่า ประกาศตัวแปรมารองรับได้พร้อมกันสองตัว
    result, err := divide(5, 2)

    if err != nil {
        os.Exit(1)
    }
    fmt.Println(result)


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, All!")
	})

	app.Get("/guys", func(c *fiber.Ctx) error {
		return c.SendString("Hello, guys 👋!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, test 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}

//function
func printFullName(firstName string, lastName string) {
    fmt.Println(firstName + " " + lastName)
}

// printInfo จะผูกติดกับ human
// ระบุ human เข้าไปก่อนหน้าชื่อเมธอด
func (h human) printInfo() {
  fmt.Println(h.name, h.age)
}

// คืนค่า float และ error ออกไปพร้อมกันจากฟังก์ชัน
func divide(dividend float32, divisor float32) (float32, error) {
    if divisor == 0.0 {
        err := errors.New("Division by zero")
        return 0.0, err
    }

    return dividend / divisor, nil
}

// พอยเตอร์ในภาษา Go ใช้ * แทนการ dereference หรือการถอดเอาค่าที่แท้จริงออกมา
func setAdult(h *human) {
    h.isAdult = h.age >= 18
}