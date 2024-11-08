import { is } from '@electron-toolkit/utils';
import { app, BrowserWindow, BrowserWindowConstructorOptions, shell } from 'electron';
import { join } from 'path';
import icon from '../../../../resources/icon.png?asset';

app.enableSandbox();

const electronRenderUrl = process.env['ELECTRON_RENDERER_URL'];

/**
 * 默认构建参数属性
 */
export const defaultWindowConstructorOptions: BrowserWindowConstructorOptions = {
  show: false,
  frame: true,
  transparent: false,
  autoHideMenuBar: true,
  // titleBarStyle: 'hidden',
  resizable: false,
  backgroundColor: '#00000000',
  ...(process.platform === 'linux' ? { icon } : {}),
  webPreferences: {
    preload: join(__dirname, '../preload/index.js'),
    devTools: is.dev,
  },
};

export interface CreateOptions {
  readyToHidden?: boolean;
}

export function createWithDefaultOptions(
  routeName: string,
  options: BrowserWindowConstructorOptions = {},
  createOption: CreateOptions = {},
): BrowserWindow {
  const bw = new BrowserWindow({
    ...defaultWindowConstructorOptions,
    ...options,
  });

  bw.on('ready-to-show', () => {
    if (!createOption.readyToHidden) {
      bw.show();
    } else if (bw.isVisible()) {
      bw.hide();
    }

    if (is.dev) {
      bw.webContents.openDevTools();
    }
  });

  bw.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url);
    return { action: 'deny' };
  });

  if (is.dev && electronRenderUrl) {
    bw.loadURL(electronRenderUrl + '/#/' + routeName);
  } else {
    bw.loadFile(join(__dirname, '../renderer/index.html'), { hash: routeName });
  }

  return bw;
}
