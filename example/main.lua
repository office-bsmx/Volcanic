
function Main(route)
    local a = string.len(route)
    a = a*12/(100-20)^4.2
    return tostring("Route's length is "..a.." symbols"), 1
end