// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package ports

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// DeletedReading defines model for DeletedReading.
type DeletedReading struct {
	Id string `json:"id"`
}

// PostReading defines model for PostReading.
type PostReading struct {
	Japanese    string `json:"japanese"`
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

// Reading defines model for Reading.
type Reading struct {
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Japanese    string              `json:"japanese"`
	Title       string              `json:"title"`
	Translation string              `json:"translation"`
}

// Readings defines model for Readings.
type Readings struct {
	Readings []Reading `json:"readings"`
}

// UpdateReadingJSONRequestBody defines body for UpdateReading for application/json ContentType.
type UpdateReadingJSONRequestBody = Reading

// CreateReadingJSONRequestBody defines body for CreateReading for application/json ContentType.
type CreateReadingJSONRequestBody = PostReading
