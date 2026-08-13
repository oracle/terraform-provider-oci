// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicySubsettingSchemasDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicySubsettingSchemasWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_derived_schema": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"subsetting_schema_collection": {
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
									"is_derived": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"schema_name": {
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

func readDataSafeSubsettingPolicySubsettingSchemasWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingSchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingSchemasResponse
}

func (s *DataSafeSubsettingPolicySubsettingSchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingSchemasDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingSchemasRequest{}

	if isDerivedSchema, ok := s.D.GetOkExists("is_derived_schema"); ok {
		tmp := isDerivedSchema.(bool)
		request.IsDerivedSchema = &tmp
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

	response, err := s.Client.ListSubsettingSchemas(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingSchemas(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicySubsettingSchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingSchemasDataSource-", DataSafeSubsettingPolicySubsettingSchemasDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicySubsettingSchema := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingSchemaSummaryToMap(item))
	}
	subsettingPolicySubsettingSchema["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicySubsettingSchemasDataSource().Schema["subsetting_schema_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicySubsettingSchema["items"] = items
	}

	resources = append(resources, subsettingPolicySubsettingSchema)
	if err := s.D.Set("subsetting_schema_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingSchemaSummaryToMap(obj oci_data_safe.SubsettingSchemaSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDerived != nil {
		result["is_derived"] = bool(*obj.IsDerived)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}
