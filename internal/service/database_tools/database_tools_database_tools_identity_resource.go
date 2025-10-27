// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

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
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsIdentityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseToolsDatabaseToolsIdentity,
		Read:     readDatabaseToolsDatabaseToolsIdentity,
		Update:   updateDatabaseToolsDatabaseToolsIdentity,
		Delete:   deleteDatabaseToolsDatabaseToolsIdentity,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credential_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ORACLE_DATABASE_RESOURCE_PRINCIPAL",
				}, true),
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
							ForceNew: true,
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

func createDatabaseToolsDatabaseToolsIdentity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsIdentityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseToolsDatabaseToolsIdentity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsIdentityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseToolsDatabaseToolsIdentity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsIdentityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseToolsDatabaseToolsIdentity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsIdentityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseToolsDatabaseToolsIdentityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsIdentity
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) ID() string {
	databaseToolsIdentity := *s.Res
	return *databaseToolsIdentity.GetId()
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateCreating),
	}
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateActive),
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateNeedsAttention),
	}
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateDeleting),
	}
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) Create() error {
	request := oci_database_tools.CreateDatabaseToolsIdentityRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsIdentityRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsIdentity(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDatabaseToolsIdentityFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) getDatabaseToolsIdentityFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsIdentityId, err := databaseToolsIdentityWaitForWorkRequest(workId, "databasetoolsidentity",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseToolsIdentityId)

	return s.Get()
}

func databaseToolsIdentityWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_tools", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_tools.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func databaseToolsIdentityWaitForWorkRequest(wId *string, entityType string, action oci_database_tools.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_tools.DatabaseToolsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_tools")
	retryPolicy.ShouldRetryOperation = databaseToolsIdentityWorkRequestShouldRetryFunc(timeout)

	response := oci_database_tools.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_tools.OperationStatusInProgress),
			string(oci_database_tools.OperationStatusAccepted),
			string(oci_database_tools.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_tools.OperationStatusSucceeded),
			string(oci_database_tools.OperationStatusFailed),
			string(oci_database_tools.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_tools.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_database_tools.OperationStatusFailed || response.Status == oci_database_tools.OperationStatusCanceled {
		return nil, getErrorFromDatabaseToolsDatabaseToolsIdentityWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseToolsDatabaseToolsIdentityWorkRequest(client *oci_database_tools.DatabaseToolsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_tools.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_tools.ListWorkRequestErrorsRequest{
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

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) Get() error {
	request := oci_database_tools.GetDatabaseToolsIdentityRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsIdentityId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsIdentity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsIdentity
	return nil
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsIdentityRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsIdentityRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsIdentity(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsIdentityFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) Delete() error {
	request := oci_database_tools.DeleteDatabaseToolsIdentityRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsIdentityId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.DeleteDatabaseToolsIdentity(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseToolsIdentityWaitForWorkRequest(workId, "databasetoolsidentity",
		oci_database_tools.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsIdentityOracleDatabaseResourcePrincipal:
		s.D.Set("type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL")

		if v.CredentialKey != nil {
			s.D.Set("credential_key", *v.CredentialKey)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

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

func DatabaseToolsIdentitySummaryToMap(obj oci_database_tools.DatabaseToolsIdentitySummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_tools.DatabaseToolsIdentityOracleDatabaseResourcePrincipalSummary:
		result["type"] = "ORACLE_DATABASE_RESOURCE_PRINCIPAL"

		if v.CredentialKey != nil {
			result["credential_key"] = string(*v.CredentialKey)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseToolsConnectionId != nil {
			result["database_tools_connection_id"] = string(*v.DatabaseToolsConnectionId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.FreeformTags != nil {
			result["freeform_tags"] = v.FreeformTags
		}

		result["state"] = string(v.LifecycleState)

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
	result := oci_database_tools.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database_tools.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func ResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsIdentityRequest(request *oci_database_tools.CreateDatabaseToolsIdentityRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE_RESOURCE_PRINCIPAL"):
		details := oci_database_tools.CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails{}
		if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
			tmp := credentialKey.(string)
			details.CredentialKey = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
			tmp := databaseToolsConnectionId.(string)
			details.DatabaseToolsConnectionId = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		request.CreateDatabaseToolsIdentityDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsIdentityRequest(request *oci_database_tools.UpdateDatabaseToolsIdentityRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE_RESOURCE_PRINCIPAL"):
		details := oci_database_tools.UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails{}
		tmp := s.D.Id()
		request.DatabaseToolsIdentityId = &tmp
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
		request.UpdateDatabaseToolsIdentityDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsIdentityResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsIdentityCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsIdentityId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.ChangeDatabaseToolsIdentityCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsIdentityFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
