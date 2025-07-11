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

func DatabaseManagementCloudDbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudDbHome,
		Read:     readDatabaseManagementCloudDbHome,
		Update:   updateDatabaseManagementCloudDbHome,
		Delete:   deleteDatabaseManagementCloudDbHome,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_db_home_id": {
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
			"cloud_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"component_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dbaas_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
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

func createDatabaseManagementCloudDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementCloudDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudDbHome(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementCloudDbHomeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.CloudDbHome
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.CloudDbHomeLifecycleStateCreating),
	}
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbHomeLifecycleStateActive),
	}
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.CloudDbHomeLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbHomeLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) Create() error {
	request := oci_database_management.UpdateCloudDbHomeRequest{}

	if cloudDbHomeId, ok := s.D.GetOkExists("cloud_db_home_id"); ok {
		tmp := cloudDbHomeId.(string)
		request.CloudDbHomeId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbHome
	return nil
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) Get() error {
	request := oci_database_management.GetCloudDbHomeRequest{}

	tmp := s.D.Id()
	request.CloudDbHomeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetCloudDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbHome
	return nil
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) Update() error {
	request := oci_database_management.UpdateCloudDbHomeRequest{}

	tmp := s.D.Id()
	request.CloudDbHomeId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbHome
	return nil
}

func (s *DatabaseManagementCloudDbHomeResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CloudDbSystemId != nil {
		s.D.Set("cloud_db_system_id", *s.Res.CloudDbSystemId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DbaasId != nil {
		s.D.Set("dbaas_id", *s.Res.DbaasId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
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

func CloudDbHomeSummaryToMap(obj oci_database_management.CloudDbHomeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudDbSystemId != nil {
		result["cloud_db_system_id"] = string(*obj.CloudDbSystemId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.DbaasId != nil {
		result["dbaas_id"] = string(*obj.DbaasId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
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
