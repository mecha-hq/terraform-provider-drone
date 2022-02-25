package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"drone_user.test_user1", "login", "test_user1",
					),
				),
			},
		},
	})
}

const testAccResourceUser = `
resource "drone_user" "test_user1" {
  login = "test_user1"
}
`
