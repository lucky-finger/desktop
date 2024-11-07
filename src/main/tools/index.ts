const _isDevelopment = process.env.NODE_ENV === 'development';

/**
 * 获取当前环境是否为开发环境
 * @returns {boolean} 是/否
 */
export function isDevelopment(): boolean {
  return _isDevelopment;
}
