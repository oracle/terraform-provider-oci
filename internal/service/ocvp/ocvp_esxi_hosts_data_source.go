// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v58/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OcvpEsxiHostsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpEsxiHosts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"sddc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"esxi_host_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OcvpEsxiHostResource()),
			},
		},
	}
}

func readOcvpEsxiHosts(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()

	return tfresource.ReadResource(sync)
}

type OcvpEsxiHostsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.EsxiHostClient
	Res    *oci_ocvp.ListEsxiHostsResponse
}

func (s *OcvpEsxiHostsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpEsxiHostsDataSourceCrud) Get() error {
	request := oci_ocvp.ListEsxiHostsRequest{}

	if computeInstanceId, ok := s.D.GetOkExists("compute_instance_id"); ok {
		tmp := computeInstanceId.(string)
		request.ComputeInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ListEsxiHostsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListEsxiHosts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpEsxiHostsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpEsxiHostsDataSource-", OcvpEsxiHostsDataSource(), s.D))

	resources := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		resources = append(resources, EsxiHostSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OcvpEsxiHostsDataSource().Schema["esxi_host_collection"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("esxi_host_collection", resources)

	return nil
}
