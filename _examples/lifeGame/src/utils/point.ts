/**
 * 判断 target ，在 source 的方向
 * @param source 中心
 * @param target 目标
 * @returns 0: 上; 1: 下; 2: 左; 3: 右
 */
export function direction(source: number[], target: number[]) {
  // 横向
  const v = Math.abs(source[0] - target[0]) > Math.abs(source[1] - target[1]) ? 1 : 0
  if (v) {
    return source[0] > target[0] ? 2 : 3
  } else {
    return source[1] > target[1] ? 0 : 1
  }
}
