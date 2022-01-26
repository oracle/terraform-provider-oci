// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_bds "github.com/oracle/oci-go-sdk/v56/bds"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func BdsBdsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.ThreeHours,
			Update: &tfresource.ThreeHours,
			Delete: &tfresource.ThreeHours,
		},
		Create: createBdsBdsInstance,
		Read:   readBdsBdsInstance,
		Update: updateBdsBdsInstance,
		Delete: deleteBdsBdsInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"cluster_public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_high_availability": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"is_secure": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"master_node": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 2),
						},
					},
				},
			},
			"util_node": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 2),
						},
					},
				},
			},
			"worker_node": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(3),
						},
					},
				},
			},

			// Optional
			"cloud_sql_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_volume_size_in_gbs": {
							Type:     schema.TypeString,
							Required: true,
						},

						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_kerberos_mapped_to_database_users": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"kerberos_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"keytab_file": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"principal_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
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
			"network_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"is_nat_gateway_required": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"cluster_details": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ambari_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bd_cell_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bda_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bdm_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bds_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"big_data_manager_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloudera_manager_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"csql_cell_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hue_server_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_refreshed": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cloud_sql_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"node_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attached_block_volumes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"volume_attachment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"volume_size_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssh_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"number_of_nodes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
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

func createBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	cloudSqlRequest := oci_bds.AddCloudSqlRequest{}
	cloudSql := false

	if cloudSqlConfigured, ok := sync.D.GetOkExists("is_cloud_sql_configured"); ok {
		if cloudSqlConfigured.(bool) {
			cloudSql = true
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_sql_details", 0)
			if blockVolumeSizeInGBs, ok := sync.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
				tmp := blockVolumeSizeInGBs.(string)
				tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
				}
				cloudSqlRequest.BlockVolumeSizeInGBs = &tmpInt64
			} else {
				return fmt.Errorf("block_volume_size_in_gbs is required in cloud_sql_details")
			}

			if shape, ok := sync.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
				tmp := shape.(string)
				cloudSqlRequest.Shape = &tmp
			} else {
				return fmt.Errorf("shape is required in cloud_sql_details")
			}
		}
	}

	if err := tfresource.CreateResource(d, sync); err != nil {
		return err
	}

	if cloudSql {
		id := sync.D.Id()
		cloudSqlRequest.BdsInstanceId = &id
		if clusterAdminPassword, ok := sync.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			cloudSqlRequest.ClusterAdminPassword = &tmp
		}
		if err := sync.addCloudSql(cloudSqlRequest); err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}
	return nil
}

func readBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsInstance
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BdsBdsInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
		string(oci_bds.BdsInstanceLifecycleStateFailed),
	}
}

func (s *BdsBdsInstanceResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateUpdating),
	}
}

func (s *BdsBdsInstanceResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceResourceCrud) Create() error {
	request := oci_bds.CreateBdsInstanceRequest{}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if clusterPublicKey, ok := s.D.GetOkExists("cluster_public_key"); ok {
		tmp := clusterPublicKey.(string)
		request.ClusterPublicKey = &tmp
	}

	if clusterVersion, ok := s.D.GetOkExists("cluster_version"); ok {
		request.ClusterVersion = oci_bds.BdsInstanceClusterVersionEnum(clusterVersion.(string))
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isHighAvailability, ok := s.D.GetOkExists("is_high_availability"); ok {
		tmp := isHighAvailability.(bool)
		request.IsHighAvailability = &tmp
	}

	if isSecure, ok := s.D.GetOkExists("is_secure"); ok {
		tmp := isSecure.(bool)
		request.IsSecure = &tmp
	}

	if networkConfig, ok := s.D.GetOkExists("network_config"); ok {
		if tmpList := networkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_config", 0)
			tmp, err := s.mapToNetworkConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfig = &tmp
		}
	}

	numOfNode := 0
	if _, ok := s.D.GetOkExists("master_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "master_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}
	if _, ok := s.D.GetOkExists("util_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "util_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}
	if _, ok := s.D.GetOkExists("worker_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "worker_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}

	createNodeDetails := make([]oci_bds.CreateNodeDetails, numOfNode)
	currentPos := 0

	if nodes, ok := s.D.GetOkExists("master_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "master_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "MASTER")
			if err != nil {
				return err
			}
			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}
	if nodes, ok := s.D.GetOkExists("util_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "util_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "UTILITY")
			if err != nil {
				return err
			}

			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}

	if nodes, ok := s.D.GetOkExists("worker_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "worker_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "WORKER")
			if err != nil {
				return err
			}

			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}

	request.Nodes = createNodeDetails

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceResourceCrud) getBdsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceId, err := bdsInstanceWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceId)

	return s.Get()
}

func bdsInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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
			if res.ActionType == action || res.ActionType == oci_bds.ActionTypesInProgress {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceResourceCrud) Get() error {
	request := oci_bds.GetBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsInstance
	return nil
}

func (s *BdsBdsInstanceResourceCrud) Update() error {
	if cloudSqlConfigured, ok := s.D.GetOkExists("is_cloud_sql_configured"); ok && s.D.HasChange("is_cloud_sql_configured") {
		oldRaw, newRaw := s.D.GetChange("is_cloud_sql_configured")
		if newRaw != "" && oldRaw != "" {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_sql_details", 0)
			if cloudSqlConfigured.(bool) {
				request := oci_bds.AddCloudSqlRequest{}
				id := s.D.Id()
				request.BdsInstanceId = &id
				if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
					tmp := clusterAdminPassword.(string)
					request.ClusterAdminPassword = &tmp
				}
				if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
					tmp := shape.(string)
					request.Shape = &tmp
				} else {
					return fmt.Errorf("shape is required in cloud_sql_details")
				}
				if blockVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
					tmp := blockVolumeSizeInGBs.(string)
					tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
					if err != nil {
						return fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
					}
					request.BlockVolumeSizeInGBs = &tmpInt64
				} else {
					return fmt.Errorf("block_volume_size_in_gbs is required in cloud_sql_details")
				}

				err := s.addCloudSql(request)
				if err != nil {
					return err
				}
			} else {
				request := oci_bds.RemoveCloudSqlRequest{}
				id := s.D.Id()
				request.BdsInstanceId = &id
				if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
					tmp := clusterAdminPassword.(string)
					request.ClusterAdminPassword = &tmp

				}
				err := s.deleteCloudSql(request)
				if err != nil {
					return err
				}

			}
		}
	}

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	workerNodeFieldKeyFormat := "worker_node.0.%s"
	masterNodeFieldKeyFormat := "master_node.0.%s"
	utilNodeFieldKeyFormat := "util_node.0.%s"
	cloudSqlNodeFieldKeyFormat := "cloud_sql_details.0.%s"

	_, blockVolumeSizeInGbsPresent := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs"))
	if blockVolumeSizeInGbsPresent && s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs"))

		tmpOld := oldRaw.(string)
		tmpInt64Old, err := strconv.ParseInt(tmpOld, 10, 64)
		if err != nil {
			return err
		}

		tmpNew := newRaw.(string)
		tmpInt64New, err := strconv.ParseInt(tmpNew, 10, 64)

		if err != nil {
			return err
		}

		if tmpInt64New > tmpInt64Old {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				dif := tmpInt64New - tmpInt64Old
				err := s.updateWorkerBlockStorage(s.D.Id(), clusterAdminPassword, dif)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}

	_, numOfWorkersPresent := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes"))
	if numOfWorkersPresent && s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes"))
		tmpOld := oldRaw.(int)
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				err := s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}

	result := oci_bds.ChangeShapeNodes{}

	changeShapeRequest := oci_bds.ChangeShapeRequest{}
	workerNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "shape"))
	if ok && s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "shape")) {
		tmp := workerNodeShape.(string)
		result.Worker = &tmp
	}
	masterNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(masterNodeFieldKeyFormat, "shape"))
	if ok && s.D.HasChange(fmt.Sprintf(masterNodeFieldKeyFormat, "shape")) {
		tmp := masterNodeShape.(string)
		result.Master = &tmp
	}

	utilNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(utilNodeFieldKeyFormat, "shape"))
	if ok && s.D.HasChange(fmt.Sprintf(utilNodeFieldKeyFormat, "shape")) {
		tmp := utilNodeShape.(string)
		result.Utility = &tmp
	}

	if _, ok := s.D.GetOkExists("is_cloud_sql_configured"); ok {
		cloudSqlNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(cloudSqlNodeFieldKeyFormat, "shape"))
		if ok && s.D.HasChange(fmt.Sprintf(cloudSqlNodeFieldKeyFormat, "shape")) {
			tmp := cloudSqlNodeShape.(string)
			result.Cloudsql = &tmp
		}
	}
	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		clusterAdminPasswordTmp := clusterAdminPassword.(string)
		changeShapeRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

		changeShapeRequest.Nodes = &result
		if !reflect.DeepEqual(result, oci_bds.ChangeShapeNodes{}) {
			tmp := s.D.Id()
			changeShapeRequest.BdsInstanceId = &tmp

			response, err := s.Client.ChangeShape(context.Background(), changeShapeRequest)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId
			err = s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		}
	}

	request := oci_bds.UpdateBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) Delete() error {
	request := oci_bds.DeleteBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceWaitForWorkRequest(workId, "bds",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceResourceCrud) SetData() error {
	if s.Res.IsCloudSqlConfigured != nil {
		s.D.Set("is_cloud_sql_configured", *s.Res.IsCloudSqlConfigured)
	}

	if s.Res.CloudSqlDetails != nil {
		s.D.Set("cloud_sql_details", []interface{}{CloudSqlDetailsToMap(s.Res.CloudSqlDetails)})
	} else {
		s.D.Set("cloud_sql_details", []interface{}{})
		s.D.Set("is_cloud_sql_configured", false)
	}

	if s.Res.ClusterDetails != nil {
		s.D.Set("cluster_details", []interface{}{ClusterDetailsToMap(s.Res.ClusterDetails)})
	} else {
		s.D.Set("cluster_details", nil)
	}

	s.D.Set("cluster_version", s.Res.ClusterVersion)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsHighAvailability != nil {
		s.D.Set("is_high_availability", *s.Res.IsHighAvailability)
	}

	if s.Res.IsSecure != nil {
		s.D.Set("is_secure", *s.Res.IsSecure)
	}

	if s.Res.NetworkConfig != nil {
		s.D.Set("network_config", []interface{}{NetworkConfigToMap(s.Res.NetworkConfig)})
	} else {
		s.D.Set("network_config", nil)
	}

	nodes := []interface{}{}
	nodeMap := make(map[string]map[string]interface{})
	for _, item := range s.Res.Nodes {
		node := BdsNodeToMap(item)
		nodes = append(nodes, node)
		PopulateNodeTemplate(item, nodeMap)
	}
	s.D.Set("nodes", nodes)
	s.D.Set("master_node", []interface{}{nodeMap["MASTER"]})
	s.D.Set("util_node", []interface{}{nodeMap["UTILITY"]})
	s.D.Set("worker_node", []interface{}{nodeMap["WORKER"]})

	if s.Res.NumberOfNodes != nil {
		s.D.Set("number_of_nodes", *s.Res.NumberOfNodes)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func CloudSqlDetailsToMap(obj *oci_bds.CloudSqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockVolumeSizeInGBs != nil {
		result["block_volume_size_in_gbs"] = strconv.FormatInt(*obj.BlockVolumeSizeInGBs, 10)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsKerberosMappedToDatabaseUsers != nil {
		result["is_kerberos_mapped_to_database_users"] = bool(*obj.IsKerberosMappedToDatabaseUsers)
	}

	kerberosDetails := []interface{}{}
	for _, item := range obj.KerberosDetails {
		kerberosDetails = append(kerberosDetails, KerberosDetailsToMap(item))
	}
	result["kerberos_details"] = kerberosDetails

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	return result
}

func ClusterDetailsToMap(obj *oci_bds.ClusterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AmbariUrl != nil {
		result["ambari_url"] = string(*obj.AmbariUrl)
	}

	if obj.BdCellVersion != nil {
		result["bd_cell_version"] = string(*obj.BdCellVersion)
	}

	if obj.BdaVersion != nil {
		result["bda_version"] = string(*obj.BdaVersion)
	}

	if obj.BdmVersion != nil {
		result["bdm_version"] = string(*obj.BdmVersion)
	}

	if obj.BdsVersion != nil {
		result["bds_version"] = string(*obj.BdsVersion)
	}

	if obj.BigDataManagerUrl != nil {
		result["big_data_manager_url"] = string(*obj.BigDataManagerUrl)
	}

	if obj.ClouderaManagerUrl != nil {
		result["cloudera_manager_url"] = string(*obj.ClouderaManagerUrl)
	}

	if obj.CsqlCellVersion != nil {
		result["csql_cell_version"] = string(*obj.CsqlCellVersion)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.HueServerUrl != nil {
		result["hue_server_url"] = string(*obj.HueServerUrl)
	}

	if obj.OsVersion != nil {
		result["os_version"] = string(*obj.OsVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeRefreshed != nil {
		result["time_refreshed"] = obj.TimeRefreshed.String()
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) mapToCreateNodeDetails(fieldKeyFormat, nodeType string) (oci_bds.CreateNodeDetails, error) {
	result := oci_bds.CreateNodeDetails{}

	if blockVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
		tmp := blockVolumeSizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.BlockVolumeSizeInGBs = &tmpInt64
	}

	result.NodeType = oci_bds.NodeNodeTypeEnum(nodeType)

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func BdsNodeToMap(obj oci_bds.Node) map[string]interface{} {
	result := map[string]interface{}{}

	attachedBlockVolumes := []interface{}{}
	for _, item := range obj.AttachedBlockVolumes {
		attachedBlockVolumes = append(attachedBlockVolumes, VolumeAttachmentDetailToMap(item))
	}

	result["attached_block_volumes"] = attachedBlockVolumes

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	result["node_type"] = string(obj.NodeType)

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SshFingerprint != nil {
		result["ssh_fingerprint"] = string(*obj.SshFingerprint)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func KerberosDetailsToMap(obj oci_bds.KerberosDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeytabFile != nil {
		result["keytab_file"] = string(*obj.KeytabFile)
	}

	if obj.PrincipalName != nil {
		result["principal_name"] = string(*obj.PrincipalName)
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) mapToNetworkConfig(fieldKeyFormat string) (oci_bds.NetworkConfig, error) {
	result := oci_bds.NetworkConfig{}

	if cidrBlock, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block")); ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	if isNatGatewayRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_nat_gateway_required")); ok {
		tmp := isNatGatewayRequired.(bool)
		result.IsNatGatewayRequired = &tmp
	}

	return result, nil
}

func NetworkConfigToMap(obj *oci_bds.NetworkConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.IsNatGatewayRequired != nil {
		result["is_nat_gateway_required"] = bool(*obj.IsNatGatewayRequired)
	}

	return result
}

func VolumeAttachmentDetailToMap(obj oci_bds.VolumeAttachmentDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VolumeAttachmentId != nil {
		result["volume_attachment_id"] = string(*obj.VolumeAttachmentId)
	}

	if obj.VolumeSizeInGBs != nil {
		result["volume_size_in_gbs"] = strconv.FormatInt(*obj.VolumeSizeInGBs, 10)
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_bds.ChangeBdsInstanceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BdsInstanceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ChangeBdsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateWorkerBlockStorage(id string, clusterAdminPassword interface{}, blockVolumeSizeInGBs int64) error {
	addBlockStorageRequest := oci_bds.AddBlockStorageRequest{}

	addBlockStorageRequest.BdsInstanceId = &id

	tmpClusterAdminPassword := clusterAdminPassword.(string)
	addBlockStorageRequest.ClusterAdminPassword = &tmpClusterAdminPassword

	addBlockStorageRequest.BlockVolumeSizeInGBs = &blockVolumeSizeInGBs

	addBlockStorageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddBlockStorage(context.Background(), addBlockStorageRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateWorkerNode(id string, clusterAdminPassword interface{}, numberOfWorker int) error {
	addWorkerNodesRequest := oci_bds.AddWorkerNodesRequest{}

	addWorkerNodesRequest.BdsInstanceId = &id

	clusterAdminPasswordTmp := clusterAdminPassword.(string)
	addWorkerNodesRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

	addWorkerNodesRequest.NumberOfWorkerNodes = &numberOfWorker

	addWorkerNodesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddWorkerNodes(context.Background(), addWorkerNodesRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) addCloudSql(request oci_bds.AddCloudSqlRequest) error {
	response, err := s.Client.AddCloudSql(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) deleteCloudSql(request oci_bds.RemoveCloudSqlRequest) error {
	response, err := s.Client.RemoveCloudSql(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func PopulateNodeTemplate(obj oci_bds.Node, nodeMap map[string]map[string]interface{}) {
	switch nodeType := string(obj.NodeType); nodeType {
	case "MASTER":
		if node, ok := nodeMap["MASTER"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["MASTER"] = BdsNodeToTemplateMap(obj)
		}
	case "UTILITY":
		if node, ok := nodeMap["UTILITY"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["UTILITY"] = BdsNodeToTemplateMap(obj)
		}
	case "WORKER":
		if node, ok := nodeMap["WORKER"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["WORKER"] = BdsNodeToTemplateMap(obj)
		}
	}
}

func BdsNodeToTemplateMap(obj oci_bds.Node) map[string]interface{} {
	result := map[string]interface{}{}

	totalSize := int64(0)
	for _, item := range obj.AttachedBlockVolumes {
		if item.VolumeSizeInGBs != nil {
			totalSize += *item.VolumeSizeInGBs
		}
	}
	result["block_volume_size_in_gbs"] = strconv.FormatInt(totalSize, 10)

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	result["number_of_nodes"] = 1

	return result
}
