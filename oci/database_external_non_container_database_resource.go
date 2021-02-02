// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v35/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v35/workrequests"
)

func init() {
	RegisterResource("oci_database_external_non_container_database", DatabaseExternalNonContainerDatabaseResource())
}

func DatabaseExternalNonContainerDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExternalNonContainerDatabase,
		Read:     readDatabaseExternalNonContainerDatabase,
		Update:   updateDatabaseExternalNonContainerDatabase,
		Delete:   deleteDatabaseExternalNonContainerDatabase,
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"database_management_connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_packs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ncharacter_set": {
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
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return CreateResource(d, sync)
}

func readDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return ReadResource(sync)
}

func updateDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return UpdateResource(d, sync)
}

func deleteDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.DisableNotFoundRetries = true
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return DeleteResource(d, sync)
}

type DatabaseExternalNonContainerDatabaseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.ExternalNonContainerDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateAvailable),
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateNotConnected),
	}
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) Create() error {
	request := oci_database.CreateExternalNonContainerDatabaseRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.ExternalNonContainerDatabase
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) Get() error {
	request := oci_database.GetExternalNonContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalNonContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalNonContainerDatabase
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateExternalNonContainerDatabaseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	/*	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}*/

	tmp := s.D.Id()
	request.ExternalNonContainerDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.ExternalNonContainerDatabase
	return s.Get()
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteExternalNonContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalNonContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExternalNonContainerDatabase(context.Background(), request)

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) SetData() error {
	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{DatabaseManagementConfigToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbId != nil {
		s.D.Set("db_id", *s.Res.DbId)
	}

	if s.Res.DbPacks != nil {
		s.D.Set("db_packs", *s.Res.DbPacks)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

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

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}

func DatabaseManagementConfigToMap(obj *oci_database.DatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseManagementConnectionId != nil {
		result["database_management_connection_id"] = string(*obj.DatabaseManagementConnectionId)
	}

	result["database_management_status"] = string(obj.DatabaseManagementStatus)

	result["license_model"] = string(obj.LicenseModel)

	return result
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExternalNonContainerDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExternalNonContainerDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExternalNonContainerDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
