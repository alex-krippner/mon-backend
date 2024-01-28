package reading

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Reading struct {
	ID          string `json:"id,omitempty"`
	Translation string `json:"translation,omitempty"`
	Japanese    string `json:"japanese,omitempty"`
	Title       string `json:"title,omitempty"`
	Username    string `json:"username,omitempty"`
	Tokens      []Token
}

type Token struct {

	// Token attributes
	Text         string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Lemma        string `protobuf:"bytes,7,opt,name=lemma_,json=lemma,proto3" json:"lemma_,omitempty"`
	Normalized   string `protobuf:"bytes,9,opt,name=norm_str,json=normStr,proto3" json:"norm_str,omitempty"`
	Prefix       string `protobuf:"bytes,14,opt,name=prefix_,json=prefix,proto3" json:"prefix_,omitempty"`
	Suffix       string `protobuf:"bytes,15,opt,name=suffix_,json=suffix,proto3" json:"suffix_,omitempty"`
	PartOfSpeech string `protobuf:"bytes,37,opt,name=pos_str,json=posStr,proto3" json:"pos_str,omitempty"`
	Tag          string `protobuf:"bytes,39,opt,name=tag_str,json=tagStr,proto3" json:"tag_str,omitempty"`
	Dependency   string `protobuf:"bytes,41,opt,name=dep_str,json=depStr,proto3" json:"dep_str,omitempty"`
	Lang         string `protobuf:"bytes,43,opt,name=lang_str,json=langStr,proto3" json:"lang_str,omitempty"`
}

type ReadingRepository interface {
	CreateReading(ctx context.Context, req *Reading) (*Reading, error)
	GetAllReading(ctx context.Context, username string) ([]*Reading, error)
	UpdateReading(ctx context.Context, req Reading) (*Reading, error)
	DeleteReading(id string) error
}

func NewReading(translation string, japanese string, title string, username string) (*Reading, error) {
	if translation == "" {
		return nil, errors.New("reading translation missing")
	}
	if japanese == "" {
		return nil, errors.New("japanese text missing")
	}
	if title == "" {
		return nil, errors.New("reading title is missing")
	}
	if username == "" {
		return nil, errors.New("username is missing")
	}

	return &Reading{
		ID:          uuid.New().String(),
		Translation: translation,
		Japanese:    japanese,
		Title:       title,
		Username:    username,
	}, nil
}
