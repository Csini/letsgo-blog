/*
 * Let's go Blog API
 *
 * Application providing Blog.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Amount struct {
	Blog int32 `json:"blog"`

	Comment int32 `json:"comment"`
}

// AssertAmountRequired checks if the required fields are not zero-ed
func AssertAmountRequired(obj Amount) error {
	elements := map[string]interface{}{
		"blog":    obj.Blog,
		"comment": obj.Comment,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAmountConstraints checks if the values respects the defined constraints
func AssertAmountConstraints(obj Amount) error {
	return nil
}
