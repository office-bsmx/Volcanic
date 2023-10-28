
function Main(route)
    local a = string.len(route)
    a = a*12/(100-20)^4.2
    return tostring("The result is "..a)
end