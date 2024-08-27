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

func DatabaseManagementManagedDatabaseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabase,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Optional: true,
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
	}
}

func readSingularDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetManagedDatabaseResponse
}

func (s *DatabaseManagementManagedDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseDataSourceCrud) Get() error {
	request := oci_database_management.GetManagedDatabaseRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabasePlatformName != nil {
		s.D.Set("database_platform_name", *s.Res.DatabasePlatformName)
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

	dbmgmtFeatureConfigs := []interface{}{}
	for _, item := range s.Res.DbmgmtFeatureConfigs {
		dbmgmtFeatureConfigs = append(dbmgmtFeatureConfigs, DatabaseFeatureConfigurationToMap(item))
	}
	s.D.Set("dbmgmt_feature_configs", dbmgmtFeatureConfigs)

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
