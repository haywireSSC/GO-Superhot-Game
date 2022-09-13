package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"fmt"
	"reflect"
)

func NewPlayer() Player {
  inst := Player{}
  inst.speed = 200
	inst.colour = rl.Color{0,0,255,255}
	inst.size = 20
	inst.pos.X = float32(rl.GetScreenWidth())/2
	inst.pos.Y = float32(rl.GetScreenHeight())/2

  return inst
}

type Player struct {
  pos rl.Vector2
  speed float32
	colour rl.Color
	size float32
	speed_speed float32
}

func (self *Player) draw() {
  rl.DrawCircle(int32(self.pos.X), int32(self.pos.Y), self.size, self.colour)
}

func (self *Player) update() bool{
  var delta rl.Vector2
  if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown('D') {
    delta.X += 1
  }
  if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown('A') {
    delta.X -= 1
  }
  if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown('W') {
    delta.Y -= 1
  }
  if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown('S') {
    delta.Y += 1
  }
  self.pos.X += delta.X * self.speed * delta_time
  self.pos.Y += delta.Y * self.speed * delta_time


	defer self.do_speed_change()
	move_amount := rl.Vector2Length(delta)
	if move_amount != 0 {
		self.speed_speed = 0.01
	}else {
		self.speed_speed = (min_speed_mult - speed_mult) * 0.1
		fmt.Println(min_speed_mult, speed_mult)
	}

	if rl.IsMouseButtonPressed(0) {
		self.speed_speed = (10 - speed_mult) * 0.1
		fmt.Println("shot")
	  bullet := NewBullet(self.pos, rl.GetMousePosition())
	  bullet.colour = self.colour
		player_bullet := PlayerBullet{bullet}

		bullet_handler.add_bullet(&player_bullet)
	}
	if bullet_handler.is_hit(self.pos, self.size, reflect.TypeOf(&EnemyBullet{})) {
		return false
	}
	return true
}

func (self *Player) do_speed_change() {
	speed_mult += self.speed_speed
}
