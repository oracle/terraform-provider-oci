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

func DatabaseManagementCloudAsmDiskGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudAsmDiskGroups,
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
			"cloud_asm_disk_group_collection": {
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

func readDatabaseManagementCloudAsmDiskGroups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudAsmDiskGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudAsmDiskGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudAsmDiskGroupsResponse
}

func (s *DatabaseManagementCloudAsmDiskGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudAsmDiskGroupsDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudAsmDiskGroupsRequest{}

	if cloudAsmId, ok := s.D.GetOkExists("cloud_asm_id"); ok {
		tmp := cloudAsmId.(string)
		request.CloudAsmId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudAsmDiskGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudAsmDiskGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudAsmDiskGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudAsmDiskGroupsDataSource-", DatabaseManagementCloudAsmDiskGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudAsmDiskGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudAsmDiskGroupSummaryToMap(item))
	}
	cloudAsmDiskGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudAsmDiskGroupsDataSource().Schema["cloud_asm_disk_group_collection"].Elem.(*schema.Resource).Schema)
		cloudAsmDiskGroup["items"] = items
	}

	resources = append(resources, cloudAsmDiskGroup)
	if err := s.D.Set("cloud_asm_disk_group_collection", resources); err != nil {
		return err
	}

	return nil
}

func CloudAsmDiskGroupSummaryToMap(obj oci_database_management.CloudAsmDiskGroupSummary) map[string]interface{} {
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
