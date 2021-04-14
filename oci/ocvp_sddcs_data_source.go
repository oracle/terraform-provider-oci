// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v39/ocvp"
)

func init() {
	RegisterDatasource("oci_ocvp_sddcs", OcvpSddcsDataSource())
}

func OcvpSddcsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpSddcs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sddc_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(OcvpSddcResource()),
			},
		},
	}
}

func readOcvpSddcs(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()

	return ReadResource(sync)
}

type OcvpSddcsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.ListSddcsResponse
}

func (s *OcvpSddcsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSddcsDataSourceCrud) Get() error {
	request := oci_ocvp.ListSddcsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeAvailabilityDomain, ok := s.D.GetOkExists("compute_availability_domain"); ok {
		tmp := computeAvailabilityDomain.(string)
		request.ComputeAvailabilityDomain = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ListSddcsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ocvp")

	response, err := s.Client.ListSddcs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpSddcsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("OcvpSddcsDataSource-", OcvpSddcsDataSource(), s.D))

	resources := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		resources = append(resources, SddcSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, OcvpSddcsDataSource().Schema["sddc_collection"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("sddc_collection", resources); err != nil {
		return err
	}

	return nil
}
