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

func DatabaseManagementManagedDatabaseUserRoleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseUserRole,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
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
	}
}

func readSingularDatabaseManagementManagedDatabaseUserRole(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserRoleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserRoleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListRolesResponse
}

func (s *DatabaseManagementManagedDatabaseUserRoleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserRoleDataSourceCrud) Get() error {
	request := oci_database_management.ListRolesRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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
	return nil
}

func (s *DatabaseManagementManagedDatabaseUserRoleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserRoleDataSource-", DatabaseManagementManagedDatabaseUserRoleDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RoleSummaryToMap(item))
	}
	s.D.Set("items", items)

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
