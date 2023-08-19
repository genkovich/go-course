package orange

type SizeCounter struct {
	Small  int
	Medium int
	Large  int
}

func NewSizeCounter() *SizeCounter {
	return &SizeCounter{}
}

func (o *SizeCounter) ClassifyAndCount(size int) {
	switch {
	case size < 5:
		o.Small++
	case size < 10:
		o.Medium++
	default:
		o.Large++
	}
}
