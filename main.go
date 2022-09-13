package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	player Player
	delta_time float32
	enemy_handler EnemyHandler
	bullet_handler BulletHandler
	speed_mult float32
	min_speed_mult float32
)

func main() {
	rl.InitWindow(500, 500, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	min_speed_mult = 0.05
	player = NewPlayer()
	enemy_handler = NewEnemyHandler()
	bullet_handler = NewBulletHandler()


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		delta_time = rl.GetFrameTime() * speed_mult

		rl.ClearBackground(rl.RayWhite)

		enemy_handler.update()
		enemy_handler.draw()

		bullet_handler.update()
		bullet_handler.draw()

		if !player.update(){
			enemy_handler.reset_wave()
		}
		player.draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
