package statuscake

import (
	"fmt"
	"testing"

	"github.com/DreamItGetIT/statuscake"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccStatusCake_basic(t *testing.T) {
	var test statuscake.Test

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccTestCheckDestroy(&test),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTestConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccTestCheckExists("statuscake_test.google", &test),
				),
			},
		},
	})
}

func testAccTestCheckExists(rn string, project *statuscake.Test) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("TestID not set")
		}

		//		client := testAccProvider.Meta().(*statuscake.Client)
		//		gotProject, err := client.GetProject(rs.Primary.ID)
		//		if err != nil {
		//			return fmt.Errorf("error getting project: %s", err)
		//		}
		//
		//		*project = *gotProject

		return nil
	}
}

func testAccTestCheckDestroy(project *statuscake.Test) resource.TestCheckFunc {
	//	return func(s *terraform.State) error {
	//		client := testAccProvider.Meta().(*statuscake.Client)
	//		//		_, err := client.Tests().All()
	//		//		if err == nil {
	//		//			return fmt.Errorf("test still exists")
	//		//		}
	//		//		if _, ok := err.(*statuscake.NotFoundError); !ok {
	//		//			return fmt.Errorf("got something other than NotFoundError (%v) when getting test", err)
	//		//		}
	//
	//		return nil
	//	}
	return nil
}

const testAccTestConfig_basic = `
resource "statuscake_test" "google" {
  website_name = "google.com"
  website_url = "www.google.com"
  test_type = "HTTP"
  check_rate = 300
}
`
