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

func DatabaseManagementManagedDatabaseUserRolesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseUserRoles,
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
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_collection": {
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
									"admin_option": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"common": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"delegate_option": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited": {
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
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabaseUserRoles(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserRolesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserRolesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListRolesResponse
}

func (s *DatabaseManagementManagedDatabaseUserRolesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserRolesDataSourceCrud) Get() error {
	request := oci_database_management.ListRolesRequest{}

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

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListRoles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRoles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseUserRolesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserRolesDataSource-", DatabaseManagementManagedDatabaseUserRolesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseUserRole := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RoleSummaryToMap(item))
	}
	managedDatabaseUserRole["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseUserRolesDataSource().Schema["role_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseUserRole["items"] = items
	}

	resources = append(resources, managedDatabaseUserRole)
	if err := s.D.Set("role_collection", resources); err != nil {
		return err
	}

	return nil
}

func RoleSummaryToMap(obj oci_database_management.RoleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["admin_option"] = string(obj.AdminOption)

	result["common"] = string(obj.Common)

	result["default_role"] = string(obj.DefaultRole)

	result["delegate_option"] = string(obj.DelegateOption)

	result["inherited"] = string(obj.Inherited)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
