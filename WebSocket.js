(function () {
    if ("WebSocket" in window) {
        var ws = new WebSocket("ws://127.0.0.1:2048/ws")
        ws.onopen = function () {
            ws.send("ping")
        }
        ws.onmessage = function (evt) {
            if (evt.data === "ok") {
                ws.close()
                locatoin.reload()
            } else if (evt.data === "err") {
                ws.close()
                alert("WebSocket 服务器出错了")
            }
        }
        ws.onclose = function () {
            ws = null
            console.log("close")
        }
        ws.onerror = function () {
            ws = null
            alert("WebSocket 出错了")
        }

        setInterval(function () {
            if (ws !== null) {
                ws.send("ping")
            }
        }, 10 * 1000)
    } else {
        alert("您的浏览器不支持 WebSocket")
    }
})()
