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

func JmsFleetDrsFilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetDrsFiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drs_file_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"bucket": {
										Type:     schema.TypeString,
										Required: true,
									},
									"drs_file_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"fleet_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"checksum_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"checksum_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"drs_file_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
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

func readJmsFleetDrsFiles(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetDrsFilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetDrsFilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListDrsFilesResponse
}

func (s *JmsFleetDrsFilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetDrsFilesDataSourceCrud) Get() error {
	request := oci_jms.ListDrsFilesRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListDrsFiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrsFiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetDrsFilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetDrsFilesDataSource-", JmsFleetDrsFilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetDrsFile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DrsFileSummaryToMap(item))
	}
	fleetDrsFile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetDrsFilesDataSource().Schema["drs_file_collection"].Elem.(*schema.Resource).Schema)
		fleetDrsFile["items"] = items
	}

	resources = append(resources, fleetDrsFile)
	if err := s.D.Set("drs_file_collection", resources); err != nil {
		return err
	}

	return nil
}

func DrsFileSummaryToMap(obj oci_jms.DrsFileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	result["checksum_type"] = string(obj.ChecksumType)

	if obj.ChecksumValue != nil {
		result["checksum_value"] = string(*obj.ChecksumValue)
	}

	if obj.DrsFileKey != nil {
		result["drs_file_key"] = string(*obj.DrsFileKey)
	}

	if obj.DrsFileName != nil {
		result["drs_file_name"] = string(*obj.DrsFileName)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	return result
}
