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

func DatabaseManagementManagedDatabaseUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_collection": {
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
									"all_shared": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"authentication": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"common": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"consumer_group": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_collation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_tablespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"editions_enabled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_shared": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"implicit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"local_temp_tablespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_maintained": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password_versions": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"profile": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_connect": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"temp_tablespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_expiring": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_login": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_locked": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_password_changed": {
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

func readDatabaseManagementManagedDatabaseUsers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListUsersResponse
}

func (s *DatabaseManagementManagedDatabaseUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUsersDataSourceCrud) Get() error {
	request := oci_database_management.ListUsersRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUsersDataSource-", DatabaseManagementManagedDatabaseUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UserSummaryToMap(item))
	}
	managedDatabaseUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseUsersDataSource().Schema["user_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseUser["items"] = items
	}

	resources = append(resources, managedDatabaseUser)
	if err := s.D.Set("user_collection", resources); err != nil {
		return err
	}

	return nil
}

func UserSummaryToMap(obj oci_database_management.UserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultTablespace != nil {
		result["default_tablespace"] = string(*obj.DefaultTablespace)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Profile != nil {
		result["profile"] = string(*obj.Profile)
	}

	result["status"] = string(obj.Status)

	if obj.TempTablespace != nil {
		result["temp_tablespace"] = string(*obj.TempTablespace)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExpiring != nil {
		result["time_expiring"] = obj.TimeExpiring.String()
	}

	if obj.TimeLocked != nil {
		result["time_locked"] = obj.TimeLocked.String()
	}

	return result
}
