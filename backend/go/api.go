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
	"context"
	"net/http"
)



// ControllerAPIRouter defines the required methods for binding the api requests to a responses for the ControllerAPI
// The ControllerAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ControllerAPIServicer to perform the required actions, then write the service results to the http response.
type ControllerAPIRouter interface { 
	AddController(http.ResponseWriter, *http.Request)
	GetControllers(http.ResponseWriter, *http.Request)
}
// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	GetHealthcheck(http.ResponseWriter, *http.Request)
}
// DeploymentAPIRouter defines the required methods for binding the api requests to a responses for the DeploymentAPI
// The DeploymentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DeploymentAPIServicer to perform the required actions, then write the service results to the http response.
type DeploymentAPIRouter interface { 
	CreateDeployment(http.ResponseWriter, *http.Request)
	DeleteDeploymentById(http.ResponseWriter, *http.Request)
	GetDeploymentById(http.ResponseWriter, *http.Request)
	GetDeployments(http.ResponseWriter, *http.Request)
	StartDeploymentById(http.ResponseWriter, *http.Request)
	StopDeploymentById(http.ResponseWriter, *http.Request)
	UpdateDeployment(http.ResponseWriter, *http.Request)
}
// PredictAPIRouter defines the required methods for binding the api requests to a responses for the PredictAPI
// The PredictAPIRouter implementation should parse necessary information from the http request,
// pass the data to a PredictAPIServicer to perform the required actions, then write the service results to the http response.
type PredictAPIRouter interface { 
	PredictMetrics(http.ResponseWriter, *http.Request)
}
// ResourceAPIRouter defines the required methods for binding the api requests to a responses for the ResourceAPI
// The ResourceAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ResourceAPIServicer to perform the required actions, then write the service results to the http response.
type ResourceAPIRouter interface { 
	GetResourceById(http.ResponseWriter, *http.Request)
	GetResources(http.ResponseWriter, *http.Request)
}
// TrainAPIRouter defines the required methods for binding the api requests to a responses for the TrainAPI
// The TrainAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TrainAPIServicer to perform the required actions, then write the service results to the http response.
type TrainAPIRouter interface { 
	TrainMetrics(http.ResponseWriter, *http.Request)
}
// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface { 
	LoginUser(http.ResponseWriter, *http.Request)
	LogoutUser(http.ResponseWriter, *http.Request)
}


// ControllerAPIServicer defines the api actions for the ControllerAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ControllerAPIServicer interface { 
	AddController(context.Context, Controller, string) (ImplResponse, error)
	GetControllers(context.Context) (ImplResponse, error)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	GetHealthcheck(context.Context) (ImplResponse, error)
}


// DeploymentAPIServicer defines the api actions for the DeploymentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DeploymentAPIServicer interface { 
	CreateDeployment(context.Context, map[string]interface{}, string) (ImplResponse, error)
	DeleteDeploymentById(context.Context, string, string) (ImplResponse, error)
	GetDeploymentById(context.Context, string, string) (ImplResponse, error)
	GetDeployments(context.Context, string) (ImplResponse, error)
	StartDeploymentById(context.Context, string, string) (ImplResponse, error)
	StopDeploymentById(context.Context, string, string) (ImplResponse, error)
	UpdateDeployment(context.Context, string, map[string]interface{}, string) (ImplResponse, error)
}


// PredictAPIServicer defines the api actions for the PredictAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PredictAPIServicer interface { 
	PredictMetrics(context.Context, map[string]interface{}, string) (ImplResponse, error)
}


// ResourceAPIServicer defines the api actions for the ResourceAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ResourceAPIServicer interface { 
	GetResourceById(context.Context, int64) (ImplResponse, error)
	GetResources(context.Context, string) (ImplResponse, error)
}


// TrainAPIServicer defines the api actions for the TrainAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TrainAPIServicer interface { 
	TrainMetrics(context.Context, map[string]interface{}, string) (ImplResponse, error)
}


// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface { 
	LoginUser(context.Context, string, string) (ImplResponse, error)
	LogoutUser(context.Context) (ImplResponse, error)
}
