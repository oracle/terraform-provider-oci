// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_blockchain "github.com/oracle/oci-go-sdk/v65/blockchain"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BlockchainBlockchainPlatformResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Update: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createBlockchainBlockchainPlatform,
		Read:   readBlockchainBlockchainPlatform,
		Update: updateBlockchainBlockchainPlatform,
		Delete: deleteBlockchainBlockchainPlatform,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_shape": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: blockchainPlatformComputeShapeDiffSuppressFunction,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"idcs_access_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"platform_role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ca_cert_archive_text": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"federated_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_byol": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"platform_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"load_balancer_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replicas": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ca_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"console_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"proxy_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"total_ocpu_capacity": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"component_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"osns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ad": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocpu_allocation_param": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"ocpu_allocation_number": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"osn_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"peers": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ad": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"alias": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocpu_allocation_param": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"ocpu_allocation_number": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"peer_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"host_ocpu_utilization_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"host": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocpu_capacity_number": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"ocpu_utilization_number": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"is_multi_ad": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_shape_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_used_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBlockchainBlockchainPlatform(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.CreateResource(d, sync)
}

func readBlockchainBlockchainPlatform(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.ReadResource(sync)
}

func updateBlockchainBlockchainPlatform(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBlockchainBlockchainPlatform(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BlockchainBlockchainPlatformResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_blockchain.BlockchainPlatformClient
	Res                    *oci_blockchain.BlockchainPlatform
	DisableNotFoundRetries bool
}

func (s *BlockchainBlockchainPlatformResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BlockchainBlockchainPlatformResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateCreating),
		string(oci_blockchain.BlockchainPlatformLifecycleStateScaling),
	}
}

func (s *BlockchainBlockchainPlatformResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateActive),
	}
}

func (s *BlockchainBlockchainPlatformResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateDeleting),
	}
}

func (s *BlockchainBlockchainPlatformResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateDeleted),
	}
}

func (s *BlockchainBlockchainPlatformResourceCrud) Create() error {
	request := oci_blockchain.CreateBlockchainPlatformRequest{}

	if caCertArchiveText, ok := s.D.GetOkExists("ca_cert_archive_text"); ok {
		tmp := caCertArchiveText.(string)
		request.CaCertArchiveText = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeShape, ok := s.D.GetOkExists("compute_shape"); ok {
		request.ComputeShape = oci_blockchain.BlockchainPlatformComputeShapeEnum(computeShape.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if federatedUserId, ok := s.D.GetOkExists("federated_user_id"); ok {
		tmp := federatedUserId.(string)
		request.FederatedUserId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if isByol, ok := s.D.GetOkExists("is_byol"); ok {
		tmp := isByol.(bool)
		request.IsByol = &tmp
	}

	if platformRole, ok := s.D.GetOkExists("platform_role"); ok {
		request.PlatformRole = oci_blockchain.BlockchainPlatformPlatformRoleEnum(platformRole.(string))
	}

	if platformVersion, ok := s.D.GetOkExists("platform_version"); ok {
		tmp := platformVersion.(string)
		request.PlatformVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.CreateBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_blockchain.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_blockchain.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "instance") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	err = s.getBlockchainPlatformFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	_, replicasExists := s.D.GetOkExists("replicas")
	_, storageSizeInTbsExists := s.D.GetOkExists("storage_size_in_tbs")
	_, totalOpcuCapacityExists := s.D.GetOkExists("total_ocpu_capacity")
	_, loadBalancerShapeExists := s.D.GetOkExists("load_balancer_shape")

	if replicasExists || storageSizeInTbsExists || totalOpcuCapacityExists || loadBalancerShapeExists {
		return s.Update()
	}

	return nil
}

func (s *BlockchainBlockchainPlatformResourceCrud) getBlockchainPlatformFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_blockchain.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	blockchainPlatformId, err := blockchainPlatformWaitForWorkRequest(workId, "instance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, blockchainPlatformId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_blockchain.DeleteWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*blockchainPlatformId)

	return s.Get()
}

func blockchainPlatformWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "blockchain", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_blockchain.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func blockchainPlatformWaitForWorkRequest(wId *string, entityType string, action oci_blockchain.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_blockchain.BlockchainPlatformClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "blockchain")
	retryPolicy.ShouldRetryOperation = blockchainPlatformWorkRequestShouldRetryFunc(timeout)

	response := oci_blockchain.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_blockchain.WorkRequestStatusInProgress),
			string(oci_blockchain.WorkRequestStatusAccepted),
			string(oci_blockchain.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_blockchain.WorkRequestStatusSucceeded),
			string(oci_blockchain.WorkRequestStatusFailed),
			string(oci_blockchain.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_blockchain.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if response.Status == oci_blockchain.WorkRequestStatusFailed {
		errorMessage := getErrorFromBlockchainPlatformWorkRequest(client, wId, retryPolicy, entityType, action)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromBlockchainPlatformWorkRequest(client *oci_blockchain.BlockchainPlatformClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_blockchain.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_blockchain.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
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

func (s *BlockchainBlockchainPlatformResourceCrud) Get() error {
	request := oci_blockchain.GetBlockchainPlatformRequest{}

	tmp := s.D.Id()
	request.BlockchainPlatformId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.GetBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BlockchainPlatform
	return nil
}

func (s *BlockchainBlockchainPlatformResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	defer func() {
		// get latest state of the instance
		err := s.Get()
		if err != nil {
			log.Printf("[ERROR] unable to invoke GET() after UPDATE '%v'", err)
		}
		// write latest state
		if err := s.SetData(); err != nil {
			log.Printf("[ERROR] unable to invoke setData() '%v'", err)
		}
	}()

	request := oci_blockchain.UpdateBlockchainPlatformRequest{}

	tmp := s.D.Id()
	// Service limitation only allow Update 1 field per API call
	request.BlockchainPlatformId = &tmp
	if replicas, ok := s.D.GetOkExists("replicas"); ok && s.D.HasChange("replicas") {
		if tmpList := replicas.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replicas", 0)
			tmp, err := s.mapToReplicaDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Replicas = &tmp
		}
		if err := sendUpdateBlockchainPlatformRequest(s, request); err != nil {
			return err
		}
		request.Replicas = nil
	}

	if storageSizeInTBs, ok := s.D.GetOkExists("storage_size_in_tbs"); ok && s.D.HasChange("storage_size_in_tbs") {
		tmp := storageSizeInTBs.(float64)
		request.StorageSizeInTBs = &tmp
		if err := sendUpdateBlockchainPlatformRequest(s, request); err != nil {
			return err
		}
		request.StorageSizeInTBs = nil
	}

	if totalOcpuCapacity, ok := s.D.GetOkExists("total_ocpu_capacity"); ok && s.D.HasChange("total_ocpu_capacity") {
		tmp := totalOcpuCapacity.(int)
		request.TotalOcpuCapacity = &tmp
		if err := sendUpdateBlockchainPlatformRequest(s, request); err != nil {
			return err
		}
		request.TotalOcpuCapacity = nil
	}

	if loadBalancerShape, ok := s.D.GetOkExists("load_balancer_shape"); ok && s.D.HasChange("load_balancer_shape") {
		request.LoadBalancerShape = oci_blockchain.BlockchainPlatformLoadBalancerShapeEnum(loadBalancerShape.(string))
		if err := sendUpdateBlockchainPlatformRequest(s, request); err != nil {
			return err
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.UpdateBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBlockchainPlatformFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BlockchainBlockchainPlatformResourceCrud) Delete() error {
	request := oci_blockchain.DeleteBlockchainPlatformRequest{}

	tmp := s.D.Id()
	request.BlockchainPlatformId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.DeleteBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := blockchainPlatformWaitForWorkRequest(workId, "instance",
		oci_blockchain.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BlockchainBlockchainPlatformResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentDetails != nil {
		s.D.Set("component_details", []interface{}{BlockchainPlatformComponentDetailsToMap(s.Res.ComponentDetails)})
	} else {
		s.D.Set("component_details", nil)
	}

	s.D.Set("compute_shape", s.Res.ComputeShape)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	hostOcpuUtilizationInfo := []interface{}{}
	for _, item := range s.Res.HostOcpuUtilizationInfo {
		hostOcpuUtilizationInfo = append(hostOcpuUtilizationInfo, OcpuUtilizationInfoToMap(item))
	}
	s.D.Set("host_ocpu_utilization_info", hostOcpuUtilizationInfo)

	if s.Res.IsByol != nil {
		s.D.Set("is_byol", *s.Res.IsByol)
	}

	if s.Res.IsMultiAD != nil {
		s.D.Set("is_multi_ad", *s.Res.IsMultiAD)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("load_balancer_shape", s.Res.LoadBalancerShape)

	s.D.Set("platform_role", s.Res.PlatformRole)

	s.D.Set("platform_shape_type", s.Res.PlatformShapeType)

	if s.Res.PlatformVersion != nil {
		s.D.Set("platform_version", *s.Res.PlatformVersion)
	}

	if s.Res.Replicas != nil {
		s.D.Set("replicas", []interface{}{ReplicaDetailsToMap(s.Res.Replicas)})
	} else {
		s.D.Set("replicas", nil)
	}

	if s.Res.ServiceEndpoint != nil {
		s.D.Set("service_endpoint", *s.Res.ServiceEndpoint)
	}

	if s.Res.ServiceVersion != nil {
		s.D.Set("service_version", *s.Res.ServiceVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageSizeInTBs != nil {
		s.D.Set("storage_size_in_tbs", *s.Res.StorageSizeInTBs)
	}

	if s.Res.StorageUsedInTBs != nil {
		s.D.Set("storage_used_in_tbs", *s.Res.StorageUsedInTBs)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalOcpuCapacity != nil {
		s.D.Set("total_ocpu_capacity", *s.Res.TotalOcpuCapacity)
	}

	return nil
}

func BlockchainPlatformComponentDetailsToMap(obj *oci_blockchain.BlockchainPlatformComponentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	osns := []interface{}{}
	for _, item := range obj.Osns {
		osns = append(osns, OsnToMap(item))
	}
	result["osns"] = osns

	peers := []interface{}{}
	for _, item := range obj.Peers {
		peers = append(peers, PeerToMap(item))
	}
	result["peers"] = peers

	return result
}

func BlockchainPlatformSummaryToMap(obj oci_blockchain.BlockchainPlatformSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compute_shape"] = string(obj.ComputeShape)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["platform_role"] = string(obj.PlatformRole)

	if obj.ServiceEndpoint != nil {
		result["service_endpoint"] = string(*obj.ServiceEndpoint)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func OcpuAllocationNumberParamToMap(obj *oci_blockchain.OcpuAllocationNumberParam) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OcpuAllocationNumber != nil {
		result["ocpu_allocation_number"] = float32(*obj.OcpuAllocationNumber)
	}

	return result
}

func OcpuUtilizationInfoToMap(obj oci_blockchain.OcpuUtilizationInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.OcpuCapacityNumber != nil {
		result["ocpu_capacity_number"] = float32(*obj.OcpuCapacityNumber)
	}

	if obj.OcpuUtilizationNumber != nil {
		result["ocpu_utilization_number"] = float32(*obj.OcpuUtilizationNumber)
	}

	return result
}

func OsnToMap(obj oci_blockchain.Osn) map[string]interface{} {
	result := map[string]interface{}{}

	result["ad"] = string(obj.Ad)

	if obj.OcpuAllocationParam != nil {
		result["ocpu_allocation_param"] = []interface{}{OcpuAllocationNumberParamToMap(obj.OcpuAllocationParam)}
	}

	if obj.OsnKey != nil {
		result["osn_key"] = string(*obj.OsnKey)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func PeerToMap(obj oci_blockchain.Peer) map[string]interface{} {
	result := map[string]interface{}{}

	result["ad"] = string(obj.Ad)

	if obj.Alias != nil {
		result["alias"] = string(*obj.Alias)
	}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.OcpuAllocationParam != nil {
		result["ocpu_allocation_param"] = []interface{}{OcpuAllocationNumberParamToMap(obj.OcpuAllocationParam)}
	}

	if obj.PeerKey != nil {
		result["peer_key"] = string(*obj.PeerKey)
	}

	result["role"] = string(obj.Role)

	result["state"] = string(obj.LifecycleState)

	return result
}

func (s *BlockchainBlockchainPlatformResourceCrud) mapToReplicaDetails(fieldKeyFormat string) (oci_blockchain.ReplicaDetails, error) {
	result := oci_blockchain.ReplicaDetails{}

	if caCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ca_count")); ok {
		tmp := caCount.(int)
		result.CaCount = &tmp
	}

	if consoleCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "console_count")); ok {
		tmp := consoleCount.(int)
		result.ConsoleCount = &tmp
	}

	if proxyCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "proxy_count")); ok {
		tmp := proxyCount.(int)
		result.ProxyCount = &tmp
	}

	return result, nil
}

func ReplicaDetailsToMap(obj *oci_blockchain.ReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CaCount != nil {
		result["ca_count"] = int(*obj.CaCount)
	}

	if obj.ConsoleCount != nil {
		result["console_count"] = int(*obj.ConsoleCount)
	}

	if obj.ProxyCount != nil {
		result["proxy_count"] = int(*obj.ProxyCount)
	}

	return result
}

func (s *BlockchainBlockchainPlatformResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_blockchain.ChangeBlockchainPlatformCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BlockchainPlatformId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.ChangeBlockchainPlatformCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBlockchainPlatformFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
