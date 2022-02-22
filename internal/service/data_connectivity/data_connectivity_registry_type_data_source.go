// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v58/dataconnectivity"
)

func DataConnectivityRegistryTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataConnectivityRegistryType,
		Schema: map[string]*schema.Schema{
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"registry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"connection_attributes": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"data_asset_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"attribute_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_base64encoded": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_generated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_mandatory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_sensitive": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_key_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readSingularDataConnectivityRegistryType(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.GetTypeResponse
}

func (s *DataConnectivityRegistryTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryTypeDataSourceCrud) Get() error {
	request := oci_data_connectivity.GetTypeRequest{}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if typeKey, ok := s.D.GetOkExists("type_key"); ok {
		tmp := typeKey.(string)
		request.TypeKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.GetType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataConnectivityRegistryTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryTypeDataSource-", DataConnectivityRegistryTypeDataSource(), s.D))

	s.D.Set("connection_attributes", s.Res.ConnectionAttributes)

	dataAssetAttributes := []interface{}{}
	for _, item := range s.Res.DataAssetAttributes {
		dataAssetAttributes = append(dataAssetAttributes, DataConnectivityAttributeToMap(item))
	}
	s.D.Set("data_asset_attributes", dataAssetAttributes)

	return nil
}

func DataConnectivityAttributeToMap(obj oci_data_connectivity.Attribute) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttributeType != nil {
		result["attribute_type"] = string(*obj.AttributeType)
	}

	if obj.IsBase64Encoded != nil {
		result["is_base64encoded"] = bool(*obj.IsBase64Encoded)
	}

	if obj.IsGenerated != nil {
		result["is_generated"] = bool(*obj.IsGenerated)
	}

	if obj.IsMandatory != nil {
		result["is_mandatory"] = bool(*obj.IsMandatory)
	}

	if obj.IsSensitive != nil {
		result["is_sensitive"] = bool(*obj.IsSensitive)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["valid_key_list"] = obj.ValidKeyList

	return result
}

func DataConnectivityTypeSummaryToMap(obj oci_data_connectivity.TypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
