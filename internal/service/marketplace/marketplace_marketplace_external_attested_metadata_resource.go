// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MarketplaceMarketplaceExternalAttestedMetadataResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMarketplaceMarketplaceExternalAttestedMetadata,
		Read:     readMarketplaceMarketplaceExternalAttestedMetadata,
		Delete:   deleteMarketplaceMarketplaceExternalAttestedMetadata,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"serialized_jwt": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMarketplaceMarketplaceExternalAttestedMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceMarketplaceExternalAttestedMetadataResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.CreateResource(d, sync)
}

func readMarketplaceMarketplaceExternalAttestedMetadata(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteMarketplaceMarketplaceExternalAttestedMetadata(d *schema.ResourceData, m interface{}) error {
	return nil
}

type MarketplaceMarketplaceExternalAttestedMetadataResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_marketplace.MarketplaceClient
	Res                    *oci_marketplace.MarketplaceExternalAttestedMetadata
	DisableNotFoundRetries bool
}

func (s *MarketplaceMarketplaceExternalAttestedMetadataResourceCrud) ID() string {
	return fmt.Sprintf("%s/%s", s.D.Get("compartment_id").(string), s.D.Get("instance_id").(string))
}

func (s *MarketplaceMarketplaceExternalAttestedMetadataResourceCrud) Create() error {
	request := oci_marketplace.CreateMarketplaceExternalAttestedMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.CreateMarketplaceExternalAttestedMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MarketplaceExternalAttestedMetadata
	return nil
}

func (s *MarketplaceMarketplaceExternalAttestedMetadataResourceCrud) SetData() error {
	if s.Res.SerializedJwt != nil {
		s.D.Set("serialized_jwt", *s.Res.SerializedJwt)
	}

	return nil
}
