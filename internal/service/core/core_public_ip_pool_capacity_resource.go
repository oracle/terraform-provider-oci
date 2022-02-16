// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func PublicIpPoolCapacityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createPublicIpPoolCapacity,
		Read:     readPublicIpPoolCapacity,
		Delete:   deletePublicIpPoolCapacity,
		Schema: map[string]*schema.Schema{
			// Required
			"public_ip_pool_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"byoip_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPublicIpPoolCapacity(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpPoolCapacityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readPublicIpPoolCapacity(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpPoolCapacityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func deletePublicIpPoolCapacity(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpPoolCapacityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CorePublicIpPoolCapacityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.PublicIpPool
	DisableNotFoundRetries bool
}

func (s *CorePublicIpPoolCapacityResourceCrud) ID() string {
	publicIpPoolId := s.D.Get("public_ip_pool_id").(string)
	byoipId := s.D.Get("byoip_id").(string)
	cidrBlock := s.D.Get("cidr_block").(string)

	id := "publicIpPoolId/" + publicIpPoolId + "/byoipId/" + byoipId + "/cidrBlock/" + cidrBlock

	return id
}

func (s *CorePublicIpPoolCapacityResourceCrud) Create() error {
	request := oci_core.AddPublicIpPoolCapacityRequest{}

	if publicIpPoolId, ok := s.D.GetOkExists("public_ip_pool_id"); ok {
		tmp := publicIpPoolId.(string)
		request.PublicIpPoolId = &tmp
	}

	if byoipId, ok := s.D.GetOkExists("byoip_id"); ok {
		tmp := byoipId.(string)
		request.ByoipRangeId = &tmp
	}

	if cidrBlock, ok := s.D.GetOkExists("cidr_block"); ok {
		tmp := cidrBlock.(string)
		request.CidrBlock = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AddPublicIpPoolCapacity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIpPool
	return nil
}

func (s *CorePublicIpPoolCapacityResourceCrud) Get() error {
	request := oci_core.GetPublicIpPoolRequest{}

	publicIpPoolId, _, cidrBlock, err := parseCapacityCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.PublicIpPoolId = &publicIpPoolId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetPublicIpPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIpPool

	cidrBlocks := s.Res.CidrBlocks

	exists := false

	for _, item := range cidrBlocks {
		if item == cidrBlock {
			exists = true
			break
		}
	}

	if !exists {
		return fmt.Errorf("cidr block  %s not found in response of PublicIpPool", cidrBlock)
	}

	return nil
}

func (s *CorePublicIpPoolCapacityResourceCrud) SetData() error {
	publicIpPoolId, byoipId, cidrBlock, err := parseCapacityCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("public_ip_pool_id", publicIpPoolId)
	s.D.Set("byoip_id", byoipId)
	s.D.Set("cidr_block", cidrBlock)

	return nil
}

func (s *CorePublicIpPoolCapacityResourceCrud) Delete() error {
	request := oci_core.RemovePublicIpPoolCapacityRequest{}

	if publicIpPoolId, ok := s.D.GetOkExists("public_ip_pool_id"); ok {
		tmp := publicIpPoolId.(string)
		request.PublicIpPoolId = &tmp
	}

	if cidrBlock, ok := s.D.GetOkExists("cidr_block"); ok {
		tmp := cidrBlock.(string)
		request.CidrBlock = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.RemovePublicIpPoolCapacity(context.Background(), request)
	return err
}

func parseCapacityCompositeId(compositeId string) (publicIpPoolId string, byoipId string, cidrBlock string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("publicIpPoolId/.*/byoipId/.*/cidrBlock/.*", compositeId)
	if !match || len(parts) != 7 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	publicIpPoolId, _ = url.PathUnescape(parts[1])
	byoipId, _ = url.PathUnescape(parts[3])
	cidrBlockPrefix, _ := url.PathUnescape(parts[5])
	cidrBlockSuffix, _ := url.PathUnescape(parts[6])
	cidrBlock = cidrBlockPrefix + "/" + cidrBlockSuffix

	return
}
