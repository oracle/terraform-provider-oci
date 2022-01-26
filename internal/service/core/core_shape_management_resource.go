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
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreShapeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreShape,
		Read:     readCoreShape,
		Delete:   deleteCoreShape,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCoreShape(d *schema.ResourceData, m interface{}) error {
	sync := &CoreShapeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreShape(d *schema.ResourceData, m interface{}) error {
	sync := &CoreShapeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func deleteCoreShape(d *schema.ResourceData, m interface{}) error {
	sync := &CoreShapeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreShapeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ImageShapeCompatibilityEntry
	DisableNotFoundRetries bool
}

func (s *CoreShapeResourceCrud) ID() string {
	return getShapeCompositeId(s.D.Get("compartment_id").(string), s.D.Get("image_id").(string), s.D.Get("shape_name").(string))
}

func (s *CoreShapeResourceCrud) Create() error {
	request := oci_core.AddImageShapeCompatibilityEntryRequest{}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AddImageShapeCompatibilityEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ImageShapeCompatibilityEntry
	return nil
}

func (s *CoreShapeResourceCrud) Get() error {
	request := oci_core.ListShapesRequest{}

	compartmentId, imageId, shape, err := parseShapeCompositeId(s.D.Id())

	if err == nil {
		request.ImageId = &imageId
		request.CompartmentId = &compartmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}
	for _, item := range response.Items {
		if *item.Shape == shape {
			s.Res = &oci_core.ImageShapeCompatibilityEntry{
				Shape:   item.Shape,
				ImageId: request.ImageId,
			}
			return nil
		}
	}
	return fmt.Errorf("requested compatible shape %s does not exist", shape)
}

func (s *CoreShapeResourceCrud) Delete() error {
	request := oci_core.RemoveImageShapeCompatibilityEntryRequest{}

	_, imageId, shape, err := parseShapeCompositeId(s.D.Id())

	if err == nil {
		request.ImageId = &imageId
		request.ShapeName = &shape
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err = s.Client.RemoveImageShapeCompatibilityEntry(context.Background(), request)
	return err
}

func (s *CoreShapeResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape_name", *s.Res.Shape)
	}

	return nil
}

func getShapeCompositeId(compartmentId string, imageId string, shape string) string {
	compartmentId = url.PathEscape(compartmentId)
	imageId = url.PathEscape(imageId)
	compositeId := "c/" + compartmentId + "/i/" + imageId + "/s/" + shape
	return compositeId
}

func parseShapeCompositeId(compositeId string) (compartmentId string, imageId string, shape string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("c/.*/i/.*/s/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	compartmentId, _ = url.PathUnescape(parts[1])
	imageId, _ = url.PathUnescape(parts[3])
	shape, _ = url.PathUnescape(parts[5])

	return
}
