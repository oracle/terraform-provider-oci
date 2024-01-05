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

func DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSdmMaskingPolicyDifferenceDifferenceColumns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"difference_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"planned_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sdm_masking_policy_difference_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sync_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdm_masking_policy_difference_column_collection": {
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
									"column_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"difference_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"masking_columnkey": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"planned_action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensitive_columnkey": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensitive_type_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sync_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_synced": {
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

func readDataSafeSdmMaskingPolicyDifferenceDifferenceColumns(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDifferenceColumnsResponse
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSourceCrud) Get() error {
	request := oci_data_safe.ListDifferenceColumnsRequest{}

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

	if differenceType, ok := s.D.GetOkExists("difference_type"); ok {
		request.DifferenceType = oci_data_safe.SdmMaskingPolicyDifferenceDifferenceTypeEnum(differenceType.(string))
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

	if plannedAction, ok := s.D.GetOkExists("planned_action"); ok {
		request.PlannedAction = oci_data_safe.DifferenceColumnPlannedActionEnum(plannedAction.(string))
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

	if sdmMaskingPolicyDifferenceId, ok := s.D.GetOkExists("sdm_masking_policy_difference_id"); ok {
		tmp := sdmMaskingPolicyDifferenceId.(string)
		request.SdmMaskingPolicyDifferenceId = &tmp
	}

	if syncStatus, ok := s.D.GetOkExists("sync_status"); ok {
		request.SyncStatus = oci_data_safe.DifferenceColumnSyncStatusEnum(syncStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDifferenceColumns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDifferenceColumns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource-", DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sdmMaskingPolicyDifferenceDifferenceColumn := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DifferenceColumnSummaryToMap(item))
	}
	sdmMaskingPolicyDifferenceDifferenceColumn["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource().Schema["sdm_masking_policy_difference_column_collection"].Elem.(*schema.Resource).Schema)
		sdmMaskingPolicyDifferenceDifferenceColumn["items"] = items
	}

	resources = append(resources, sdmMaskingPolicyDifferenceDifferenceColumn)
	if err := s.D.Set("sdm_masking_policy_difference_column_collection", resources); err != nil {
		return err
	}

	return nil
}

func DifferenceColumnSummaryToMap(obj oci_data_safe.DifferenceColumnSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	result["difference_type"] = string(obj.DifferenceType)

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.MaskingColumnkey != nil {
		result["masking_columnkey"] = string(*obj.MaskingColumnkey)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["planned_action"] = string(obj.PlannedAction)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SensitiveColumnkey != nil {
		result["sensitive_columnkey"] = string(*obj.SensitiveColumnkey)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	result["sync_status"] = string(obj.SyncStatus)

	if obj.TimeLastSynced != nil {
		result["time_last_synced"] = obj.TimeLastSynced.String()
	}

	return result
}
