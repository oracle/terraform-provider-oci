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

func DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseUserConsumerGroupPrivilege,
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
						"grant_option": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initial_group": {
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

func readSingularDatabaseManagementManagedDatabaseUserConsumerGroupPrivilege(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListConsumerGroupPrivilegesResponse
}

func (s *DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSourceCrud) Get() error {
	request := oci_database_management.ListConsumerGroupPrivilegesRequest{}

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

	response, err := s.Client.ListConsumerGroupPrivileges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource-", DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConsumerGroupPrivilegeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func ConsumerGroupPrivilegeSummaryToMap(obj oci_database_management.ConsumerGroupPrivilegeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["grant_option"] = string(obj.GrantOption)

	result["initial_group"] = string(obj.InitialGroup)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
