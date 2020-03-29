import './css/style.css'
import { Game } from './game'
const game = new Game('game', 40, 40)
game.init()
game.play(50)

// stage.show(3, 6)
// for (let x = 0; x < 100; x++) {
//   for (let y = 0; y < 100; y++) {
//     if (x % 3 && y % 7 && (x + y) % 3) {
//       stage.show(x, y)
//     }
//   }
// }
