package gap_buffer

// GapBuffer represents a gap buffer. As cursor is moved or the gap is filled
// the gap is automatically resized or shifted over
type GapBuffer struct {
	gapSize, cursor int
	data            []rune
	//maybe keep track of the default gapSize
	debug bool
}

func New(init string, debug bool) *GapBuffer {
	initAsRunes := []rune(init)
	return &GapBuffer{
		gapSize: 20,
		cursor:  len(init), // cursor starts at the end
		data:    append(initAsRunes, make([]rune, 20)...),
		debug:   debug,
	}
}

// MoveCursor changes the location of the gap
// requires O(k) time, where k is the difference between the
// current cursor and p
func (gb *GapBuffer) MoveCursor(p int) (ok bool) {
	// prevent moves before or after the string
	if p < 0 || p > len(gb.data)-gb.gapSize {
		return false
	} else {
		// if p is greater than cursor
		if p > gb.cursor {
			for i := gb.cursor; i < p; i++ {
				gb.data[i] = gb.data[i+gb.gapSize]
			}
		} else {
			for i := gb.cursor; i > p; i-- {
				gb.data[i+gb.gapSize-1] = gb.data[i-1]
			}
		}
		gb.cursor = p
		return true
	}
}

func (gb *GapBuffer) InsertRune(r rune) bool {
	if gb.gapSize == 0 {
		// recreate the gap by shifting runes at cursor out by 20 again
		data := make([]rune, len(gb.data)+20)
		copy(data, gb.data)
		gb.data = data
		oldCursor := gb.cursor
		gb.cursor = len(data) - 20
		gb.gapSize = 20
		gb.MoveCursor(oldCursor)
	}

	gb.data[gb.cursor] = r
	gb.cursor++
	gb.gapSize--
	return true

}

func (gb *GapBuffer) DeleteRune() bool {
	if gb.cursor-1 < 0 {
		return false
	} else {
		// gap gets larger by one and cursor goes left
		gb.cursor--
		gb.gapSize++
		return true
	}
}

// String returns the string version of the GapBuffer without the gap represented
func (gb *GapBuffer) String() string {
	if !gb.debug {
		line := make([]rune, len(gb.data)-gb.gapSize)
		copy(line, gb.data[:gb.cursor])
		copy(line[gb.cursor:], gb.data[gb.cursor+gb.gapSize:])

		return string(line)
	} else {
		line := make([]rune, len(gb.data))
		copy(line, gb.data[:gb.cursor])
		for i := 0; i < gb.gapSize; i++ {
			line[gb.cursor+i] = '_'
		}
		copy(line[gb.cursor+gb.gapSize:], gb.data[gb.cursor+gb.gapSize:])
		return string(line)
	}
}

func (gb *GapBuffer) Cursor() int {
	return gb.cursor
}

func (gb *GapBuffer) Begin() int {
	return 0
}

func (gb *GapBuffer) End() int {
	return len(gb.data) - gb.gapSize
}
