// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

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

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeDataSafeConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeDataSafeConfiguration,
		Read:     readDataSafeDataSafeConfiguration,
		Update:   updateDataSafeDataSafeConfiguration,
		Delete:   deleteDataSafeDataSafeConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"data_safe_nat_gateway_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"global_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_paid_usage": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"offline_retention_period": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"online_retention_period": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeDataSafeConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafeConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeDataSafeConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafeConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeDataSafeConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafeConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeDataSafeConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeDataSafeConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.DataSafeConfiguration
	DisableNotFoundRetries bool
}

func (s *DataSafeDataSafeConfigurationResourceCrud) ID() string {
	return s.Res.TimeEnabled.Format(time.RFC3339Nano) + "datasafe-configuration"
}

func (s *DataSafeDataSafeConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.LifecycleStateCreating),
	}
}

func (s *DataSafeDataSafeConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.LifecycleStateActive),
	}
}

func (s *DataSafeDataSafeConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.LifecycleStateDeleting),
	}
}

func (s *DataSafeDataSafeConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.LifecycleStateDeleted),
	}
}

func (s *DataSafeDataSafeConfigurationResourceCrud) Create() error {
	request := oci_data_safe.EnableDataSafeConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.EnableDataSafeConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && (strings.Contains(strings.ToLower(*res.EntityType), "datasafeinstance") || strings.Contains(strings.ToLower(*res.EntityType), "configuration")) &&
				res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getDataSafeConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeDataSafeConfigurationResourceCrud) getDataSafeConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dataSafeConfigurationId, err := dataSafeConfigurationWaitForWorkRequest(workId, "configuration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, dataSafeConfigurationId)
		return err
	}
	s.D.SetId(*dataSafeConfigurationId)

	return s.Get()
}

func dataSafeConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func dataSafeConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = dataSafeConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), "datasafeinstance") || strings.Contains(strings.ToLower(*res.EntityType), "configuration") {
			if res.ActionType == oci_data_safe.WorkRequestResourceActionTypeCreated || res.ActionType == oci_data_safe.WorkRequestResourceActionTypeUpdated {
				identifier = res.Identifier
				break
			}
		}
	}
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeDataSafeConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}
	return identifier, nil
}

func getErrorFromDataSafeDataSafeConfigurationWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeDataSafeConfigurationResourceCrud) Get() error {
	request := oci_data_safe.GetDataSafeConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetDataSafeConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataSafeConfiguration
	return nil
}

func (s *DataSafeDataSafeConfigurationResourceCrud) Update() error {
	request := oci_data_safe.EnableDataSafeConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.EnableDataSafeConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSafeConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeDataSafeConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSafeNatGatewayIpAddress != nil {
		s.D.Set("data_safe_nat_gateway_ip_address", *s.Res.DataSafeNatGatewayIpAddress)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GlobalSettings != nil {
		s.D.Set("global_settings", []interface{}{GlobalSettingsToMap(s.Res.GlobalSettings)})
	} else {
		s.D.Set("global_settings", nil)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnabled != nil {
		s.D.Set("time_enabled", s.Res.TimeEnabled.String())
	}

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}

func GlobalSettingsToMap(obj *oci_data_safe.GlobalSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPaidUsage != nil {
		result["is_paid_usage"] = bool(*obj.IsPaidUsage)
	}

	if obj.OfflineRetentionPeriod != nil {
		result["offline_retention_period"] = int(*obj.OfflineRetentionPeriod)
	}

	if obj.OnlineRetentionPeriod != nil {
		result["online_retention_period"] = int(*obj.OnlineRetentionPeriod)
	}

	return result
}
