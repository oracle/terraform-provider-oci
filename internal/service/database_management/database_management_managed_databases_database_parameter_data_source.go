// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v56/databasemanagement"
)

func DatabaseManagementManagedDatabasesDatabaseParameterDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabasesDatabaseParameter,
		Schema: map[string]*schema.Schema{
			"is_allowed_values_included": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"database_name": {
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
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"allowed_values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"ordinal": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"constraint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_id": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_adjusted": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_basic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_deprecated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_instance_modifiable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pdb_modifiable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_session_modifiable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_specified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_system_modifiable": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"number": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"ordinal": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabasesDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesDatabaseParameterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabasesDatabaseParameterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListDatabaseParametersResponse
}

func (s *DatabaseManagementManagedDatabasesDatabaseParameterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasesDatabaseParameterDataSourceCrud) Get() error {
	request := oci_database_management.ListDatabaseParametersRequest{}

	if isAllowedValuesIncluded, ok := s.D.GetOkExists("is_allowed_values_included"); ok {
		tmp := isAllowedValuesIncluded.(bool)
		request.IsAllowedValuesIncluded = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		request.Source = oci_database_management.ListDatabaseParametersSourceEnum(source.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListDatabaseParameters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabasesDatabaseParameterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabasesDatabaseParameterDataSource-", DatabaseManagementManagedDatabasesDatabaseParameterDataSource(), s.D))

	if s.Res.DatabaseName != nil {
		s.D.Set("database_name", *s.Res.DatabaseName)
	}

	s.D.Set("database_sub_type", s.Res.DatabaseSubType)

	s.D.Set("database_type", s.Res.DatabaseType)

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseParameterSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AllowedParameterValueToMap(obj oci_database_management.AllowedParameterValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.Ordinal != nil {
		result["ordinal"] = float32(*obj.Ordinal)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func DatabaseParameterSummaryToMap(obj oci_database_management.DatabaseParameterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	allowedValues := []interface{}{}
	for _, item := range obj.AllowedValues {
		allowedValues = append(allowedValues, AllowedParameterValueToMap(item))
	}
	result["allowed_values"] = allowedValues

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	result["constraint"] = string(obj.Constraint)

	if obj.ContainerId != nil {
		result["container_id"] = float32(*obj.ContainerId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayValue != nil {
		result["display_value"] = string(*obj.DisplayValue)
	}

	if obj.IsAdjusted != nil {
		result["is_adjusted"] = bool(*obj.IsAdjusted)
	}

	if obj.IsBasic != nil {
		result["is_basic"] = bool(*obj.IsBasic)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.IsDeprecated != nil {
		result["is_deprecated"] = bool(*obj.IsDeprecated)
	}

	if obj.IsInstanceModifiable != nil {
		result["is_instance_modifiable"] = bool(*obj.IsInstanceModifiable)
	}

	result["is_modified"] = string(obj.IsModified)

	if obj.IsPdbModifiable != nil {
		result["is_pdb_modifiable"] = bool(*obj.IsPdbModifiable)
	}

	if obj.IsSessionModifiable != nil {
		result["is_session_modifiable"] = bool(*obj.IsSessionModifiable)
	}

	if obj.IsSpecified != nil {
		result["is_specified"] = bool(*obj.IsSpecified)
	}

	result["is_system_modifiable"] = string(obj.IsSystemModifiable)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Number != nil {
		result["number"] = float32(*obj.Number)
	}

	if obj.Ordinal != nil {
		result["ordinal"] = float32(*obj.Ordinal)
	}

	if obj.Sid != nil {
		result["sid"] = string(*obj.Sid)
	}

	result["type"] = string(obj.Type)

	if obj.UpdateComment != nil {
		result["update_comment"] = string(*obj.UpdateComment)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
