package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockResourceProvider struct {
	vals map[string]interface{}
}

func (m *mockResourceProvider) GetOk(key string) (val interface{}, ok bool) {
	val, ok = m.vals[key]
	return
}

func (m *mockResourceProvider) Set(key string, val interface{}) {
	m.vals[key] = val
}

type ResourceCoreHelperTestSuite struct {
	suite.Suite
	resource *mockResourceProvider
}

func (s *ResourceCoreHelperTestSuite) SetupTest() {
	s.resource = &mockResourceProvider{
		vals: make(map[string]interface{}),
	}

	s.resource.Set("availability_domain", "availabilityid")
	s.resource.Set("image_id", "imageid")
	s.resource.Set("instance_id", "instanceid")
	s.resource.Set("vnic_id", "vnicid")
	s.resource.Set("limit", 100)
}

func (s *ResourceCoreHelperTestSuite) TestGetCoreOptionsFromResourceData() {

	opts := getCoreOptionsFromResourceData(s.resource, "availability_domain", "vnic_id", "limit")
	s.NotNil(opts)
	s.Equal(len(opts), 1)
	s.Equal("availabilityid", opts[0].AvailabilityDomain)
	s.Equal("vnicid", opts[0].VnicID)
	s.Equal(uint64(100), opts[0].Limit)
	s.NotEqual("instanceid", opts[0].InstanceID)
}

func (s *ResourceCoreHelperTestSuite) TestGetCoreOptionsFromResourceDataMissingKey() {
	s.resource.Set("whacky", "foo")
	s.Panics(func() {
		getCoreOptionsFromResourceData(s.resource, "availability_domain", "whacky")
	},
		"Unknown key 'whacky' supplied for CoreOptions",
	)

}

func TestResourceCoreHelpers(t *testing.T) {
	suite.Run(t, new(ResourceCoreHelperTestSuite))
}
