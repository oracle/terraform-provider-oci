// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveDataModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveDataModels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
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
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_data_model_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeSensitiveDataModelResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSensitiveDataModels(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveDataModelsResponse
}

func (s *DataSafeSensitiveDataModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveDataModelsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSensitiveDataModelsAccessLevelEnum(accessLevel.(string))
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

	if sensitiveDataModelId, ok := s.D.GetOkExists("id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListSensitiveDataModelsLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveDataModels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveDataModels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveDataModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelsDataSource-", DataSafeSensitiveDataModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveDataModel := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveDataModelSummaryToMap(item))
	}
	sensitiveDataModel["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveDataModelsDataSource().Schema["sensitive_data_model_collection"].Elem.(*schema.Resource).Schema)
		sensitiveDataModel["items"] = items
	}

	resources = append(resources, sensitiveDataModel)
	if err := s.D.Set("sensitive_data_model_collection", resources); err != nil {
		return err
	}

	return nil
}
