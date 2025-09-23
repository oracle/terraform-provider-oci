// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceAnalyticsResourceAnalyticsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwoHours,
			Update: &tfresource.TwoHours,
			Delete: &tfresource.TwoHours,
			Read:   &tfresource.TwoHours,
		},
		Create: createResourceAnalyticsResourceAnalyticsInstance,
		Read:   readResourceAnalyticsResourceAnalyticsInstance,
		Update: updateResourceAnalyticsResourceAnalyticsInstance,
		Delete: deleteResourceAnalyticsResourceAnalyticsInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"adw_admin_password": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							Sensitive:        true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PLAIN_TEXT",
								"VAULT_SECRET",
							}, true),
						},

						// Optional
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
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
			"display_name": {
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
			"is_mutual_tls_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"adw_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oac_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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

func createResourceAnalyticsResourceAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceAnalyticsInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readResourceAnalyticsResourceAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceAnalyticsInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateResourceAnalyticsResourceAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceAnalyticsInstanceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteResourceAnalyticsResourceAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceAnalyticsInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourceAnalyticsResourceAnalyticsInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resource_analytics.ResourceAnalyticsInstanceClient
	Res                    *oci_resource_analytics.ResourceAnalyticsInstance
	DisableNotFoundRetries bool
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateCreating),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateActive),
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateNeedsAttention),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateDeleting),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateDeleted),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) Create() error {
	request := oci_resource_analytics.CreateResourceAnalyticsInstanceRequest{}

	if adwAdminPassword, ok := s.D.GetOkExists("adw_admin_password"); ok {
		if tmpList := adwAdminPassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adw_admin_password", 0)
			tmp, err := s.mapToAdwAdminPasswordDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdwAdminPassword = tmp
		}
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isMutualTlsRequired, ok := s.D.GetOkExists("is_mutual_tls_required"); ok {
		tmp := isMutualTlsRequired.(bool)
		request.IsMutualTlsRequired = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_resource_analytics.CreateResourceAnalyticsInstanceDetailsLicenseModelEnum(licenseModel.(string))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.CreateResourceAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_resource_analytics.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_resource_analytics.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "resourceanalyticsinstance") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getResourceAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) getResourceAnalyticsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resource_analytics.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	resourceAnalyticsInstanceId, err := resourceAnalyticsInstanceWaitForWorkRequest(workId, "resourceanalyticsinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, resourceAnalyticsInstanceId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_resource_analytics.CancelWorkRequestRequest{
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
	s.D.SetId(*resourceAnalyticsInstanceId)

	return s.Get()
}

func resourceAnalyticsInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "resource_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_resource_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func resourceAnalyticsInstanceWaitForWorkRequest(wId *string, entityType string, action oci_resource_analytics.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_resource_analytics.ResourceAnalyticsInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resource_analytics")
	retryPolicy.ShouldRetryOperation = resourceAnalyticsInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_resource_analytics.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_resource_analytics.OperationStatusInProgress),
			string(oci_resource_analytics.OperationStatusAccepted),
			string(oci_resource_analytics.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_resource_analytics.OperationStatusSucceeded),
			string(oci_resource_analytics.OperationStatusFailed),
			string(oci_resource_analytics.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_resource_analytics.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_resource_analytics.OperationStatusFailed || response.Status == oci_resource_analytics.OperationStatusCanceled {
		return nil, getErrorFromResourceAnalyticsResourceAnalyticsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourceAnalyticsResourceAnalyticsInstanceWorkRequest(client *oci_resource_analytics.ResourceAnalyticsInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resource_analytics.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_resource_analytics.ListWorkRequestErrorsRequest{
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

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) Get() error {
	request := oci_resource_analytics.GetResourceAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.ResourceAnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.GetResourceAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourceAnalyticsInstance
	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_resource_analytics.UpdateResourceAnalyticsInstanceRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ResourceAnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.UpdateResourceAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getResourceAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) Delete() error {
	request := oci_resource_analytics.DeleteResourceAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.ResourceAnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.DeleteResourceAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := resourceAnalyticsInstanceWaitForWorkRequest(workId, "resourceanalyticsinstance",
		oci_resource_analytics.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) SetData() error {
	if s.Res.AdwId != nil {
		s.D.Set("adw_id", *s.Res.AdwId)
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OacId != nil {
		s.D.Set("oac_id", *s.Res.OacId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) mapToAdwAdminPasswordDetails(fieldKeyFormat string) (oci_resource_analytics.AdwAdminPasswordDetails, error) {
	var baseObject oci_resource_analytics.AdwAdminPasswordDetails
	//discriminator
	passwordTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_type"))
	var passwordType string
	if ok {
		passwordType = passwordTypeRaw.(string)
	} else {
		passwordType = "" // default value
	}
	switch strings.ToLower(passwordType) {
	case strings.ToLower("PLAIN_TEXT"):
		details := oci_resource_analytics.PlainTextPasswordDetails{}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		baseObject = details
	case strings.ToLower("VAULT_SECRET"):
		details := oci_resource_analytics.VaultSecretPasswordDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown password_type '%v' was specified", passwordType)
	}
	return baseObject, nil
}

func AdwAdminPasswordDetailsToMap(obj *oci_resource_analytics.AdwAdminPasswordDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_resource_analytics.PlainTextPasswordDetails:
		result["password_type"] = "PLAIN_TEXT"

		if v.Password != nil {
			result["password"] = string(*v.Password)
		}
	case oci_resource_analytics.VaultSecretPasswordDetails:
		result["password_type"] = "VAULT_SECRET"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'password_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ResourceAnalyticsInstanceSummaryToMap(obj oci_resource_analytics.ResourceAnalyticsInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

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

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_resource_analytics.ChangeResourceAnalyticsInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ResourceAnalyticsInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.ChangeResourceAnalyticsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getResourceAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
