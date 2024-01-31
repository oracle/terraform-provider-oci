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

func DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseUserObjectPrivileges,
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
			"object_privilege_collection": {
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
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabaseUserObjectPrivileges(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListObjectPrivilegesResponse
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSourceCrud) Get() error {
	request := oci_database_management.ListObjectPrivilegesRequest{}

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

	response, err := s.Client.ListObjectPrivileges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListObjectPrivileges(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource-", DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseUserObjectPrivilege := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ObjectPrivilegeSummaryToMap(item))
	}
	managedDatabaseUserObjectPrivilege["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource().Schema["object_privilege_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseUserObjectPrivilege["items"] = items
	}

	resources = append(resources, managedDatabaseUserObjectPrivilege)
	if err := s.D.Set("object_privilege_collection", resources); err != nil {
		return err
	}

	return nil
}

func ObjectPrivilegeSummaryToMap(obj oci_database_management.ObjectPrivilegeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["common"] = string(obj.Common)

	result["grant_option"] = string(obj.GrantOption)

	if obj.Grantor != nil {
		result["grantor"] = string(*obj.Grantor)
	}

	result["hierarchy"] = string(obj.Hierarchy)

	result["inherited"] = string(obj.Inherited)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Object != nil {
		result["object"] = string(*obj.Object)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.SchemaType != nil {
		result["schema_type"] = string(*obj.SchemaType)
	}

	return result
}
