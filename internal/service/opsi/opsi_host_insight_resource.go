// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

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
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
)

func OpsiHostInsightResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiHostInsight,
		Read:     readOpsiHostInsight,
		Update:   updateOpsiHostInsight,
		Delete:   deleteOpsiHostInsight,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_source": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"EM_MANAGED_EXTERNAL_HOST",
					"MACS_MANAGED_EXTERNAL_HOST",
				}, true),
			},

			// Optional
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_bridge_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_entity_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exadata_insight_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"enterprise_manager_entity_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_manager_entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_manager_entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"processor_count": {
				Type:     schema.TypeInt,
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

func createOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiHostInsightResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.HostInsight
	DisableNotFoundRetries bool
}

func (s *OpsiHostInsightResourceCrud) ID() string {
	hostInsight := *s.Res
	return *hostInsight.GetId()
}

func (s *OpsiHostInsightResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiHostInsightResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
	}
}

func (s *OpsiHostInsightResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiHostInsightResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiHostInsightResourceCrud) Create() error {
	request := oci_opsi.CreateHostInsightRequest{}
	err := s.populateTopLevelPolymorphicCreateHostInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostInsightId)

	if status, ok := s.D.GetOkExists("status"); ok {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			request := oci_opsi.DisableHostInsightRequest{}
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
			tmp := s.D.Id()
			request.HostInsightId = &tmp
			response, err := s.Client.DisableHostInsight(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId

			// Wait until it finishes
			hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
				oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
			if err != nil {
				return err
			}
			s.D.SetId(*hostInsightId)
		}
	}

	return s.Get()
}

func (s *OpsiHostInsightResourceCrud) getHostInsightFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostInsightId)

	return s.Get()
}

func hostInsightWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func hostInsightWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = hostInsightWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiHostInsightWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiHostInsightWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiHostInsightResourceCrud) Get() error {
	request := oci_opsi.GetHostInsightRequest{}

	tmp := s.D.Id()
	request.HostInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HostInsight
	return nil
}

func (s *OpsiHostInsightResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateHostInsightRequest{}
	err := s.populateTopLevelPolymorphicUpdateHostInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostInsightId)

	disableHostInsight, enableHostInsight := false, false

	if status, ok := s.D.GetOkExists("status"); ok && s.D.HasChange("status") {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			disableHostInsight = true
		} else if oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(wantedState) {
			enableHostInsight = true
		}
	}

	if disableHostInsight {
		request := oci_opsi.DisableHostInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.HostInsightId = &tmp
		response, err := s.Client.DisableHostInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*hostInsightId)
	}

	if enableHostInsight {
		request := oci_opsi.EnableHostInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.HostInsightId = &tmp
		err := s.populateTopLevelPolymorphicEnableHostInsightRequest(&request)
		if err != nil {
			return err
		}

		response, err := s.Client.EnableHostInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		hostInsightId, err := hostInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*hostInsightId)
	}

	return s.Get()
}

func (s *OpsiHostInsightResourceCrud) Delete() error {

	status, ok := s.D.GetOkExists("status")
	if ok && oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(strings.ToUpper(status.(string))) {
		request := oci_opsi.DisableHostInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.HostInsightId = &tmp
		response, err := s.Client.DisableHostInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		_, disableWorkRequestErr := hostInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if disableWorkRequestErr != nil {
			return disableWorkRequestErr
		}

	}

	request := oci_opsi.DeleteHostInsightRequest{}

	tmp := s.D.Id()
	request.HostInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := hostInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiHostInsightResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_opsi.EmManagedExternalHostInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_HOST")

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			s.D.Set("enterprise_manager_entity_display_name", *v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			s.D.Set("enterprise_manager_entity_identifier", *v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			s.D.Set("enterprise_manager_entity_name", *v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			s.D.Set("enterprise_manager_entity_type", *v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			s.D.Set("enterprise_manager_identifier", *v.EnterpriseManagerIdentifier)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
		}

		if v.PlatformName != nil {
			s.D.Set("platform_name", *v.PlatformName)
		}

		s.D.Set("platform_type", v.PlatformType)

		if v.PlatformVersion != nil {
			s.D.Set("platform_version", *v.PlatformVersion)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.HostDisplayName != nil {
			s.D.Set("host_display_name", *v.HostDisplayName)
		}

		if v.HostName != nil {
			s.D.Set("host_name", *v.HostName)
		}

		if v.HostType != nil {
			s.D.Set("host_type", *v.HostType)
		}

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("platform_type", v.PlatformType)

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.MacsManagedExternalHostInsight:
		s.D.Set("entity_source", "MACS_MANAGED_EXTERNAL_HOST")

		if v.ManagementAgentId != nil {
			s.D.Set("management_agent_id", *v.ManagementAgentId)
		}

		if v.PlatformName != nil {
			s.D.Set("platform_name", *v.PlatformName)
		}

		s.D.Set("platform_type", v.PlatformType)

		if v.PlatformVersion != nil {
			s.D.Set("platform_version", *v.PlatformVersion)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.HostDisplayName != nil {
			s.D.Set("host_display_name", *v.HostDisplayName)
		}

		if v.HostName != nil {
			s.D.Set("host_name", *v.HostName)
		}

		if v.HostType != nil {
			s.D.Set("host_type", *v.HostType)
		}

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

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
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func HostInsightSummaryToMap(obj oci_opsi.HostInsightSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.EmManagedExternalHostInsightSummary:
		result["entity_source"] = "EM_MANAGED_EXTERNAL_HOST"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.HostDisplayName != nil {
			result["host_display_name"] = string(*v.HostDisplayName)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.HostType != nil {
			result["host_type"] = string(*v.HostType)
		}

		if v.EnterpriseManagerBridgeId != nil {
			result["enterprise_manager_bridge_id"] = string(*v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			result["enterprise_manager_entity_display_name"] = string(*v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			result["enterprise_manager_entity_identifier"] = string(*v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			result["enterprise_manager_entity_name"] = string(*v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			result["enterprise_manager_entity_type"] = string(*v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			result["enterprise_manager_identifier"] = string(*v.EnterpriseManagerIdentifier)
		}

		if v.ExadataInsightId != nil {
			result["exadata_insight_id"] = string(*v.ExadataInsightId)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["platform_type"] = string(v.PlatformType)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.ProcessorCount != nil {
			result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

	case oci_opsi.MacsManagedExternalHostInsightSummary:
		result["entity_source"] = "MACS_MANAGED_EXTERNAL_HOST"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.ManagementAgentId != nil {
			result["management_agent_id"] = string(*v.ManagementAgentId)
		}

		if v.HostDisplayName != nil {
			result["host_display_name"] = string(*v.HostDisplayName)
		}

		if v.HostType != nil {
			result["host_type"] = string(*v.HostType)
		}

		if v.ProcessorCount != nil {
			result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["platform_type"] = string(v.PlatformType)
		result["state"] = string(v.LifecycleState)
		result["status"] = string(v.Status)

	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OpsiHostInsightResourceCrud) populateTopLevelPolymorphicCreateHostInsightRequest(request *oci_opsi.CreateHostInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.CreateEmManagedExternalHostInsightDetails{}
		if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
			tmp := enterpriseManagerBridgeId.(string)
			details.EnterpriseManagerBridgeId = &tmp
		}
		if enterpriseManagerEntityIdentifier, ok := s.D.GetOkExists("enterprise_manager_entity_identifier"); ok {
			tmp := enterpriseManagerEntityIdentifier.(string)
			details.EnterpriseManagerEntityIdentifier = &tmp
		}
		if enterpriseManagerIdentifier, ok := s.D.GetOkExists("enterprise_manager_identifier"); ok {
			tmp := enterpriseManagerIdentifier.(string)
			details.EnterpriseManagerIdentifier = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateHostInsightDetails = details
	case strings.ToLower("MACS_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.CreateMacsManagedExternalHostInsightDetails{}
		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			details.ManagementAgentId = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateHostInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiHostInsightResourceCrud) populateTopLevelPolymorphicUpdateHostInsightRequest(request *oci_opsi.UpdateHostInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.UpdateEmManagedExternalHostInsightDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.HostInsightId = &tmp
		request.UpdateHostInsightDetails = details
	case strings.ToLower("MACS_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.UpdateMacsManagedExternalHostInsightDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.HostInsightId = &tmp
		request.UpdateHostInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiHostInsightResourceCrud) populateTopLevelPolymorphicEnableHostInsightRequest(request *oci_opsi.EnableHostInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.EnableEmManagedExternalHostInsightDetails{}
		request.EnableHostInsightDetails = details
	case strings.ToLower("MACS_MANAGED_EXTERNAL_HOST"):
		details := oci_opsi.EnableMacsManagedExternalHostInsightDetails{}
		request.EnableHostInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiHostInsightResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeHostInsightCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.HostInsightId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeHostInsightCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHostInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
