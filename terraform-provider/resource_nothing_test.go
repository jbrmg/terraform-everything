package main

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var something, _ = uuid.GenerateUUID()
var anything, _ = uuid.GenerateUUID()

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		CheckDestroy:      testAccResourceNothingDestroy,
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceNothing("invalid!§§$%", anything),
				ExpectError: regexp.MustCompile("Something is in invalid format"),
			},
			{
				Config: testAccResourceNothing(something, anything),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("everything_nothing.test", "something", something),
					resource.TestCheckResourceAttr("everything_nothing.test", "anything", anything),
				),
			},
		},
	})
}

func testAccResourceNothing(something string, anything string) string {
	return fmt.Sprintf(`
		resource "everything_nothing" "test" {
		  something = "%s"
		  anything = "%s"
		}`, something, anything)
}

func testAccResourceNothingDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	client := testAccProvider.Meta().(*apiClient)

	// loop through the resources in state, verifying each widget
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "everything_nothing" {
			continue
		}
		id := rs.Primary.ID

		// check if the resource with the ID from the state still exists in the backend
		// if so, the destroy operation was not successful
		nothing, _ := client.GetNothing(id)
		if nothing != nil {
			return fmt.Errorf("nothing %s still exists", nothing.Id)
		}
	}

	return nil
}
