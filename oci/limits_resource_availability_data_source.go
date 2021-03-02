// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v36/limits"
)

func init() {
	RegisterDatasource("oci_limits_resource_availability", LimitsResourceAvailabilityDataSource())
}

func LimitsResourceAvailabilityDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLimitsResourceAvailability,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"available": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularLimitsResourceAvailability(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsResourceAvailabilityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).limitsClient()

	return ReadResource(sync)
}

type LimitsResourceAvailabilityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.LimitsClient
	Res    *oci_limits.GetResourceAvailabilityResponse
}

func (s *LimitsResourceAvailabilityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsResourceAvailabilityDataSourceCrud) Get() error {
	request := oci_limits.GetResourceAvailabilityRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limitName, ok := s.D.GetOkExists("limit_name"); ok {
		tmp := limitName.(string)
		request.LimitName = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "limits")

	response, err := s.Client.GetResourceAvailability(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LimitsResourceAvailabilityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LimitsResourceAvailabilityDataSource-", LimitsResourceAvailabilityDataSource(), s.D))

	if s.Res.Available != nil {
		s.D.Set("available", strconv.FormatInt(*s.Res.Available, 10))
	}

	if s.Res.Used != nil {
		s.D.Set("used", strconv.FormatInt(*s.Res.Used, 10))
	}

	return nil
}
