package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()

	if err := L.DoFile("main.lua"); err != nil {
		log.Fatal(err)
	}
	entry_point := L.GetGlobal("Main")
	if entry_point == lua.LNil {
		log.Fatal("undefined lua entry point")
	}
	L.Close()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		CaseSensitive:         true,
	})

	app.All("/*", func(c *fiber.Ctx) error {
		L := lua.NewState()
		defer L.Close()
		if err := L.DoFile("main.lua"); err != nil {
			return c.SendStatus(500)
		}
		pass_route := lua.LString(string(c.Path()))
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("Main"),
			NRet:    1,
			Protect: true,
		}, pass_route); err != nil {
			log.Fatal(err)
		}
		ret := L.Get(-1) // returned value
		L.Pop(1)
		if table, ok := ret.(*lua.LTable); ok {
			c.Context().SetContentType(string(table.RawGetString("type").(lua.LString)))
			return c.SendString(string(table.RawGetString("body").(lua.LString)))
		} else {
			return c.SendStatus(500)
		}
	})
	app.Listen(":8000")
}
