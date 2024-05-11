HomePage = {
    html = [[
        <h1>Main page</h1>
        </br>
        <a style="background:#1110BB;border-radius:6px;color:#fff;padding:10px;font-weight:bold;" href="/info">Info</a>
    ]],
    render = function (self)
        return {body = self.html, type = "html"} -- Response is a table with "body" and "content-type"
    end
}

InfoPage = {
    html = [[
        <h1>Main page</h1>
        </br>
        <a style="background:#11BB11;border-radius:6px;color:#fff;padding:10px;font-weight:bold;" href="/">Back to main</a>
    ]],
    render = function (self)
        return {body = self.html, type = "html"}
    end
}
