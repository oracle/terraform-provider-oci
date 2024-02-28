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

func DatabaseManagementExternalAsmInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalAsmInstance,
		Read:     readDatabaseManagementExternalAsmInstance,
		Update:   updateDatabaseManagementExternalAsmInstance,
		Delete:   deleteDatabaseManagementExternalAsmInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"external_asm_instance_id": {
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
			"adr_home_directory": {
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
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_asm_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
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

func createDatabaseManagementExternalAsmInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalAsmInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalAsmInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalAsmInstance(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalAsmInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalAsmInstance
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalAsmInstanceLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalAsmInstanceLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalAsmInstanceLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalAsmInstanceLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalAsmInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if externalAsmInstanceId, ok := s.D.GetOkExists("external_asm_instance_id"); ok {
		tmp := externalAsmInstanceId.(string)
		request.ExternalAsmInstanceId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalAsmInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalAsmInstance
	return nil
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) Get() error {
	request := oci_database_management.GetExternalAsmInstanceRequest{}

	tmp := s.D.Id()
	request.ExternalAsmInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalAsmInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalAsmInstance
	return nil
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalAsmInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExternalAsmInstanceId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalAsmInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalAsmInstance
	return nil
}

func (s *DatabaseManagementExternalAsmInstanceResourceCrud) SetData() error {
	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

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

	if s.Res.ExternalAsmId != nil {
		s.D.Set("external_asm_id", *s.Res.ExternalAsmId)
	}

	if s.Res.ExternalDbNodeId != nil {
		s.D.Set("external_db_node_id", *s.Res.ExternalDbNodeId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ExternalAsmInstanceSummaryToMap(obj oci_database_management.ExternalAsmInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

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

	if obj.ExternalAsmId != nil {
		result["external_asm_id"] = string(*obj.ExternalAsmId)
	}

	if obj.ExternalDbNodeId != nil {
		result["external_db_node_id"] = string(*obj.ExternalDbNodeId)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
