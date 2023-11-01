//需要使用到的元素
let Box = document.querySelector(".box");
let colorItems = document.querySelectorAll(".RGB > div");

// 按键
let keyValue = [
    "ESC", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "←BackSpace", "~",
    "Tab", "Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "{ [", "} ]", "| \\", "Delete",
    "CapsLock", "A", "S", "D", "F", "G", "H", "J", "K", "L", ": ;", "\" \'", "←Enter", "PgUp",
    "LShift", "Z", "X", "C", "V", "B", "N", "M", "< ,", "> .", "? /", "RShift", "↑", "PgDn",
    "LCtrl", "Win", "LAlt", "Space", "RAlt", "Fn", "RCtrl", "←", "↓", "→"
]

// 按键的添加
for (let i = 0; i < keyValue.length; i++) {
    let items = document.createElement("div");
    Box.append(items);
    items.innerText = keyValue[i];
    items.className = keyValue[i];
}

// RGB灯光效果
for (let i = 0; i < colorItems.length; i++) {
    colorItems[i].onclick = function () {
        Box.style.color = this.className + "";
        Box.style.textShadow = "0 0 10px " + this.className;
        const LightModel = this.id;
        const R = document.getElementById("rgb_r").value;
        const G = document.getElementById("rgb_g").value;
        const B = document.getElementById("rgb_b").value;
        const Lightness = document.getElementById("rgb_l").value;
        const xhr = new XMLHttpRequest();
        xhr.open("POST", "http://127.0.0.1:8000/color", true);
        xhr.setRequestHeader('content-type', 'application/json');
        const sendData = {color_type: LightModel, r: R, g: G, b: B, lightness: Lightness};//将用户输入值序列化成字符串\n
        xhr.send(JSON.stringify(sendData));
    }
}
