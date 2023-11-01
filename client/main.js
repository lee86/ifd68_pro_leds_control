const path = require('node:path');
const {Tray, Menu, app, BrowserWindow} = require('electron');

let tray = null;  // 用来存放系统托盘
const createWindow = () => {
    const win = new BrowserWindow({
        width: 1440,
        height: 860,
        show: false,
        icon: "icon.png",
        maximizable: false,
        minimizable: true,
        frame: false,
        resizable: false, //禁止改变主窗口尺寸,设置过后，win.isMaximized()总是返回false
        webPreferences: {
            preload: path.join(__dirname, 'preload.js')
        }
    });
    // 创建系统托盘
    createMenu(win);

    win.loadFile('index.html').then(r => {
        console.log(r)
    })
    // 隐藏顶部菜单
    win.setMenu(null);
    // 资源加载完成后显示
    win.once('ready-to-show', () => {
        win.show();
    });
}
const createMenu = (win) => {
    tray = new Tray(path.join(__dirname, 'icon.png'));
    // 菜单模板
    let menu = [
        {
            label: '显示主界面',
            id: 'show-window',
            enabled: !win.show,
            click() {
                win.show();
            }
        },
        {
            label: '隐藏',
            id: 'hide-windows',
            click() {
                win.hide();
            }
        },
        {
            label: '关于',
            id: 'about'
        },
        {
            label: '退出',
            click() {
                win.close();
                app.quit()
            }
        }
    ];
    // 构建菜单
    menu = Menu.buildFromTemplate(menu);
    // 给系统托盘设置菜单
    tray.setContextMenu(menu);
    // 给托盘图标设置气球提示
    tray.setToolTip('ifd68驱动');
    // 窗口最小化直接缩放到托盘
    // win.on('minimize', ev => {
    //     // 阻止最小化
    //     ev.preventDefault();
    //     // 隐藏窗口
    //     win.hide();
    // });
    win.on('close', ev => {
        ev.preventDefault();
        win.hide();
    });
    // 托盘图标被双击
    tray.on('click', () => {
        win.hide()
    });
    // 托盘图标被双击
    tray.on('double-click', () => {
        win.show()
    });

    // 窗口隐藏
    win.on('hide', () => {
        // 启用菜单的显示主窗口项
        menu.getMenuItemById('show-window').enabled = true;
        // 重新设置系统托盘菜单
        tray.setContextMenu(menu);
    });

    // 窗口显示
    win.on('show', () => {
        // 禁用显示主窗口项
        menu.getMenuItemById('show-window').enabled = true;
        // 重新设置系统托盘菜单
        tray.setContextMenu(menu);
    });
}
app.whenReady().then(() => {
    // ipcMain.handle('ping', () => 'pong')
    createWindow()
    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow()
        }
    })
})

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit()
    }
})