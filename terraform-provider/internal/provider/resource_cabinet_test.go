package provider

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"regexp"
	"terraform-provider-everything/internal/ikea"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var color = "#FFFFFF"
var front = "RINGHULT"
var kitchenId, _ = uuid.GenerateUUID()

func TestAccResourceCabinet(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		CheckDestroy:      testAccResourceCabinetDestroy,
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceCabinet("invalid!§§$%", front, kitchenId),
				ExpectError: regexp.MustCompile("color must be a valid hex color code"),
			},
			{
				Config: testAccResourceCabinet(color, front, kitchenId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ikea_cabinet.test", "color", color),
					resource.TestCheckResourceAttr("ikea_cabinet.test", "front", front),
				),
			},
		},
	})
}

func testAccResourceCabinet(color string, front string, kitchenId string) string {
	return fmt.Sprintf(`
		resource "ikea_cabinet" "test" {
		  color = "%s"
		  front = "%s"
		  kitchen_id = "%s"
		}`, color, front, kitchenId)
}

func testAccResourceCabinetDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	client := testAccProvider.Meta().(*ikea.ApiClient)

	// loop through the resources in state, verifying each widget
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ikea_cabinet" {
			continue
		}
		id := rs.Primary.ID

		// check if the resource with the ID from the state still exists in the backend
		// if so, the destroy operation was not successful
		cabinet, _ := client.GetCabinet(id)
		if cabinet != nil {
			return fmt.Errorf("cabinet %s still exists", cabinet.Id)
		}
	}

	return nil
}
