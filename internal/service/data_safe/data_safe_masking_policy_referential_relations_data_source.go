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

func DataSafeMaskingPolicyReferentialRelationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingPolicyReferentialRelations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"masking_policy_referential_relation_collection": {
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
									"child": {
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
												"referential_column_group": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"schema_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"masking_format": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"masking_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parent": {
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
												"referential_column_group": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"schema_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"relation_type": {
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

func readDataSafeMaskingPolicyReferentialRelations(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyReferentialRelationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPolicyReferentialRelationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingPolicyReferentialRelationsResponse
}

func (s *DataSafeMaskingPolicyReferentialRelationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPolicyReferentialRelationsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingPolicyReferentialRelationsRequest{}

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

	if relationType, ok := s.D.GetOkExists("relation_type"); ok {
		interfaces := relationType.([]interface{})
		tmp := make([]oci_data_safe.ListMaskingPolicyReferentialRelationsRelationTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListMaskingPolicyReferentialRelationsRelationTypeEnum(interfaces[i].(string))
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
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingPolicyReferentialRelations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingPolicyReferentialRelations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingPolicyReferentialRelationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingPolicyReferentialRelationsDataSource-", DataSafeMaskingPolicyReferentialRelationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingPolicyReferentialRelation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingPolicyReferentialRelationSummaryToMap(item))
	}
	maskingPolicyReferentialRelation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingPolicyReferentialRelationsDataSource().Schema["masking_policy_referential_relation_collection"].Elem.(*schema.Resource).Schema)
		maskingPolicyReferentialRelation["items"] = items
	}

	resources = append(resources, maskingPolicyReferentialRelation)
	if err := s.D.Set("masking_policy_referential_relation_collection", resources); err != nil {
		return err
	}

	return nil
}

func MaskingPolicyColumnsInfoToMap(obj *oci_data_safe.MaskingPolicyColumnsInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	result["referential_column_group"] = obj.ReferentialColumnGroup

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}

func MaskingPolicyReferentialRelationSummaryToMap(obj oci_data_safe.MaskingPolicyReferentialRelationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Child != nil {
		result["child"] = []interface{}{MaskingPolicyColumnsInfoToMap(obj.Child)}
	}

	result["masking_format"] = obj.MaskingFormat

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.Parent != nil {
		result["parent"] = []interface{}{MaskingPolicyColumnsInfoToMap(obj.Parent)}
	}

	result["relation_type"] = string(obj.RelationType)

	return result
}
