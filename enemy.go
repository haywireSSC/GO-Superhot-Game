package main

import (
	"github.com/gen2brain/raylib-go/raylib"
  "math"
	"math/rand"
	"reflect"
)

func NewEnemy() Enemy {
  inst := Enemy{}
  inst.speed = 100
  inst.pos.X = rand.Float32() * float32(rl.GetScreenWidth())
  inst.pos.Y = rand.Float32() * float32(rl.GetScreenHeight())
  inst.new_target()
	inst.colour = rl.Color{255,0,0,255}
	inst.size = 30
	inst.min_dist = 5
	inst.timer = NewTimer(0.5)

  return inst
}

type Enemy struct {
  pos rl.Vector2
  target rl.Vector2
  speed float32
  min_dist float32
	colour rl.Color
	size float32
	timer Timer
}

func (self *Enemy) new_target() {
  self.target.X = rand.Float32() * float32(rl.GetScreenWidth())
  self.target.Y = rand.Float32() * float32(rl.GetScreenHeight())
}

func (self *Enemy) shoot() {
	bullet := NewBullet(self.pos, player.pos)
	bullet.colour = self.colour
	enemy_bullet := EnemyBullet{bullet}
	bullet_handler.add_bullet(&enemy_bullet)
}

func (self *Enemy) draw() {
  rl.DrawCircle(int32(self.pos.X), int32(self.pos.Y), self.size, self.colour)
}

func (self *Enemy) update() bool {
	if self.timer.trigger_total % 2 == 0 {

		var delta rl.Vector2
		delta = rl.Vector2Normalize( rl.Vector2Subtract(self.target, self.pos) )
		delta.X *= self.speed
		delta.Y *= self.speed

		if math.IsNaN(float64(delta.X)){
			delta.X = 0
		}
		if math.IsNaN(float64(delta.Y)){
			delta.Y = 0
		}

		if rl.Vector2Distance(self.target, self.pos) < self.min_dist{
			self.new_target()
		}
		self.pos.X += delta.X * delta_time
		self.pos.Y += delta.Y * delta_time
	}

	if self.timer.tick(){
		self.shoot()
	}
	if bullet_handler.is_hit(self.pos, self.size, reflect.TypeOf(&PlayerBullet{})) {
		return false
	}

	return true
}
