package types

type Gedcom struct {
	Header     Header
	Submission *Submission
	Family     []*Family
	Individual []*Individual
	Media      []*MultiMedia
	Repository []*Repository
	Source     []*Source
	Submitter  []*Submitter
	Trailer    *Trailer
}

func (g Gedcom) IsValid() bool {
	valid := true

	return valid
}
