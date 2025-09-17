package gap_buffer

import "testing"

func TestNew(t *testing.T) {
	g := New("hello")
	if len(g.data) != 25 {
		t.Errorf("Error, expected slice of size 25, got %d", len(g.data))
	}
}

func TestString(t *testing.T) {

	g := New("hello")

	if g.String() != "hello" {
		t.Errorf(" Error, String should not inlude a gap, got %s", g.String())
	}

	g.MoveCursor(3)

	if g.String() != "hello" {
		t.Errorf(" Error, String should not inlude a gap, got %s", g.String())
	}

	g.MoveCursor(0)
	if g.String() != "hello" {
		t.Errorf(" Error, String should not inlude a gap, got %s", g.String())
	}

	g.MoveCursor(5)
	if g.String() != "hello" {
		t.Errorf(" Error, String should not inlude a gap, got %s", g.String())
	}
}

func TestGapBuffer_MoveCursor(t *testing.T) {
	g := New("hello")

	if g.MoveCursor(-1) {
		t.Errorf("Error, cursor move to negative")
	}

	if g.MoveCursor(6) {
		t.Errorf("Error, cursor move after string")
	}

	if !g.MoveCursor(2) {
		t.Errorf("Error, should be able to move to the middle")
	}

	if !g.MoveCursor(0) {
		t.Errorf("Error, should be able to move to the front")
	}

	if !g.MoveCursor(5) {
		t.Errorf("Error, should be able to move to the end")
	}
}

func TestGabBuffer_DeleteSimple(t *testing.T) {
	g := New("hello")

	g.DeleteRune()
	if g.String() != "hell" {
		t.Errorf(" Error, should be hell, got %s", g.String())
	}

	g.DeleteRune()
	if g.String() != "hel" {
		t.Errorf(" Error, should be hel, got %s", g.String())
	}

	g = New("hello")

	g.MoveCursor(0)
	if g.DeleteRune() {
		t.Errorf("Error: cannot delete from the before nothing")
	}

	g.MoveCursor(1)
	g.DeleteRune()

	if g.String() != "ello" {
		t.Errorf(" Error, should be ello, got %s", g.String())
	}
}

func TestGabBuffer_InsertSimple(t *testing.T) {
	g := New("hello")

	g.InsertRune('!')
	if g.String() != "hello!" {
		t.Errorf(" Error, should be hello!, got %s", g.String())
	}

	g.MoveCursor(0)
	g.InsertRune('!')
	if g.String() != "!hello!" {
		t.Errorf(" Error, should be !hello!, got %s", g.String())
	}
}

func TestGapBuffer_Expansion(t *testing.T) {
	gb := New("1")
	for i := 0; i < 20; i++ {
		gb.InsertRune(rune(i % 10))
	}

	gb.InsertRune('1')
}
