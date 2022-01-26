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

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
)

func OpsiExadataInsightResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiExadataInsight,
		Read:     readOpsiExadataInsight,
		Update:   updateOpsiExadataInsight,
		Delete:   deleteOpsiExadataInsight,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enterprise_manager_bridge_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enterprise_manager_entity_identifier": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enterprise_manager_identifier": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entity_source": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"EM_MANAGED_EXTERNAL_EXADATA",
				}, true),
			},

			//Optional
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
				Computed: true,
				Optional: true,
			},
			"is_auto_sync_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"exadata_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_rack_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_virtualized_exadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
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

func createOpsiExadataInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiExadataInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiExadataInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiExadataInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiExadataInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiExadataInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiExadataInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiExadataInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiExadataInsightResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.ExadataInsight
	DisableNotFoundRetries bool
}

func (s *OpsiExadataInsightResourceCrud) ID() string {
	exadataInsight := *s.Res
	return *exadataInsight.GetId()
}

func (s *OpsiExadataInsightResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiExadataInsightResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
	}
}

func (s *OpsiExadataInsightResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiExadataInsightResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiExadataInsightResourceCrud) Create() error {
	request := oci_opsi.CreateExadataInsightRequest{}
	err := s.populateTopLevelPolymorphicCreateExadataInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateExadataInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*exadataInsightId)

	if status, ok := s.D.GetOkExists("status"); ok {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			request := oci_opsi.DisableExadataInsightRequest{}
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
			tmp := s.D.Id()
			request.ExadataInsightId = &tmp
			response, err := s.Client.DisableExadataInsight(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId

			// Wait until it finishes
			exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
				oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
			if err != nil {
				return err
			}
			s.D.SetId(*exadataInsightId)
		}
	}

	return s.Get()
}

func (s *OpsiExadataInsightResourceCrud) getExadataInsightFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*exadataInsightId)

	return s.Get()
}

func exadataInsightWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func exadataInsightWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = exadataInsightWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOpsiExadataInsightWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiExadataInsightWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
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

func (s *OpsiExadataInsightResourceCrud) Get() error {
	request := oci_opsi.GetExadataInsightRequest{}

	tmp := s.D.Id()
	request.ExadataInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetExadataInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInsight
	return nil
}

func (s *OpsiExadataInsightResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateExadataInsightRequest{}
	err := s.populateTopLevelPolymorphicUpdateExadataInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateExadataInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*exadataInsightId)

	disableExadataInsight, enableExadataInsight := false, false

	if status, ok := s.D.GetOkExists("status"); ok && s.D.HasChange("status") {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			disableExadataInsight = true
		} else if oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(wantedState) {
			enableExadataInsight = true
		}
	}

	if disableExadataInsight {
		request := oci_opsi.DisableExadataInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.ExadataInsightId = &tmp
		response, err := s.Client.DisableExadataInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*exadataInsightId)
	}

	if enableExadataInsight {
		request := oci_opsi.EnableExadataInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.ExadataInsightId = &tmp
		err := s.populateTopLevelPolymorphicEnableExadataInsightRequest(&request)
		if err != nil {
			return err
		}

		response, err := s.Client.EnableExadataInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		exadataInsightId, err := exadataInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*exadataInsightId)
	}

	return s.Get()
}

func (s *OpsiExadataInsightResourceCrud) Delete() error {
	status, ok := s.D.GetOkExists("status")
	if ok && oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(strings.ToUpper(status.(string))) {
		request := oci_opsi.DisableExadataInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.ExadataInsightId = &tmp
		response, err := s.Client.DisableExadataInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		_, disableWorkRequestErr := exadataInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if disableWorkRequestErr != nil {
			return disableWorkRequestErr
		}
	}

	request := oci_opsi.DeleteExadataInsightRequest{}

	tmp := s.D.Id()
	request.ExadataInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteExadataInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := exadataInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiExadataInsightResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_opsi.EmManagedExternalExadataInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_EXADATA")

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

		if v.IsAutoSyncEnabled != nil {
			s.D.Set("is_auto_sync_enabled", *v.IsAutoSyncEnabled)
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

		if v.ExadataDisplayName != nil {
			s.D.Set("exadata_display_name", *v.ExadataDisplayName)
		}

		if v.ExadataName != nil {
			s.D.Set("exadata_name", *v.ExadataName)
		}

		s.D.Set("exadata_rack_type", v.ExadataRackType)

		s.D.Set("exadata_type", v.ExadataType)

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsVirtualizedExadata != nil {
			s.D.Set("is_virtualized_exadata", *v.IsVirtualizedExadata)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
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

func (s *OpsiExadataInsightResourceCrud) mapToCreateEmManagedExternalExadataMemberEntityDetails(fieldKeyFormat string) (oci_opsi.CreateEmManagedExternalExadataMemberEntityDetails, error) {
	result := oci_opsi.CreateEmManagedExternalExadataMemberEntityDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if enterpriseManagerEntityIdentifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enterprise_manager_entity_identifier")); ok {
		tmp := enterpriseManagerEntityIdentifier.(string)
		result.EnterpriseManagerEntityIdentifier = &tmp
	}

	return result, nil
}

func CreateEmManagedExternalExadataMemberEntityDetailsToMap(obj oci_opsi.CreateEmManagedExternalExadataMemberEntityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.EnterpriseManagerEntityIdentifier != nil {
		result["enterprise_manager_entity_identifier"] = string(*obj.EnterpriseManagerEntityIdentifier)
	}

	return result
}

func ExadataInsightSummaryToMap(obj oci_opsi.ExadataInsightSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.EmManagedExternalExadataInsightSummary:
		result["entity_source"] = "EM_MANAGED_EXTERNAL_EXADATA"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
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

		if v.ExadataDisplayName != nil {
			result["exadata_display_name"] = string(*v.ExadataDisplayName)
		}

		if v.ExadataName != nil {
			result["exadata_name"] = string(*v.ExadataName)
		}

		result["exadata_rack_type"] = string(v.ExadataRackType)

		result["exadata_type"] = string(v.ExadataType)

		result["freeform_tags"] = v.FreeformTags

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeCreated.String()
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OpsiExadataInsightResourceCrud) populateTopLevelPolymorphicCreateExadataInsightRequest(request *oci_opsi.CreateExadataInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_EXADATA"):
		details := oci_opsi.CreateEmManagedExternalExadataInsightDetails{}
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
		if isAutoSyncEnabled, ok := s.D.GetOkExists("is_auto_sync_enabled"); ok {
			tmp := isAutoSyncEnabled.(bool)
			details.IsAutoSyncEnabled = &tmp
		}
		if memberEntityDetails, ok := s.D.GetOkExists("member_entity_details"); ok {
			interfaces := memberEntityDetails.([]interface{})
			tmp := make([]oci_opsi.CreateEmManagedExternalExadataMemberEntityDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "member_entity_details", stateDataIndex)
				converted, err := s.mapToCreateEmManagedExternalExadataMemberEntityDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("member_entity_details") {
				details.MemberEntityDetails = tmp
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
		if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
			tmp := enterpriseManagerBridgeId.(string)
			details.EnterpriseManagerBridgeId = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateExadataInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiExadataInsightResourceCrud) populateTopLevelPolymorphicUpdateExadataInsightRequest(request *oci_opsi.UpdateExadataInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_EXADATA"):
		details := oci_opsi.UpdateEmManagedExternalExadataInsightDetails{}
		if isAutoSyncEnabled, ok := s.D.GetOkExists("is_auto_sync_enabled"); ok {
			tmp := isAutoSyncEnabled.(bool)
			details.IsAutoSyncEnabled = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.ExadataInsightId = &tmp
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateExadataInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiExadataInsightResourceCrud) populateTopLevelPolymorphicEnableExadataInsightRequest(request *oci_opsi.EnableExadataInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_EXADATA"):
		details := oci_opsi.EnableEmManagedExternalExadataInsightDetails{}
		request.EnableExadataInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiExadataInsightResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeExadataInsightCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExadataInsightId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeExadataInsightCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExadataInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
