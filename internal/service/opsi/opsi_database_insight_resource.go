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

func OpsiDatabaseInsightResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiDatabaseInsight,
		Read:     readOpsiDatabaseInsight,
		Update:   updateOpsiDatabaseInsight,
		Delete:   deleteOpsiDatabaseInsight,
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
					"EM_MANAGED_EXTERNAL_DATABASE",
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
				Computed: true,
				Optional: true,
			},
			"exadata_insight_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"database_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"lifecycle_details": {
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

func createOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiDatabaseInsightResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.DatabaseInsight
	DisableNotFoundRetries bool
}

func (s *OpsiDatabaseInsightResourceCrud) ID() string {
	databaseInsight := *s.Res
	return *databaseInsight.GetId()
}

func (s *OpsiDatabaseInsightResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) Create() error {
	request := oci_opsi.CreateDatabaseInsightRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	if status, ok := s.D.GetOkExists("status"); ok {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			request := oci_opsi.DisableDatabaseInsightRequest{}
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
			tmp := s.D.Id()
			request.DatabaseInsightId = &tmp
			response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId

			// Wait until it finishes
			databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
				oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
			if err != nil {
				return err
			}
			s.D.SetId(*databaseInsightId)
		}
	}

	return s.Get()
}

func (s *OpsiDatabaseInsightResourceCrud) getDatabaseInsightFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	return s.Get()
}

func databaseInsightWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func databaseInsightWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = databaseInsightWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOpsiDatabaseInsightWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiDatabaseInsightWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
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

func (s *OpsiDatabaseInsightResourceCrud) Get() error {
	request := oci_opsi.GetDatabaseInsightRequest{}

	tmp := s.D.Id()
	request.DatabaseInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseInsight
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateDatabaseInsightRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	disableDatabaseInsight, enableDatabaseInsight := false, false

	if status, ok := s.D.GetOkExists("status"); ok && s.D.HasChange("status") {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			disableDatabaseInsight = true
		} else if oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(wantedState) {
			enableDatabaseInsight = true
		}
	}

	if disableDatabaseInsight {
		request := oci_opsi.DisableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*databaseInsightId)
	}

	if enableDatabaseInsight {
		request := oci_opsi.EnableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		err := s.populateTopLevelPolymorphicEnableDatabaseInsightRequest(&request)
		if err != nil {
			return err
		}

		response, err := s.Client.EnableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*databaseInsightId)
	}

	return s.Get()
}

func (s *OpsiDatabaseInsightResourceCrud) Delete() error {
	status, ok := s.D.GetOkExists("status")
	if ok && oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(strings.ToUpper(status.(string))) {
		request := oci_opsi.DisableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		_, disableWorkRequestErr := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if disableWorkRequestErr != nil {
			return disableWorkRequestErr
		}
	}

	request := oci_opsi.DeleteDatabaseInsightRequest{}

	tmp := s.D.Id()
	request.DatabaseInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiDatabaseInsightResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_opsi.EmManagedExternalDatabaseInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_DATABASE")

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

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
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

func (s *OpsiDatabaseInsightResourceCrud) mapToConnectionDetails(fieldKeyFormat string) (oci_opsi.ConnectionDetails, error) {
	result := oci_opsi.ConnectionDetails{}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_opsi.ConnectionDetailsProtocolEnum(protocol.(string))
	}

	if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
		tmp := serviceName.(string)
		result.ServiceName = &tmp
	}

	return result, nil
}

func ConnectionDetailsToMap(obj *oci_opsi.ConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) mapToCredentialDetails(fieldKeyFormat string) (oci_opsi.CredentialDetails, error) {
	var baseObject oci_opsi.CredentialDetails
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("CREDENTIALS_BY_SOURCE"):
		details := oci_opsi.CredentialsBySource{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func CredentialDetailsToMap(obj *oci_opsi.CredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.CredentialsBySource:
		result["credential_type"] = "CREDENTIALS_BY_SOURCE"
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %T", v)
		return nil
	}

	return result
}

func DatabaseInsightSummaryToMap(obj oci_opsi.DatabaseInsightSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.EmManagedExternalDatabaseInsightSummary:
		result["entity_source"] = "EM_MANAGED_EXTERNAL_DATABASE"
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.ExadataInsightId != nil {
			result["exadata_insight_id"] = string(*v.ExadataInsightId)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		}

		if v.DatabaseHostNames != nil {
			result["database_host_names"] = v.DatabaseHostNames
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

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
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicCreateDatabaseInsightRequest(request *oci_opsi.CreateDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.CreateEmManagedExternalDatabaseInsightDetails{}
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
		if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
			tmp := enterpriseManagerBridgeId.(string)
			details.EnterpriseManagerBridgeId = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicUpdateDatabaseInsightRequest(request *oci_opsi.UpdateDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.UpdateEmManagedExternalDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
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
		request.UpdateDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicEnableDatabaseInsightRequest(request *oci_opsi.EnableDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.EnableEmManagedExternalDatabaseInsightDetails{}
		request.EnableDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeDatabaseInsightCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseInsightId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeDatabaseInsightCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
