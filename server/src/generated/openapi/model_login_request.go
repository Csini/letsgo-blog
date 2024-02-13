/*
 * Let's go Blog API
 *
 * Application providing Blog.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginRequest struct {
	User string `json:"user"`

	Pw string `json:"pw"`
}

// AssertLoginRequestRequired checks if the required fields are not zero-ed
func AssertLoginRequestRequired(obj LoginRequest) error {
	elements := map[string]interface{}{
		"user": obj.User,
		"pw":   obj.Pw,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertLoginRequestConstraints checks if the values respects the defined constraints
func AssertLoginRequestConstraints(obj LoginRequest) error {
	return nil
}
