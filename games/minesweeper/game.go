package minesweeper

import (
	"github.com/dragon1672/go-collections/vector"
	"github.com/dragon1672/go-get-games/games/minesweeper/common/events"
	"github.com/dragon1672/go-get-games/games/minesweeper/common/sliceutls"
	"github.com/dragon1672/go-get-games/games/minesweeper/common/vectorutil"
	"github.com/golang/glog"
	"sync"
)

type annotation int64

//goland:noinspection GoUnusedConst
const (
	annotationUnset annotation = iota
	annotationFlag
	annotationMaybe
	annotationSafe
)

func (e annotation) State() CellState {
	switch e {
	case annotationFlag:
		return CellFlag
	case annotationMaybe:
		return CellMaybeBomb
	case annotationSafe:
		return CellSafe
	default:
		return CellUnset
	}
}

type GameStateGenerator interface {
	GenerateBombs(width, height int, discouragedPositions ...vector.IntVec2) map[vector.IntVec2]bool
}

type GameGenerator struct {
	Width, Height int
	BigOpening    bool
	Gen           GameStateGenerator
}

type Game struct {
	width, height int
	bombs         map[vector.IntVec2]bool
	annotations   map[vector.IntVec2]annotation
	annotationsMu sync.RWMutex
	gen           GameStateGenerator
	bigOpening    bool
	revealed      map[vector.IntVec2]CellState
	revealedMu    sync.RWMutex
	ChangeEvent   events.Feed[ChangeEventData] // notably isn't a pointer so each event gets an independent copy
}

func (g *Game) Width() int  { return g.width }
func (g *Game) Height() int { return g.height }
func (g *Game) ValidPos(pos vector.IntVec2) bool {
	return 0 <= pos.X && pos.X < g.Width() &&
		0 <= pos.Y && pos.Y < g.Height()
}
func (g *Game) NumBombs() int { return len(g.bombs) }
func (g *Game) String() string {
	return toString(g)
}
func (g *Game) SnapshotReadonly() ReadOnlyGame {
	return &readonlyCopy{
		width:       g.width,
		height:      g.height,
		bombCount:   len(g.bombs),
		annotations: g.copyAnnotations(),
		revealed:    g.GetAllRevealed(),
	}
}

func (g *Game) HasLost() bool {
	g.revealedMu.RLock()
	defer g.revealedMu.RUnlock()
	for _, v := range g.revealed {
		if v == CellBomb {
			return true
		}
	}
	return false
}

func (g *Game) HasWon() bool {
	g.revealedMu.RLock()
	defer g.revealedMu.RUnlock()
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			pos := vector.Of(x, y)
			if _, revealed := g.revealed[pos]; !revealed && !g.bombs[pos] {
				// still have positions that aren't bombs to reveal
				return false
			}
		}
	}
	return true
}

func (g *Game) GameOver() bool {
	return g.HasLost() || g.HasWon()
}

func (g *Game) Get(pos vector.IntVec2) CellState {
	g.revealedMu.RLock()
	defer g.revealedMu.RUnlock()
	if val, ok := g.revealed[pos]; ok {
		return val
	}
	g.annotationsMu.RLock()
	defer g.annotationsMu.RUnlock()
	if f, ok := g.annotations[pos]; ok {
		return f.State()
	}
	return CellEmpty
}

func (g *Game) Reveal(pos vector.IntVec2) []ChangeEventData {
	g.ensureGen(pos)
	ret := g.silentReveal(pos)
	// Emit after full updated
	for _, val := range ret {
		g.ChangeEvent.Send(val)
	}

	return ret
}

func (g *Game) setFlag(pos vector.IntVec2, a annotation) {
	if current := g.Get(pos); current.Revealed() {
		glog.Warningf("Cannot flag already revealed location at %v. Already a %v.", pos, current)
		return
	}
	if f, ok := g.annotations[pos]; ok && f == a {
		return //  already flagged with same value
	}
	g.annotationsMu.Lock()
	changed := false
	if f, ok := g.annotations[pos]; !ok || f != a {
		changed = true
		if g.annotations == nil {
			g.annotations = make(map[vector.IntVec2]annotation)
		}
		g.annotations[pos] = a
	}
	g.annotationsMu.Unlock()
	if changed {
		g.ChangeEvent.Send(ChangeEventData{Pos: pos, Val: g.Get(pos)})
	}
}

func (g *Game) removeAnnotations(pos vector.IntVec2) {
	if _, ok := g.annotations[pos]; !ok {
		return
	}
	g.annotationsMu.Lock()
	changed := false
	if _, ok := g.annotations[pos]; ok {
		changed = true
		delete(g.annotations, pos)
	}
	g.annotationsMu.Unlock()
	if changed {
		g.ChangeEvent.Send(ChangeEventData{Pos: pos, Val: g.Get(pos)})
	}
}

// SetFlagged will mark position as flagged.
func (g *Game) SetFlagged(pos vector.IntVec2) {
	g.setFlag(pos, annotationFlag)
}

// SetSafe will mark position as safe.
func (g *Game) SetSafe(pos vector.IntVec2) {
	g.setFlag(pos, annotationSafe)
}

// SetMaybe will mark position as maybe.
func (g *Game) SetMaybe(pos vector.IntVec2) {
	g.setFlag(pos, annotationMaybe)
}

// ToggleFlag will toggle the square as flagged, returns true if newly flagged.
func (g *Game) ToggleFlag(pos vector.IntVec2) {
	if _, flagged := g.annotations[pos]; !flagged {
		g.SetFlagged(pos)
	} else {
		g.removeAnnotations(pos)
	}
}

func (g *Game) copyAnnotations() map[vector.IntVec2]annotation {
	// make a copy
	ret := make(map[vector.IntVec2]annotation)
	g.annotationsMu.RLock()
	for key, val := range g.annotations {
		ret[key] = val
	}
	g.annotationsMu.RUnlock()
	return ret
}

func (g *Game) GetAllRevealed() map[vector.IntVec2]CellState {
	// make a copy
	ret := make(map[vector.IntVec2]CellState)
	g.revealedMu.RLock()
	for key, val := range g.revealed {
		ret[key] = val
	}
	g.revealedMu.RUnlock()
	return ret
}

// silentlySet will not emit an event
func (g *Game) silentlySet(pos vector.IntVec2, s CellState) ChangeEventData {
	g.revealedMu.Lock()
	if g.revealed == nil {
		g.revealed = make(map[vector.IntVec2]CellState)
	}
	g.revealed[pos] = s
	g.revealedMu.Unlock()
	g.removeAnnotations(pos)
	return ChangeEventData{pos, s}
}

// silentReveal will appropriately reveal all adjacent squares
func (g *Game) silentReveal(toCheck ...vector.IntVec2) []ChangeEventData {
	var ret []ChangeEventData
	for len(toCheck) > 0 {
		var ok bool
		var pos vector.IntVec2
		toCheck, pos, ok = sliceutls.PopLast(toCheck)
		if !ok {
			glog.Warning("Unexpected error when popping from list, as pop should always exceed ")
			continue // unexpected
		}
		if !g.ValidPos(pos) {
			continue // invalid pos
		}
		if _, revealed := g.revealed[pos]; revealed {
			continue // already revealed or invalid
		}
		state := g.calculateAsState(pos)
		ret = append(ret, g.silentlySet(pos, state))
		if state == CellN0 {
			vectorutil.IterateSurroundingInclusive(pos, func(newPos vector.IntVec2) {
				if g.ValidPos(newPos) {
					// add everything including self and let loop handle
					toCheck = append(toCheck, newPos)
				}
			})
		}
	}
	return ret
}

// ensureGen makes sure there are bombs. Otherwise, noop.
func (g *Game) ensureGen(pos vector.IntVec2) {
	if g.bombs == nil {
		discouragedPositions := []vector.IntVec2{pos}
		if g.bigOpening {
			vectorutil.IterateSurroundingExclusive(pos, func(pos vector.IntVec2) {
				discouragedPositions = append(discouragedPositions, pos)
			})
		}
		g.bombs = g.gen.GenerateBombs(g.width, g.height, discouragedPositions...)
	}
}

// calculateAsState calculates the CellState (readonly)
func (g *Game) calculateAsState(pos vector.IntVec2) CellState {
	if isBomb := g.bombs[pos]; isBomb {
		return CellBomb
	}
	return CellStateFromBombCount(g.calcNum(pos))
}

// calcNum calculates adjacent bombs  (readonly)
func (g *Game) calcNum(pos vector.IntVec2) int {
	g.ensureGen(pos)
	bombCount := 0
	vectorutil.IterateSurroundingExclusive(pos, func(pos vector.IntVec2) {
		if b := g.bombs[pos]; b {
			bombCount++
		}
	})
	return bombCount
}

func MakeFromGenerator(gen *GameGenerator) *Game {
	return &Game{
		width:      gen.Width,
		height:     gen.Height,
		gen:        gen.Gen,
		bigOpening: gen.BigOpening,
	}
}
