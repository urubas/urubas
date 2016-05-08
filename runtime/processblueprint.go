package runtime

type ProcessBlueprint struct {
	InputMap [][]ThreadVertex
	OutputCount int

	Threads []*ThreadBlueprint
}

func (p *ProcessBlueprint) ThreadById(id int) (*ThreadBlueprint, bool) {
	for _, t := range p.Threads {
		if t.ID == id {
			return t, true
		}
	}

	return nil, false
}
