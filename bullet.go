package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)


func NewBullet(pos rl.Vector2, target rl.Vector2) Bullet {
  inst := Bullet{}
  inst.pos = pos
  inst.dir = rl.Vector2Normalize( rl.Vector2Subtract(target, pos) )
  inst.speed = 600
  inst.size = 10

  return inst
}

type Bullet struct {
  pos rl.Vector2
  dir rl.Vector2
  speed float32
	colour rl.Color
  size float32
}

func (self *Bullet) draw() {
  rl.DrawCircle(int32(self.pos.X), int32(self.pos.Y), self.size, self.colour)
}

func (self *Bullet) update() bool {
  self.pos.X += self.dir.X * self.speed * delta_time
  self.pos.Y += self.dir.Y * self.speed * delta_time

  return IsOnScreen(self.pos)
}

type PlayerBullet struct {
  bullet Bullet
}

func (self *PlayerBullet) update() bool {
  out := self.bullet.update()
  return out
}

func (self *PlayerBullet) draw() {
  self.bullet.draw()
}
func (self *PlayerBullet) get_inner() Bullet {
  return self.bullet
}

type EnemyBullet struct {
  bullet Bullet
}

func (self *EnemyBullet) update() bool {
  out := self.bullet.update()
  return out
}

func (self *EnemyBullet) draw() {
  self.bullet.draw()
}
func (self *EnemyBullet) get_inner() Bullet {
  return self.bullet
}
