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

func DataSafeSensitiveDataModelSensitiveObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveDataModelSensitiveObjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_object_collection": {
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

func readDataSafeSensitiveDataModelSensitiveObjects(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelSensitiveObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelSensitiveObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveObjectsResponse
}

func (s *DataSafeSensitiveDataModelSensitiveObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelSensitiveObjectsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveObjectsRequest{}

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
		tmp := make([]oci_data_safe.ListSensitiveObjectsObjectTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListSensitiveObjectsObjectTypeEnum)
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

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveObjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveObjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveDataModelSensitiveObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelSensitiveObjectsDataSource-", DataSafeSensitiveDataModelSensitiveObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveDataModelSensitiveObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveObjectSummaryToMap(item))
	}
	sensitiveDataModelSensitiveObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveDataModelSensitiveObjectsDataSource().Schema["sensitive_object_collection"].Elem.(*schema.Resource).Schema)
		sensitiveDataModelSensitiveObject["items"] = items
	}

	resources = append(resources, sensitiveDataModelSensitiveObject)
	if err := s.D.Set("sensitive_object_collection", resources); err != nil {
		return err
	}

	return nil
}

func SensitiveObjectSummaryToMap(obj oci_data_safe.SensitiveObjectSummary) map[string]interface{} {
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
