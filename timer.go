package main

import (
	//"github.com/gen2brain/raylib-go/raylib"
  "math/rand"
)

type Timer struct {
  trigger_gap float32
  time float32
  trigger_total int
}

func NewTimer(trigger_gap float32) Timer {
  inst := Timer{}
  inst.trigger_gap = trigger_gap
  inst.time = rand.Float32()*trigger_gap

  return inst
}

func (self *Timer) tick() bool {
  self.time += delta_time
  if self.time > self.trigger_gap {
    self.trigger_total += 1
    self.time = 0
    return true
  }
  return false
}
