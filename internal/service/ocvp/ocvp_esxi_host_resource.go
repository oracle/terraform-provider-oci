// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v58/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func OcvpEsxiHostResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createOcvpEsxiHost,
		Read:   readOcvpEsxiHost,
		Update: updateOcvpEsxiHost,
		Delete: deleteOcvpEsxiHost,
		Schema: map[string]*schema.Schema{
			// Required
			"sddc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compute_availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"current_sku": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failed_esxi_host_id": {
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
			"next_sku": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"billing_contract_end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grace_period_end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacement_esxi_host_id": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpEsxiHostResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.EsxiHostClient
	Res                    *oci_ocvp.EsxiHost
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpEsxiHostResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpEsxiHostResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpEsxiHostResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpEsxiHostResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpEsxiHostResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpEsxiHostResourceCrud) Create() error {
	request := oci_ocvp.CreateEsxiHostRequest{}

	if computeAvailabilityDomain, ok := s.D.GetOkExists("compute_availability_domain"); ok {
		tmp := computeAvailabilityDomain.(string)
		request.ComputeAvailabilityDomain = &tmp
	}

	if currentSku, ok := s.D.GetOkExists("current_sku"); ok {
		request.CurrentSku = oci_ocvp.SkuEnum(currentSku.(string))
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

	if failedEsxiHostId, ok := s.D.GetOkExists("failed_esxi_host_id"); ok {
		tmp := failedEsxiHostId.(string)
		request.FailedEsxiHostId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nextSku, ok := s.D.GetOkExists("next_sku"); ok {
		request.NextSku = oci_ocvp.SkuEnum(nextSku.(string))
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEsxiHostFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpEsxiHostResourceCrud) getEsxiHostFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	esxiHostId, err := esxiHostWaitForWorkRequest(workId, "esxihost",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*esxiHostId)

	return s.Get()
}

func esxiHostWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func esxiHostWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = esxiHostWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpEsxiHostWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpEsxiHostWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpEsxiHostResourceCrud) Get() error {
	request := oci_ocvp.GetEsxiHostRequest{}

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EsxiHost
	return nil
}

func (s *OcvpEsxiHostResourceCrud) Update() error {
	request := oci_ocvp.UpdateEsxiHostRequest{}

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

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nextSku, ok := s.D.GetOkExists("next_sku"); ok {
		request.NextSku = oci_ocvp.SkuEnum(nextSku.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EsxiHost
	return nil
}

func (s *OcvpEsxiHostResourceCrud) Delete() error {
	request := oci_ocvp.DeleteEsxiHostRequest{}

	tmp := s.D.Id()
	request.EsxiHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := esxiHostWaitForWorkRequest(workId, "esxihost",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpEsxiHostResourceCrud) SetData() error {
	if s.Res.BillingContractEndDate != nil {
		s.D.Set("billing_contract_end_date", s.Res.BillingContractEndDate.String())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	s.D.Set("current_sku", s.Res.CurrentSku)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedEsxiHostId != nil {
		s.D.Set("failed_esxi_host_id", *s.Res.FailedEsxiHostId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracePeriodEndDate != nil {
		s.D.Set("grace_period_end_date", s.Res.GracePeriodEndDate.String())
	}

	s.D.Set("next_sku", s.Res.NextSku)

	if s.Res.ReplacementEsxiHostId != nil {
		s.D.Set("replacement_esxi_host_id", *s.Res.ReplacementEsxiHostId)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
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

func EsxiHostSummaryToMap(obj oci_ocvp.EsxiHostSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillingContractEndDate != nil {
		result["billing_contract_end_date"] = obj.BillingContractEndDate.String()
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = string(*obj.ComputeAvailabilityDomain)
	}

	if obj.ComputeInstanceId != nil {
		result["compute_instance_id"] = string(*obj.ComputeInstanceId)
	}

	result["current_sku"] = string(obj.CurrentSku)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FailedEsxiHostId != nil {
		result["failed_esxi_host_id"] = string(*obj.FailedEsxiHostId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GracePeriodEndDate != nil {
		result["grace_period_end_date"] = obj.GracePeriodEndDate.String()
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["next_sku"] = string(obj.NextSku)

	if obj.ReplacementEsxiHostId != nil {
		result["replacement_esxi_host_id"] = string(*obj.ReplacementEsxiHostId)
	}

	if obj.SddcId != nil {
		result["sddc_id"] = string(*obj.SddcId)
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
