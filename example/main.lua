
function Main(route) --default entry point
    local a = string.len(route)
    a = a*12/(100-20)^4.2
    return tostring("The result is "..a) --returned value must be always a string
end
