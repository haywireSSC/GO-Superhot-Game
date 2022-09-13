package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func IsOnScreen(pos rl.Vector2) bool{
  return (float32(rl.GetScreenWidth()) > pos.X && float32(rl.GetScreenHeight()) > pos.Y && pos.X > 0 && pos.Y > 0)
}

func Lerp(num, target, amount float32) float32 {
  return num + (target - num) * amount
}
