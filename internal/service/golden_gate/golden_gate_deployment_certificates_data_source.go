// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeploymentCertificates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GoldenGateDeploymentCertificateResource(),
						},
					},
				},
			},
		},
	}
}

func readGoldenGateDeploymentCertificates(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentCertificatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListCertificatesResponse
}

func (s *GoldenGateDeploymentCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentCertificatesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListCertificatesRequest{}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_golden_gate.ListCertificatesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentCertificatesDataSource-", GoldenGateDeploymentCertificatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	deploymentCertificate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificateSummaryToMap(item))
	}
	deploymentCertificate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentCertificatesDataSource().Schema["certificate_collection"].Elem.(*schema.Resource).Schema)
		deploymentCertificate["items"] = items
	}

	resources = append(resources, deploymentCertificate)
	if err := s.D.Set("certificate_collection", resources); err != nil {
		return err
	}

	return nil
}
