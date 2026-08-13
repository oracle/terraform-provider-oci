// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingReportSubsettedObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingReportSubsettedObjectsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"object": {
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
			"subsetting_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetted_object_collection": {
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
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"row_count_after_subsetting": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"row_count_before_subsetting": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_after_subsetting_in_kbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_before_subsetting_in_kbs": {
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

func readDataSafeSubsettingReportSubsettedObjectsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportSubsettedObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingReportSubsettedObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettedObjectsResponse
}

func (s *DataSafeSubsettingReportSubsettedObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingReportSubsettedObjectsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettedObjectsRequest{}

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

	if subsettingReportId, ok := s.D.GetOkExists("subsetting_report_id"); ok {
		tmp := subsettingReportId.(string)
		request.SubsettingReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettedObjects(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettedObjects(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingReportSubsettedObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingReportSubsettedObjectsDataSource-", DataSafeSubsettingReportSubsettedObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingReportSubsettedObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettedObjectSummaryToMap(item))
	}
	subsettingReportSubsettedObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingReportSubsettedObjectsDataSource().Schema["subsetted_object_collection"].Elem.(*schema.Resource).Schema)
		subsettingReportSubsettedObject["items"] = items
	}

	resources = append(resources, subsettingReportSubsettedObject)
	if err := s.D.Set("subsetted_object_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettedObjectSummaryToMap(obj oci_data_safe.SubsettedObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	if obj.RowCountAfterSubsetting != nil {
		result["row_count_after_subsetting"] = strconv.FormatInt(*obj.RowCountAfterSubsetting, 10)
	}

	if obj.RowCountBeforeSubsetting != nil {
		result["row_count_before_subsetting"] = strconv.FormatInt(*obj.RowCountBeforeSubsetting, 10)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SizeAfterSubsettingInKBs != nil {
		result["size_after_subsetting_in_kbs"] = string(*obj.SizeAfterSubsettingInKBs)
	}

	if obj.SizeBeforeSubsettingInKBs != nil {
		result["size_before_subsetting_in_kbs"] = string(*obj.SizeBeforeSubsettingInKBs)
	}

	return result
}
