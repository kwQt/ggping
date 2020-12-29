package ggping

type Status struct {
	Max float64

	data     []float64
	capacity int
	head     int
}

func NewStatus(capacity int) *Status {
	status := &Status{
		0,
		[]float64{},
		capacity - 5,
		0,
	}
	return status
}

func (s *Status) Update(value float64) {
	s.data = append(s.data, value)
	if len(s.data) > s.capacity {
		s.head++
	}
	if value > s.Max {
		s.Max = value
	}
}

func (s *Status) GetAll() []float64 {
	return s.data[s.head:]
}
