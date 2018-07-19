// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"time"

	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

var nodePoolOperationMaxTime = 20 * time.Minute

func NodePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createNodePool,
		Read:     readNodePool,
		Update:   updateNodePool,
		Delete:   deleteNodePool,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_image_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"initial_node_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"quantity_per_subnet": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssh_public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"message": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_pool_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.CreateResource(d, sync)
}

func readNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.ReadResource(sync)
}

func updateNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.UpdateResource(d, sync)
}

func deleteNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type NodePoolResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.NodePool
	DisableNotFoundRetries bool
}

func (s *NodePoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NodePoolResourceCrud) Create() error {
	request := oci_containerengine.CreateNodePoolRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.InitialNodeLabels = []oci_containerengine.KeyValue{}
	if initialNodeLabels, ok := s.D.GetOkExists("initial_node_labels"); ok {
		interfaces := initialNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.KeyValue, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToKeyValue(toBeConverted.(map[string]interface{}))
		}
		request.InitialNodeLabels = tmp
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nodeImageName, ok := s.D.GetOkExists("node_image_name"); ok {
		tmp := nodeImageName.(string)
		request.NodeImageName = &tmp
	}

	if nodeShape, ok := s.D.GetOkExists("node_shape"); ok {
		tmp := nodeShape.(string)
		request.NodeShape = &tmp
	}

	if quantityPerSubnet, ok := s.D.GetOkExists("quantity_per_subnet"); ok {
		tmp := quantityPerSubnet.(int)
		request.QuantityPerSubnet = &tmp
	}

	if sshPublicKey, ok := s.D.GetOkExists("ssh_public_key"); ok {
		tmp := sshPublicKey.(string)
		request.SshPublicKey = &tmp
	}

	request.SubnetIds = []string{}
	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.SubnetIds = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	//Trigger create request
	response, err := s.Client.CreateNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId

	//Wait until it finishes
	nodePoolID, err := containerEngineWaitForWorkRequest(workID, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeCreated, nodePoolOperationMaxTime, s.DisableNotFoundRetries,
		s.Client)
	if err != nil {
		if nodePoolID != nil {
			//Try to clean up
			delReq := oci_containerengine.DeleteNodePoolRequest{}
			delReq.NodePoolId = nodePoolID
			delReq.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

			//Issues delete delRequest
			delRes, delErr := s.Client.DeleteNodePool(context.Background(), delReq)
			if delErr != nil {
				return err
			}
			delWorkRequest := delRes.OpcWorkRequestId

			//Wait until delRequest finishes
			_, _ = containerEngineWaitForWorkRequest(delWorkRequest, "nodepool",
				oci_containerengine.WorkRequestResourceActionTypeDeleted,
				nodePoolOperationMaxTime, s.DisableNotFoundRetries, s.Client)
		}
		return err
	}

	//Fetch nodepool object
	requestGet := oci_containerengine.GetNodePoolRequest{}
	requestGet.NodePoolId = nodePoolID
	requestGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetNodePool(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.NodePool
	return nil
}

func (s *NodePoolResourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolRequest{}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NodePool
	return nil
}

func (s *NodePoolResourceCrud) Update() error {
	request := oci_containerengine.UpdateNodePoolRequest{}

	request.InitialNodeLabels = []oci_containerengine.KeyValue{}
	if initialNodeLabels, ok := s.D.GetOkExists("initial_node_labels"); ok {
		interfaces := initialNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.KeyValue, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToKeyValue(toBeConverted.(map[string]interface{}))
		}
		request.InitialNodeLabels = tmp
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	if quantityPerSubnet, ok := s.D.GetOkExists("quantity_per_subnet"); ok {
		tmp := quantityPerSubnet.(int)
		request.QuantityPerSubnet = &tmp
	}

	request.SubnetIds = []string{}
	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.SubnetIds = tmp
	}

	//Issue update request
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	response, err := s.Client.UpdateNodePool(context.Background(), request)
	if err != nil {
		return err
	}
	workRequest := response.OpcWorkRequestId

	//Wait until request finishes
	nodePoolID, err := containerEngineWaitForWorkRequest(workRequest, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeUpdated,
		nodePoolOperationMaxTime, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	//Refresh data
	requestGet := oci_containerengine.GetNodePoolRequest{}
	requestGet.NodePoolId = nodePoolID
	requestGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetNodePool(context.Background(), requestGet)
	if err != nil {
		return err
	}

	s.Res = &responseGet.NodePool
	return nil
}

func (s *NodePoolResourceCrud) Delete() error {
	request := oci_containerengine.DeleteNodePoolRequest{}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	//Issues delete request
	response, err := s.Client.DeleteNodePool(context.Background(), request)
	workRequest := response.OpcWorkRequestId

	//Wait until request finishes
	_, err = containerEngineWaitForWorkRequest(workRequest, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeDeleted,
		nodePoolOperationMaxTime, s.DisableNotFoundRetries, s.Client)

	return err

}

func (s *NodePoolResourceCrud) SetData() {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	initialNodeLabels := []interface{}{}
	for _, item := range s.Res.InitialNodeLabels {
		initialNodeLabels = append(initialNodeLabels, KeyValueToMap(item))
	}
	s.D.Set("initial_node_labels", initialNodeLabels)

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NodeImageId != nil {
		s.D.Set("node_image_id", *s.Res.NodeImageId)
	}

	if s.Res.NodeImageName != nil {
		s.D.Set("node_image_name", *s.Res.NodeImageName)
	}

	if s.Res.NodeShape != nil {
		s.D.Set("node_shape", *s.Res.NodeShape)
	}

	nodes := []interface{}{}
	for _, item := range s.Res.Nodes {
		nodes = append(nodes, NodeToMap(item))
	}
	s.D.Set("nodes", nodes)

	if s.Res.QuantityPerSubnet != nil {
		s.D.Set("quantity_per_subnet", *s.Res.QuantityPerSubnet)
	}

	if s.Res.SshPublicKey != nil {
		s.D.Set("ssh_public_key", *s.Res.SshPublicKey)
	}

	s.D.Set("subnet_ids", s.Res.SubnetIds)

}

func ErrorToMap(obj *oci_containerengine.NodeError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	return result
}

func mapToKeyValue(raw map[string]interface{}) oci_containerengine.KeyValue {
	result := oci_containerengine.KeyValue{}

	if key, ok := raw["key"]; ok && key != "" {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := raw["value"]; ok && value != "" {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result
}

func KeyValueToMap(obj oci_containerengine.KeyValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func NodeToMap(obj oci_containerengine.Node) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.NodeError != nil {
		result["error"] = []interface{}{ErrorToMap(obj.NodeError)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NodePoolId != nil {
		result["node_pool_id"] = string(*obj.NodePoolId)
	}

	if obj.PublicIp != nil {
		result["public_ip"] = string(*obj.PublicIp)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}
