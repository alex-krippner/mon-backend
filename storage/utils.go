package storage

import "encoding/json"

type VocabStructs interface {
	*VocabDefinitions | *VocabSentences | *PartsOfSpeech
}

func HandleNil[S VocabStructs](s S) error {
	b := []byte(`[]`)
	return json.Unmarshal(b, &s)
}
