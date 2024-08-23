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

func OcvpSupportedVmwareSoftwareVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpSupportedVmwareSoftwareVersions,
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
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version_to_upgrade": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"esxi_software_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_host_shape_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOcvpSupportedVmwareSoftwareVersions(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSupportedVmwareSoftwareVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpSupportedVmwareSoftwareVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.ListSupportedVmwareSoftwareVersionsResponse
}

func (s *OcvpSupportedVmwareSoftwareVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSupportedVmwareSoftwareVersionsDataSourceCrud) Get() error {
	request := oci_ocvp.ListSupportedVmwareSoftwareVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if hostShapeName, ok := s.D.GetOkExists("host_shape_name"); ok {
		tmp := hostShapeName.(string)
		request.HostShapeName = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	if versionToUpgrade, ok := s.D.GetOkExists("version_to_upgrade"); ok {
		tmp := versionToUpgrade.(string)
		request.VersionToUpgrade = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListSupportedVmwareSoftwareVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpSupportedVmwareSoftwareVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpSupportedVmwareSoftwareVersionsDataSource-", OcvpSupportedVmwareSoftwareVersionsDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SupportedVmwareSoftwareVersionSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func SupportedEsxiSoftwareVersionSummaryToMap(obj oci_ocvp.SupportedEsxiSoftwareVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["supported_host_shape_names"] = obj.SupportedHostShapeNames

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func SupportedVmwareSoftwareVersionSummaryToMap(obj oci_ocvp.SupportedVmwareSoftwareVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	esxiSoftwareVersions := []interface{}{}
	for _, item := range obj.EsxiSoftwareVersions {
		esxiSoftwareVersions = append(esxiSoftwareVersions, SupportedEsxiSoftwareVersionSummaryToMap(item))
	}
	result["esxi_software_versions"] = esxiSoftwareVersions

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
