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

func DataSafeAttributeSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAttributeSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_set_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_use": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_user_defined": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_set_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeAttributeSetResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeAttributeSets(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAttributeSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAttributeSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAttributeSetsResponse
}

func (s *DataSafeAttributeSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAttributeSetsDataSourceCrud) Get() error {
	request := oci_data_safe.ListAttributeSetsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAttributeSetsAccessLevelEnum(accessLevel.(string))
	}

	if attributeSetId, ok := s.D.GetOkExists("id"); ok {
		tmp := attributeSetId.(string)
		request.AttributeSetId = &tmp
	}

	if attributeSetType, ok := s.D.GetOkExists("attribute_set_type"); ok {
		request.AttributeSetType = oci_data_safe.AttributeSetAttributeSetTypeEnum(attributeSetType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inUse, ok := s.D.GetOkExists("in_use"); ok {
		request.InUse = oci_data_safe.ListAttributeSetsInUseEnum(inUse.(string))
	}

	if isUserDefined, ok := s.D.GetOkExists("is_user_defined"); ok {
		tmp := isUserDefined.(bool)
		request.IsUserDefined = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.AttributeSetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAttributeSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAttributeSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAttributeSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAttributeSetsDataSource-", DataSafeAttributeSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	attributeSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AttributeSetSummaryToMap(item))
	}
	attributeSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAttributeSetsDataSource().Schema["attribute_set_collection"].Elem.(*schema.Resource).Schema)
		attributeSet["items"] = items
	}

	resources = append(resources, attributeSet)
	if err := s.D.Set("attribute_set_collection", resources); err != nil {
		return err
	}

	return nil
}
