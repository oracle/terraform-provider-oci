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

func DatabaseManagementExternalDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_database_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_platform_name": {
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
									"db_management_config": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"connector_id": {
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
									"db_system_info": {
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
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"exadata_infra_info": {
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
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"id": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"db_unique_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"dbmgmt_feature_configs": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"connector_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"connector_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"database_connector_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"management_agent_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"private_end_point_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"database_connection_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"connection_credentials": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"credential_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"credential_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"named_credential_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"password_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"role": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"ssl_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"user_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"connection_string": {
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
																		"port": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"protocol": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"service": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"feature": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"feature_status": {
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_db_home_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"host_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"instance_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"instance_number": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalDatabasesResponse
}

func (s *DatabaseManagementExternalDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalDatabasesDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalDatabaseId, ok := s.D.GetOkExists("external_database_id"); ok {
		tmp := externalDatabaseId.(string)
		request.ExternalDatabaseId = &tmp
	}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalDatabasesDataSource-", DatabaseManagementExternalDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalDatabaseSummaryToMap(item))
	}
	externalDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalDatabasesDataSource().Schema["external_database_collection"].Elem.(*schema.Resource).Schema)
		externalDatabase["items"] = items
	}

	resources = append(resources, externalDatabase)
	if err := s.D.Set("external_database_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseManagementConfigToMap(obj *oci_database_management.DatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	result["database_management_status"] = string(obj.DatabaseManagementStatus)

	result["license_model"] = string(obj.LicenseModel)

	return result
}

func ExternalDatabaseInstanceToMap(obj oci_database_management.ExternalDatabaseInstance) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.InstanceName != nil {
		result["instance_name"] = string(*obj.InstanceName)
	}

	if obj.InstanceNumber != nil {
		result["instance_number"] = int(*obj.InstanceNumber)
	}

	return result
}

func ExternalDatabaseSummaryToMap(obj oci_database_management.ExternalDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabasePlatformName != nil {
		result["database_platform_name"] = string(*obj.DatabasePlatformName)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	if obj.DbManagementConfig != nil {
		result["db_management_config"] = []interface{}{DatabaseManagementConfigToMap(obj.DbManagementConfig)}
	}

	if obj.DbSystemInfo != nil {
		result["db_system_info"] = []interface{}{ExternalDbSystemBasicInfoToMap(obj.DbSystemInfo)}
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	dbmgmtFeatureConfigs := []interface{}{}
	for _, item := range obj.DbmgmtFeatureConfigs {
		dbmgmtFeatureConfigs = append(dbmgmtFeatureConfigs, DatabaseFeatureConfigurationToMap(item))
	}
	result["dbmgmt_feature_configs"] = dbmgmtFeatureConfigs

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalContainerDatabaseId != nil {
		result["external_container_database_id"] = string(*obj.ExternalContainerDatabaseId)
	}

	if obj.ExternalDbHomeId != nil {
		result["external_db_home_id"] = string(*obj.ExternalDbHomeId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	instanceDetails := []interface{}{}
	for _, item := range obj.InstanceDetails {
		instanceDetails = append(instanceDetails, ExternalDatabaseInstanceToMap(item))
	}
	result["instance_details"] = instanceDetails

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func ExternalDbSystemBasicInfoToMap(obj *oci_database_management.ExternalDbSystemBasicInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExadataInfraInfo != nil {
		result["exadata_infra_info"] = []interface{}{ExternalExadataInfraBasicInfoToMap(obj.ExadataInfraInfo)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func ExternalExadataInfraBasicInfoToMap(obj *oci_database_management.ExternalExadataInfraBasicInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
