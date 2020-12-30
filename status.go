package ggping

import "math"

type Status struct {
	data     []float64
	capacity int
	head     int
}

func NewStatus(capacity int) *Status {
	status := &Status{
		[]float64{},
		capacity,
		0,
	}
	return status
}

func (s *Status) Update(value float64) {
	s.data = append(s.data, value)

	if len(s.data) > s.capacity {
		s.head++
	}
}

func (s *Status) GetAll() []float64 {
	return s.data[s.head:]
}

func (s *Status) GetMax() float64 {
	var max float64 = 0
	for _, value := range s.GetAll() {
		max = math.Max(max, value)
	}
	return max
}
