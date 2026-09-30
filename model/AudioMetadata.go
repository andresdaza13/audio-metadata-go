package model

// MetadataAudio represents the metadata of an audio file.
type MetadataAudio struct {
	title     string
	duration  int
	typeSong  string
	available bool
}

func (m *MetadataAudio) setTitle(title string) {
	m.title = title
}

func (m *MetadataAudio) setDuration(duration int) {
	m.duration = duration
}

func (m *MetadataAudio) setTypeSong(typeSong string) {
	m.typeSong = typeSong
}

func (m *MetadataAudio) setAvailable(available bool) {
	m.available = available
}
