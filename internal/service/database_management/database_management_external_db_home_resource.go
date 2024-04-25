// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalDbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalDbHome,
		Read:     readDatabaseManagementExternalDbHome,
		Update:   updateDatabaseManagementExternalDbHome,
		Delete:   deleteDatabaseManagementExternalDbHome,
		Schema: map[string]*schema.Schema{
			// Required
			"external_db_home_id": {
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

			// Computed
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"component_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_directory": {
				Type:     schema.TypeString,
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

func createDatabaseManagementExternalDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalDbHome(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalDbHomeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalDbHome
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbHomeLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbHomeLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbHomeLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbHomeLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalDbHomeRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if externalDbHomeId, ok := s.D.GetOkExists("external_db_home_id"); ok {
		tmp := externalDbHomeId.(string)
		request.ExternalDbHomeId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbHome
	return nil
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) Get() error {
	request := oci_database_management.GetExternalDbHomeRequest{}

	tmp := s.D.Id()
	request.ExternalDbHomeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbHome
	return nil
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalDbHomeRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExternalDbHomeId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbHome
	return nil
}

func (s *DatabaseManagementExternalDbHomeResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeDirectory != nil {
		s.D.Set("home_directory", *s.Res.HomeDirectory)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ExternalDbHomeSummaryToMap(obj oci_database_management.ExternalDbHomeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HomeDirectory != nil {
		result["home_directory"] = string(*obj.HomeDirectory)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
