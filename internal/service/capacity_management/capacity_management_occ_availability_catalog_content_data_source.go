// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccAvailabilityCatalogContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCapacityManagementOccAvailabilityCatalogContent,
		Schema: map[string]*schema.Schema{
			"occ_availability_catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularCapacityManagementOccAvailabilityCatalogContent(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccAvailabilityCatalogContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.GetOccAvailabilityCatalogContentResponse
}

func (s *CapacityManagementOccAvailabilityCatalogContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccAvailabilityCatalogContentDataSourceCrud) Get() error {
	request := oci_capacity_management.GetOccAvailabilityCatalogContentRequest{}

	if occAvailabilityCatalogId, ok := s.D.GetOkExists("occ_availability_catalog_id"); ok {
		tmp := occAvailabilityCatalogId.(string)
		request.OccAvailabilityCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetOccAvailabilityCatalogContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccAvailabilityCatalogContentDataSource-", CapacityManagementOccAvailabilityCatalogContentDataSource(), s.D))

	return nil
}
