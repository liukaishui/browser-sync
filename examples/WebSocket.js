(function () {
    if ("WebSocket" in window) {
        var ws = new WebSocket("ws://localhost:2345/ws")
        ws.onopen = function () {
            ws.send("ping")
        }
        ws.onmessage = function (evt) {
            if (evt.data === "ok" || evt.data === "err") {
                ws.close()
                location.reload()
            }
        }
        ws.onclose = function () {
            console.log("WebSocket 已关闭")
        }
        ws.onerror = function () {
            console.log("WebSocket 出错了")
        }

        var ping = setInterval(function () {
            console.log("s")
            if (ws.readyState === 1) {
                ws.send("ping")
            } else {
                clearInterval(ping)
            }
        }, 10 * 1000)
    } else {
        alert("您的浏览器不支持 WebSocket")
    }
})()
