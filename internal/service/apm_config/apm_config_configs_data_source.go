// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_config "github.com/oracle/oci-go-sdk/v56/apmconfig"
)

func ApmConfigConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmConfigConfigs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApmConfigConfigResource()),
						},
					},
				},
			},
		},
	}
}

func readApmConfigConfigs(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.ReadResource(sync)
}

type ApmConfigConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_config.ConfigClient
	Res    *oci_apm_config.ListConfigsResponse
}

func (s *ApmConfigConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmConfigConfigsDataSourceCrud) Get() error {
	request := oci_apm_config.ListConfigsRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if configType, ok := s.D.GetOkExists("config_type"); ok {
		tmp := configType.(string)
		request.ConfigType = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_config")

	response, err := s.Client.ListConfigs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConfigs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmConfigConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmConfigConfigsDataSource-", ApmConfigConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}
	config := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConfigSummaryToMap(item))
	}
	config["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmConfigConfigsDataSource().Schema["config_collection"].Elem.(*schema.Resource).Schema)
		config["items"] = items
	}

	resources = append(resources, config)
	if err := s.D.Set("config_collection", resources); err != nil {
		return err
	}

	return nil
}
