package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
)

func literalTypeHashCodeForSets(m interface{}) int {
	return hashcode.String(fmt.Sprintf("%v", m))
}
