/*
 * Let's go Blog API
 *
 * Application providing Blog.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// AuthenticationAPIRouter defines the required methods for binding the api requests to a responses for the AuthenticationAPI
// The AuthenticationAPIRouter implementation should parse necessary information from the http request,
// pass the data to a AuthenticationAPIServicer to perform the required actions, then write the service results to the http response.
type AuthenticationAPIRouter interface {
	PostLogin(http.ResponseWriter, *http.Request)
}

// StatisticsAPIRouter defines the required methods for binding the api requests to a responses for the StatisticsAPI
// The StatisticsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a StatisticsAPIServicer to perform the required actions, then write the service results to the http response.
type StatisticsAPIRouter interface {
	GetStatistics(http.ResponseWriter, *http.Request)
}

// AuthenticationAPIServicer defines the api actions for the AuthenticationAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AuthenticationAPIServicer interface {
	PostLogin(context.Context, LoginRequest) (ImplResponse, error)
}

// StatisticsAPIServicer defines the api actions for the StatisticsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type StatisticsAPIServicer interface {
	GetStatistics(context.Context, int32) (ImplResponse, error)
}
