// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func ContainerengineNodePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &TwentyMinutes,
			Update: &TwentyMinutes,
			Delete: &TwentyMinutes,
		},
		Create: createContainerengineNodePool,
		Read:   readContainerengineNodePool,
		Update: updateContainerengineNodePool,
		Delete: deleteContainerengineNodePool,
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
			"node_shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"node_config_details": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"quantity_per_subnet", "subnet_ids"},
				MaxItems:      1,
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"placement_configs": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"availability_domain": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"node_image_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"node_image_name"},
			},
			"node_image_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"node_image_id"},
			},
			"node_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"quantity_per_subnet": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"node_config_details"},
			},
			"ssh_public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"node_config_details"},
				Set:           literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
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
									"status": {
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

func createContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return CreateResource(d, sync)
}

func readContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

func updateContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return UpdateResource(d, sync)
}

func deleteContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ContainerengineNodePoolResourceCrud struct {
	BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.NodePool
	DisableNotFoundRetries bool
}

func (s *ContainerengineNodePoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineNodePoolResourceCrud) Create() error {
	request := oci_containerengine.CreateNodePoolRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if initialNodeLabels, ok := s.D.GetOkExists("initial_node_labels"); ok {
		interfaces := initialNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.KeyValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_node_labels", stateDataIndex)
			converted, err := s.mapToKeyValue(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("initial_node_labels") {
			request.InitialNodeLabels = tmp
		}
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nodeConfigDetails, ok := s.D.GetOkExists("node_config_details"); ok {
		if tmpList := nodeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_config_details", 0)
			tmp, err := s.mapToCreateNodePoolNodeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NodeConfigDetails = &tmp
		}
	}

	if nodeImageId, ok := s.D.GetOkExists("node_image_id"); ok {
		tmp := nodeImageId.(string)
		request.NodeImageName = &tmp
	}

	if nodeImageName, ok := s.D.GetOkExists("node_image_name"); ok {
		tmp := nodeImageName.(string)
		request.NodeImageName = &tmp
	}

	if nodeMetadata, ok := s.D.GetOkExists("node_metadata"); ok {
		request.NodeMetadata = objectMapToStringMap(nodeMetadata.(map[string]interface{}))
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

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		set := subnetIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
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
		oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries,
		s.Client)
	if err != nil {
		if nodePoolID != nil {
			//Try to clean up
			log.Printf("[DEBUG] creation failed, attempting to delete the node pool: %v\n", nodePoolID)

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
			_, delErr = containerEngineWaitForWorkRequest(delWorkRequest, "nodepool",
				oci_containerengine.WorkRequestResourceActionTypeDeleted,
				s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
			if delErr != nil {
				log.Printf("[DEBUG] cleanup delWorkRequest failed with the error: %v\n", delErr)
			}
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

func (s *ContainerengineNodePoolResourceCrud) Get() error {
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

func (s *ContainerengineNodePoolResourceCrud) Update() error {
	request := oci_containerengine.UpdateNodePoolRequest{}

	if initialNodeLabels, ok := s.D.GetOkExists("initial_node_labels"); ok {
		interfaces := initialNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.KeyValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_node_labels", stateDataIndex)
			converted, err := s.mapToKeyValue(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("initial_node_labels") {
			request.InitialNodeLabels = tmp
		}
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nodeConfigDetails, ok := s.D.GetOkExists("node_config_details"); ok {
		if s.D.HasChange("node_config_details") {
			if tmpList := nodeConfigDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_config_details", 0)
				tmp, err := s.mapToUpdateNodePoolNodeConfigDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.NodeConfigDetails = &tmp
			}
		}
	}
	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		if s.D.HasChange("subnet_ids") {
			set := subnetIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			request.SubnetIds = tmp
		}
	}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	if quantityPerSubnet, ok := s.D.GetOkExists("quantity_per_subnet"); ok {
		if s.D.HasChange("quantity_per_subnet") {
			tmp := quantityPerSubnet.(int)
			request.QuantityPerSubnet = &tmp
		}
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
		s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
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

func (s *ContainerengineNodePoolResourceCrud) Delete() error {
	request := oci_containerengine.DeleteNodePoolRequest{}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	//Issue delete request
	response, err := s.Client.DeleteNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workRequest := response.OpcWorkRequestId

	//Wait until request finishes
	_, err = containerEngineWaitForWorkRequest(workRequest, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeDeleted,
		s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)

	return err

}

func (s *ContainerengineNodePoolResourceCrud) SetData() error {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.NodeConfigDetails != nil {
		s.D.Set("node_config_details", []interface{}{NodePoolNodeConfigDetailsToMap(s.Res.NodeConfigDetails)})
	} else {
		s.D.Set("node_config_details", nil)
	}

	if s.Res.NodeImageId != nil {
		s.D.Set("node_image_id", *s.Res.NodeImageId)
	}

	if s.Res.NodeImageName != nil {
		s.D.Set("node_image_name", *s.Res.NodeImageName)
	}

	s.D.Set("node_metadata", s.Res.NodeMetadata)

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

	if s.Res.SubnetIds != nil {
		subnetIds := []interface{}{}
		for _, item := range s.Res.SubnetIds {
			subnetIds = append(subnetIds, item)
		}
		s.D.Set("subnet_ids", schema.NewSet(literalTypeHashCodeForSets, subnetIds))
	} else {
		s.D.Set("subnet_ids", nil)
	}

	return nil
}

func (s *ContainerengineNodePoolResourceCrud) mapToCreateNodePoolNodeConfigDetails(fieldKeyFormat string) (oci_containerengine.CreateNodePoolNodeConfigDetails, error) {
	result := oci_containerengine.CreateNodePoolNodeConfigDetails{}

	if placementConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "placement_configs")); ok {
		interfaces := placementConfigs.([]interface{})
		tmp := make([]oci_containerengine.NodePoolPlacementConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "placement_configs"), stateDataIndex)
			converted, err := s.mapToNodePoolPlacementConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "placement_configs")) {
			result.PlacementConfigs = tmp
		}
	}

	if size, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size")); ok {
		tmp := size.(int)
		result.Size = &tmp
	}

	return result, nil
}

func (s *ContainerengineNodePoolResourceCrud) mapToUpdateNodePoolNodeConfigDetails(fieldKeyFormat string) (oci_containerengine.UpdateNodePoolNodeConfigDetails, error) {
	result := oci_containerengine.UpdateNodePoolNodeConfigDetails{}

	result.PlacementConfigs = []oci_containerengine.NodePoolPlacementConfigDetails{}
	if placementConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "placement_configs")); ok {
		interfaces := placementConfigs.([]interface{})
		tmp := make([]oci_containerengine.NodePoolPlacementConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "placement_configs"), stateDataIndex)
			converted, err := s.mapToNodePoolPlacementConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.PlacementConfigs = tmp
	}

	if size, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size")); ok {
		tmp := size.(int)
		result.Size = &tmp
	}

	return result, nil
}

func NodePoolNodeConfigDetailsToMap(obj *oci_containerengine.NodePoolNodeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	placementConfigs := []interface{}{}
	for _, item := range obj.PlacementConfigs {
		placementConfigs = append(placementConfigs, NodePoolPlacementConfigDetailsToMap(item))
	}
	result["placement_configs"] = placementConfigs

	if obj.Size != nil {
		result["size"] = int(*obj.Size)
	}

	return result
}

func (s *ContainerengineNodePoolResourceCrud) mapToKeyValue(fieldKeyFormat string) (oci_containerengine.KeyValue, error) {
	result := oci_containerengine.KeyValue{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
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
		result["error"] = []interface{}{NodeErrorToMap(obj.NodeError)}
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

func NodeErrorToMap(obj *oci_containerengine.NodeError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.OpcRequestId != nil {
		result["opc_request_id"] = string(*obj.OpcRequestId)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	return result
}

func (s *ContainerengineNodePoolResourceCrud) mapToNodePoolPlacementConfigDetails(fieldKeyFormat string) (oci_containerengine.NodePoolPlacementConfigDetails, error) {
	result := oci_containerengine.NodePoolPlacementConfigDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NodePoolPlacementConfigDetailsToMap(obj oci_containerengine.NodePoolPlacementConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}
