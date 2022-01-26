// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm "github.com/oracle/oci-go-sdk/v56/apmcontrolplane"
)

func ApmDataKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmDataKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_key_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_keys": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readApmDataKeys(d *schema.ResourceData, m interface{}) error {
	sync := &ApmDataKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmDomainClient()

	return tfresource.ReadResource(sync)
}

type ApmDataKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm.ApmDomainClient
	Res    *oci_apm.ListDataKeysResponse
}

func (s *ApmDataKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmDataKeysDataSourceCrud) Get() error {
	request := oci_apm.ListDataKeysRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if dataKeyType, ok := s.D.GetOkExists("data_key_type"); ok {
		request.DataKeyType = oci_apm.ListDataKeysDataKeyTypeEnum(dataKeyType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm")

	response, err := s.Client.ListDataKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmDataKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmDataKeysDataSource-", ApmDataKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dataKey := map[string]interface{}{}

		if r.Name != nil {
			dataKey["name"] = *r.Name
		}

		dataKey["type"] = r.Type

		if r.Value != nil {
			dataKey["value"] = *r.Value
		}

		resources = append(resources, dataKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ApmDataKeysDataSource().Schema["data_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("data_keys", resources); err != nil {
		return err
	}

	return nil
}
