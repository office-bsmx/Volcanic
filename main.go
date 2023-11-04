package main

import (
	"errors"
	"strings"
	"volcanic/internal"

	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	if err := L.DoFile("main.lua"); err != nil {
		internal.Throw(2)
	}
	entry_point := L.GetGlobal("Main")
	if entry_point == lua.LNil {
		internal.Throw(3, errors.New("undefined lua entry point"))
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
			return c.SendString("error!")
		}
		trim_path := strings.Trim(c.Path(), "/")
		res_arr := strings.Split(trim_path, "/")
		pass_table := lua.LTable{}
		for _, v := range res_arr {
			pass_table.Append(lua.LString(string(v)))
		}
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("Main"),
			NRet:    1,
			Protect: true,
		}, &pass_table); err != nil {
			internal.Throw(3, err)
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
