// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v58/containerengine"
)

func ContainerengineNodePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("50m"),
			Update: tfresource.GetTimeoutDuration("50m"),
			Delete: tfresource.GetTimeoutDuration("50m"),
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
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
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
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"quantity_per_subnet", "subnet_ids"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"placement_configs": {
							Type:     schema.TypeSet,
							Required: true,
							Set:      placementConfigsHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"availability_domain": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"capacity_reservation_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      utils.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"node_image_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"node_image_name", "node_source_details"},
				Deprecated:    tfresource.FieldDeprecatedAndOverridenByAnother("node_image_id", "node_source_details"),
			},
			"node_image_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"node_image_id", "node_source_details"},
				Deprecated:    tfresource.FieldDeprecatedAndOverridenByAnother("node_image_name", "node_source_details"),
			},
			"node_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"node_shape_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"ocpus": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"node_source_details": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"node_image_id", "node_image_name"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"image_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IMAGE",
							}, true),
						},

						// Optional
						"boot_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
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
			},
			"subnet_ids": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"node_config_details"},
				Set:           utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"node_source": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
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
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kubernetes_version": {
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
						"private_ip": {
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
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func createContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineNodePoolResourceCrud struct {
	tfresource.BaseCrud
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

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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
		request.NodeMetadata = utils.ObjectMapToStringMap(nodeMetadata.(map[string]interface{}))
	}

	if nodeShape, ok := s.D.GetOkExists("node_shape"); ok {
		tmp := nodeShape.(string)
		request.NodeShape = &tmp
	}

	if nodeShapeConfig, ok := s.D.GetOkExists("node_shape_config"); ok {
		if tmpList := nodeShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_shape_config", 0)
			tmp, err := s.mapToCreateNodeShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NodeShapeConfig = &tmp
		}
	}

	if nodeSourceDetails, ok := s.D.GetOkExists("node_source_details"); ok {
		if tmpList := nodeSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_source_details", 0)
			tmp, err := s.mapToNodeSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NodeSourceDetails = tmp
		}
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	nodePoolID, err := nodePoolWaitForWorkRequest(workId, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		if nodePoolID != nil {

			log.Printf("[DEBUG] creation failed, attempting to delete the node pool: %v\n", nodePoolID)

			delReq := oci_containerengine.DeleteNodePoolRequest{}
			delReq.NodePoolId = nodePoolID
			delReq.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

			delRes, delErr := s.Client.DeleteNodePool(context.Background(), delReq)
			if delErr != nil {
				return err
			}
			delWorkRequest := delRes.OpcWorkRequestId

			_, delErr = nodePoolWaitForWorkRequest(delWorkRequest, "nodepool",
				oci_containerengine.WorkRequestResourceActionTypeDeleted,
				s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
			if delErr != nil {
				log.Printf("[DEBUG] cleanup delWorkRequest failed with the error: %v\n", delErr)
			}
		}
		return err
	}

	requestGet := oci_containerengine.GetNodePoolRequest{}
	requestGet.NodePoolId = nodePoolID
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetNodePool(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.NodePool
	return nil
}

func (s *ContainerengineNodePoolResourceCrud) getNodePoolFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_containerengine.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	nodePoolId, err := nodePoolWaitForWorkRequest(workId, "nodepool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*nodePoolId)

	return s.Get()
}

func nodePoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerengine", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func nodePoolWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = nodePoolWorkRequestShouldRetryFunc(timeout)

	response := oci_containerengine.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_containerengine.WorkRequestStatusInProgress),
			string(oci_containerengine.WorkRequestStatusAccepted),
			string(oci_containerengine.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_containerengine.WorkRequestStatusSucceeded),
			string(oci_containerengine.WorkRequestStatusFailed),
			string(oci_containerengine.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_containerengine.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	// Set PollInterval to 1 for replay mode.
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			if res.ActionType == action {
				return res.Identifier, nil
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found.
	if identifier == nil {
		return nil, getErrorFromContainerengineNodePoolWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	// The workrequest may have failed, check for errors if work failed or got cancelled
	if response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return identifier, getErrorFromContainerengineNodePoolWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineNodePoolWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_containerengine.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			CompartmentId: compartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *ContainerengineNodePoolResourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolRequest{}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NodePool
	return nil
}

func (s *ContainerengineNodePoolResourceCrud) Update() error {
	request := oci_containerengine.UpdateNodePoolRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if nodeConfigDetails, ok := s.D.GetOkExists("node_config_details"); ok && s.D.HasChange("node_config_details") {
		if tmpList := nodeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_config_details", 0)
			_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "placement_configs"))
			_, exists := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size"))
			if (ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "placement_configs"))) || (exists && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "size"))) {
				tmp, err := s.mapToUpdateNodePoolNodeConfigDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.NodeConfigDetails = &tmp
			}
		}
	}

	if nodeMetadata, ok := s.D.GetOkExists("node_metadata"); ok {
		request.NodeMetadata = utils.ObjectMapToStringMap(nodeMetadata.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	if nodeShape, ok := s.D.GetOkExists("node_shape"); ok {
		tmp := nodeShape.(string)
		request.NodeShape = &tmp
	}

	if nodeShapeConfig, ok := s.D.GetOkExists("node_shape_config"); ok {
		if tmpList := nodeShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_shape_config", 0)
			tmp, err := s.mapToUpdateNodeShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NodeShapeConfig = &tmp
		}
	}

	if nodeSourceDetails, ok := s.D.GetOkExists("node_source_details"); ok {
		if tmpList := nodeSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "node_source_details", 0)
			tmp, err := s.mapToNodeSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NodeSourceDetails = tmp
		}
	}

	if quantityPerSubnet, ok := s.D.GetOkExists("quantity_per_subnet"); ok && s.D.HasChange("quantity_per_subnet") {
		tmp := quantityPerSubnet.(int)
		request.QuantityPerSubnet = &tmp
	}

	if sshPublicKey, ok := s.D.GetOkExists("ssh_public_key"); ok {
		tmp := sshPublicKey.(string)
		request.SshPublicKey = &tmp
	}

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok && s.D.HasChange("subnet_ids") {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNodePoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineNodePoolResourceCrud) Delete() error {
	request := oci_containerengine.DeleteNodePoolRequest{}

	tmp := s.D.Id()
	request.NodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DeleteNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := nodePoolWaitForWorkRequest(workId, "nodepool",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineNodePoolResourceCrud) SetData() error {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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
		s.D.Set("node_config_details", []interface{}{NodePoolNodeConfigDetailsToMap(s.Res.NodeConfigDetails, false)})
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

	if s.Res.NodeShapeConfig != nil {
		s.D.Set("node_shape_config", []interface{}{NodeShapeConfigToMap(s.Res.NodeShapeConfig)})
	} else {
		s.D.Set("node_shape_config", nil)
	}

	if s.Res.NodeSource != nil {
		nodeSourceArray := []interface{}{}
		if nodeSourceMap := NodeSourceOptionToMap(&s.Res.NodeSource); nodeSourceMap != nil {
			nodeSourceArray = append(nodeSourceArray, nodeSourceMap)
		}
		s.D.Set("node_source", nodeSourceArray)
	} else {
		s.D.Set("node_source", nil)
	}

	if s.Res.NodeSourceDetails != nil {
		nodeSourceDetailsArray := []interface{}{}
		if nodeSourceDetailsMap := NodeSourceDetailsToMap(&s.Res.NodeSourceDetails); nodeSourceDetailsMap != nil {
			nodeSourceDetailsArray = append(nodeSourceDetailsArray, nodeSourceDetailsMap)
		}
		s.D.Set("node_source_details", nodeSourceDetailsArray)
	} else {
		s.D.Set("node_source_details", nil)
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
		s.D.Set("subnet_ids", schema.NewSet(utils.LiteralTypeHashCodeForSets, subnetIds))
	} else {
		s.D.Set("subnet_ids", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	return nil
}

func (s *ContainerengineNodePoolResourceCrud) mapToCreateNodePoolNodeConfigDetails(fieldKeyFormat string) (oci_containerengine.CreateNodePoolNodeConfigDetails, error) {
	result := oci_containerengine.CreateNodePoolNodeConfigDetails{}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}
	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if placementConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "placement_configs")); ok {
		set := placementConfigs.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NodePoolPlacementConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := placementConfigsHashCodeForSets(interfaces[i])
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

func NodePoolNodeConfigDetailsToMap(obj *oci_containerengine.NodePoolNodeConfigDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}
	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds)
	}

	placementConfigs := []interface{}{}
	for _, item := range obj.PlacementConfigs {
		placementConfigs = append(placementConfigs, NodePoolPlacementConfigDetailsToMap(item))
	}
	if datasource {
		result["placement_configs"] = placementConfigs
	} else {
		result["placement_configs"] = schema.NewSet(placementConfigsHashCodeForSets, placementConfigs)
	}

	if obj.Size != nil {
		result["size"] = int(*obj.Size)
	}

	return result
}

func (s *ContainerengineNodePoolResourceCrud) mapToCreateNodeShapeConfigDetails(fieldKeyFormat string) (oci_containerengine.CreateNodeShapeConfigDetails, error) {
	result := oci_containerengine.CreateNodeShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func (s *ContainerengineNodePoolResourceCrud) mapToUpdateNodeShapeConfigDetails(fieldKeyFormat string) (oci_containerengine.UpdateNodeShapeConfigDetails, error) {
	result := oci_containerengine.UpdateNodeShapeConfigDetails{}

	if memory_in_gbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memory_in_gbs.(float64))
		// prevent Update with 0 value
		if tmp != 0 {
			result.MemoryInGBs = &tmp
		}
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		// prevent Update with 0 value
		if tmp != 0 {
			result.Ocpus = &tmp
		}
	}

	return result, nil
}

func NodeShapeConfigToMap(obj *oci_containerengine.NodeShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
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

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.NodeError != nil {
		result["error"] = []interface{}{NodeErrorToMap(obj.NodeError)}
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.KubernetesVersion != nil {
		result["kubernetes_version"] = string(*obj.KubernetesVersion)
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

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.PublicIp != nil {
		result["public_ip"] = string(*obj.PublicIp)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
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

	if capacityReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_reservation_id")); ok {
		tmp := capacityReservationId.(string)
		result.CapacityReservationId = &tmp
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

	if obj.CapacityReservationId != nil {
		result["capacity_reservation_id"] = string(*obj.CapacityReservationId)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *ContainerengineNodePoolResourceCrud) mapToNodeSourceDetails(fieldKeyFormat string) (oci_containerengine.NodeSourceDetails, error) {
	var baseObject oci_containerengine.NodeSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("IMAGE"):
		details := oci_containerengine.NodeSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			if tmp != "" {
				tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return details, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
				}
				details.BootVolumeSizeInGBs = &tmpInt64
			}
		}
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func NodeSourceDetailsToMap(obj *oci_containerengine.NodeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_containerengine.NodeSourceViaImageDetails:
		result["source_type"] = "IMAGE"

		if v.BootVolumeSizeInGBs != nil {
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*v.BootVolumeSizeInGBs, 10)
		}

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func NodeSourceOptionToMap(obj *oci_containerengine.NodeSourceOption) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_containerengine.NodeSourceViaImageOption:
		result["source_type"] = "IMAGE"

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}

		if v.SourceName != nil {
			result["source_name"] = string(*v.SourceName)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func placementConfigsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if availabilityDomain, ok := m["availability_domain"]; ok && availabilityDomain != "" {
		buf.WriteString(fmt.Sprintf("%v-", availabilityDomain))
	}
	if capacityReservationId, ok := m["capacity_reservation_id"]; ok && capacityReservationId != "" {
		buf.WriteString(fmt.Sprintf("%v-", capacityReservationId))
	}
	if subnetId, ok := m["subnet_id"]; ok && subnetId != "" {
		buf.WriteString(fmt.Sprintf("%v-", subnetId))
	}
	return hashcode.String(buf.String())
}

func (s *ContainerengineNodePoolResourceCrud) mapToUpdateNodePoolNodeConfigDetails(fieldKeyFormat string) (oci_containerengine.UpdateNodePoolNodeConfigDetails, error) {
	result := oci_containerengine.UpdateNodePoolNodeConfigDetails{}

	result.PlacementConfigs = []oci_containerengine.NodePoolPlacementConfigDetails{}
	if placementConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "placement_configs")); ok {
		set := placementConfigs.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NodePoolPlacementConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := placementConfigsHashCodeForSets(interfaces[i])
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
