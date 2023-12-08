// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationUsageCarbonEmissionsConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMeteringComputationUsageCarbonEmissionsConfig,
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"values": {
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
	}
}

func readSingularMeteringComputationUsageCarbonEmissionsConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationUsageCarbonEmissionsConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.RequestUsageCarbonEmissionConfigResponse
}

func (s *MeteringComputationUsageCarbonEmissionsConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationUsageCarbonEmissionsConfigDataSourceCrud) Get() error {
	request := oci_metering_computation.RequestUsageCarbonEmissionConfigRequest{}

	if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
		tmp := tenantId.(string)
		request.TenantId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.RequestUsageCarbonEmissionConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationUsageCarbonEmissionsConfigDataSource-", MeteringComputationUsageCarbonEmissionsConfigDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConfigurationToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
