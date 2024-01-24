// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseExternalNonContainerDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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

func createDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalNonContainerDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExternalNonContainerDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExternalNonContainerDatabase

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

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
	request.ExternalNonContainerDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteExternalNonContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.ExternalNonContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
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

	s.D.Set("database_configuration", s.Res.DatabaseConfiguration)

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{DatabaseManagementNonContainerConfigToMap(s.Res.DatabaseManagementConfig)})
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

	if s.Res.OperationsInsightsConfig != nil {
		s.D.Set("operations_insights_config", []interface{}{OperationsInsightsConfigToMap(s.Res.OperationsInsightsConfig)})
	} else {
		s.D.Set("operations_insights_config", nil)
	}

	if s.Res.StackMonitoringConfig != nil {
		s.D.Set("stack_monitoring_config", []interface{}{StackMonitoringNonContainerConfigToMap(s.Res.StackMonitoringConfig)})
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

func DatabaseManagementNonContainerConfigToMap(obj *oci_database.DatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseManagementConnectionId != nil {
		result["database_management_connection_id"] = string(*obj.DatabaseManagementConnectionId)
	}

	result["database_management_status"] = string(obj.DatabaseManagementStatus)

	result["license_model"] = string(obj.LicenseModel)

	return result
}

func OperationsInsightsConfigToMap(obj *oci_database.OperationsInsightsConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OperationsInsightsConnectorId != nil {
		result["operations_insights_connector_id"] = string(*obj.OperationsInsightsConnectorId)
	}

	result["operations_insights_status"] = string(obj.OperationsInsightsStatus)

	return result
}

func StackMonitoringNonContainerConfigToMap(obj *oci_database.StackMonitoringConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StackMonitoringConnectorId != nil {
		result["stack_monitoring_connector_id"] = string(*obj.StackMonitoringConnectorId)
	}

	result["stack_monitoring_status"] = string(obj.StackMonitoringStatus)

	return result
}

func (s *DatabaseExternalNonContainerDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExternalNonContainerDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExternalNonContainerDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExternalNonContainerDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
}
