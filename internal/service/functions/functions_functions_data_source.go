// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v58/functions"
)

func FunctionsFunctionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFunctionsFunctions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FunctionsFunctionResource()),
			},
		},
	}
}

func readFunctionsFunctions(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsFunctionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListFunctionsResponse
}

func (s *FunctionsFunctionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsFunctionsDataSourceCrud) Get() error {
	request := oci_functions.ListFunctionsRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.FunctionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListFunctions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFunctions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FunctionsFunctionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsFunctionsDataSource-", FunctionsFunctionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		function := map[string]interface{}{
			"application_id": *r.ApplicationId,
		}

		if r.CompartmentId != nil {
			function["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			function["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			function["display_name"] = *r.DisplayName
		}

		function["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			function["id"] = *r.Id
		}

		if r.Image != nil {
			function["image"] = *r.Image
		}

		if r.ImageDigest != nil {
			function["image_digest"] = *r.ImageDigest
		}

		if r.InvokeEndpoint != nil {
			function["invoke_endpoint"] = *r.InvokeEndpoint
		}

		if r.MemoryInMBs != nil {
			function["memory_in_mbs"] = strconv.FormatInt(*r.MemoryInMBs, 10)
		}

		function["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			function["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			function["time_updated"] = r.TimeUpdated.String()
		}

		if r.TimeoutInSeconds != nil {
			function["timeout_in_seconds"] = *r.TimeoutInSeconds
		}

		if r.TraceConfig != nil {
			function["trace_config"] = []interface{}{FunctionTraceConfigToMap(r.TraceConfig)}
		} else {
			function["trace_config"] = nil
		}

		resources = append(resources, function)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FunctionsFunctionsDataSource().Schema["functions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("functions", resources); err != nil {
		return err
	}

	return nil
}
