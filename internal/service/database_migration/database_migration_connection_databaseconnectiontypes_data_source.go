// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationConnectionDatabaseconnectiontypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationConnectionDatabaseconnectiontypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"technology_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_connection_type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"connection_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"technology_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"database_versions": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"technology_sub_types": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"database_versions": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"technology_sub_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"technology_sub_type_display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"technology_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationConnectionDatabaseconnectiontypes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionDatabaseconnectiontypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationConnectionDatabaseconnectiontypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListDatabaseConnectionTypeResponse
}

func (s *DatabaseMigrationConnectionDatabaseconnectiontypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationConnectionDatabaseconnectiontypesDataSourceCrud) Get() error {
	request := oci_database_migration.ListDatabaseConnectionTypeRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectionType, ok := s.D.GetOkExists("connection_type"); ok {
		interfaces := connectionType.([]interface{})
		tmp := make([]oci_database_migration.ConnectionTypeEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp = append(tmp, oci_database_migration.ConnectionTypeEnum(interfaces[i].(string)))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("connection_type") {
			request.ConnectionType = tmp
		}
	}

	if sourceConnectionId, ok := s.D.GetOkExists("source_connection_id"); ok {
		tmp := sourceConnectionId.(string)
		request.SourceConnectionId = &tmp
	}

	if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
		interfaces := technologyType.([]interface{})
		tmp := make([]oci_database_migration.TechnologyTypeEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp = append(tmp, oci_database_migration.TechnologyTypeEnum(interfaces[i].(string)))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("technology_type") {
			request.TechnologyType = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListDatabaseConnectionType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseConnectionType(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationConnectionDatabaseconnectiontypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationConnectionDatabaseconnectiontypesDataSource-", DatabaseMigrationConnectionDatabaseconnectiontypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	connectionDatabaseconnectiontype := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseConnectionTypeSummaryToMap(item))
	}
	connectionDatabaseconnectiontype["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationConnectionDatabaseconnectiontypesDataSource().Schema["database_connection_type_collection"].Elem.(*schema.Resource).Schema)
		connectionDatabaseconnectiontype["items"] = items
	}

	resources = append(resources, connectionDatabaseconnectiontype)
	if err := s.D.Set("database_connection_type_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseConnectionTypeSummaryToMap(obj oci_database_migration.DatabaseConnectionTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["connection_type"] = string(obj.ConnectionType)

	technologyTypes := []interface{}{}
	for _, item := range obj.TechnologyTypes {
		technologyTypes = append(technologyTypes, DatabaseTechnologyTypeToMap(item))
	}
	result["technology_types"] = technologyTypes

	return result
}

func DatabaseTechnologySubTypeToMap(obj oci_database_migration.DatabaseTechnologySubType) map[string]interface{} {
	result := map[string]interface{}{}

	result["database_versions"] = obj.DatabaseVersions

	if obj.TechnologySubType != nil {
		result["technology_sub_type"] = string(*obj.TechnologySubType)
	}

	if obj.TechnologySubTypeDisplayName != nil {
		result["technology_sub_type_display_name"] = string(*obj.TechnologySubTypeDisplayName)
	}

	return result
}

func DatabaseTechnologyTypeToMap(obj oci_database_migration.DatabaseTechnologyType) map[string]interface{} {
	result := map[string]interface{}{}

	result["database_versions"] = obj.DatabaseVersions

	technologySubTypes := []interface{}{}
	for _, item := range obj.TechnologySubTypes {
		technologySubTypes = append(technologySubTypes, DatabaseTechnologySubTypeToMap(item))
	}
	result["technology_sub_types"] = technologySubTypes

	result["technology_type"] = string(obj.TechnologyType)

	return result
}
