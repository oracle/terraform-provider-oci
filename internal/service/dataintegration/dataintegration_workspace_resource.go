// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"
)

func DataintegrationWorkspaceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		CreateContext: createDataintegrationWorkspaceWithContext,
		ReadContext:   readDataintegrationWorkspaceWithContext,
		UpdateContext: updateDataintegrationWorkspaceWithContext,
		DeleteContext: deleteDataintegrationWorkspaceWithContext,
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
			"dns_server_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dns_server_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"endpoint_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"endpoint_name": {
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
			"is_force_operation": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_private_network_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"registry_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"registry_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"registry_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"quiesce_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_message": {
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

func createDataintegrationWorkspaceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataintegrationWorkspaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataintegrationWorkspaceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataintegrationWorkspaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataintegrationWorkspaceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataintegrationWorkspaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataintegrationWorkspaceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataintegrationWorkspaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataintegrationWorkspaceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Workspace
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataintegrationWorkspaceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dataintegration.WorkspaceLifecycleStateCreating),
		string(oci_dataintegration.WorkspaceLifecycleStateStarting),
	}
}

func (s *DataintegrationWorkspaceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataintegration.WorkspaceLifecycleStateActive),
	}
}

func (s *DataintegrationWorkspaceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataintegration.WorkspaceLifecycleStateDeleting),
	}
}

func (s *DataintegrationWorkspaceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataintegration.WorkspaceLifecycleStateDeleted),
	}
}

func (s *DataintegrationWorkspaceResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_dataintegration.CreateWorkspaceRequest{}

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

	if dnsServerIp, ok := s.D.GetOkExists("dns_server_ip"); ok {
		tmp := dnsServerIp.(string)
		request.DnsServerIp = &tmp
	}

	if dnsServerZone, ok := s.D.GetOkExists("dns_server_zone"); ok {
		tmp := dnsServerZone.(string)
		request.DnsServerZone = &tmp
	}

	if endpointCompartmentId, ok := s.D.GetOkExists("endpoint_compartment_id"); ok {
		tmp := endpointCompartmentId.(string)
		request.EndpointCompartmentId = &tmp
	}

	if endpointId, ok := s.D.GetOkExists("endpoint_id"); ok {
		tmp := endpointId.(string)
		request.EndpointId = &tmp
	}

	if endpointName, ok := s.D.GetOkExists("endpoint_name"); ok {
		tmp := endpointName.(string)
		request.EndpointName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isPrivateNetworkEnabled, ok := s.D.GetOkExists("is_private_network_enabled"); ok {
		tmp := isPrivateNetworkEnabled.(bool)
		request.IsPrivateNetworkEnabled = &tmp
	}

	if registryCompartmentId, ok := s.D.GetOkExists("registry_compartment_id"); ok {
		tmp := registryCompartmentId.(string)
		request.RegistryCompartmentId = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if registryName, ok := s.D.GetOkExists("registry_name"); ok {
		tmp := registryName.(string)
		request.RegistryName = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateWorkspace(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_dataintegration.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(ctx,
		oci_dataintegration.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "disworkspace") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getWorkspaceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disworkspace"), oci_dataintegration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataintegrationWorkspaceResourceCrud) getWorkspaceFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dataintegration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	workspaceId, err := workspaceWaitForWorkRequest(ctx, workId, "disworkspace",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, workspaceId)
		return err
	}
	s.D.SetId(*workspaceId)

	return s.GetWithContext(ctx)
}

func workspaceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dataintegration", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dataintegration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func workspaceWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_dataintegration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dataintegration.DataIntegrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dataintegration")
	retryPolicy.ShouldRetryOperation = workspaceWorkRequestShouldRetryFunc(timeout)

	response := oci_dataintegration.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_dataintegration.WorkRequestStatusInProgress),
			string(oci_dataintegration.WorkRequestStatusAccepted),
			string(oci_dataintegration.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_dataintegration.WorkRequestStatusSucceeded),
			string(oci_dataintegration.WorkRequestStatusFailed),
			string(oci_dataintegration.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_dataintegration.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_dataintegration.WorkRequestStatusFailed || response.Status == oci_dataintegration.WorkRequestStatusCanceled {
		return nil, getErrorFromDataintegrationWorkspaceWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataintegrationWorkspaceWorkRequest(ctx context.Context, client *oci_dataintegration.DataIntegrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dataintegration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_dataintegration.ListWorkRequestErrorsRequest{
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

func (s *DataintegrationWorkspaceResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dataintegration.GetWorkspaceRequest{}

	tmp := s.D.Id()
	request.WorkspaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetWorkspace(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Workspace
	return nil
}

func (s *DataintegrationWorkspaceResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dataintegration.UpdateWorkspaceRequest{}

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
	request.WorkspaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.UpdateWorkspace(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWorkspaceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disworkspace"), oci_dataintegration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataintegrationWorkspaceResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_dataintegration.DeleteWorkspaceRequest{}

	if isForceOperation, ok := s.D.GetOkExists("is_force_operation"); ok {
		tmp := isForceOperation.(bool)
		request.IsForceOperation = &tmp
	}

	if quiesceTimeout, ok := s.D.GetOkExists("quiesce_timeout"); ok {
		tmp := quiesceTimeout.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert quiesceTimeout string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.QuiesceTimeout = &tmpInt64
	}

	tmp := s.D.Id()
	request.WorkspaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.DeleteWorkspace(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := workspaceWaitForWorkRequest(ctx, workId, "disworkspace",
		oci_dataintegration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataintegrationWorkspaceResourceCrud) SetData() error {
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

	if s.Res.DnsServerIp != nil {
		s.D.Set("dns_server_ip", *s.Res.DnsServerIp)
	}

	if s.Res.DnsServerZone != nil {
		s.D.Set("dns_server_zone", *s.Res.DnsServerZone)
	}

	if s.Res.EndpointId != nil {
		s.D.Set("endpoint_id", *s.Res.EndpointId)
	}

	if s.Res.EndpointName != nil {
		s.D.Set("endpoint_name", *s.Res.EndpointName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrivateNetworkEnabled != nil {
		s.D.Set("is_private_network_enabled", *s.Res.IsPrivateNetworkEnabled)
	}

	if s.Res.RegistryId != nil {
		s.D.Set("registry_id", *s.Res.RegistryId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *DataintegrationWorkspaceResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_dataintegration.ChangeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.WorkspaceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.ChangeCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWorkspaceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disworkspace"), oci_dataintegration.WorkRequestResourceActionTypeMoved, s.D.Timeout(schema.TimeoutUpdate))
}
