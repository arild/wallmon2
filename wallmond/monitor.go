package main


type Monitor struct {
    cpu, ram int
}

func (m *Monitor) getCpu() float64 {
    return 50.0
}
