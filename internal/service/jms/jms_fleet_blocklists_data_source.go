// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetBlocklistsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetBlocklists,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation": {
				Type:     schema.TypeString,
				Optional: true,
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
						"operation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"fleet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"installation_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
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

func readJmsFleetBlocklists(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetBlocklistsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetBlocklistsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListBlocklistsResponse
}

func (s *JmsFleetBlocklistsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetBlocklistsDataSourceCrud) Get() error {
	request := oci_jms.ListBlocklistsRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if operation, ok := s.D.GetOkExists("operation"); ok {
		request.Operation = oci_jms.ListBlocklistsOperationEnum(operation.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListBlocklists(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetBlocklistsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetBlocklistsDataSource-", JmsFleetBlocklistsDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BlocklistToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func BlocklistToMap(obj oci_jms.Blocklist) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["operation"] = string(obj.Operation)

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	if obj.Target != nil {
		result["target"] = []interface{}{BlocklistTargetToMap(obj.Target)}
	}

	return result
}

func BlocklistTargetToMap(obj *oci_jms.BlocklistTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	if obj.InstallationKey != nil {
		result["installation_key"] = string(*obj.InstallationKey)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	return result
}
