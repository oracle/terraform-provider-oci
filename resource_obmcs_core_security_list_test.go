// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"




	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.SecurityList
	DeletingRes  *baremetal.SecurityList
	DeletedRes   *baremetal.SecurityList
}

func extraWait(ew crud.ExtraWaitPostDelete) {
	return
}

func (s *ResourceCoreSecurityListTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_core_security_list" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
      egress_security_rules {
				destination = "destination"
				icmp_options {
					"code" = 1
					"type" = 2
				}
				protocol = "protocol"
				stateless = true
			}
      ingress_security_rules {
				tcp_options {
					"max" = 2
					"min" = 1
				}
				protocol = "protocol"
				source = "source"
			}
			vcn_id = "vcn_id"
		}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_core_security_list.t"

	egressRules := []baremetal.EgressSecurityRule{
		{
			Destination: "destination",
			ICMPOptions: &baremetal.ICMPOptions{Code: 1, Type: 2},
			Protocol:    "protocol",
			IsStateless: true,
		},
	}
	ingressRules := []baremetal.IngressSecurityRule{
		{
			TCPOptions: &baremetal.TCPOptions{
				baremetal.PortRange{Max: 2, Min: 1},
			},
			Protocol: "protocol",
			Source:   "source",
		},
	}

	s.Res = &baremetal.SecurityList{
		CompartmentID:        "compartment_id",
		DisplayName:          "display_name",
		EgressSecurityRules:  egressRules,
		ID:                   "id",
		IngressSecurityRules: ingressRules,
		State:                baremetal.ResourceAvailable,
		TimeCreated:          s.TimeCreated,
		VcnID:                "vcn_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	deletingRes := *s.Res
	s.DeletingRes = &deletingRes
	s.DeletingRes.State = baremetal.ResourceTerminating

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"

	s.Client.On("CreateSecurityList",
		"compartment_id",
		"vcn_id",
		egressRules,
		ingressRules,
		opts,
	).Return(s.Res, nil)

	s.Client.On("DeleteSecurityList", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreSecurityListTestSuite) TestCreateResourceCoreSecurityList() {
	s.Client.On("GetSecurityList", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetSecurityList", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.0.code", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "2"),
				),
			},
		},
	})
}

func (s ResourceCoreSecurityListTestSuite) TestUpdateSecurityList() {
	s.Client.On("GetSecurityList", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_security_list" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
      egress_security_rules {
				destination = "destination"
				icmp_options {
					"code" = 1
					"type" = 2
				}
				protocol = "protocol"
				stateless = true
			}
      ingress_security_rules {
				tcp_options {
					"max" = 3
					"min" = 1
				}
				protocol = "protocol"
				source = "source"
			}
			vcn_id = "vcn_id"
		}
	`
	config += testProviderConfig()

	ingressRules := []baremetal.IngressSecurityRule{
		{
			TCPOptions: &baremetal.TCPOptions{
				baremetal.PortRange{Max: 3, Min: 1},
			},
			Protocol: "protocol",
			Source:   "source",
		},
	}

	res := &baremetal.SecurityList{
		CompartmentID:        "compartment_id",
		DisplayName:          "display_name",
		EgressSecurityRules:  s.Res.EgressSecurityRules,
		ID:                   "id",
		IngressSecurityRules: ingressRules,
		State:                baremetal.ResourceAvailable,
		TimeCreated:          s.TimeCreated,
		VcnID:                "vcn_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateSecurityListOptions{
		EgressRules:  s.Res.EgressSecurityRules,
		IngressRules: ingressRules,
	}

	s.Client.On("UpdateSecurityList", "id", opts).Return(res, nil)

	s.Client.On("GetSecurityList", "id").Return(res, nil).Times(1)
	s.Client.On("GetSecurityList", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.0.code", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "3"),
				),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestDeleteSecurityList() {
	s.Client.On("GetSecurityList", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetSecurityList", "id").Return(s.DeletingRes, nil).Times(2)
	s.Client.On("GetSecurityList", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteSecurityList", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreSecurityListTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSecurityListTestSuite))
}
