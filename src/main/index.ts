import { electronApp, optimizer } from '@electron-toolkit/utils';
import { app, BrowserWindow } from 'electron';
import { createWithDefaultOptions } from './tools/windows';

// app.enableSandbox();
function createWindow(): void {
  createWithDefaultOptions('main');
}

app.whenReady().then(() => {
  electronApp.setAppUserModelId('cn.luckyfinger');

  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window);
  });

  createWindow();

  app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});
