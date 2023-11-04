require("controller")

local Routes = {
    {
        "/hello/:name",
        index
    }
}

function string.explode(str, delimiter)
    assert(type(str) == "string" and type(delimiter) == "string", "invalid arguments")
    local o = {}
    while true do
        local pos1,pos2 = str:find(delimiter)
        if not pos1 then
            o[#o+1] = str
            break
        end
        o[#o+1],str = str:sub(1,pos1-1),str:sub(pos2+1)
    end
    return o
end

function Match(route)
    for k, v in pairs(Routes) do
        if(v[0]) then
            
        end
      end
end