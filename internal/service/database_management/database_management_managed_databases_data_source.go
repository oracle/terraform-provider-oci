// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
)

func DatabaseManagementManagedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_option": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"deployment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
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
									"database_platform_name": {
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
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"workload_type": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListManagedDatabasesResponse
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) Get() error {
	request := oci_database_management.ListManagedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_database_management.ListManagedDatabasesDeploymentTypeEnum(deploymentType.(string))
	}

	if externalExadataInfrastructureId, ok := s.D.GetOkExists("external_exadata_infrastructure_id"); ok {
		tmp := externalExadataInfrastructureId.(string)
		request.ExternalExadataInfrastructureId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if managementOption, ok := s.D.GetOkExists("management_option"); ok {
		request.ManagementOption = oci_database_management.ListManagedDatabasesManagementOptionEnum(managementOption.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListManagedDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabasesDataSource-", DatabaseManagementManagedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedDatabaseSummaryToMap(item))
	}
	managedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabasesDataSource().Schema["managed_database_collection"].Elem.(*schema.Resource).Schema)
		managedDatabase["items"] = items
	}

	resources = append(resources, managedDatabase)
	if err := s.D.Set("managed_database_collection", resources); err != nil {
		return err
	}

	return nil
}
