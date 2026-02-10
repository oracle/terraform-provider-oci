// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

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
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationAssessmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationAssessment,
		Read:     readDatabaseMigrationAssessment,
		Update:   updateDatabaseMigrationAssessment,
		Delete:   deleteDatabaseMigrationAssessment,
		Schema: map[string]*schema.Schema{
			// Required
			"acceptable_downtime": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_combination": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MYSQL",
					"ORACLE",
				}, true),
			},
			"database_data_size": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ddl_expectation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_speed_megabit_per_second": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_database_connection": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"target_database_connection": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"connection_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"database_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"technology_sub_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"technology_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"bulk_include_exclude_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"creation_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exclude_objects": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"is_omit_excluded_table_from_replication": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"include_objects": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"is_omit_excluded_table_from_replication": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"assessment_migration_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cdb_supported": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"migration_id": {
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

func createDatabaseMigrationAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMigrationAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMigrationAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseMigrationAssessmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.Assessment
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationAssessmentResourceCrud) ID() string {
	assessment := *s.Res
	return *assessment.GetId()
}

func (s *DatabaseMigrationAssessmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.AssessmentLifecycleStatesCreating),
		string(oci_database_migration.AssessmentLifecycleStatesInProgress),
	}
}

func (s *DatabaseMigrationAssessmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.AssessmentLifecycleStatesActive),
		string(oci_database_migration.AssessmentLifecycleStatesSucceeded),
		string(oci_database_migration.AssessmentLifecycleStatesNeedsAttention),
	}
}

func (s *DatabaseMigrationAssessmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_migration.AssessmentLifecycleStatesDeleting),
	}
}

func (s *DatabaseMigrationAssessmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.AssessmentLifecycleStatesDeleted),
	}
}

func (s *DatabaseMigrationAssessmentResourceCrud) Create() error {
	request := oci_database_migration.CreateAssessmentRequest{}
	err := s.populateTopLevelPolymorphicCreateAssessmentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.CreateAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseMigrationAssessmentResourceCrud) getAssessmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_migration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait for work request to reach a terminal state and surface any errors.
	if err := assessmentWaitForWorkRequest(workId, "assessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client); err != nil {
		return err
	}

	return s.Get()
}

func assessmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_migration", startTime) {
			return true
		}

		// Only stop if status of work request response is succeeded
		if workRequestResponse, ok := response.Response.(oci_database_migration.GetWorkRequestResponse); ok {
			return workRequestResponse.Status != oci_database_migration.OperationStatusSucceeded
		}
		return false
	}
}

func assessmentWaitForWorkRequest(wId *string, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_migration.DatabaseMigrationClient) error {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_migration")
	retryPolicy.ShouldRetryOperation = assessmentWorkRequestShouldRetryFunc(timeout)

	response := oci_database_migration.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_migration.OperationStatusInProgress),
			string(oci_database_migration.OperationStatusAccepted),
			string(oci_database_migration.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_migration.OperationStatusSucceeded),
			string(oci_database_migration.OperationStatusFailed),
			string(oci_database_migration.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_migration.GetWorkRequestRequest{
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
		return e
	}

	// If work request did not succeed, surface backend error details.
	if response.Status == oci_database_migration.OperationStatusFailed || response.Status == oci_database_migration.OperationStatusCanceled {
		return getErrorFromDatabaseMigrationAssessmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	// Succeeded: do not require Resources list to contain identifiers.
	return nil
}

func getErrorFromDatabaseMigrationAssessmentWorkRequest(client *oci_database_migration.DatabaseMigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_migration.ListWorkRequestErrorsRequest{
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

func (s *DatabaseMigrationAssessmentResourceCrud) Get() error {
	request := oci_database_migration.GetAssessmentRequest{}

	tmp := s.D.Id()
	request.AssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.GetAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Assessment
	return nil
}

func (s *DatabaseMigrationAssessmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_migration.UpdateAssessmentRequest{}
	err := s.populateTopLevelPolymorphicUpdateAssessmentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseMigrationAssessmentResourceCrud) Delete() error {
	request := oci_database_migration.DeleteAssessmentRequest{}

	tmp := s.D.Id()
	request.AssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.DeleteAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	delWorkRequestErr := assessmentWaitForWorkRequest(workId, "assessment",
		oci_database_migration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseMigrationAssessmentResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_migration.MySqlAssessment:
		s.D.Set("database_combination", "MYSQL")

		s.D.Set("acceptable_downtime", v.AcceptableDowntime)

		s.D.Set("assessment_migration_type", v.AssessmentMigrationType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		s.D.Set("creation_type", v.CreationType)

		s.D.Set("database_data_size", v.DatabaseDataSize)

		s.D.Set("ddl_expectation", v.DdlExpectation)

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

		if v.MigrationId != nil {
			s.D.Set("migration_id", *v.MigrationId)
		}

		s.D.Set("network_speed_megabit_per_second", v.NetworkSpeedMegabitPerSecond)

		if v.SourceDatabaseConnection != nil {
			s.D.Set("source_database_connection", []interface{}{SourceAssessmentConnectionToMap(v.SourceDatabaseConnection)})
		} else {
			s.D.Set("source_database_connection", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnection != nil {
			s.D.Set("target_database_connection", []interface{}{TargetAssessmentConnectionToMap(v.TargetDatabaseConnection)})
		} else {
			s.D.Set("target_database_connection", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_database_migration.OracleAssessment:
		s.D.Set("database_combination", "ORACLE")

		if v.IsCdbSupported != nil {
			s.D.Set("is_cdb_supported", *v.IsCdbSupported)
		}

		s.D.Set("acceptable_downtime", v.AcceptableDowntime)

		s.D.Set("assessment_migration_type", v.AssessmentMigrationType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		s.D.Set("creation_type", v.CreationType)

		s.D.Set("database_data_size", v.DatabaseDataSize)

		s.D.Set("ddl_expectation", v.DdlExpectation)

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

		if v.MigrationId != nil {
			s.D.Set("migration_id", *v.MigrationId)
		}

		s.D.Set("network_speed_megabit_per_second", v.NetworkSpeedMegabitPerSecond)

		if v.SourceDatabaseConnection != nil {
			s.D.Set("source_database_connection", []interface{}{SourceAssessmentConnectionToMap(v.SourceDatabaseConnection)})
		} else {
			s.D.Set("source_database_connection", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnection != nil {
			s.D.Set("target_database_connection", []interface{}{TargetAssessmentConnectionToMap(v.TargetDatabaseConnection)})
		} else {
			s.D.Set("target_database_connection", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func AssessmentSummaryToMap(obj oci_database_migration.AssessmentSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MySqlAssessmentSummary:
		result["database_combination"] = "MYSQL"
	case oci_database_migration.OracleAssessmentSummary:
		result["database_combination"] = "ORACLE"

		if v.IsCdbSupported != nil {
			result["is_cdb_supported"] = bool(*v.IsCdbSupported)
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationAssessmentResourceCrud) mapToMySqlDatabaseObject(fieldKeyFormat string) (oci_database_migration.MySqlDatabaseObject, error) {
	result := oci_database_migration.MySqlDatabaseObject{}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	if schema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema")); ok {
		tmp := schema.(string)
		result.Schema = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationAssessmentResourceCrud) mapToOracleDatabaseObject(fieldKeyFormat string) (oci_database_migration.OracleDatabaseObject, error) {
	result := oci_database_migration.OracleDatabaseObject{}

	if isOmitExcludedTableFromReplication, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_omit_excluded_table_from_replication")); ok {
		tmp := isOmitExcludedTableFromReplication.(bool)
		result.IsOmitExcludedTableFromReplication = &tmp
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	if owner, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "owner")); ok {
		tmp := owner.(string)
		result.Owner = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationAssessmentResourceCrud) mapToSourceAssessmentConnection(fieldKeyFormat string) (oci_database_migration.SourceAssessmentConnection, error) {
	result := oci_database_migration.SourceAssessmentConnection{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func SourceAssessmentConnectionToMap(obj *oci_database_migration.SourceAssessmentConnection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *DatabaseMigrationAssessmentResourceCrud) mapToTargetAssessmentConnection(fieldKeyFormat string) (oci_database_migration.TargetAssessmentConnection, error) {
	result := oci_database_migration.TargetAssessmentConnection{}

	if connectionType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type")); ok {
		result.ConnectionType = oci_database_migration.ConnectionTypeEnum(connectionType.(string))
	}

	if databaseVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_version")); ok {
		tmp := databaseVersion.(string)
		result.DatabaseVersion = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if technologySubType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "technology_sub_type")); ok {
		tmp := technologySubType.(string)
		result.TechnologySubType = &tmp
	}

	if technologyType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "technology_type")); ok {
		result.TechnologyType = oci_database_migration.TechnologyTypeEnum(technologyType.(string))
	}

	return result, nil
}

func TargetAssessmentConnectionToMap(obj *oci_database_migration.TargetAssessmentConnection) map[string]interface{} {
	result := map[string]interface{}{}

	result["connection_type"] = string(obj.ConnectionType)

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.TechnologySubType != nil {
		result["technology_sub_type"] = string(*obj.TechnologySubType)
	}

	result["technology_type"] = string(obj.TechnologyType)

	return result
}

func (s *DatabaseMigrationAssessmentResourceCrud) populateTopLevelPolymorphicCreateAssessmentRequest(request *oci_database_migration.CreateAssessmentRequest) error {
	//discriminator
	databaseCombinationRaw, ok := s.D.GetOkExists("database_combination")
	var databaseCombination string
	if ok {
		databaseCombination = databaseCombinationRaw.(string)
	} else {
		databaseCombination = "" // default value
	}
	switch strings.ToLower(databaseCombination) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.CreateMySqlAssessmentDetails{}
		if bulkIncludeExcludeData, ok := s.D.GetOkExists("bulk_include_exclude_data"); ok {
			tmp := bulkIncludeExcludeData.(string)
			details.BulkIncludeExcludeData = &tmp
		}
		if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
			interfaces := excludeObjects.([]interface{})
			tmp := make([]oci_database_migration.MySqlDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
				converted, err := s.mapToMySqlDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
				details.ExcludeObjects = tmp
			}
		}
		if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
			interfaces := includeObjects.([]interface{})
			tmp := make([]oci_database_migration.MySqlDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
				converted, err := s.mapToMySqlDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("include_objects") {
				details.IncludeObjects = tmp
			}
		}
		if acceptableDowntime, ok := s.D.GetOkExists("acceptable_downtime"); ok {
			details.AcceptableDowntime = oci_database_migration.AcceptableDowntimeEnum(acceptableDowntime.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if creationType, ok := s.D.GetOkExists("creation_type"); ok {
			details.CreationType = oci_database_migration.CreationTypeEnum(creationType.(string))
		}
		if databaseDataSize, ok := s.D.GetOkExists("database_data_size"); ok {
			details.DatabaseDataSize = oci_database_migration.DatabaseDataSizeEnum(databaseDataSize.(string))
		}
		if ddlExpectation, ok := s.D.GetOkExists("ddl_expectation"); ok {
			details.DdlExpectation = oci_database_migration.DdlExpectationEnum(ddlExpectation.(string))
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
		if networkSpeedMegabitPerSecond, ok := s.D.GetOkExists("network_speed_megabit_per_second"); ok {
			details.NetworkSpeedMegabitPerSecond = oci_database_migration.NetworkSpeedMegabitPerSecondEnum(networkSpeedMegabitPerSecond.(string))
		}
		if sourceDatabaseConnection, ok := s.D.GetOkExists("source_database_connection"); ok {
			if tmpList := sourceDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_database_connection", 0)
				tmp, err := s.mapToSourceAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceDatabaseConnection = &tmp
			}
		}
		if targetDatabaseConnection, ok := s.D.GetOkExists("target_database_connection"); ok {
			if tmpList := targetDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_database_connection", 0)
				tmp, err := s.mapToTargetAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TargetDatabaseConnection = &tmp
			}
		}
		request.CreateAssessmentDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.CreateOracleAssessmentDetails{}
		if bulkIncludeExcludeData, ok := s.D.GetOkExists("bulk_include_exclude_data"); ok {
			tmp := bulkIncludeExcludeData.(string)
			details.BulkIncludeExcludeData = &tmp
		}
		if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
			interfaces := excludeObjects.([]interface{})
			tmp := make([]oci_database_migration.OracleDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
				converted, err := s.mapToOracleDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
				details.ExcludeObjects = tmp
			}
		}
		if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
			interfaces := includeObjects.([]interface{})
			tmp := make([]oci_database_migration.OracleDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
				converted, err := s.mapToOracleDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("include_objects") {
				details.IncludeObjects = tmp
			}
		}
		if acceptableDowntime, ok := s.D.GetOkExists("acceptable_downtime"); ok {
			details.AcceptableDowntime = oci_database_migration.AcceptableDowntimeEnum(acceptableDowntime.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if creationType, ok := s.D.GetOkExists("creation_type"); ok {
			details.CreationType = oci_database_migration.CreationTypeEnum(creationType.(string))
		}
		if databaseDataSize, ok := s.D.GetOkExists("database_data_size"); ok {
			details.DatabaseDataSize = oci_database_migration.DatabaseDataSizeEnum(databaseDataSize.(string))
		}
		if ddlExpectation, ok := s.D.GetOkExists("ddl_expectation"); ok {
			details.DdlExpectation = oci_database_migration.DdlExpectationEnum(ddlExpectation.(string))
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
		if networkSpeedMegabitPerSecond, ok := s.D.GetOkExists("network_speed_megabit_per_second"); ok {
			details.NetworkSpeedMegabitPerSecond = oci_database_migration.NetworkSpeedMegabitPerSecondEnum(networkSpeedMegabitPerSecond.(string))
		}
		if sourceDatabaseConnection, ok := s.D.GetOkExists("source_database_connection"); ok {
			if tmpList := sourceDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_database_connection", 0)
				tmp, err := s.mapToSourceAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceDatabaseConnection = &tmp
			}
		}
		if targetDatabaseConnection, ok := s.D.GetOkExists("target_database_connection"); ok {
			if tmpList := targetDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_database_connection", 0)
				tmp, err := s.mapToTargetAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TargetDatabaseConnection = &tmp
			}
		}
		request.CreateAssessmentDetails = details
	default:
		return fmt.Errorf("unknown database_combination '%v' was specified", databaseCombination)
	}
	return nil
}

func (s *DatabaseMigrationAssessmentResourceCrud) populateTopLevelPolymorphicUpdateAssessmentRequest(request *oci_database_migration.UpdateAssessmentRequest) error {
	//discriminator
	databaseCombinationRaw, ok := s.D.GetOkExists("database_combination")
	var databaseCombination string
	if ok {
		databaseCombination = databaseCombinationRaw.(string)
	} else {
		databaseCombination = "" // default value
	}
	switch strings.ToLower(databaseCombination) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.UpdateMySqlAssessmentDetails{}
		if acceptableDowntime, ok := s.D.GetOkExists("acceptable_downtime"); ok {
			details.AcceptableDowntime = oci_database_migration.AcceptableDowntimeEnum(acceptableDowntime.(string))
		}
		tmp := s.D.Id()
		request.AssessmentId = &tmp
		if creationType, ok := s.D.GetOkExists("creation_type"); ok {
			details.CreationType = oci_database_migration.CreationTypeEnum(creationType.(string))
		}
		if databaseDataSize, ok := s.D.GetOkExists("database_data_size"); ok {
			details.DatabaseDataSize = oci_database_migration.DatabaseDataSizeEnum(databaseDataSize.(string))
		}
		if ddlExpectation, ok := s.D.GetOkExists("ddl_expectation"); ok {
			details.DdlExpectation = oci_database_migration.DdlExpectationEnum(ddlExpectation.(string))
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
		if networkSpeedMegabitPerSecond, ok := s.D.GetOkExists("network_speed_megabit_per_second"); ok {
			details.NetworkSpeedMegabitPerSecond = oci_database_migration.NetworkSpeedMegabitPerSecondEnum(networkSpeedMegabitPerSecond.(string))
		}
		if sourceDatabaseConnection, ok := s.D.GetOkExists("source_database_connection"); ok {
			if tmpList := sourceDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_database_connection", 0)
				tmp, err := s.mapToSourceAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceDatabaseConnection = &tmp
			}
		}
		if targetDatabaseConnection, ok := s.D.GetOkExists("target_database_connection"); ok {
			if tmpList := targetDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_database_connection", 0)
				tmp, err := s.mapToTargetAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TargetDatabaseConnection = &tmp
			}
		}
		request.UpdateAssessmentDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.UpdateOracleAssessmentDetails{}
		if acceptableDowntime, ok := s.D.GetOkExists("acceptable_downtime"); ok {
			details.AcceptableDowntime = oci_database_migration.AcceptableDowntimeEnum(acceptableDowntime.(string))
		}
		tmp := s.D.Id()
		request.AssessmentId = &tmp
		if creationType, ok := s.D.GetOkExists("creation_type"); ok {
			details.CreationType = oci_database_migration.CreationTypeEnum(creationType.(string))
		}
		if databaseDataSize, ok := s.D.GetOkExists("database_data_size"); ok {
			details.DatabaseDataSize = oci_database_migration.DatabaseDataSizeEnum(databaseDataSize.(string))
		}
		if ddlExpectation, ok := s.D.GetOkExists("ddl_expectation"); ok {
			details.DdlExpectation = oci_database_migration.DdlExpectationEnum(ddlExpectation.(string))
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
		if networkSpeedMegabitPerSecond, ok := s.D.GetOkExists("network_speed_megabit_per_second"); ok {
			details.NetworkSpeedMegabitPerSecond = oci_database_migration.NetworkSpeedMegabitPerSecondEnum(networkSpeedMegabitPerSecond.(string))
		}
		if sourceDatabaseConnection, ok := s.D.GetOkExists("source_database_connection"); ok {
			if tmpList := sourceDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_database_connection", 0)
				tmp, err := s.mapToSourceAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceDatabaseConnection = &tmp
			}
		}
		if targetDatabaseConnection, ok := s.D.GetOkExists("target_database_connection"); ok {
			if tmpList := targetDatabaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_database_connection", 0)
				tmp, err := s.mapToTargetAssessmentConnection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TargetDatabaseConnection = &tmp
			}
		}
		request.UpdateAssessmentDetails = details
	default:
		return fmt.Errorf("unknown database_combination '%v' was specified", databaseCombination)
	}
	return nil
}

func (s *DatabaseMigrationAssessmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_migration.ChangeAssessmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.AssessmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.ChangeAssessmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
