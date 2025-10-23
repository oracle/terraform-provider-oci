// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerPrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createResourcemanagerPrivateEndpoint,
		Read:     readResourcemanagerPrivateEndpoint,
		Update:   updateResourcemanagerPrivateEndpoint,
		Delete:   deleteResourcemanagerPrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_zones": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_used_with_configuration_source_provider": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nsg_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"source_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		},
	}
}

func createResourcemanagerPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.CreateResource(d, sync)
}

func readResourcemanagerPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

func updateResourcemanagerPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteResourcemanagerPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourcemanagerPrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resourcemanager.ResourceManagerClient
	Res                    *oci_resourcemanager.PrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resourcemanager.PrivateEndpointLifecycleStateCreating),
	}
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resourcemanager.PrivateEndpointLifecycleStateActive),
	}
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resourcemanager.PrivateEndpointLifecycleStateDeleting),
	}
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resourcemanager.PrivateEndpointLifecycleStateDeleted),
	}
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) Create() error {
	request := oci_resourcemanager.CreatePrivateEndpointRequest{}

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

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isUsedWithConfigurationSourceProvider, ok := s.D.GetOkExists("is_used_with_configuration_source_provider"); ok {
		tmp := isUsedWithConfigurationSourceProvider.(bool)
		request.IsUsedWithConfigurationSourceProvider = &tmp
	}

	if nsgIdList, ok := s.D.GetOkExists("nsg_id_list"); ok {
		interfaces := nsgIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_id_list") {
			request.NsgIdList = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.CreatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	return s.getPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager"), oci_resourcemanager.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) getPrivateEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resourcemanager.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privateEndpointId, err := privateEndpointWaitForWorkRequest(workId, "ormprivateendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*privateEndpointId)

	return s.Get()
}

func privateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "resourcemanager", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_resourcemanager.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func privateEndpointWaitForWorkRequest(wId *string, entityType string, action oci_resourcemanager.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_resourcemanager.ResourceManagerClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resourcemanager")
	retryPolicy.ShouldRetryOperation = privateEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_resourcemanager.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_resourcemanager.WorkRequestStatusInProgress),
			string(oci_resourcemanager.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_resourcemanager.WorkRequestStatusSucceeded),
			string(oci_resourcemanager.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_resourcemanager.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_resourcemanager.WorkRequestStatusFailed {
		return nil, getErrorFromResourcemanagerPrivateEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourcemanagerPrivateEndpointWorkRequest(client *oci_resourcemanager.ResourceManagerClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resourcemanager.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_resourcemanager.ListWorkRequestErrorsRequest{
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

func (s *ResourcemanagerPrivateEndpointResourceCrud) Get() error {
	request := oci_resourcemanager.GetPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateEndpoint
	return nil
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_resourcemanager.UpdatePrivateEndpointRequest{}

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

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isUsedWithConfigurationSourceProvider, ok := s.D.GetOkExists("is_used_with_configuration_source_provider"); ok {
		tmp := isUsedWithConfigurationSourceProvider.(bool)
		request.IsUsedWithConfigurationSourceProvider = &tmp
	}

	if nsgIdList, ok := s.D.GetOkExists("nsg_id_list"); ok {
		interfaces := nsgIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_id_list") {
			request.NsgIdList = tmp
		}
	}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.UpdatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager"), oci_resourcemanager.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) Delete() error {
	request := oci_resourcemanager.DeletePrivateEndpointRequest{}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.DeletePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := privateEndpointWaitForWorkRequest(workId, "ormprivateendpoint",
		oci_resourcemanager.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) SetData() error {
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

	s.D.Set("dns_zones", s.Res.DnsZones)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsUsedWithConfigurationSourceProvider != nil {
		s.D.Set("is_used_with_configuration_source_provider", *s.Res.IsUsedWithConfigurationSourceProvider)
	}

	s.D.Set("nsg_id_list", s.Res.NsgIdList)

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	s.D.Set("source_ips", s.Res.SourceIps)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func PrivateEndpointSummaryToMap(obj oci_resourcemanager.PrivateEndpointSummary) map[string]interface{} {
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

	result["dns_zones"] = obj.DnsZones

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsUsedWithConfigurationSourceProvider != nil {
		result["is_used_with_configuration_source_provider"] = bool(*obj.IsUsedWithConfigurationSourceProvider)
	}

	result["security_attributes"] = tfresource.SecurityAttributesToMap(obj.SecurityAttributes)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *ResourcemanagerPrivateEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_resourcemanager.ChangePrivateEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PrivateEndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	_, err := s.Client.ChangePrivateEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
