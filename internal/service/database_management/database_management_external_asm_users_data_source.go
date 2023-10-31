// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalAsmUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalAsmUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"external_asm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_asm_user_collection": {
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
									"asm_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"privileges": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
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

func readDatabaseManagementExternalAsmUsers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalAsmUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalAsmUsersResponse
}

func (s *DatabaseManagementExternalAsmUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalAsmUsersDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalAsmUsersRequest{}

	if externalAsmId, ok := s.D.GetOkExists("external_asm_id"); ok {
		tmp := externalAsmId.(string)
		request.ExternalAsmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalAsmUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalAsmUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalAsmUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalAsmUsersDataSource-", DatabaseManagementExternalAsmUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalAsmUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalAsmUserSummaryToMap(item))
	}
	externalAsmUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalAsmUsersDataSource().Schema["external_asm_user_collection"].Elem.(*schema.Resource).Schema)
		externalAsmUser["items"] = items
	}

	resources = append(resources, externalAsmUser)
	if err := s.D.Set("external_asm_user_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalAsmUserSummaryToMap(obj oci_database_management.ExternalAsmUserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AsmId != nil {
		result["asm_id"] = string(*obj.AsmId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["privileges"] = obj.Privileges

	return result
}
