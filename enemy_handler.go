package main

import (
  "golang.org/x/exp/slices"
  "github.com/gen2brain/raylib-go/raylib"
  "fmt"
)

type EnemyHandler struct {
  enemies []*Enemy
  wave int
  font_size int32
}

func NewEnemyHandler() EnemyHandler {
  inst := EnemyHandler{}
  inst.font_size = 128
  return inst
}


func (self *EnemyHandler) spawn_wave() {
  self.wave += 1
  bullet_handler.bullets = []bullet{}
  speed_mult = -min_speed_mult
  for i := 0; i <= self.wave;i ++ {//doesnt need loop
    enemy := NewEnemy()
    self.enemies = append(self.enemies, &enemy)
  }
}

func (self *EnemyHandler) reset_wave() {
  self.wave = 0
  self.enemies = []*Enemy{}
  self.spawn_wave()
}

func (self *EnemyHandler) draw() {
  wave_text := fmt.Sprintf("current wave: %v", self.wave)
  width := rl.MeasureText(wave_text, self.font_size)
  rl.DrawText(wave_text, (int32(rl.GetScreenWidth())-width)/2, (int32(rl.GetScreenHeight())-self.font_size)/2, self.font_size, rl.Color{150,165,170,255})

  for _, enemy := range self.enemies {
    enemy.draw()
  }
}

func (self *EnemyHandler) update() {
  if len(self.enemies) == 0 {
    self.spawn_wave()
  }
  for i := len(self.enemies)-1; i >= 0; i-- {
    enemy := self.enemies[i]
    if !enemy.update() {
      self.enemies = slices.Delete(self.enemies, i,i+1)
    }
  }
}
