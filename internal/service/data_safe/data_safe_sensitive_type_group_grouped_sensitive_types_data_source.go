// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveTypeGroupGroupedSensitiveTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"sensitive_type_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"grouped_sensitive_type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeSensitiveTypeGroupGroupedSensitiveTypeResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSensitiveTypeGroupGroupedSensitiveTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListGroupedSensitiveTypesResponse
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSourceCrud) Get() error {
	request := oci_data_safe.ListGroupedSensitiveTypesRequest{}

	if sensitiveTypeGroupId, ok := s.D.GetOkExists("sensitive_type_group_id"); ok {
		tmp := sensitiveTypeGroupId.(string)
		request.SensitiveTypeGroupId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListGroupedSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGroupedSensitiveTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSource-", DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveTypeGroupGroupedSensitiveType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, GroupedSensitiveTypeSummaryToMap(item))
	}
	sensitiveTypeGroupGroupedSensitiveType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSource().Schema["grouped_sensitive_type_collection"].Elem.(*schema.Resource).Schema)
		sensitiveTypeGroupGroupedSensitiveType["items"] = items
	}

	resources = append(resources, sensitiveTypeGroupGroupedSensitiveType)
	if err := s.D.Set("grouped_sensitive_type_collection", resources); err != nil {
		return err
	}

	return nil
}
