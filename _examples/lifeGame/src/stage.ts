/**
 * 舞台
 */
export class Stage {
  canvas: HTMLCanvasElement
  ctx: CanvasRenderingContext2D

  constructor(
    private width = 100,
    private height = 100,
    private size = 10,
  ) {
  }
  init(id: string) {
    const div = document.getElementById(id);
    if (!div) {
      throw `找不到id: ${id}，无法创建画板`
    }
    for (const node of div.childNodes) {
      div.removeChild(node)
    }
    this.canvas = document.createElement('canvas');
    this.canvas.setAttribute('class', 'stage')
    this.canvas.width = this.width * this.size;
    this.canvas.height = this.height * this.size;
    // this.canvas.style.position = "absolute";
    div.appendChild(this.canvas);
    this.ctx = this.canvas.getContext('2d');
  }
  show(x: number, y: number, color = 'chocolate') {
    this.ctx.fillStyle = color;
    this.ctx.fillRect(x * this.size, y * this.size, this.size, this.size)
  }
  clean() {
    this.ctx.fillStyle = 'white'
    this.ctx.fillRect(0, 0, this.width * this.size, this.height * this.size);
  }
}
