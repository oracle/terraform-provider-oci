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

func DatabaseManagementManagedDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementManagedDatabase,
		Read:     readDatabaseManagementManagedDatabase,
		Update:   updateDatabaseManagementManagedDatabase,
		Delete:   deleteDatabaseManagementManagedDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_database_id": {
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
			"database_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_sub_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"managed_database_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"management_option": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_system_id": {
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
			"workload_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementManagedDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ManagedDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Create() error {
	request := oci_database_management.UpdateManagedDatabaseRequest{}

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

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Get() error {
	request := oci_database_management.GetManagedDatabaseRequest{}

	tmp := s.D.Id()
	request.ManagedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Update() error {
	request := oci_database_management.UpdateManagedDatabaseRequest{}

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

	tmp := s.D.Id()
	request.ManagedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_status", s.Res.DatabaseStatus)

	s.D.Set("database_sub_type", s.Res.DatabaseSubType)

	s.D.Set("database_type", s.Res.DatabaseType)

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	managedDatabaseGroups := []interface{}{}
	for _, item := range s.Res.ManagedDatabaseGroups {
		managedDatabaseGroups = append(managedDatabaseGroups, ParentGroupToMap(item))
	}
	s.D.Set("managed_database_groups", managedDatabaseGroups)

	s.D.Set("management_option", s.Res.ManagementOption)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentContainerId != nil {
		s.D.Set("parent_container_id", *s.Res.ParentContainerId)
	}

	if s.Res.StorageSystemId != nil {
		s.D.Set("storage_system_id", *s.Res.StorageSystemId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("workload_type", s.Res.WorkloadType)

	return nil
}

func ManagedDatabaseSummaryToMap(obj oci_database_management.ManagedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["management_option"] = string(obj.ManagementOption)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentContainerId != nil {
		result["parent_container_id"] = string(*obj.ParentContainerId)
	}

	if obj.StorageSystemId != nil {
		result["storage_system_id"] = string(*obj.StorageSystemId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["workload_type"] = string(obj.WorkloadType)

	return result
}

func ParentGroupToMap(obj oci_database_management.ParentGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
