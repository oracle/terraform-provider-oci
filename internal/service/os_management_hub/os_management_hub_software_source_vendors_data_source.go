// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceVendorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSourceVendors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_vendor_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"arch_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_families": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubSoftwareSourceVendors(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceVendorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceVendorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListSoftwareSourceVendorsResponse
}

func (s *OsManagementHubSoftwareSourceVendorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceVendorsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListSoftwareSourceVendorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListSoftwareSourceVendors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubSoftwareSourceVendorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceVendorsDataSource-", OsManagementHubSoftwareSourceVendorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSourceVendor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SoftwareSourceVendorSummaryToMap(item))
	}
	softwareSourceVendor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourceVendorsDataSource().Schema["software_source_vendor_collection"].Elem.(*schema.Resource).Schema)
		softwareSourceVendor["items"] = items
	}

	resources = append(resources, softwareSourceVendor)
	if err := s.D.Set("software_source_vendor_collection", resources); err != nil {
		return err
	}

	return nil
}

func SoftwareSourceVendorSummaryToMap(obj oci_os_management_hub.SoftwareSourceVendorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_types"] = obj.ArchTypes

	result["name"] = string(obj.Name)

	result["os_families"] = obj.OsFamilies

	return result
}
