const { app, BrowserWindow } = require('electron')
const path = require('node:path')
const { updateElectronApp } = require('update-electron-app')
updateElectronApp()

const createWindow = () => {
    const win = new BrowserWindow({
      width: 800,
      height: 600,
      webPreferences: {
        nodeIntegration: false,
        contextIsolation: true,
        sandbox: true,
        preload: path.join(__dirname, 'preload.js')
      }
    })
  
    win.loadFile('dist/index.html')
    win.webContents.openDevTools()
  }

  // for Windows and Linux
  app.whenReady().then(() => {
    createWindow()

    // for mac os
    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) createWindow()
      })
  })

  // for windows and linux
  app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') app.quit()
  })
