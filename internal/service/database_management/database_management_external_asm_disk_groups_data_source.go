// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalAsmDiskGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalAsmDiskGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"external_asm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_asm_disk_group_collection": {
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
									"databases": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dismounting_instance_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"is_sparse": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mounting_instance_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"redundancy_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_size_in_mbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"used_percent": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"used_size_in_mbs": {
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

func readDatabaseManagementExternalAsmDiskGroups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmDiskGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalAsmDiskGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalAsmDiskGroupsResponse
}

func (s *DatabaseManagementExternalAsmDiskGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalAsmDiskGroupsDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalAsmDiskGroupsRequest{}

	if externalAsmId, ok := s.D.GetOkExists("external_asm_id"); ok {
		tmp := externalAsmId.(string)
		request.ExternalAsmId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalAsmDiskGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalAsmDiskGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalAsmDiskGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalAsmDiskGroupsDataSource-", DatabaseManagementExternalAsmDiskGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalAsmDiskGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalAsmDiskGroupSummaryToMap(item))
	}
	externalAsmDiskGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalAsmDiskGroupsDataSource().Schema["external_asm_disk_group_collection"].Elem.(*schema.Resource).Schema)
		externalAsmDiskGroup["items"] = items
	}

	resources = append(resources, externalAsmDiskGroup)
	if err := s.D.Set("external_asm_disk_group_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalAsmDiskGroupSummaryToMap(obj oci_database_management.ExternalAsmDiskGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["databases"] = obj.Databases

	if obj.DismountingInstanceCount != nil {
		result["dismounting_instance_count"] = int(*obj.DismountingInstanceCount)
	}

	if obj.IsSparse != nil {
		result["is_sparse"] = bool(*obj.IsSparse)
	}

	if obj.MountingInstanceCount != nil {
		result["mounting_instance_count"] = int(*obj.MountingInstanceCount)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["redundancy_type"] = string(obj.RedundancyType)

	if obj.TotalSizeInMBs != nil {
		result["total_size_in_mbs"] = strconv.FormatInt(*obj.TotalSizeInMBs, 10)
	}

	if obj.UsedPercent != nil {
		result["used_percent"] = float32(*obj.UsedPercent)
	}

	if obj.UsedSizeInMBs != nil {
		result["used_size_in_mbs"] = strconv.FormatInt(*obj.UsedSizeInMBs, 10)
	}

	return result
}
