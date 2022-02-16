// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oce

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_oce "github.com/oracle/oci-go-sdk/v58/oce"
)

func OceOceInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createOceOceInstance,
		Read:   readOceOceInstance,
		Update: updateOceOceInstance,
		Delete: deleteOceOceInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_access_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				StateFunc: utils.GetMd5Hash,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object_storage_namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"instance_access_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_license_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_usage_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"waf_primary_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"guid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_tenancy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OceOceInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oce.OceInstanceClient
	Res                    *oci_oce.OceInstance
	DisableNotFoundRetries bool
}

func (s *OceOceInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OceOceInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateCreating),
	}
}

func (s *OceOceInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateActive),
	}
}

func (s *OceOceInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateDeleting),
	}
}

func (s *OceOceInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateDeleted),
	}
}

func (s *OceOceInstanceResourceCrud) Create() error {
	request := oci_oce.CreateOceInstanceRequest{}

	if adminEmail, ok := s.D.GetOkExists("admin_email"); ok {
		tmp := adminEmail.(string)
		request.AdminEmail = &tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if instanceAccessType, ok := s.D.GetOkExists("instance_access_type"); ok {
		request.InstanceAccessType = oci_oce.CreateOceInstanceDetailsInstanceAccessTypeEnum(instanceAccessType.(string))
	}

	if instanceLicenseType, ok := s.D.GetOkExists("instance_license_type"); ok {
		request.InstanceLicenseType = oci_oce.LicenseTypeEnum(instanceLicenseType.(string))
	}

	if instanceUsageType, ok := s.D.GetOkExists("instance_usage_type"); ok {
		request.InstanceUsageType = oci_oce.CreateOceInstanceDetailsInstanceUsageTypeEnum(instanceUsageType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStorageNamespace, ok := s.D.GetOkExists("object_storage_namespace"); ok {
		tmp := objectStorageNamespace.(string)
		request.ObjectStorageNamespace = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	if tenancyName, ok := s.D.GetOkExists("tenancy_name"); ok {
		tmp := tenancyName.(string)
		request.TenancyName = &tmp
	}

	if upgradeSchedule, ok := s.D.GetOkExists("upgrade_schedule"); ok {
		request.UpgradeSchedule = oci_oce.OceInstanceUpgradeScheduleEnum(upgradeSchedule.(string))
	}

	if wafPrimaryDomain, ok := s.D.GetOkExists("waf_primary_domain"); ok {
		tmp := wafPrimaryDomain.(string)
		request.WafPrimaryDomain = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.CreateOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOceInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce"), oci_oce.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OceOceInstanceResourceCrud) getOceInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_oce.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	oceInstanceId, err := oceInstanceWaitForWorkRequest(workId, "oce",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, oceInstanceId)
		return err
	}

	if oceInstanceId == nil {
		return fmt.Errorf("operation failed: %v for identifier: %v\n", workId, oceInstanceId)
	}

	s.D.SetId(*oceInstanceId)

	return s.Get()
}

func oceInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "oce", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_oce.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func oceInstanceWaitForWorkRequest(wId *string, entityType string, action oci_oce.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_oce.OceInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "oce")
	retryPolicy.ShouldRetryOperation = oceInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_oce.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_oce.WorkRequestStatusInProgress),
			string(oci_oce.WorkRequestStatusAccepted),
			string(oci_oce.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_oce.WorkRequestStatusSucceeded),
			string(oci_oce.WorkRequestStatusFailed),
			string(oci_oce.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_oce.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_oce.WorkRequestStatusFailed || response.Status == oci_oce.WorkRequestStatusCanceled {
		return nil, getErrorFromOceOceInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOceOceInstanceWorkRequest(client *oci_oce.OceInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_oce.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_oce.ListWorkRequestErrorsRequest{
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

func (s *OceOceInstanceResourceCrud) Get() error {
	request := oci_oce.GetOceInstanceRequest{}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.GetOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OceInstance
	return nil
}

func (s *OceOceInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_oce.UpdateOceInstanceRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceLicenseType, ok := s.D.GetOkExists("instance_license_type"); ok {
		request.InstanceLicenseType = oci_oce.LicenseTypeEnum(instanceLicenseType.(string))
	}

	if instanceUsageType, ok := s.D.GetOkExists("instance_usage_type"); ok {
		request.InstanceUsageType = oci_oce.UpdateOceInstanceDetailsInstanceUsageTypeEnum(instanceUsageType.(string))
	}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	if wafPrimaryDomain, ok := s.D.GetOkExists("waf_primary_domain"); ok {
		tmp := wafPrimaryDomain.(string)
		request.WafPrimaryDomain = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.UpdateOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOceInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce"), oci_oce.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OceOceInstanceResourceCrud) Delete() error {
	request := oci_oce.DeleteOceInstanceRequest{}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.DeleteOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oceInstanceWaitForWorkRequest(workId, "oce",
		oci_oce.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OceOceInstanceResourceCrud) SetData() error {
	if s.Res.AdminEmail != nil {
		s.D.Set("admin_email", *s.Res.AdminEmail)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Guid != nil {
		s.D.Set("guid", *s.Res.Guid)
	}

	if s.Res.IdcsTenancy != nil {
		s.D.Set("idcs_tenancy", *s.Res.IdcsTenancy)
	}

	s.D.Set("instance_access_type", s.Res.InstanceAccessType)

	s.D.Set("instance_license_type", s.Res.InstanceLicenseType)

	s.D.Set("instance_usage_type", s.Res.InstanceUsageType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStorageNamespace != nil {
		s.D.Set("object_storage_namespace", *s.Res.ObjectStorageNamespace)
	}

	s.D.Set("service", tfresource.GenericMapToJsonMap(s.Res.Service))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TenancyName != nil {
		s.D.Set("tenancy_name", *s.Res.TenancyName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("upgrade_schedule", s.Res.UpgradeSchedule)

	if s.Res.WafPrimaryDomain != nil {
		s.D.Set("waf_primary_domain", *s.Res.WafPrimaryDomain)
	}

	return nil
}

func (s *OceOceInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_oce.ChangeOceInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OceInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.ChangeOceInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOceInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oce"), oci_oce.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
