// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/datasafe"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingPoliciesMaskingColumnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingPoliciesMaskingColumns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_masking_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_seed_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"masking_column_group": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"masking_column_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sensitive_type_id": {
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
			"time_updated_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_updated_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_column_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeMaskingPoliciesMaskingColumnResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeMaskingPoliciesMaskingColumns(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPoliciesMaskingColumnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingColumnsResponse
}

func (s *DataSafeMaskingPoliciesMaskingColumnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPoliciesMaskingColumnsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingColumnsRequest{}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		interfaces := columnName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("column_name") {
			request.ColumnName = tmp
		}
	}

	if dataType, ok := s.D.GetOkExists("data_type"); ok {
		interfaces := dataType.([]interface{})
		tmp := make([]datasafe.ListMaskingColumnsDataTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = datasafe.ListMaskingColumnsDataTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("data_type") {
			request.DataType = tmp
		}
	}

	if isMaskingEnabled, ok := s.D.GetOkExists("is_masking_enabled"); ok {
		tmp := isMaskingEnabled.(bool)
		request.IsMaskingEnabled = &tmp
	}

	if isSeedRequired, ok := s.D.GetOkExists("is_seed_required"); ok {
		tmp := isSeedRequired.(bool)
		request.IsSeedRequired = &tmp
	}

	if maskingColumnGroup, ok := s.D.GetOkExists("masking_column_group"); ok {
		interfaces := maskingColumnGroup.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("masking_column_group") {
			request.MaskingColumnGroup = tmp
		}
	}

	if maskingColumnLifecycleState, ok := s.D.GetOkExists("masking_column_lifecycle_state"); ok {
		request.MaskingColumnLifecycleState = oci_data_safe.ListMaskingColumnsMaskingColumnLifecycleStateEnum(maskingColumnLifecycleState.(string))
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		interfaces := object.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object") {
			request.ObjectName = tmp
		}
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		interfaces := objectType.([]interface{})
		tmp := make([]datasafe.ListMaskingColumnsObjectTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = datasafe.ListMaskingColumnsObjectTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object_type") {
			request.ObjectType = tmp
		}
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
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

	if timeUpdatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_updated_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdatedLessThan, ok := s.D.GetOkExists("time_updated_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingColumns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingColumns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingPoliciesMaskingColumnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingPoliciesMaskingColumnsDataSource-", DataSafeMaskingPoliciesMaskingColumnsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingPoliciesMaskingColumn := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingColumnSummaryToMap(item))
	}
	maskingPoliciesMaskingColumn["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingPoliciesMaskingColumnsDataSource().Schema["masking_column_collection"].Elem.(*schema.Resource).Schema)
		maskingPoliciesMaskingColumn["items"] = items
	}

	resources = append(resources, maskingPoliciesMaskingColumn)
	if err := s.D.Set("masking_column_collection", resources); err != nil {
		return err
	}

	return nil
}
