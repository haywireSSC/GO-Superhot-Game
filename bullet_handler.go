package main

import (
	"github.com/gen2brain/raylib-go/raylib"
  "reflect"
	"golang.org/x/exp/slices"
)

type BulletHandler struct {
  bullets []bullet
}

type bullet interface {
  draw()
  update() bool
  get_inner() Bullet
}

func NewBulletHandler() BulletHandler {
  inst := BulletHandler{}
  return inst
}

func (self *BulletHandler) add_bullet(bullet bullet) {
  self.bullets = append(self.bullets, bullet)
}

func (self *BulletHandler) draw() {
  for _, bullet := range self.bullets {
    bullet.draw()
  }
}

func (self *BulletHandler) update() {
  for i := len(self.bullets)-1; i >= 0; i-- {
    bullet := self.bullets[i]
    if !bullet.update() {
      self.bullets = slices.Delete(self.bullets, i,i+1)
    }
  }
}

func (self *BulletHandler) is_hit(pos rl.Vector2, size float32, check_with reflect.Type) bool {// change to interface
  for _, bullet := range self.bullets {
    bullet_in := bullet.get_inner()
    if reflect.TypeOf(bullet) == check_with && rl.Vector2Distance(bullet_in.pos, pos) < bullet_in.size + size{
      return true
    }
  }
  return false
}
