import { uniqueId, random } from 'lodash'
/**
 * 生命
 */
export class Life {
  private ox: number
  private oy: number
  private _x: number
  private _y: number
  r: number
  g: number
  b: number
  constructor(x: number, y: number, public id = uniqueId()) {
    this._x = x
    this._y = y
    this.ox = x
    this.oy = y
  }
  // 繁殖
  breed(l: Life) {
    const c = new Life(l.x, l.y)
    /*
    c.r = l.r
    c.g = this.g
    // c.r = Math.floor((this.r + l.r) / 2)
    // c.g = Math.floor((this.g + l.g) / 2)
    c.b = Math.floor((this.b + l.b) / 2)
    */
    c.r = random(0, 254)
    c.g = random(0, 254)
    c.b = random(0, 254)
    return c
  }
  private bak() {
    this.ox = this._x
    this.oy = this._y
  }
  set x(x: number) {
    this.bak()
    this._x = x
  }
  get x() {
    return this._x
  }
  set y(y: number) {
    this.bak()
    this._y = y
  }
  get y() {
    return this._y
  }
  break() {
    this._x = this.ox
    this._y = this.oy
  }
}
