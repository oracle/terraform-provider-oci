// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudAsmUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudAsmUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_asm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_asm_user_collection": {
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

func readDatabaseManagementCloudAsmUsers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudAsmUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudAsmUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudAsmUsersResponse
}

func (s *DatabaseManagementCloudAsmUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudAsmUsersDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudAsmUsersRequest{}

	if cloudAsmId, ok := s.D.GetOkExists("cloud_asm_id"); ok {
		tmp := cloudAsmId.(string)
		request.CloudAsmId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudAsmUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudAsmUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudAsmUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudAsmUsersDataSource-", DatabaseManagementCloudAsmUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudAsmUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudAsmUserSummaryToMap(item))
	}
	cloudAsmUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudAsmUsersDataSource().Schema["cloud_asm_user_collection"].Elem.(*schema.Resource).Schema)
		cloudAsmUser["items"] = items
	}

	resources = append(resources, cloudAsmUser)
	if err := s.D.Set("cloud_asm_user_collection", resources); err != nil {
		return err
	}

	return nil
}

func CloudAsmUserSummaryToMap(obj oci_database_management.CloudAsmUserSummary) map[string]interface{} {
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
