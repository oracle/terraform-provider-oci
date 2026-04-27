// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsSqlReportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsDatabaseToolsSqlReportWithContext,
		ReadContext:   readDatabaseToolsDatabaseToolsSqlReportWithContext,
		UpdateContext: updateDatabaseToolsDatabaseToolsSqlReportWithContext,
		DeleteContext: deleteDatabaseToolsDatabaseToolsSqlReportWithContext,
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
			"source": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"columns": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
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
			"instructions": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"purpose": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"variables": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

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

func createDatabaseToolsDatabaseToolsSqlReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsDatabaseToolsSqlReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsDatabaseToolsSqlReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsDatabaseToolsSqlReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsDatabaseToolsSqlReportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsSqlReport
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) ID() string {
	databaseToolsSqlReport := *s.Res
	return *databaseToolsSqlReport.GetId()
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsSqlReportLifecycleStateActive),
	}
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsSqlReportLifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools.CreateDatabaseToolsSqlReportRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsSqlReportRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsSqlReport(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	if workId == nil {
		return s.GetWithContext(ctx)
	}

	return s.getDatabaseToolsSqlReportFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) getDatabaseToolsSqlReportFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsSqlReportId, err := databaseToolsSqlReportWaitForWorkRequest(ctx, workId, "databasetoolssqlreport",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseToolsSqlReportId)

	return s.GetWithContext(ctx)
}

func databaseToolsSqlReportWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func databaseToolsSqlReportWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_database_tools.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_tools.DatabaseToolsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_tools")
	retryPolicy.ShouldRetryOperation = databaseToolsSqlReportWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromDatabaseToolsDatabaseToolsSqlReportWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseToolsDatabaseToolsSqlReportWorkRequest(ctx context.Context, client *oci_database_tools.DatabaseToolsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_tools.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsSqlReportRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsSqlReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsSqlReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsSqlReport
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsSqlReportRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsSqlReportRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsSqlReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsSqlReport
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools.DeleteDatabaseToolsSqlReportRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsSqlReportId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	_, err := s.Client.DeleteDatabaseToolsSqlReport(ctx, request)
	return err
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsSqlReportOracleDatabase:
		s.D.Set("type", "ORACLE_DATABASE")

		columns := []interface{}{}
		for _, item := range v.Columns {
			columns = append(columns, DatabaseToolsSqlReportColumnToMap(item))
		}
		s.D.Set("columns", columns)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Instructions != nil {
			s.D.Set("instructions", *v.Instructions)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.Format(time.RFC3339Nano))
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.Format(time.RFC3339Nano))
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsSqlReportResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.Purpose != nil {
			s.D.Set("purpose", *v.Purpose)
		}

		if v.Source != nil {
			s.D.Set("source", *v.Source)
		}

		s.D.Set("state", v.LifecycleState)

		variables := []interface{}{}
		for _, item := range v.Variables {
			variables = append(variables, DatabaseToolsSqlReportVariableToMap(item))
		}
		s.D.Set("variables", variables)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) mapToDatabaseToolsSqlReportColumn(fieldKeyFormat string) (oci_database_tools.DatabaseToolsSqlReportColumn, error) {
	result := oci_database_tools.DatabaseToolsSqlReportColumn{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func DatabaseToolsSqlReportColumnToMap(obj oci_database_tools.DatabaseToolsSqlReportColumn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func DatabaseToolsSqlReportSummaryToMap(obj oci_database_tools.DatabaseToolsSqlReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetId() != nil {
		result["id"] = *obj.GetId()
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = *obj.GetCompartmentId()
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = *obj.GetDisplayName()
	}

	result["state"] = string(obj.GetLifecycleState())

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	switch (obj).(type) {
	case oci_database_tools.DatabaseToolsSqlReportSummaryOracleDatabase:
		result["type"] = "ORACLE_DATABASE"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) mapToDatabaseToolsSqlReportVariable(fieldKeyFormat string) (oci_database_tools.DatabaseToolsSqlReportVariable, error) {
	result := oci_database_tools.DatabaseToolsSqlReportVariable{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func DatabaseToolsSqlReportVariableToMap(obj oci_database_tools.DatabaseToolsSqlReportVariable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
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

func DbtoolsSqlReportResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
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

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsSqlReportRequest(request *oci_database_tools.CreateDatabaseToolsSqlReportRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.CreateDatabaseToolsSqlReportOracleDatabaseDetails{}
		if columns, ok := s.D.GetOkExists("columns"); ok {
			interfaces := columns.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsSqlReportColumn, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "columns", stateDataIndex)
				converted, err := s.mapToDatabaseToolsSqlReportColumn(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("columns") {
				details.Columns = tmp
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if instructions, ok := s.D.GetOkExists("instructions"); ok {
			tmp := instructions.(string)
			details.Instructions = &tmp
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
		if purpose, ok := s.D.GetOkExists("purpose"); ok {
			tmp := purpose.(string)
			details.Purpose = &tmp
		}
		if source, ok := s.D.GetOkExists("source"); ok {
			tmp := source.(string)
			details.Source = &tmp
		}
		if variables, ok := s.D.GetOkExists("variables"); ok {
			interfaces := variables.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsSqlReportVariable, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "variables", stateDataIndex)
				converted, err := s.mapToDatabaseToolsSqlReportVariable(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("variables") {
				details.Variables = tmp
			}
		}
		request.CreateDatabaseToolsSqlReportDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsSqlReportRequest(request *oci_database_tools.UpdateDatabaseToolsSqlReportRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.UpdateDatabaseToolsSqlReportDetailsOracleDatabase{}
		if columns, ok := s.D.GetOkExists("columns"); ok {
			interfaces := columns.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsSqlReportColumn, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "columns", stateDataIndex)
				converted, err := s.mapToDatabaseToolsSqlReportColumn(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("columns") {
				details.Columns = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsSqlReportId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if instructions, ok := s.D.GetOkExists("instructions"); ok {
			tmp := instructions.(string)
			details.Instructions = &tmp
		}
		if purpose, ok := s.D.GetOkExists("purpose"); ok {
			tmp := purpose.(string)
			details.Purpose = &tmp
		}
		if source, ok := s.D.GetOkExists("source"); ok {
			tmp := source.(string)
			details.Source = &tmp
		}
		if variables, ok := s.D.GetOkExists("variables"); ok {
			interfaces := variables.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsSqlReportVariable, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "variables", stateDataIndex)
				converted, err := s.mapToDatabaseToolsSqlReportVariable(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("variables") {
				details.Variables = tmp
			}
		}
		request.UpdateDatabaseToolsSqlReportDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsSqlReportCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsSqlReportId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	_, err := s.Client.ChangeDatabaseToolsSqlReportCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
