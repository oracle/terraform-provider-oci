// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database_migration "github.com/oracle/oci-go-sdk/v46/databasemigration"
)

func init() {
	RegisterResource("oci_database_migration_job", DatabaseMigrationJobResource())
}

func DatabaseMigrationJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseMigrationJob,
		Read:     readDatabaseMigrationJob,
		Update:   updateDatabaseMigrationJob,
		Delete:   deleteDatabaseMigrationJob,
		Schema: map[string]*schema.Schema{
			// Required
			"job_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"migration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"current_phase": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phases": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"duration_in_ms": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"progress": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unsupported_objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDatabaseMigrationJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return CreateResource(d, sync)
}

func readDatabaseMigrationJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return ReadResource(sync)
}

func updateDatabaseMigrationJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return UpdateResource(d, sync)
}

func deleteDatabaseMigrationJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseMigrationJobResourceCrud struct {
	BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.Job
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseMigrationJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.JobLifecycleStatesInProgress),
	}
}

func (s *DatabaseMigrationJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.JobLifecycleStatesSucceeded),
	}
}

func (s *DatabaseMigrationJobResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseMigrationJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.JobLifecycleStatesTerminated),
	}
}

func (s *DatabaseMigrationJobResourceCrud) Create() error {
	request := oci_database_migration.UpdateJobRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database_migration")
	response, err := s.Client.UpdateJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *DatabaseMigrationJobResourceCrud) Get() error {
	request := oci_database_migration.GetJobRequest{}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.GetJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *DatabaseMigrationJobResourceCrud) Update() error {
	request := oci_database_migration.UpdateJobRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *DatabaseMigrationJobResourceCrud) Delete() error {
	request := oci_database_migration.DeleteJobRequest{}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.DeleteJob(context.Background(), request)
	return err
}

func (s *DatabaseMigrationJobResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MigrationId != nil {
		s.D.Set("migration_id", *s.Res.MigrationId)
	}

	if s.Res.Progress != nil {
		s.D.Set("progress", []interface{}{MigrationJobProgressResourceToMap(s.Res.Progress)})
	} else {
		s.D.Set("progress", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	unsupportedObjects := []interface{}{}
	for _, item := range s.Res.UnsupportedObjects {
		unsupportedObjects = append(unsupportedObjects, UnsupportedDatabaseObjectToMap(item))
	}
	s.D.Set("unsupported_objects", unsupportedObjects)

	return nil
}

func JobSummaryToMap(obj oci_database_migration.JobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
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

	if obj.MigrationId != nil {
		result["migration_id"] = string(*obj.MigrationId)
	}

	if obj.Progress != nil {
		result["progress"] = []interface{}{MigrationJobProgressSummaryToMap(obj.Progress)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = systemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func MigrationJobProgressResourceToMap(obj *oci_database_migration.MigrationJobProgressResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["current_phase"] = string(obj.CurrentPhase)

	result["current_status"] = string(obj.CurrentStatus)

	phases := []interface{}{}
	for _, item := range obj.Phases {
		phases = append(phases, PhaseStatusToMap(item))
	}
	result["phases"] = phases

	return result
}

func MigrationJobProgressSummaryToMap(obj *oci_database_migration.MigrationJobProgressSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["current_phase"] = string(obj.CurrentPhase)

	result["current_status"] = string(obj.CurrentStatus)

	if obj.JobProgress != nil {
		result["job_progress"] = int(*obj.JobProgress)
	}

	return result
}

func PhaseStatusToMap(obj oci_database_migration.PhaseStatus) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DurationInMs != nil {
		result["duration_in_ms"] = int(*obj.DurationInMs)
	}

	result["name"] = string(obj.Name)

	if obj.Progress != nil {
		result["progress"] = int(*obj.Progress)
	}

	result["status"] = string(obj.Status)

	return result
}

func UnsupportedDatabaseObjectToMap(obj oci_database_migration.UnsupportedDatabaseObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	result["type"] = string(obj.Type)

	return result
}
