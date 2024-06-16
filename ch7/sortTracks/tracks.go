package sorttracks

type Track struct {
	Title  string
	Author string
	Year   int16
}

type CustomSort struct {
	Tracks     []*Track
	CustomLess func(t1, t2 *Track) bool
}

func (s CustomSort) Len() int {
	return len(s.Tracks)
}

func (s CustomSort) Less(i, j int) bool {
	return s.CustomLess(s.Tracks[i], s.Tracks[j])
}

func (s CustomSort) Swap(i, j int) {}

type columnSortFunc func(a, b *Track) bool

type selectionSort struct {
	t                []*Track
	columnsSortFuncs []columnSortFunc
}
