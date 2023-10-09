package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ExecutionEnvironmentService implements awx execution environment apis.
type ExecutionEnvironmentService struct {
	client *Client
}

// ListExecutionEnvironmentsResponse represents `ListExecutionEnvironments` endpoint response.
type ListExecutionEnvironmentsResponse struct {
	Pagination
	Results []*ExecutionEnvironment `json:"results"`
}

const ExecutionEnvironmentAPIEndpoint = "/api/v2/execution_environments/"

// GetExecutionEnvironmentByID shows the details of a execution environment.
func (jt *ExecutionEnvironmentService) GetExecutionEnvironmentByID(id int, params map[string]string) (*ExecutionEnvironment, error) {
	result := new(ExecutionEnvironment)
	endpoint := fmt.Sprintf("%s%d/", ExecutionEnvironmentAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListExecutionEnvironments shows a list of execution environments.
func (jt *ExecutionEnvironmentService) ListExecutionEnvironments(params map[string]string) ([]*ExecutionEnvironment, *ListExecutionEnvironmentsResponse, error) {
	result := new(ListExecutionEnvironmentsResponse)
	resp, err := jt.client.Requester.GetJSON(ExecutionEnvironmentAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateExecutionEnvironment creates an execution environment
func (jt *ExecutionEnvironmentService) CreateExecutionEnvironment(data map[string]interface{}, params map[string]string) (*ExecutionEnvironment, error) {
	result := new(ExecutionEnvironment)
	mandatoryFields = []string{"name", "image"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(ExecutionEnvironmentAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateExecutionEnvironment updates an execution environment
func (jt *ExecutionEnvironmentService) UpdateExecutionEnvironment(id int, data map[string]interface{}, params map[string]string) (*ExecutionEnvironment, error) {
	result := new(ExecutionEnvironment)
	endpoint := fmt.Sprintf("%s%d", ExecutionEnvironmentAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteExecutionEnvironment deletes an execution environment
func (jt *ExecutionEnvironmentService) DeleteExecutionEnvironment(id int) (*ExecutionEnvironment, error) {
	result := new(ExecutionEnvironment)
	endpoint := fmt.Sprintf("%s%d", ExecutionEnvironmentAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
