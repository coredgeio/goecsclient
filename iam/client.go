package iam

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	client "github.com/coredgeio/goecsclient"
)

type IamClient interface {
	AttachPolicy(namespace string, param *AttachPolicyParameters) (*AttachPolicyResp, error)
	CreateAccessKey(namespace string, param *CreateAccessKeyParameters) (*CreateAccessKeyResp, error)
	CreatePolicy(namespace string, param *CreatePolicyParameters) (*CreatePolicyResp, error)
	CreatePolicyVersion(namespace string, param *CreatePolicyVersionParameters) (*CreatePolicyVersionResp, error)
	CreateUser(namespace string, param *CreateUserParameters) (*CreateUserResp, error)
	DeleteAccessKey(namespace string, param *DeleteAccessKeyParameters) (*DeleteAccessKeyResp, error)
	DeletePolicy(namespace string, param *DeletePolicyParameters) (*DeletePolicyResp, error)
	DeletePolicyVersion(namespace string, param *DeletePolicyVersionParameters) (*DeletePolicyVersionResp, error)
	DeleteUser(namespace string, param *DeleteUserParameters) (*DeleteUserResp, error)
	DetachPolicy(namespace string, param *DetachPolicyParameters) (*DetachPolicyResp, error)
	GetPolicy(namespace string, param *GetPolicyParameters) (*GetPolicyResp, error)
	GetUserAttachedPolicyList(namespace string, param *GetUserAttachedPolicyListParameters) (*GetUserAttachedPolicyListResp, error)
	ListPolicies(namespace string, param *ListPoliciesParameters) (*ListPoliciesResp, error)
	UpdateAccessKey(namespace string, param *UpdateAccessKeyParameters) (*UpdateAccessKeyResp, error)
	
}

type iamClient struct {
	apiClient client.EcsClient
}

// Create Policy
func (c *iamClient) CreatePolicy(namespace string, param *CreatePolicyParameters) (*CreatePolicyResp, error) {
	var query url.Values
	if param != nil && (param.Description != "" || param.Path != "" || param.PolicyDocument != "" ||
		param.PolicyName != "") && param.Action == "CreatePolicy" {
		query = url.Values{}
		if param.Description != "" {
			query.Add("Description", param.Description)
		}
		if param.Path != "" {
			query.Add("Path", param.Path)
		}
		if param.PolicyDocument != "" {
			query.Set("PolicyDocument", param.PolicyDocument)
		}
		if param.PolicyName != "" {
			query.Add("PolicyName", param.PolicyName)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Create Policy", err)
	}
	return resp, err
}

// Get managed Policy
func (c *iamClient) GetPolicy(namespace string, param *GetPolicyParameters) (*GetPolicyResp, error) {
	var query url.Values
	if param != nil && (param.PolicyArn != "") && param.Action == "GetPolicy" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &GetPolicyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for get specific user policy", err)
	}
	return resp, err
}

// List Policies
func (c *iamClient) ListPolicies(namespace string, param *ListPoliciesParameters) (*ListPoliciesResp, error) {
	var query url.Values
	if param != nil && (param.Marker != "" || param.MaxItems != 0 || param.OnlyAttached != "" || param.PathPrefix != "" ||
		param.PolicyScope != "" || param.PolicyUsageFilter != "") || param.Action == "ListPolicies" {
		query = url.Values{}
		if param.Marker != "" {
			query.Add("Marker", param.Marker)
		}
		if param.OnlyAttached != "" {
			query.Add("OnlyAttached", param.OnlyAttached)
		}
		if param.PathPrefix != "" {
			query.Add("PathPrefix", param.PathPrefix)
		}
		if param.PolicyScope != "" {
			query.Add("PolicyScope", param.PolicyScope)
		}
		if param.PolicyUsageFilter != "" {
			query.Add("PolicyUsageFilter", param.PolicyUsageFilter)
		}
		if param.MaxItems != 0 {
			query.Add("MaxItems", strconv.Itoa(param.MaxItems))
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}
	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for List Policies", err)
	}
	return resp, err
}

// Delete User Policy
func (c *iamClient) DeletePolicy(namespace string, param *DeletePolicyParameters) (*DeletePolicyResp, error) {
	var query url.Values
	if param != nil && param.PolicyArn != "" && param.Action == "DeletePolicy" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		query.Add("Action", param.Action)
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for delete user policy", err)
	}
	return resp, err
}

// Create Policy Version
func (c *iamClient) CreatePolicyVersion(namespace string, param *CreatePolicyVersionParameters) (*CreatePolicyVersionResp, error) {
	var query url.Values
	if param != nil && (param.PolicyArn != "" || param.PolicyDocument != "") &&
		param.Action == "CreatePolicyVersion" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		if param.PolicyDocument != "" {
			query.Set("PolicyDocument", param.PolicyDocument)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}
	query.Add("SetAsDefault", strconv.FormatBool(param.SetAsDefault))

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyVersionResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for create user policy version", err)
	}
	return resp, err
}

// Delete Policy Version
func (c *iamClient) DeletePolicyVersion(namespace string, param *DeletePolicyVersionParameters) (*DeletePolicyVersionResp, error) {
	var query url.Values
	if param != nil && (param.PolicyArn != "" || param.VersionId != "") && param.Action == "DeletePolicyVersion" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		if param.VersionId != "" {
			query.Add("VersionId", param.VersionId)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyVersionResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for delete policy version", err)
	}
	return resp, err
}

// Attach User Policy
func (c *iamClient) AttachPolicy(namespace string, param *AttachPolicyParameters) (*AttachPolicyResp, error) {
	var query url.Values
	if param != nil && (param.PolicyArn != "" || param.UserName != "") && param.Action == "AttachUserPolicy" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &AttachPolicyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Attach User Policy", err)
	}
	return resp, err
}

// Detach User Policy
func (c *iamClient) DetachPolicy(namespace string, param *DetachPolicyParameters) (*DetachPolicyResp, error) {
	var query url.Values
	if param != nil && (param.PolicyArn != "" || param.UserName != "") && param.Action == "DetachUserPolicy" {
		query = url.Values{}
		if param.PolicyArn != "" {
			query.Add("PolicyArn", param.PolicyArn)
		}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &DetachPolicyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Detach User Policy", err)
	}
	return resp, err
}

// Create Access Key
func (c *iamClient) CreateAccessKey(namespace string, param *CreateAccessKeyParameters) (*CreateAccessKeyResp, error) {
	var query url.Values
	if param != nil && param.UserName != "" && param.Action == "CreateAccessKey" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &CreateAccessKeyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Create Access Key", err)
	}
	return resp, err
}

// Delete Access Key
func (c *iamClient) DeleteAccessKey(namespace string, param *DeleteAccessKeyParameters) (*DeleteAccessKeyResp, error) {
	var query url.Values
	if param != nil && (param.UserName != "" || param.AccessKeyId != "") && param.Action == "DeleteAccessKey" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.AccessKeyId != "" {
			query.Add("AccessKeyId", param.AccessKeyId)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessKeyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Delete Access Key", err)
	}
	return resp, err
}

// Update Access Key
func (c *iamClient) UpdateAccessKey(namespace string, param *UpdateAccessKeyParameters) (*UpdateAccessKeyResp, error) {
	var query url.Values
	if param != nil && (param.UserName != "" || param.AccessKeyId != "" || param.Status != "") && param.Action == "UpdateAccessKey" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.AccessKeyId != "" {
			query.Add("AccessKeyId", param.AccessKeyId)
		}
		if param.Status != "" {
			query.Add("Status", param.Status)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccessKeyResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for Update Access Key", err)
	}
	return resp, err
}

// iam Create User
func (c *iamClient) CreateUser(namespace string, param *CreateUserParameters) (*CreateUserResp, error) {
	var query url.Values
	if param != nil && (param.UserName != "" || param.Path != "" || param.PermissionsBoundary != "") &&
		param.Action == "CreateUser" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if param.Path != "" {
			query.Add("Path", param.Path)
		}
		if param.PermissionsBoundary != "" {
			query.Add("PermissionsBoundary", param.PermissionsBoundary)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
		// TODO (do something to add tags, currently don't know if below code is correct or not
		// POST https://192.168.0.0:4443/iam?UserName=payroll1&Path=/&Tags.member.1.Key=Department&Tags.member.1.Value=Finance&Action=CreateUser
		for _, tag := range param.Tags {
			query.Add(tag.Key, tag.Value)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for create user", err)
	}
	return resp, err
}

// iam Delete User
func (c *iamClient) DeleteUser(namespace string, param *DeleteUserParameters) (*DeleteUserResp, error) {
	var query url.Values
	if param != nil && param.UserName != "" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
	}
	query.Add("Action", "DeleteUser")

	h := map[string]string{
		"x-emc-namespace": namespace,
	}
	bytes, err := c.apiClient.Post("/iam", nil, query, h)
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for delete user", err)
	}
	return resp, err
}

// https://coredge-jira.atlassian.net/browse/COMPASS-4180?focusedCommentId=21543
func (c *iamClient) GetUserAttachedPolicyList(namespace string, param *GetUserAttachedPolicyListParameters) (*GetUserAttachedPolicyListResp, error) {
	var query url.Values
	if param != nil && (param.UserName != "" || namespace != "") && param.Action == "ListAttachedUserPolicies" {
		query = url.Values{}
		if param.UserName != "" {
			query.Add("UserName", param.UserName)
		}
		if namespace != "" {
			query.Add("Namespace", namespace)
		}
		if param.Action != "" {
			query.Add("Action", param.Action)
		}
	}

	h := map[string]string{
		"x-emc-namespace": namespace,
	}

	bytes, err := c.apiClient.Get("/iam", query, h)
	if err != nil {
		return nil, err
	}

	resp := &GetUserAttachedPolicyListResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for get user policies", err)
	}
	return resp, err
}

// provides EcsIamClient for give handler to EcsClient
func GetEcsIamClient(apiClient client.EcsClient) IamClient {
	return &iamClient{
		apiClient: apiClient,
	}
}
