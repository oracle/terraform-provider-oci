package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// Custom assertion for TypeList attributes.
func testCheckAttributeTypeList(resourceName, attributeName string, expecteds []string) resource.TestCheckFunc {
	return func(s *terraform.State) (e error) {
		ms := s.RootModule()
		rs, ok := ms.Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", resourceName)
		}

		for i, expected := range expecteds {
			// Keys for individual list elements are represented in Terraform as
			// key.0 key.1 ... key.(N-1)
			key := fmt.Sprintf("%s.%d", attributeName, i)
			actual := is.Attributes[key]
			if actual != expected {
				return fmt.Errorf("Expected '%s' but got '%s'", actual, expected)
			}
		}

		return
	}
}
