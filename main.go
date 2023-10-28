package main

import (
	"os"
	"volcanic/internal"

	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	if leng := len(os.Args); leng != 2 {
		internal.Throw(1)
	}
	main_path := os.Args[1]
	L := lua.NewState()
	if err := L.DoFile(main_path + "/main.lua"); err != nil {
		internal.Throw(2)
	}
	entry_point := L.GetGlobal("Main")
	if entry_point == lua.LNil {
		internal.Throw(3)
	}
	L.Close()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		CaseSensitive:         true,
	})
	app.All("/*", func(c *fiber.Ctx) error {
		L := lua.NewState()
		defer L.Close()
		if err := L.DoFile(main_path + "/main.lua"); err != nil {
			return c.SendString("error!")
		}
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("Main"),
			NRet:    1,
			Protect: true,
		}, lua.LString(c.Path())); err != nil {
			internal.Throw(3)
		}
		ret := L.Get(-1) // returned value
		L.Pop(1)
		if str, ok := ret.(lua.LString); ok {
			return c.SendString(string(str))
		} else {
			return c.SendString("error!")
		}
	})
	internal.Startup(app)
}
