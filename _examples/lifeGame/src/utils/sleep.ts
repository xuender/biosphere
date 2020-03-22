/**
 * 睡眠
 * @param millisecond 毫秒数
 */
export async function sleep(millisecond: number) {
  return new Promise(resole => {
    setTimeout(resole, millisecond)
  })
}
