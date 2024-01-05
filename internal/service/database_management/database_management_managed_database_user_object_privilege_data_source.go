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

func DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseUserObjectPrivilege,
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
						"common": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"grant_option": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"grantor": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hierarchy": {
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
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schema_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_database_management_managed_database_user_object_privilege", "oci_database_management_managed_database_user_object_privileges"),
	}
}

func readSingularDatabaseManagementManagedDatabaseUserObjectPrivilege(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListObjectPrivilegesResponse
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceCrud) Get() error {
	request := oci_database_management.ListObjectPrivilegesRequest{}

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

	response, err := s.Client.ListObjectPrivileges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource-", DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ObjectPrivilegeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
