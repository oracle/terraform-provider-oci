// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"sort"

	"fmt"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityAvailabilityDomainDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAvailabilityDomain,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ad_number": {
				Type:          schema.TypeInt,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"id"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readAvailabilityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &AvailabilityDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type AvailabilityDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAvailabilityDomainsResponse
}

func (s *AvailabilityDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AvailabilityDomainDataSourceCrud) Get() error {
	request := oci_identity.ListAvailabilityDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AvailabilityDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	items := s.Res.Items

	// sort ADs by name
	sort.Slice(items, func(i, j int) bool {
		return *items[i].Name < *items[j].Name
	})

	id, idOk := s.D.GetOkExists("id")
	adNumber, adNumberOk := s.D.GetOkExists("ad_number")

	for _, r := range items {
		adNumberFromName := adNumberFromName(*r.Name)

		if adNumberFromName == -1 {
			return fmt.Errorf("unable to get index from Availability Domain name pattern: %s", *r.Name)
		}

		// match on user supplied AD index or Availability Domain global OCID
		if adNumberOk && adNumber == adNumberFromName || idOk && id == *r.Id {
			s.D.SetId(*r.Id)

			if r.CompartmentId != nil {
				s.D.Set("compartment_id", *r.CompartmentId)
			}
			if r.Name != nil {
				s.D.Set("name", *r.Name)
			}

			s.D.Set("ad_number", adNumberFromName)

			return nil
		}
	}

	if idOk {
		return fmt.Errorf("no Availability Domain match for ID: %s", id)
	} else {
		return fmt.Errorf("no Availability Domain match for AD number: %d", adNumber)
	}

}

func adNumberFromName(adName string) int {
	// case insensitive matching
	regex := regexp.MustCompile(`(?i)AD-(\d)`)
	res := regex.FindAllStringSubmatch(adName, -1)

	// no matching AD name
	if res == nil || len(res) < 1 {
		return -1
	}

	// int coercion failure
	index, err := strconv.Atoi(res[0][1])
	if err != nil {
		return -1
	}

	return index
}
