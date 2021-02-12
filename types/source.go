package types

type Source struct {
	Xref  string
	Title string
	Media []*MultiMedia
	Note  []*Note
	// Name string
	// Data SourceData
}

func (s Source) IsValid() bool {
	valid := true

	return valid
}

type SourceData struct {
	Date      string
	Copyright string
}

func (sd SourceData) IsValid() bool {
	valid := true

	return valid
}
