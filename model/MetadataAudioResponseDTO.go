package model

// MetadataAudioResponseDTO represents the response DTO for audio metadat
type MetadataAudioResponseDTO struct {
	ObAudioMetadata MetadataAudio
	Code            int
	Message         string
}
