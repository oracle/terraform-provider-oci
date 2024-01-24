// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpEsxiHostsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpEsxiHosts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sddc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_billing_donors_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_swap_billing_only": {
				Type:     schema.TypeBool,
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

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeInstanceId, ok := s.D.GetOkExists("compute_instance_id"); ok {
		tmp := computeInstanceId.(string)
		request.ComputeInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isBillingDonorsOnly, ok := s.D.GetOkExists("is_billing_donors_only"); ok {
		tmp := isBillingDonorsOnly.(bool)
		request.IsBillingDonorsOnly = &tmp
	}

	if isSwapBillingOnly, ok := s.D.GetOkExists("is_swap_billing_only"); ok {
		tmp := isSwapBillingOnly.(bool)
		request.IsSwapBillingOnly = &tmp
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
