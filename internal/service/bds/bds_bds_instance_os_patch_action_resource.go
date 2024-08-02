// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BdsBdsInstanceOSPatchActionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		Create: createBdsBdsInstanceOSPatchAction,
		Read:   readBdsBdsInstanceOSPatchAction,
		Delete: deleteBdsBdsInstanceOSPatchAction,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"os_patch_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"patching_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"patching_config_strategy": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DOWNTIME_BASED",
								"BATCHING_BASED",
								"DOMAIN_BASED",
							}, true),
						},

						"batch_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"wait_time_between_batch_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"tolerance_threshold_per_batch": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"tolerance_threshold_per_domain": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"wait_time_between_domain_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createBdsBdsInstanceOSPatchAction(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceOSPatchActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

// Return object nil for below two func because this is an action-type update operation
func readBdsBdsInstanceOSPatchAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteBdsBdsInstanceOSPatchAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

type BdsBdsInstanceOSPatchActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceOSPatchActionResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("BdsBdsInstanceOSPatchActionResource-", BdsBdsInstanceOSPatchActionResource(), s.D)
}

func (s *BdsBdsInstanceOSPatchActionResourceCrud) Create() error {
	request := oci_bds.InstallOsPatchRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if osPatchVersion, ok := s.D.GetOkExists("os_patch_version"); ok {
		tmp := osPatchVersion.(string)
		request.OsPatchVersion = &tmp
	}

	if patchingConfigs, ok := s.D.GetOkExists("patching_configs"); ok {
		if tmpList := patchingConfigs.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patching_configs", 0)
			tmp, err := s.mapToPatchingConfigs(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PatchingConfigs = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.InstallOsPatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceOSPatchActionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceOSPatchActionResourceCrud) getBdsInstanceOSPatchActionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceOSPatchActionId, err := bdsInstanceOSPatchActionWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceOSPatchActionId)

	return nil
}

func bdsInstanceOSPatchActionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceOSPatchActionWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceOSPatchActionWorkRequestShouldRetryFunc(timeout)

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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceOSPatchActionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceOSPatchActionWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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

func (s *BdsBdsInstanceOSPatchActionResourceCrud) SetData() error {
	return nil
}

func (s *BdsBdsInstanceOSPatchActionResourceCrud) mapToPatchingConfigs(fieldKeyFormat string) (oci_bds.PatchingConfigs, error) {
	var baseObject oci_bds.PatchingConfigs
	//discriminator
	patchingConfigStrategyRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patching_config_strategy"))
	var patchingConfigStrategy string
	if ok {
		patchingConfigStrategy = patchingConfigStrategyRaw.(string)
	} else {
		patchingConfigStrategy = "" // default value
	}
	switch strings.ToLower(patchingConfigStrategy) {

	case strings.ToLower("BATCHING_BASED"):
		result := oci_bds.BatchingBasedPatchingConfigs{}

		if batchSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_size")); ok {
			tmp := batchSize.(int)
			result.BatchSize = &tmp
		}

		if waitTimeBetweenBatchInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_time_between_batch_in_seconds")); ok {
			tmp := waitTimeBetweenBatchInSeconds.(int)
			result.WaitTimeBetweenBatchInSeconds = &tmp
		}

		if toleranceThresholdPerBatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tolerance_threshold_per_batch")); ok {
			tmp := toleranceThresholdPerBatch.(int)
			result.ToleranceThresholdPerBatch = &tmp
		}

		baseObject = result
	case strings.ToLower("DOMAIN_BASED"):
		result := oci_bds.DomainBasedPatchingConfigs{}

		if waitTimeBetweenDomainInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_time_between_domain_in_seconds")); ok {
			tmp := waitTimeBetweenDomainInSeconds.(int)
			result.WaitTimeBetweenDomainInSeconds = &tmp
		}

		if toleranceThresholdPerDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tolerance_threshold_per_domain")); ok {
			tmp := toleranceThresholdPerDomain.(int)
			result.ToleranceThresholdPerDomain = &tmp
		}

		baseObject = result
	case strings.ToLower("DOWNTIME_BASED"):
		result := oci_bds.DowntimeBasedPatchingConfigs{}
		baseObject = result

	default:
		return nil, fmt.Errorf("unknown patching_config_strategy '%v' was specified", patchingConfigStrategy)
	}
	return baseObject, nil
}
