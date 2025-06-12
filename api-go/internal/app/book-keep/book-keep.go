package bookkeep

type BookKeep struct{}

func NewBookKeep() *BookKeep {
	return &BookKeep{}
}

func (b *BookKeep) Start() error {
	return nil
}
