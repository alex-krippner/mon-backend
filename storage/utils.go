package storage

import "encoding/json"

type VocabStructs interface {
	*VocabDefinitions | *VocabSentences | *PartsOfSpeech
}

type KanjiStructs interface {
	*ExampleSentences | *ExampleWords | *Meanings
}

func HandleNil[S VocabStructs | KanjiStructs](s S) error {
	b := []byte(`[]`)
	return json.Unmarshal(b, &s)
}
