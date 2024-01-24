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

func OcvpSupportedSkusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpSupportedSkus,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_ocvp_supported_skus", "oci_ocvp_supported_commitments"),
	}
}

func readOcvpSupportedSkus(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSupportedSkusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpSupportedSkusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.ListSupportedCommitmentsResponse
}

func (s *OcvpSupportedSkusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSupportedSkusDataSourceCrud) Get() error {
	request := oci_ocvp.ListSupportedCommitmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if hostShapeName, ok := s.D.GetOkExists("host_shape_name"); ok {
		tmp := hostShapeName.(string)
		request.HostShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListSupportedCommitments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSupportedCommitments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpSupportedSkusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpSupportedSkusDataSource-", OcvpSupportedSkusDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SupportedCommitmentSummaryToSkuMap(item))
	}

	if err := s.D.Set("items", items); err != nil {
		return err
	}

	return nil
}

func SupportedCommitmentSummaryToSkuMap(obj oci_ocvp.SupportedCommitmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}
