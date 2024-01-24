// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerLicenseMetricDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLicenseManagerLicenseMetric,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"license_record_expiring_soon_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_byol_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_license_included_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_product_license_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularLicenseManagerLicenseMetric(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseMetricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerLicenseMetricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.GetLicenseMetricResponse
}

func (s *LicenseManagerLicenseMetricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerLicenseMetricDataSourceCrud) Get() error {
	request := oci_license_manager.GetLicenseMetricRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.GetLicenseMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LicenseManagerLicenseMetricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerLicenseMetricDataSource-", LicenseManagerLicenseMetricDataSource(), s.D))

	if s.Res.LicenseRecordExpiringSoonCount != nil {
		s.D.Set("license_record_expiring_soon_count", *s.Res.LicenseRecordExpiringSoonCount)
	}

	if s.Res.TotalByolInstanceCount != nil {
		s.D.Set("total_byol_instance_count", *s.Res.TotalByolInstanceCount)
	}

	if s.Res.TotalLicenseIncludedInstanceCount != nil {
		s.D.Set("total_license_included_instance_count", *s.Res.TotalLicenseIncludedInstanceCount)
	}

	if s.Res.TotalProductLicenseCount != nil {
		s.D.Set("total_product_license_count", *s.Res.TotalProductLicenseCount)
	}

	return nil
}
