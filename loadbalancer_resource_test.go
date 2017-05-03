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

	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceLoadBalancerTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.LoadBalancer
	DeletedRes   *baremetal.LoadBalancer
}

func (s *ResourceLoadBalancerTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
}

func (s *ResourceLoadBalancerTestSuite) TestCreateResourceLoadBalancerMinimal() {
	s.ResourceName = "baremetal_load_balancer.t"
	s.Config = `
resource "baremetal_load_balancer" "t" {
  shape          = "stub_shape_id"
  compartment_id = "ocid1.compartment.stub_id"
  subnet_ids     = ["ocid1.subnet.stub_id"]
  display_name   = "stub_display_name"
}
`
	s.Config += testProviderConfig

	loadBalancerID := "ocid1.loadbalancer.stub_id"
	s.Res = &baremetal.LoadBalancer{
		CompartmentID: "ocid1.compartment.stub_id",
		Shape:         "stub_shape_id",
		SubnetIDs: []string{
			"ocid1.subnet.stub_id",
		},
		// Optional, Empty
		DisplayName: "",
		// Computed
		ID: loadBalancerID,
		IPAddresses: []baremetal.IPAddress{
			{
				"127.0.0.1",
			},
		},
		State:       baremetal.ResourceActive,
		TimeCreated: s.TimeCreated,
	}
	// TODO: inline above?
	s.Res.RequestID = "stub_opc_request_id"
	s.Res.DisplayName = "stub_display_name"

	s.DeletedRes = &baremetal.LoadBalancer{}
	*s.DeletedRes = *s.Res
	s.DeletedRes.State = baremetal.ResourceDeleted

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.Res.DisplayName

	workReqID := "ocid1.loadbalancerworkrequest.stub_id"
	wrSucc := &baremetal.WorkRequest{
		ID:             workReqID,
		LoadBalancerID: loadBalancerID,
		State:          baremetal.WorkRequestSucceeded,
	}

	s.Client.On(
		"CreateLoadBalancer",
		(*baremetal.BackendSet)(nil),
		(*baremetal.Certificate)(nil),
		s.Res.CompartmentID,
		(*baremetal.Listener)(nil),
		s.Res.Shape,
		s.Res.SubnetIDs,
		opts).Return(workReqID, nil)
	s.Client.On("GetWorkRequest", workReqID, (*baremetal.ClientRequestOptions)(nil)).Return(wrSucc, nil).Twice()
	s.Client.On("DeleteLoadBalancer", s.Res.ID, (*baremetal.ClientRequestOptions)(nil)).Return(workReqID, nil)
	s.Client.On("GetWorkRequest", workReqID, (*baremetal.ClientRequestOptions)(nil)).Return(wrSucc, nil).Twice()
	s.Client.On("GetLoadBalancer", loadBalancerID, (*baremetal.ClientRequestOptions)(nil)).Return(s.Res, nil).Twice()
	s.Client.On("GetLoadBalancer", loadBalancerID, (*baremetal.ClientRequestOptions)(nil)).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceLoadBalancerTestSuite) TestCreateResourceLoadBalancerMaximal() {
	s.ResourceName = "baremetal_load_balancer.t"
	s.Config = `
resource "baremetal_load_balancer" "t" {
  shape          = "stub_shape_id"
  compartment_id = "ocid1.compartment.stub_id"
  display_name   = "stub_display_name"
  subnet_ids     = ["ocid1.subnet.stub_id"]
}
`
	s.Config += testProviderConfig

	loadBalancerID := "ocid1.loadbalancer.stub_id"
	res := &baremetal.LoadBalancer{
		CompartmentID: "ocid1.compartment.stub_id",
		DisplayName:   "stub_display_name",
		Shape:         "stub_shape_id",
		SubnetIDs: []string{
			"ocid1.subnet.stub_id",
		},
		// Computed
		ID: loadBalancerID,
		IPAddresses: []baremetal.IPAddress{
			{
				"127.0.0.1",
			},
		},
		State:       baremetal.ResourceActive,
		TimeCreated: s.TimeCreated,
	}
	res.RequestID = "stub_opc_request_id"

	s.DeletedRes = &baremetal.LoadBalancer{}
	*s.DeletedRes = *res
	s.DeletedRes.State = baremetal.ResourceDeleted

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = res.DisplayName

	workReqID := "ocid1.loadbalancerworkrequest.stub_id"
	s.Client.On(
		"CreateLoadBalancer",
		(*baremetal.BackendSet)(nil),
		(*baremetal.Certificate)(nil),
		res.CompartmentID,
		(*baremetal.Listener)(nil),
		res.Shape,
		res.SubnetIDs,
		opts).Return(workReqID, nil)

	wrSucc := &baremetal.WorkRequest{
		ID:             workReqID,
		LoadBalancerID: loadBalancerID,
		State:          baremetal.WorkRequestSucceeded,
	}
	s.Client.On("GetWorkRequest", workReqID, (*baremetal.ClientRequestOptions)(nil)).Return(wrSucc, nil).Twice()

	s.Client.On("DeleteLoadBalancer", res.ID, (*baremetal.ClientRequestOptions)(nil)).Return(workReqID, nil)
	s.Client.On("GetWorkRequest", workReqID, (*baremetal.ClientRequestOptions)(nil)).Return(wrSucc, nil).Twice()
	s.Client.On("GetLoadBalancer", loadBalancerID, (*baremetal.ClientRequestOptions)(nil)).Return(res, nil).Twice()
	s.Client.On("GetLoadBalancer", loadBalancerID, (*baremetal.ClientRequestOptions)(nil)).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					// Assigned
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", res.ID),
					// resource.TestCheckResourceAttr(s.ResourceName, "state", res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", res.TimeCreated.String()),
					// Computed
					resource.TestCheckResourceAttr(s.ResourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_addresses.0", "127.0.0.1"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerTestSuite))
}
