package options

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

type HelpersTestSuite struct {
	suite.Suite
	resource *mockResourceProvider
}

func (s *HelpersTestSuite) SetupTest() {
	s.resource = &mockResourceProvider{
		vals: make(map[string]interface{}),
	}

	s.resource.Set("availability_domain", "availabilityid")
	s.resource.Set("image_id", "imageid")
	s.resource.Set("instance_id", "instanceid")
	s.resource.Set("vnic_id", "vnicid")
	s.resource.Set("limit", 100)
}

func TestHelpers(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}
