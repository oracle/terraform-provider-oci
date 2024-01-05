// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FunctionsPbfListingTriggersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFunctionsPbfListingTriggers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"triggers_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
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

func readFunctionsPbfListingTriggers(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsPbfListingTriggersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsPbfListingTriggersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListTriggersResponse
}

func (s *FunctionsPbfListingTriggersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsPbfListingTriggersDataSourceCrud) Get() error {
	request := oci_functions.ListTriggersRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListTriggers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTriggers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FunctionsPbfListingTriggersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsPbfListingTriggersDataSource-", FunctionsPbfListingTriggersDataSource(), s.D))
	resources := []map[string]interface{}{}
	pbfListingTrigger := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TriggerSummaryToMap(item))
	}
	pbfListingTrigger["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FunctionsPbfListingTriggersDataSource().Schema["triggers_collection"].Elem.(*schema.Resource).Schema)
		pbfListingTrigger["items"] = items
	}

	resources = append(resources, pbfListingTrigger)
	if err := s.D.Set("triggers_collection", resources); err != nil {
		return err
	}

	return nil
}

func TriggerSummaryToMap(obj oci_functions.TriggerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
