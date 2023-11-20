/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package shellbackend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// DeploymentAPIService is a service that implements the logic for the DeploymentAPIServicer
// This service should implement the business logic for every endpoint for the DeploymentAPI API.
// Include any external packages or services that will be required by this service.
type DeploymentAPIService struct {
}

// NewDeploymentAPIService creates a default api service
func NewDeploymentAPIService() DeploymentAPIServicer {
	return &DeploymentAPIService{}
}

func prepareToken(ctx context.Context, apiKey string, req *http.Request) *http.Request {
	_, authToken, _ := receiveAndValidateAccessToken(ctx, apiKey)
	authToken = "Bearer " + authToken
	req.Header.Add("Authorization", authToken)
	return req
}

func errorConnect(resp *http.Response, err error) (ImplResponse, error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	if resp != nil {
		return Response(500, resp.Body), nil
	}
	return Response(500, "Error while connecting to job_manager"), nil
}

// CreateDeployment - Creates a new deployment
func (s *DeploymentAPIService) CreateDeployment(ctx context.Context, body map[string]interface{}, apiKey string) (ImplResponse, error) {
	jsonData, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, "POST", viper.GetString("components.job_manager")+"/jobmanager/jobs/create", bytes.NewBuffer(jsonData))
	req = prepareToken(ctx, apiKey, req)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	resp.Body.Close()
	if err != nil {
		return errorConnect(resp, err)
	} else {
		if resp.StatusCode == 201 {
			return Response(resp.StatusCode, "Deployment successfully created!"), nil
		} else if resp.StatusCode == 405 {
			return Response(resp.StatusCode, "Invalid input"), nil
		} else {
			return Response(resp.StatusCode, resp.Body), nil
		}
	}
}

// DeleteDeploymentById - Deletes a deployment
func (s *DeploymentAPIService) DeleteDeploymentById(ctx context.Context, deploymentId int64, apiKey string) (ImplResponse, error) {
	req, _ := http.NewRequestWithContext(ctx, "DELETE", viper.GetString("components.job_manager")+"/jobmanager/jobs/"+strconv.FormatInt(deploymentId, 10), nil)
	req = prepareToken(ctx, apiKey, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	resp.Body.Close()
	if err != nil {
		return errorConnect(resp, err)
	} else {
		responseString := "Unexpected status code received"
		if resp.StatusCode == 204 {
			responseString = "No deployments found"
		} else if resp.StatusCode == 405 {
			responseString = "Invalid input"
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", responseString)
			return Response(resp.StatusCode, resp.Body), nil
		}
		return Response(resp.StatusCode, responseString), nil
	}
}

// GetDeploymentById - Find deployment by ID
func (s *DeploymentAPIService) GetDeploymentById(ctx context.Context, deploymentId int64, apiKey string) (ImplResponse, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", viper.GetString("components.job_manager")+"/jobmanager/jobs/"+strconv.FormatInt(deploymentId, 10), nil)
	req = prepareToken(ctx, apiKey, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errorConnect(resp, err)
	} else {
		if resp.StatusCode == 201 {
			resBody := resp.Body
			resp.Body.Close()
			return Response(resp.StatusCode, resBody), nil
		} else if resp.StatusCode == 204 {
			resp.Body.Close()
			return Response(resp.StatusCode, "No deployments found"), nil
		} else if resp.StatusCode == 405 {
			resp.Body.Close()
			return Response(resp.StatusCode, "Invalid input"), nil
		} else {
			resBody := resp.Body
			resp.Body.Close()
			return Response(resp.StatusCode, resBody), nil
		}
	}
}

// GetDeployments - Returns a list of deployments
func (s *DeploymentAPIService) GetDeployments(ctx context.Context, apiKey string) (ImplResponse, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", viper.GetString("components.job_manager")+"/jobmanager/jobs", nil)
	req = prepareToken(ctx, apiKey, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errorConnect(resp, err)
	} else {
		if resp.StatusCode == 200 {
			resBody := resp.Body
			resp.Body.Close()
			return Response(resp.StatusCode, resBody), nil
		} else if resp.StatusCode == 405 {
			resp.Body.Close()
			return Response(resp.StatusCode, "Invalid input"), nil
		} else {
			resBody := resp.Body
			resp.Body.Close()
			return Response(resp.StatusCode, resBody), nil
		}
	}
}

// UpdateDeployment - Updates a deployment
func (s *DeploymentAPIService) UpdateDeployment(ctx context.Context, deploymentId int64, body map[string]interface{}, apiKey string) (ImplResponse, error) {
	jsonData, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, "PUT", viper.GetString("components.job_manager")+"/jobmanager/jobs/"+strconv.FormatInt(deploymentId, 10), bytes.NewBuffer(jsonData))
	req = prepareToken(ctx, apiKey, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errorConnect(resp, err)
	} else {
		responseString := "Unexpected status code received"
		if resp.StatusCode == 204 {
			responseString = "No deployments found"
		} else if resp.StatusCode == 405 {
			responseString = "Invalid input"
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", responseString)
			resBody := resp.Body
			resp.Body.Close()
			return Response(resp.StatusCode, resBody), nil
		}
		resp.Body.Close()
		return Response(resp.StatusCode, responseString), nil
	}
}
