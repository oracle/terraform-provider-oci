// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_masking_format_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_common": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parent_category_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeSensitiveTypeResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSensitiveTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveTypesResponse
}

func (s *DataSafeSensitiveTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveTypesDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveTypesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSensitiveTypesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if defaultMaskingFormatId, ok := s.D.GetOkExists("default_masking_format_id"); ok {
		tmp := defaultMaskingFormatId.(string)
		request.DefaultMaskingFormatId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		request.EntityType = oci_data_safe.ListSensitiveTypesEntityTypeEnum(entityType.(string))
	}

	if isCommon, ok := s.D.GetOkExists("is_common"); ok {
		tmp := isCommon.(bool)
		request.IsCommon = &tmp
	}

	if parentCategoryId, ok := s.D.GetOkExists("parent_category_id"); ok {
		tmp := parentCategoryId.(string)
		request.ParentCategoryId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	if sensitiveTypeSource, ok := s.D.GetOkExists("sensitive_type_source"); ok {
		request.SensitiveTypeSource = oci_data_safe.ListSensitiveTypesSensitiveTypeSourceEnum(sensitiveTypeSource.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListSensitiveTypesLifecycleStateEnum(state.(string))
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveTypesDataSource-", DataSafeSensitiveTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveTypeSummaryToMap(item))
	}
	sensitiveType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveTypesDataSource().Schema["sensitive_type_collection"].Elem.(*schema.Resource).Schema)
		sensitiveType["items"] = items
	}

	resources = append(resources, sensitiveType)
	if err := s.D.Set("sensitive_type_collection", resources); err != nil {
		return err
	}

	return nil
}
