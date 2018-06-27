// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstances,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     InstanceResource(),
			},
		},
	}
}

func readInstances(d *schema.ResourceData, m interface{}) error {
	sync := &InstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type InstancesDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstancesResponse
}

func (s *InstancesDataSourceCrud) Get() error {
	request := oci_core.ListInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// @CODEGEN 1/2018: page & limit were never actually wired to requests

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *InstancesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			instance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			instance["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			instance["display_name"] = *r.DisplayName
		}

		if r.ExtendedMetadata != nil {
			instance["extended_metadata"] = convertNestedMapToFlatMap(r.ExtendedMetadata)
		}

		instance["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instance["id"] = *r.Id
		}

		if r.ImageId != nil {
			instance["image"] = *r.ImageId
		}

		if r.IpxeScript != nil {
			instance["ipxe_script"] = *r.IpxeScript
		}

		instance["launch_mode"] = r.LaunchMode

		if r.LaunchOptions != nil {
			instance["launch_options"] = []interface{}{LaunchOptionsToMap(r.LaunchOptions)}
		}

		if r.Metadata != nil {
			instance["metadata"] = r.Metadata
		}

		if r.Region != nil {
			instance["region"] = *r.Region
		}

		if r.Shape != nil {
			instance["shape"] = *r.Shape
		}

		if r.SourceDetails != nil {
			instance["source_details"] = []interface{}{InstanceSourceDetailsToMap(&r.SourceDetails, nil, nil)}
		}

		instance["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			instance["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, InstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		panic(err)
	}

	return
}

func convertNestedMapToFlatMap(m map[string]interface{}) map[string]string {
	flatMap := make(map[string]string)
	var ok bool
	for key, val := range m {
		if flatMap[key], ok = val.(string); !ok {
			mapValStr, err := json.Marshal(val)
			if err != nil {
				mapValStr = []byte{}
			}
			flatMap[key] = string(mapValStr)
		}
	}
	return flatMap
}
