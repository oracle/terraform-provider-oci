// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v58/databasemanagement"
)

func DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabasesUserSystemPrivileges,
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
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_privilege_collection": {
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

func readDatabaseManagementManagedDatabasesUserSystemPrivileges(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListSystemPrivilegesResponse
}

func (s *DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSourceCrud) Get() error {
	request := oci_database_management.ListSystemPrivilegesRequest{}

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

	response, err := s.Client.ListSystemPrivileges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSystemPrivileges(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource-", DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabasesUserSystemPrivilege := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SystemPrivilegeSummaryToMap(item))
	}
	managedDatabasesUserSystemPrivilege["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource().Schema["system_privilege_collection"].Elem.(*schema.Resource).Schema)
		managedDatabasesUserSystemPrivilege["items"] = items
	}

	resources = append(resources, managedDatabasesUserSystemPrivilege)
	if err := s.D.Set("system_privilege_collection", resources); err != nil {
		return err
	}

	return nil
}
