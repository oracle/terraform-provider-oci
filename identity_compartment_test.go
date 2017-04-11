// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityCompartmentTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.Compartment
}

func (s *ResourceIdentityCompartmentTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	configfn := func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	}

	s.Provider = Provider(configfn)
	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["baremetal_identity_compartment"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_compartment" "t" {
			name = "name!"
			description = "desc!"
		}
	`
	s.Config += testProviderConfig
	s.ResourceName = "baremetal_identity_compartment.t"
	s.Res = &baremetal.Compartment{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
	}
	s.Client.On("CreateCompartment", "name!", "desc!", (*baremetal.RetryTokenOptions)(nil)).Return(s.Res, nil)
}

func (s *ResourceIdentityCompartmentTestSuite) TestCreateResourceIdentityCompartment() {
	s.Client.On("GetCompartment", "id!").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "description", s.Res.Description),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityCompartmentTestSuite) TestCreateResourceIdentityCompartmentPolling() {
	s.Res.State = baremetal.ResourceCreating
	s.Client.On("GetCompartment", "id!").Return(s.Res, nil).Once()

	u := *s.Res
	u.State = baremetal.ResourceActive
	s.Client.On("GetCompartment", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
			},
		},
	})
}

func (s *ResourceIdentityCompartmentTestSuite) TestUpdateResourceIdentityCompartmentDescription() {
	s.Client.On("GetCompartment", "id!").Return(s.Res, nil).Twice()

	c := `
		resource "baremetal_identity_compartment" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	c += testProviderConfig
	u := *s.Res
	u.Description = "newdesc!"

	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateCompartment", "id!", opts).Return(&u, nil)
	s.Client.On("GetCompartment", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityCompartmentTestSuite) TestFailedUpdateResourceIdentityCompartmentDescription() {
	s.Client.On("GetCompartment", "id!").Return(s.Res, nil).Times(3)

	c := `
		resource "baremetal_identity_compartment" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	c += testProviderConfig
	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateCompartment", "id!", opts).Return(nil, errors.New("FAILED!")).Once()

	u := *s.Res
	u.Description = "newdesc!"
	s.Client.On("UpdateCompartment", "id!", opts).Return(&u, nil)
	s.Client.On("GetCompartment", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:      c,
				ExpectError: regexp.MustCompile(`FAILED`),
				Check:       resource.TestCheckResourceAttr(s.ResourceName, "description", "desc!"),
			},
			{
				Config: c,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
			},
		},
	})
}

func TestResourceIdentityCompartmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityCompartmentTestSuite))
}
