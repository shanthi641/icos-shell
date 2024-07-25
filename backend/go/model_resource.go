// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 */

package shellbackend




type Resource struct {

	// Unique identifier of the resource
	Id int64 `json:"id"`

	// Name of the resource
	Name string `json:"name,omitempty"`

	// Type of resource
	Type string `json:"type"`

	// ID of the parent resource
	ParentId int64 `json:"parentId,omitempty"`

	// Status of the resource
	Status string `json:"status,omitempty"`
}

// AssertResourceRequired checks if the required fields are not zero-ed
func AssertResourceRequired(obj Resource) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"type": obj.Type,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertResourceConstraints checks if the values respects the defined constraints
func AssertResourceConstraints(obj Resource) error {
	return nil
}
