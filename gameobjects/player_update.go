package gameobjects

import (
	"fmt"
	"strings"

	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) updateMovement() {
	if raylibIsKeyDown(system.KeyBindings["left"]) {
		p.Velocity.X = -p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["right"]) {
		p.Velocity.X = +p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["up"]) {
		p.Velocity.Y = -p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["down"]) {
		p.Velocity.Y = +p.MoveSpeed
	}
}

func (p *Player) updateAction() {
	if raylib.IsMouseButtonDown(system.KeyBindings["action"]) {
		if p.Ase.CurrentFrame == p.Ase.CurrentAnimation.Start && !p.shot {
			fmt.Println("SHOOT ~ >>>--------|>")
			p.shot = true
		}
		if p.Ase.CurrentFrame == p.Ase.CurrentAnimation.End {
			p.shot = false
		}
	}
}

func (p *Player) updateAnimation() {
	switch {
	case raylib.IsMouseButtonDown(system.KeyBindings["action"]):
		if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "action") {
			break
		} else if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "idle") {
			fix := strings.TrimSuffix(p.Ase.CurrentAnimation.Name, "idle")
			p.Ase.Play(fix + "action")
		} else {
			p.Ase.Play(p.Ase.CurrentAnimation.Name + "action")
		}
	case raylibIsKeyDown(system.KeyBindings["left"]):
		p.Ase.Play("left")
	case raylibIsKeyDown(system.KeyBindings["up"]):
		p.Ase.Play("up")
	case raylibIsKeyDown(system.KeyBindings["down"]):
		p.Ase.Play("down")
	case raylibIsKeyDown(system.KeyBindings["right"]):
		p.Ase.Play("right")
	}
}

func (p *Player) updateIdleAnimation() {
	if p.Velocity.X == 0 && p.Velocity.Y == 0 && !strings.HasSuffix(p.Ase.CurrentAnimation.Name, "idle") && !raylib.IsMouseButtonDown(system.KeyBindings["action"]) {
		if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "action") {
			fix := strings.TrimSuffix(p.Ase.CurrentAnimation.Name, "action")
			p.Ase.Play(fix + "idle")
		} else {
			p.Ase.Play(p.Ase.CurrentAnimation.Name + "idle")
		}
	}
}

func (p *Player) updateActionAnimationSpeed() {
	if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "action") {
		p.Ase.PlaySpeed = p.ShootSpeed
	} else {
		p.Ase.PlaySpeed = 1.0
	}
}
