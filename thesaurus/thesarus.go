package thesaurus

type Thesaurus interface {
	Synonims(term string) ([]string, error)
}
