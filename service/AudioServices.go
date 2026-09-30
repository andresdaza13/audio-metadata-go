package service

func loandMetadataAudio() []MetadataAudio {
	vec := make([]MetadataAudio, 0, 5)

	var ObjMetadataAudio1, ObjMetadataAudio2, ObjMetadataAudio3, ObjMetadataAudio4, ObjMetadataAudio5 MetadataAudio
	ObjMetadataAudio1.setTitle("Song 1")
	ObjMetadataAudio1.setDuration(180)
	ObjMetadataAudio1.setTypeSong("Pop")
	ObjMetadataAudio1.setAvailable(true)

	ObjMetadataAudio2.setTitle("Song 2")
	ObjMetadataAudio2.setDuration(200)
	ObjMetadataAudio2.setTypeSong("Rock")
	ObjMetadataAudio2.setAvailable(true)

	ObjMetadataAudio3.setTitle("Song 3")
	ObjMetadataAudio3.setDuration(150)
	ObjMetadataAudio3.setTypeSong("Jazz")
	ObjMetadataAudio3.setAvailable(false)

	ObjMetadataAudio4.setTitle("Song 4")
	ObjMetadataAudio4.setDuration(220)
	ObjMetadataAudio4.setTypeSong("Classical")
	ObjMetadataAudio4.setAvailable(true)

	ObjMetadataAudio5.setTitle("Song 5")
	ObjMetadataAudio5.setDuration(190)
	ObjMetadataAudio5.setTypeSong("Hip Hop")
	ObjMetadataAudio5.setAvailable(false)

	return append(vec,
		ObjMetadataAudio1,
		ObjMetadataAudio2,
		ObjMetadataAudio3,
		ObjMetadataAudio4,
		ObjMetadataAudio5,
	)
}
