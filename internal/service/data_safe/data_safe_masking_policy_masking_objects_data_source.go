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

func DataSafeMaskingPolicyMaskingObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingPolicyMaskingObjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"masking_object_collection": {
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

func readDataSafeMaskingPolicyMaskingObjects(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyMaskingObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPolicyMaskingObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingObjectsResponse
}

func (s *DataSafeMaskingPolicyMaskingObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPolicyMaskingObjectsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingObjectsRequest{}

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
		tmp := make([]oci_data_safe.ListMaskingObjectsObjectTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListMaskingObjectsObjectTypeEnum)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingObjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingObjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingPolicyMaskingObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingPolicyMaskingObjectsDataSource-", DataSafeMaskingPolicyMaskingObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingPolicyMaskingObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingObjectSummaryToMap(item))
	}
	maskingPolicyMaskingObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingPolicyMaskingObjectsDataSource().Schema["masking_object_collection"].Elem.(*schema.Resource).Schema)
		maskingPolicyMaskingObject["items"] = items
	}

	resources = append(resources, maskingPolicyMaskingObject)
	if err := s.D.Set("masking_object_collection", resources); err != nil {
		return err
	}

	return nil
}

func MaskingObjectSummaryToMap(obj oci_data_safe.MaskingObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}
