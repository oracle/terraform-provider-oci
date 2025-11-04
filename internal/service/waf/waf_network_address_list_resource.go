// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WafNetworkAddressListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createWafNetworkAddressListWithContext,
		ReadContext:   readWafNetworkAddressListWithContext,
		UpdateContext: updateWafNetworkAddressListWithContext,
		DeleteContext: deleteWafNetworkAddressListWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ADDRESSES",
					"VCN_ADDRESSES",
				}, true),
			},

			// Optional
			"addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"vcn_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"addresses": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"lifecycle_details": {
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

func createWafNetworkAddressListWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &WafNetworkAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readWafNetworkAddressListWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &WafNetworkAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateWafNetworkAddressListWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &WafNetworkAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteWafNetworkAddressListWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &WafNetworkAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type WafNetworkAddressListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waf.WafClient
	Res                    *oci_waf.NetworkAddressList
	DisableNotFoundRetries bool
}

func (s *WafNetworkAddressListResourceCrud) ID() string {
	networkAddressList := *s.Res
	return *networkAddressList.GetId()
}

func (s *WafNetworkAddressListResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waf.NetworkAddressListLifecycleStateCreating),
	}
}

func (s *WafNetworkAddressListResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waf.NetworkAddressListLifecycleStateActive),
	}
}

func (s *WafNetworkAddressListResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waf.NetworkAddressListLifecycleStateDeleting),
	}
}

func (s *WafNetworkAddressListResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waf.NetworkAddressListLifecycleStateDeleted),
	}
}

func (s *WafNetworkAddressListResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_waf.CreateNetworkAddressListRequest{}
	err := s.populateTopLevelPolymorphicCreateNetworkAddressListRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.CreateNetworkAddressList(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNetworkAddressListFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WafNetworkAddressListResourceCrud) getNetworkAddressListFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waf.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkAddressListId, err := networkAddressListWaitForWorkRequest(ctx, workId, "networkAddressList",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*networkAddressListId)

	return s.GetWithContext(ctx)
}

func networkAddressListWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waf", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waf.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func networkAddressListWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_waf.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waf.WafClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waf")
	retryPolicy.ShouldRetryOperation = networkAddressListWorkRequestShouldRetryFunc(timeout)

	response := oci_waf.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_waf.WorkRequestStatusInProgress),
			string(oci_waf.WorkRequestStatusAccepted),
			string(oci_waf.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waf.WorkRequestStatusSucceeded),
			string(oci_waf.WorkRequestStatusFailed),
			string(oci_waf.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_waf.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_waf.WorkRequestStatusFailed || response.Status == oci_waf.WorkRequestStatusCanceled {
		return nil, getErrorFromWafNetworkAddressListWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromWafNetworkAddressListWorkRequest(ctx context.Context, client *oci_waf.WafClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_waf.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_waf.ListWorkRequestErrorsRequest{
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

func (s *WafNetworkAddressListResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_waf.GetNetworkAddressListRequest{}

	tmp := s.D.Id()
	request.NetworkAddressListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.GetNetworkAddressList(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkAddressList
	return nil
}

func (s *WafNetworkAddressListResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waf.UpdateNetworkAddressListRequest{}
	err := s.populateTopLevelPolymorphicUpdateNetworkAddressListRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.UpdateNetworkAddressList(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkAddressListFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WafNetworkAddressListResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_waf.DeleteNetworkAddressListRequest{}

	tmp := s.D.Id()
	request.NetworkAddressListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.DeleteNetworkAddressList(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkAddressListWaitForWorkRequest(ctx, workId, "networkAddressList",
		oci_waf.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *WafNetworkAddressListResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_waf.NetworkAddressListAddresses:
		s.D.Set("type", "ADDRESSES")

		s.D.Set("addresses", v.Addresses)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_waf.NetworkAddressListVcnAddresses:
		s.D.Set("type", "VCN_ADDRESSES")

		vcnAddresses := []interface{}{}
		for _, item := range v.VcnAddresses {
			vcnAddresses = append(vcnAddresses, PrivateAddressesToMap(item))
		}
		s.D.Set("vcn_addresses", vcnAddresses)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func NetworkAddressListSummaryToMap(obj oci_waf.NetworkAddressListSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_waf.NetworkAddressListAddressesSummary:
		result["type"] = "ADDRESSES"

		result["addresses"] = v.Addresses
	case oci_waf.NetworkAddressListVcnAddressesSummary:
		result["type"] = "VCN_ADDRESSES"

		vcnAddresses := []interface{}{}
		for _, item := range v.VcnAddresses {
			vcnAddresses = append(vcnAddresses, PrivateAddressesToMap(item))
		}
		result["vcn_addresses"] = vcnAddresses
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil

	}
	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = string(*obj.GetCompartmentId())
	}
	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = string(*obj.GetDisplayName())
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetId() != nil {
		result["id"] = string(*obj.GetId())
	}

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = string(*obj.GetLifecycleDetails())
	}

	result["state"] = string(obj.GetLifecycleState())

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}
	return result
}

func (s *WafNetworkAddressListResourceCrud) mapToPrivateAddresses(fieldKeyFormat string) (oci_waf.PrivateAddresses, error) {
	result := oci_waf.PrivateAddresses{}

	if addresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "addresses")); ok {
		tmp := addresses.(string)
		result.Addresses = &tmp
	}

	if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcnId.(string)
		result.VcnId = &tmp
	}

	return result, nil
}

func PrivateAddressesToMap(obj oci_waf.PrivateAddresses) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Addresses != nil {
		result["addresses"] = string(*obj.Addresses)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *WafNetworkAddressListResourceCrud) populateTopLevelPolymorphicCreateNetworkAddressListRequest(request *oci_waf.CreateNetworkAddressListRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ADDRESSES"):
		details := oci_waf.CreateNetworkAddressListAddressesDetails{}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateNetworkAddressListDetails = details
	case strings.ToLower("VCN_ADDRESSES"):
		details := oci_waf.CreateNetworkAddressListVcnAddressesDetails{}
		if vcnAddresses, ok := s.D.GetOkExists("vcn_addresses"); ok {
			interfaces := vcnAddresses.([]interface{})
			tmp := make([]oci_waf.PrivateAddresses, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vcn_addresses", stateDataIndex)
				converted, err := s.mapToPrivateAddresses(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("vcn_addresses") {
				details.VcnAddresses = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateNetworkAddressListDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *WafNetworkAddressListResourceCrud) populateTopLevelPolymorphicUpdateNetworkAddressListRequest(request *oci_waf.UpdateNetworkAddressListRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ADDRESSES"):
		details := oci_waf.UpdateNetworkAddressListAddressesDetails{}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.NetworkAddressListId = &tmp
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.UpdateNetworkAddressListDetails = details
	case strings.ToLower("VCN_ADDRESSES"):
		details := oci_waf.UpdateNetworkAddressListVcnAddressesDetails{}
		if vcnAddresses, ok := s.D.GetOkExists("vcn_addresses"); ok {
			interfaces := vcnAddresses.([]interface{})
			tmp := make([]oci_waf.PrivateAddresses, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vcn_addresses", stateDataIndex)
				converted, err := s.mapToPrivateAddresses(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("vcn_addresses") {
				details.VcnAddresses = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.NetworkAddressListId = &tmp
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.UpdateNetworkAddressListDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *WafNetworkAddressListResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_waf.ChangeNetworkAddressListCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkAddressListId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.ChangeNetworkAddressListCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkAddressListFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
