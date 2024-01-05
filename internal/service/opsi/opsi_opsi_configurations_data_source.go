// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOpsiConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiOpsiConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opsi_config_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"opsi_configurations_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiOpsiConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiOpsiConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOpsiConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListOpsiConfigurationsResponse
}

func (s *OpsiOpsiConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOpsiConfigurationsDataSourceCrud) Get() error {
	request := oci_opsi.ListOpsiConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if opsiConfigType, ok := s.D.GetOkExists("opsi_config_type"); ok {
		interfaces := opsiConfigType.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.OpsiConfigurationTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingOpsiConfigurationTypeEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp2) != 0 {
			request.OpsiConfigType = tmp2
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.OpsiConfigurationLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingOpsiConfigurationLifecycleStateEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp2) != 0 {
			request.LifecycleState = tmp2
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListOpsiConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOpsiConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiOpsiConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiOpsiConfigurationsDataSource-", OpsiOpsiConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	opsiConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OpsiConfigurationSummaryToMap(item))
	}
	opsiConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiOpsiConfigurationsDataSource().Schema["opsi_configurations_collection"].Elem.(*schema.Resource).Schema)
		opsiConfiguration["items"] = items
	}

	resources = append(resources, opsiConfiguration)
	if err := s.D.Set("opsi_configurations_collection", resources); err != nil {
		return err
	}

	return nil
}
