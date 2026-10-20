// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpRetrieveVmwareBinariesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpRetrieveVmwareBinaries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"sddc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"checksum": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOcvpRetrieveVmwareBinaries(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpRetrieveVmwareBinariesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpRetrieveVmwareBinariesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.RetrieveVmwareBinariesResponse
}

func (s *OcvpRetrieveVmwareBinariesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpRetrieveVmwareBinariesDataSourceCrud) Get() error {
	request := oci_ocvp.RetrieveVmwareBinariesRequest{}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.RetrieveVmwareBinaries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.RetrieveVmwareBinaries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpRetrieveVmwareBinariesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpRetrieveVmwareBinariesDataSource-", OcvpRetrieveVmwareBinariesDataSource(), s.D))

	items := make([]interface{}, 0, len(s.Res.Items))
	for _, item := range s.Res.Items {
		items = append(items, VmwareBinaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpRetrieveVmwareBinariesDataSource().Schema["items"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("items", items); err != nil {
		return err
	}

	return nil
}

func VmwareBinaryToMap(obj oci_ocvp.VmwareBinary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Checksum != nil {
		result["checksum"] = *obj.Checksum
	}

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	if obj.FileName != nil {
		result["file_name"] = *obj.FileName
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = strconv.FormatInt(*obj.SizeInBytes, 10)
	}

	return result
}
