// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseExternalPluggableDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalPluggableDatabase,
		Read:     readDatabaseExternalPluggableDatabase,
		Update:   updateDatabaseExternalPluggableDatabase,
		Delete:   deleteDatabaseExternalPluggableDatabase,
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
			"external_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"source_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"operations_insights_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"operations_insights_connector_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"operations_insights_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"stack_monitoring_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"stack_monitoring_connector_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stack_monitoring_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDatabaseExternalPluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExternalPluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExternalPluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExternalPluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalPluggableDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExternalPluggableDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateAvailable),
		string(oci_database.ExternalPluggableDatabaseLifecycleStateNotConnected),
	}
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) Create() error {
	request := oci_database.CreateExternalPluggableDatabaseRequest{}

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

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExternalPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExternalPluggableDatabase

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) Get() error {
	request := oci_database.GetExternalPluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalPluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExternalPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalPluggableDatabase
	return nil
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateExternalPluggableDatabaseRequest{}

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
	request.ExternalPluggableDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExternalPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteExternalPluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalPluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExternalPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) SetData() error {
	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_configuration", s.Res.DatabaseConfiguration)

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{DatabaseManagementConfigurationToMap(s.Res.DatabaseManagementConfig)})
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

	if s.Res.ExternalContainerDatabaseId != nil {
		s.D.Set("external_container_database_id", *s.Res.ExternalContainerDatabaseId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	if s.Res.OperationsInsightsConfig != nil {
		s.D.Set("operations_insights_config", []interface{}{OperationsInsightsConfigurationToMap(s.Res.OperationsInsightsConfig)})
	} else {
		s.D.Set("operations_insights_config", nil)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	if s.Res.StackMonitoringConfig != nil {
		s.D.Set("stack_monitoring_config", []interface{}{StackMonitoringConfigToMap(s.Res.StackMonitoringConfig)})
	} else {
		s.D.Set("stack_monitoring_config", nil)
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

func DatabaseManagementConfigurationToMap(obj *oci_database.DatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseManagementConnectionId != nil {
		result["database_management_connection_id"] = string(*obj.DatabaseManagementConnectionId)
	}

	result["database_management_status"] = string(obj.DatabaseManagementStatus)

	result["license_model"] = string(obj.LicenseModel)

	return result
}

func OperationsInsightsConfigurationToMap(obj *oci_database.OperationsInsightsConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OperationsInsightsConnectorId != nil {
		result["operations_insights_connector_id"] = string(*obj.OperationsInsightsConnectorId)
	}

	result["operations_insights_status"] = string(obj.OperationsInsightsStatus)

	return result
}

func StackMonitoringConfigToMap(obj *oci_database.StackMonitoringConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StackMonitoringConnectorId != nil {
		result["stack_monitoring_connector_id"] = string(*obj.StackMonitoringConnectorId)
	}

	result["stack_monitoring_status"] = string(obj.StackMonitoringStatus)

	return result
}

func (s *DatabaseExternalPluggableDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExternalPluggableDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExternalPluggableDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExternalPluggableDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
}
