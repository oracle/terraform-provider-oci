// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func DatabaseExternalContainerDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalContainerDatabase,
		Read:     readDatabaseExternalContainerDatabase,
		Update:   updateDatabaseExternalContainerDatabase,
		Delete:   deleteDatabaseExternalContainerDatabase,
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
			"database_configuration": {
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

func createDatabaseExternalContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExternalContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExternalContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExternalContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalContainerDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExternalContainerDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateAvailable),
		string(oci_database.ExternalContainerDatabaseLifecycleStateNotConnected),
	}
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) Create() error {
	request := oci_database.CreateExternalContainerDatabaseRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExternalContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExternalContainerDatabase

	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) Get() error {
	request := oci_database.GetExternalContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExternalContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalContainerDatabase
	return nil
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateExternalContainerDatabaseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
	request.ExternalContainerDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExternalContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteExternalContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExternalContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalContainerDatabaseResourceCrud) SetData() error {
	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_configuration", s.Res.DatabaseConfiguration)

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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

func (s *DatabaseExternalContainerDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExternalContainerDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExternalContainerDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExternalContainerDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
}
