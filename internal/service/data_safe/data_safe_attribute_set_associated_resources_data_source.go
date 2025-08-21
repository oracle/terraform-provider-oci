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

func DataSafeAttributeSetAssociatedResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAttributeSetAssociatedResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"associated_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_set_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"associated_resource_collection": {
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
									"associated_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"associated_resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"associated_resource_type": {
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

func readDataSafeAttributeSetAssociatedResources(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAttributeSetAssociatedResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAttributeSetAssociatedResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAssociatedResourcesResponse
}

func (s *DataSafeAttributeSetAssociatedResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAttributeSetAssociatedResourcesDataSourceCrud) Get() error {
	request := oci_data_safe.ListAssociatedResourcesRequest{}

	if associatedResourceId, ok := s.D.GetOkExists("associated_resource_id"); ok {
		tmp := associatedResourceId.(string)
		request.AssociatedResourceId = &tmp
	}

	if associatedResourceType, ok := s.D.GetOkExists("associated_resource_type"); ok {
		request.AssociatedResourceType = oci_data_safe.AssociatedResourceSummaryAssociatedResourceTypeEnum(associatedResourceType.(string))
	}

	if attributeSetId, ok := s.D.GetOkExists("attribute_set_id"); ok {
		tmp := attributeSetId.(string)
		request.AttributeSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAssociatedResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssociatedResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAttributeSetAssociatedResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAttributeSetAssociatedResourcesDataSource-", DataSafeAttributeSetAssociatedResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	attributeSetAssociatedResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssociatedResourceSummaryToMap(item))
	}
	attributeSetAssociatedResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAttributeSetAssociatedResourcesDataSource().Schema["associated_resource_collection"].Elem.(*schema.Resource).Schema)
		attributeSetAssociatedResource["items"] = items
	}

	resources = append(resources, attributeSetAssociatedResource)
	if err := s.D.Set("associated_resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssociatedResourceSummaryToMap(obj oci_data_safe.AssociatedResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociatedResourceId != nil {
		result["associated_resource_id"] = string(*obj.AssociatedResourceId)
	}

	if obj.AssociatedResourceName != nil {
		result["associated_resource_name"] = string(*obj.AssociatedResourceName)
	}

	result["associated_resource_type"] = string(obj.AssociatedResourceType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
