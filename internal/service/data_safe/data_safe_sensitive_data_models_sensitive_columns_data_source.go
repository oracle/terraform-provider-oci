// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeSensitiveDataModelsSensitiveColumnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveDataModelsSensitiveColumns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"is_case_in_sensitive": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"parent_column_key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"relation_type": {
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
			"sensitive_column_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"sensitive_column_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeSensitiveDataModelsSensitiveColumnResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSensitiveDataModelsSensitiveColumns(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelsSensitiveColumnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveColumnsResponse
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveColumnsRequest{}

	if columnGroup, ok := s.D.GetOkExists("column_group"); ok {
		tmp := columnGroup.(string)
		request.ColumnGroup = &tmp
	}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		interfaces := columnName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.ColumnName = tmp
	}

	if dataType, ok := s.D.GetOkExists("data_type"); ok {
		interfaces := dataType.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.DataType = tmp
	}

	if isCaseInSensitive, ok := s.D.GetOkExists("is_case_in_sensitive"); ok {
		tmp := isCaseInSensitive.(bool)
		request.IsCaseInSensitive = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		interfaces := object.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.ObjectName = tmp
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		interfaces := objectType.([]interface{})
		tmp := make([]oci_data_safe.ListSensitiveColumnsObjectTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListSensitiveColumnsObjectTypeEnum)
			}
		}
		request.ObjectType = tmp
	}

	if parentColumnKey, ok := s.D.GetOkExists("parent_column_key"); ok {
		interfaces := parentColumnKey.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("parent_column_key") {
			request.ParentColumnKey = tmp
		}
	}

	if relationType, ok := s.D.GetOkExists("relation_type"); ok {
		interfaces := relationType.([]interface{})
		tmp := make([]oci_data_safe.ListSensitiveColumnsRelationTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListSensitiveColumnsRelationTypeEnum)
			}
		}
		request.RelationType = tmp
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.SchemaName = tmp
	}

	if sensitiveColumnLifecycleState, ok := s.D.GetOkExists("sensitive_column_lifecycle_state"); ok {
		request.SensitiveColumnLifecycleState = oci_data_safe.ListSensitiveColumnsSensitiveColumnLifecycleStateEnum(sensitiveColumnLifecycleState.(string))
	}
	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		interfaces := sensitiveTypeId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.SensitiveTypeId = tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]oci_data_safe.ListSensitiveColumnsStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListSensitiveColumnsStatusEnum)
			}
		}
		request.Status = tmp
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

	response, err := s.Client.ListSensitiveColumns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveColumns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelsSensitiveColumnsDataSource-", DataSafeSensitiveDataModelsSensitiveColumnsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveDataModelsSensitiveColumn := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveColumnSummaryToMap(item))
	}
	sensitiveDataModelsSensitiveColumn["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveDataModelsSensitiveColumnsDataSource().Schema["sensitive_column_collection"].Elem.(*schema.Resource).Schema)
		sensitiveDataModelsSensitiveColumn["items"] = items
	}

	resources = append(resources, sensitiveDataModelsSensitiveColumn)
	if err := s.D.Set("sensitive_column_collection", resources); err != nil {
		return err
	}

	return nil
}
