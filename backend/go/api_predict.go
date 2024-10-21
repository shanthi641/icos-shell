// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 */

package shellbackend

import (
	"encoding/json"
	"net/http"
	"strings"
)

// PredictAPIController binds http requests to an api service and writes the service results to the http response
type PredictAPIController struct {
	service PredictAPIServicer
	errorHandler ErrorHandler
}

// PredictAPIOption for how the controller is set up.
type PredictAPIOption func(*PredictAPIController)

// WithPredictAPIErrorHandler inject ErrorHandler into controller
func WithPredictAPIErrorHandler(h ErrorHandler) PredictAPIOption {
	return func(c *PredictAPIController) {
		c.errorHandler = h
	}
}

// NewPredictAPIController creates a default api controller
func NewPredictAPIController(s PredictAPIServicer, opts ...PredictAPIOption) *PredictAPIController {
	controller := &PredictAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PredictAPIController
func (c *PredictAPIController) Routes() Routes {
	return Routes{
		"PredictMetrics": Route{
			strings.ToUpper("Post"),
			"/api/v3/metrics/predict",
			c.PredictMetrics,
		},
	}
}

// PredictMetrics - Predict metrics development based on model and input metrics
func (c *PredictAPIController) PredictMetrics(w http.ResponseWriter, r *http.Request) {
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.PredictMetrics(r.Context(), bodyParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}