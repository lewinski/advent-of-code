package main

import (
	"container/heap"
	"fmt"
)

type fightState struct {
	playerHp, playerArmor, playerMana       int
	bossHp, bossDamage                      int
	manaSpent                               int
	shieldTimer, poisonTimer, rechargeTimer int
}

func (f *fightState) applyEffects() {
	f.playerArmor = 0
	if f.shieldTimer > 0 {
		f.playerArmor = 7
		f.shieldTimer--
	}
	if f.poisonTimer > 0 {
		f.bossHp -= 3
		f.poisonTimer--
	}
	if f.rechargeTimer > 0 {
		f.playerMana += 101
		f.rechargeTimer--
	}
}

func (f *fightState) bossAttack() {
	damage := f.bossDamage - f.playerArmor
	if damage < 1 {
		damage = 1
	}
	f.playerHp -= damage
}

func (f *fightState) done() (done bool, win bool) {
	if f.bossHp <= 0 {
		done = true
		win = true
	}
	if f.playerHp <= 0 {
		done = true
		win = false
	}
	return
}

type fightHeap []*fightState

func (h fightHeap) Len() int           { return len(h) }
func (h fightHeap) Less(i, j int) bool { return h[i].manaSpent < h[j].manaSpent }
func (h fightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *fightHeap) Push(x any) {
	*h = append(*h, x.(*fightState))
}

func (h *fightHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var state fightState
	state.playerHp = 50
	state.playerMana = 500
	state.bossHp = 51
	state.bossDamage = 9

	fmt.Println("part1:", simulateFight(state, false))
	fmt.Println("part2:", simulateFight(state, true))
}

func simulateFight(state fightState, hard bool) int {
	spells := []func(fightState) *fightState{
		castMagicMissile,
		castDrain,
		castShield,
		castPoison,
		castRecharge,
	}

	h := &fightHeap{&state}
	heap.Init(h)

	for h.Len() > 0 {
		state := heap.Pop(h).(*fightState)

		// player turn
		if hard {
			state.playerHp--
			done, win := state.done()
			if done {
				if win {
					return state.manaSpent
				}
				continue
			}
		}

		state.applyEffects()
		done, win := state.done()
		if done {
			if win {
				return state.manaSpent
			}
			continue
		}

		for s := range spells {
			newState := spells[s](*state)
			if newState == nil {
				continue
			}

			newState.applyEffects()
			done, win := newState.done()
			if done {
				if win {
					return newState.manaSpent
				}
				continue
			}

			newState.bossAttack()
			done, win = newState.done()
			if done {
				if win {
					return newState.manaSpent
				}
				continue
			}

			heap.Push(h, newState)
		}
	}
	panic("unwinnable")
}

func castMagicMissile(state fightState) *fightState {
	cost := 53
	if state.playerMana < cost {
		return nil
	}
	state.playerMana -= cost
	state.manaSpent += cost
	state.bossHp -= 4
	return &state
}

func castDrain(state fightState) *fightState {
	cost := 73
	if state.playerMana < cost {
		return nil
	}
	state.playerMana -= cost
	state.manaSpent += cost
	state.bossHp -= 2
	state.playerHp += 2
	return &state
}

func castShield(state fightState) *fightState {
	cost := 113
	if state.playerMana < cost || state.shieldTimer > 0 {
		return nil
	}
	state.playerMana -= cost
	state.manaSpent += cost
	state.shieldTimer = 6
	return &state
}

func castPoison(state fightState) *fightState {
	cost := 173
	if state.playerMana < cost || state.poisonTimer > 0 {
		return nil
	}
	state.playerMana -= cost
	state.manaSpent += cost
	state.poisonTimer = 6
	return &state
}

func castRecharge(state fightState) *fightState {
	cost := 229
	if state.playerMana < cost || state.rechargeTimer > 0 {
		return nil
	}
	state.playerMana -= cost
	state.manaSpent += cost
	state.rechargeTimer = 5
	return &state
}
