// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v48/apmsynthetics"
)

func init() {
	RegisterDatasource("oci_apm_synthetics_scripts", ApmSyntheticsScriptsDataSource())
}

func ApmSyntheticsScriptsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmSyntheticsScripts,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"script_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(ApmSyntheticsScriptResource()),
						},
					},
				},
			},
		},
	}
}

func readApmSyntheticsScripts(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).apmSyntheticClient()

	return ReadResource(sync)
}

type ApmSyntheticsScriptsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListScriptsResponse
}

func (s *ApmSyntheticsScriptsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsScriptsDataSourceCrud) Get() error {
	request := oci_apm_synthetics.ListScriptsRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		tmp := contentType.(string)
		request.ContentType = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.ListScripts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScripts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmSyntheticsScriptsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ApmSyntheticsScriptsDataSource-", ApmSyntheticsScriptsDataSource(), s.D))
	resources := []map[string]interface{}{}
	script := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScriptSummaryToMap(item))
	}
	script["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, ApmSyntheticsScriptsDataSource().Schema["script_collection"].Elem.(*schema.Resource).Schema)
		script["items"] = items
	}

	resources = append(resources, script)
	if err := s.D.Set("script_collection", resources); err != nil {
		return err
	}

	return nil
}
