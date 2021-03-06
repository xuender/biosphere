import { Stage } from "./stage";
import { times, includes, random, forEach, chain } from "lodash";
import { sleep } from "./utils/sleep";
import { Life } from "./life";
import { direction } from "./utils/point";

/**
 * 游戏
 */
export class Game {
  stage: Stage
  lifes: Life[] = []
  dna: number[] = []
  constructor(id: string, public width = 100, public height = 100) {
    // 初始化舞台
    this.stage = new Stage(width, height, 1000 / width)
    this.stage.init(id)
    // 初始化生命DNA
    forEach(`
25335515425543225005333303315223415315143515515305
05450500532542524302500554324110523541010530451020
55030204052051155002404404150131323055433224053032
05200205013135335515215535315035335113315035315435
2330421110430115224131413254013111425143425
`, n => {
      if (n.trim()) {
        this.dna.push(parseInt(n))
      }
    })
  }
  init() {
    console.debug('初始化游戏');
    times(2, () => {
      const l = new Life(random(0, this.width - 1), random(0, this.height - 1))
      l.r = random(0, 255)
      l.g = random(0, 255)
      l.b = random(0, 255)
      this.lifes.push(l)

    })
    this.show()
  }
  show() {
    this.stage.clean()
    for (const l of this.lifes) {
      this.stage.show(l.x, l.y, `rgb(${l.r},${l.g},${l.b})`)
    }
  }
  async play(s = 200) {
    for (let i = 0; i < s; i++) {
      await sleep(200)
      for (const l of this.lifes) {
        this.run(l)
      }
      this.show()
    }
  }
  private run(l: Life) {
    const state = this.state(l.x, l.y, l.id)
    let m = this.movement(state)
    // 行动
    switch (m) {
      case 5: // 捡罐头
        if (this.hasLife(l.x, l.y, l.id)) {
          console.log(l.id, '抓到了')
          l.x = random(0, this.width - 1)
          l.y = random(0, this.height - 1)
          const o = this.other(l)
          const c = l.breed(o)
          c.x = random(0, this.width - 1)
          c.y = random(0, this.height - 1)
          this.lifes.push(c)
          // l.break()
          break
        } else {
          console.log(l.id, '错误的操作,没有其他人')
        }
      case 4: // 随机移动
        m = random(0, 3)
        console.debug(l.id, '随机', m)
      case 0:
      case 1:
      case 2:
      case 3: // 上下左右
        if (this.move(m, state, l.x, l.y)) {
          switch (m) {
            case 0:
              l.y--
              break
            case 1:
              l.y++
              break
            case 2:
              l.x--
              break
            case 3:
              l.x++
              break
          }
        } else {
          console.warn(l.id, '撞墙了')
        }
        break;
    }
  }
  /**
   * 是否可以移动
   * @param movement 行动策略
   * @param state 状态
   * @param x
   * @param y
   */
  private move(movement: number, state: number[], x: number, y: number) {
    switch (movement) {
      case 0:
        // 上
        if (state[0] == 2) {
          // 碰壁
          return false
        }
        return true
      case 1:
        // 下
        if (state[1] == 2) {
          // 碰壁
          return false
        }
        return true
      case 2:
        // 左
        if (state[2] == 2) {
          // 碰壁
          return false
        }
        return true
      case 3:
        // 右
        if (state[3] == 2) {
          // 碰壁
          return false
        }
        return true
      default:
        return true
    }
  }
  /**
   * 行动策略
   * @param state 状态
   * @param */
  private movement(state: number[]) {
    let id = 0
    for (let i = 0; i < state.length; i++) {
      const a = state[i]
      if (a > 0) {
        id += (3 ** (4 - i)) * a
      }
    }
    return this.dna[id]
  }
  /**
   * 获取舞台状态
   * @param x
   * @param y
   */
  state(x: number, y: number, id: string) {
    const ret = []
    // 上
    if (y == 0) {
      ret[0] = 2
    } else {
      ret[0] = this.hasLife(x, y - 1, id)
    }
    // 下
    if (y == this.height - 1) {
      ret[1] = 2
    } else {
      ret[1] = this.hasLife(x, y + 1, id)
    }
    // 左
    if (x == 0) {
      ret[2] = 2
    } else {
      ret[2] = this.hasLife(x - 1, y, id)
    }
    // 右
    if (x == this.width - 1) {
      ret[3] = 2
    } else {
      ret[3] = this.hasLife(x + 1, y, id)
    }
    // 中
    ret[4] = this.hasLife(x, y, id)
    // 上下左右有生命
    if (includes(ret, 1)) {
      return ret
    }
    const latest = this.latest(x, y, id)
    for (let i = 0; i < ret.length; i++) {
      if (ret[i] == 1) {
        ret[i] = 0
      }
    }
    // 确定大方向
    ret[direction([x, y], [latest.x, latest.y])] = 1
    return ret
  }
  private other(m: Life) {
    for (const l of this.lifes) {
      if (l.x == m.x && l.y == m.y && l.id != m.id) {
        return l
      }
    }
    return m
  }
  // 判断位置是否有其他生命
  hasLife(x: number, y: number, id: string) {
    for (const l of this.lifes) {
      if (l.x == x && l.y == y && l.id != id) {
        return 1
      }
    }
    return 0
  }
  // 寻找最近的生命
  private latest(x: number, y: number, id: string) {
    return chain(this.lifes)
      .filter(l => l.id != id)
      .sortBy(l => (x - l.x) ** 2 + (y - l.y) ** 2)
      .head()
      .value()
  }
}
