import { is } from '@electron-toolkit/utils';
import { BrowserWindowConstructorOptions } from 'electron';
import { join } from 'path';
import icon from '../../../../resources/icon.png?asset';

export const defaultWindowConstructorOptions: BrowserWindowConstructorOptions = {
  show: false,
  frame: false,
  transparent: true,
  autoHideMenuBar: true,
  titleBarStyle: 'hidden',
  resizable: false,
  backgroundColor: '#00000000',
  ...(process.platform === 'linux' ? { icon } : {}),
  webPreferences: {
    preload: join(__dirname, '../preload/index.js'),
    devTools: is.dev,
    sandbox: false,
  },
};
