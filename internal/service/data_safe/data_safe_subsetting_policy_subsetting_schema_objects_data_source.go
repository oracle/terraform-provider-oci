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

func DataSafeSubsettingPolicySubsettingSchemaObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicySubsettingSchemaObjectsWithContext,
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
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_schema_object_collection": {
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
									"initial_row_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_stats_stale": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readDataSafeSubsettingPolicySubsettingSchemaObjectsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemaObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingSchemaObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingSchemaObjectsResponse
}

func (s *DataSafeSubsettingPolicySubsettingSchemaObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingSchemaObjectsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingSchemaObjectsRequest{}

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

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingSchemaObjects(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingSchemaObjects(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicySubsettingSchemaObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingSchemaObjectsDataSource-", DataSafeSubsettingPolicySubsettingSchemaObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicySubsettingSchemaObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingSchemaObjectSummaryToMap(item))
	}
	subsettingPolicySubsettingSchemaObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicySubsettingSchemaObjectsDataSource().Schema["subsetting_schema_object_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicySubsettingSchemaObject["items"] = items
	}

	resources = append(resources, subsettingPolicySubsettingSchemaObject)
	if err := s.D.Set("subsetting_schema_object_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingSchemaObjectSummaryToMap(obj oci_data_safe.SubsettingSchemaObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InitialRowCount != nil {
		result["initial_row_count"] = strconv.FormatInt(*obj.InitialRowCount, 10)
	}

	if obj.IsStatsStale != nil {
		result["is_stats_stale"] = bool(*obj.IsStatsStale)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
