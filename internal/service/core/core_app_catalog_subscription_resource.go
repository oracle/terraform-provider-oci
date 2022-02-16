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
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreAppCatalogSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreAppCatalogSubscription,
		Read:     readCoreAppCatalogSubscription,
		Delete:   deleteCoreAppCatalogSubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listing_resource_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oracle_terms_of_use_link": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_retrieved": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Optional
			"eula_link": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"listing_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publisher_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreAppCatalogSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreAppCatalogSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func deleteCoreAppCatalogSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreAppCatalogSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.AppCatalogSubscription
	DisableNotFoundRetries bool
}

func (s *CoreAppCatalogSubscriptionResourceCrud) ID() string {
	return getSubscriptionCompositeId(*s.Res.CompartmentId, *s.Res.ListingId, *s.Res.ListingResourceVersion)
}

func (s *CoreAppCatalogSubscriptionResourceCrud) Create() error {
	request := oci_core.CreateAppCatalogSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if eulaLink, ok := s.D.GetOkExists("eula_link"); ok {
		tmp := eulaLink.(string)
		request.EulaLink = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if listingResourceVersion, ok := s.D.GetOkExists("listing_resource_version"); ok {
		tmp := listingResourceVersion.(string)
		request.ListingResourceVersion = &tmp
	}

	if oracleTermsOfUseLink, ok := s.D.GetOkExists("oracle_terms_of_use_link"); ok {
		tmp := oracleTermsOfUseLink.(string)
		request.OracleTermsOfUseLink = &tmp
	}

	if signature, ok := s.D.GetOkExists("signature"); ok {
		tmp := signature.(string)
		request.Signature = &tmp
	}

	if timeRetrieved, ok := s.D.GetOkExists("time_retrieved"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeRetrieved.(string))
		if err != nil {
			return err
		}
		request.TimeRetrieved = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.CreateAppCatalogSubscription(context.Background(), request)
	if err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res != nil && s.Res.TimeCreated != nil }
	compositeId := getSubscriptionCompositeId(*request.CompartmentId, *request.ListingId, *request.ListingResourceVersion)
	s.D.SetId(compositeId)
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CoreAppCatalogSubscriptionResourceCrud) Get() error {
	compartmentId, listingId, listingResourceVersion, err := parseSubscriptionCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}
	request := oci_core.ListAppCatalogSubscriptionsRequest{
		CompartmentId: &compartmentId,
		ListingId:     &listingId,
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ListAppCatalogSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	isFound := false
	for _, item := range response.Items {
		if *item.ListingResourceVersion == listingResourceVersion {
			isFound = true
			s.Res = &oci_core.AppCatalogSubscription{
				CompartmentId:          item.CompartmentId,
				ListingId:              item.ListingId,
				DisplayName:            item.DisplayName,
				ListingResourceId:      item.ListingResourceId,
				ListingResourceVersion: item.ListingResourceVersion,
				PublisherName:          item.PublisherName,
				TimeCreated:            item.TimeCreated,
				Summary:                item.Summary,
			}
			break
		}
	}

	for !isFound && response.OpcNextPage != nil {
		request.Page = response.OpcNextPage
		response, err := s.Client.ListAppCatalogSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}
		for _, item := range response.Items {
			if *item.ListingResourceVersion == listingResourceVersion {
				isFound = true
				s.Res = &oci_core.AppCatalogSubscription{
					CompartmentId:          item.CompartmentId,
					ListingId:              item.ListingId,
					DisplayName:            item.DisplayName,
					ListingResourceId:      item.ListingResourceId,
					ListingResourceVersion: item.ListingResourceVersion,
					PublisherName:          item.PublisherName,
					TimeCreated:            item.TimeCreated,
				}
				break
			}
		}
	}
	if !isFound {
		s.Res = nil
	}
	return nil
}

func (s *CoreAppCatalogSubscriptionResourceCrud) Delete() error {
	request := oci_core.DeleteAppCatalogSubscriptionRequest{}
	compartmentId, listingId, listingResourceVersion, err := parseSubscriptionCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}
	request.CompartmentId = &compartmentId
	request.ListingId = &listingId
	request.ResourceVersion = &listingResourceVersion
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err = s.Client.DeleteAppCatalogSubscription(context.Background(), request)
	if err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res == nil }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutDelete))
}

func (s *CoreAppCatalogSubscriptionResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	compartmentId, listingId, listingResourceVersion, err := parseSubscriptionCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}
	s.D.Set("compartment_id", compartmentId)
	s.D.Set("listing_id", listingId)
	s.D.Set("listing_resource_version", listingResourceVersion)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ListingId != nil {
		s.D.Set("listing_id", *s.Res.ListingId)
	}

	if s.Res.ListingResourceId != nil {
		s.D.Set("listing_resource_id", *s.Res.ListingResourceId)
	}

	if s.Res.ListingResourceVersion != nil {
		s.D.Set("listing_resource_version", *s.Res.ListingResourceVersion)
	}

	if s.Res.PublisherName != nil {
		s.D.Set("publisher_name", *s.Res.PublisherName)
	}

	if s.Res.Summary != nil {
		s.D.Set("summary", *s.Res.Summary)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func parseSubscriptionCompositeId(compositeId string) (compartmentId string, listingId string, listingResourceVersion string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("compartmentId/.*/listingId/.*/listingResourceVersion/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	compartmentId, _ = url.PathUnescape(parts[1])
	listingId, _ = url.PathUnescape(parts[3])
	listingResourceVersion, _ = url.PathUnescape(parts[5])
	return
}

func getSubscriptionCompositeId(compartmentId string, listingId string, listingResourceVersion string) string {
	compartmentId = url.PathEscape(compartmentId)
	listingId = url.PathEscape(listingId)
	listingResourceVersion = url.PathEscape(listingResourceVersion)
	compositeId := "compartmentId/" + compartmentId + "/listingId/" + listingId + "/listingResourceVersion/" + listingResourceVersion
	return compositeId
}
