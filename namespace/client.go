package namespace

import (
	"encoding/json"
	"log"

	client "github.com/coredgeio/goecsclient"
)

type NamespaceClient interface {
	CreateNamespace(req *CreateNamespaceReq) (*CreateNamespaceResp, error)
	DeleteNamespace(namespace string) error
	UpdateNamespace(namespace string, req *UpdateNamespaceReq) error
	SetNamespaceQuota(namespace string, req *SetNamespaceQuotaReq) error
}

type namespaceClient struct {
	apiClient client.EcsClient
}

// Create Namespace
func (c *namespaceClient) CreateNamespace(req *CreateNamespaceReq) (*CreateNamespaceResp, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	bytes, err := c.apiClient.Post("/object/namespaces/namespace", data, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := &CreateNamespaceResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for create namespace", err)
	}
	return resp, err
}

// Update Namespace
func (c *namespaceClient) UpdateNamespace(namespace string, req *UpdateNamespaceReq) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = c.apiClient.Put("/object/namespaces/namespace/"+namespace, data, nil)
	return err
}

func (c *namespaceClient) DeleteNamespace(namespace string) error {
	_, err := c.apiClient.Post("/object/namespaces/namespace/"+namespace+"/deactivate", nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *namespaceClient) SetNamespaceQuota(namespace string, req *SetNamespaceQuotaReq) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = c.apiClient.Put("/object/namespaces/namespace/"+namespace+"/quota", data, nil)
	if err != nil {
		return err
	}
	return nil
}

// provides EcsNamespaceClient for give handler to EcsClient
func GetEcsNamespaceClient(apiClient client.EcsClient) NamespaceClient {
	return &namespaceClient{
		apiClient: apiClient,
	}
}
