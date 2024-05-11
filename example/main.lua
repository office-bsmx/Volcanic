require "views"

function Main(route) --application entrypoint
    --here you can add your own router

    if route == "/" then
        return HomePage:render()
    elseif route == "/info" then
        return InfoPage:render()
    end
end
