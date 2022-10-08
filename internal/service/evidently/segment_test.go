package evidently_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudwatchevidently"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfcloudwatchevidently "github.com/hashicorp/terraform-provider-aws/internal/service/evidently"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func TestAccEvidentlySegment_basic(t *testing.T) {
	var segment cloudwatchevidently.Segment

	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_evidently_segment.test"
	pattern := "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(cloudwatchevidently.EndpointsID, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, cloudwatchevidently.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSegmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentConfig_basic(rName, pattern),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSegmentExists(resourceName, &segment),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "evidently", fmt.Sprintf("segment/%s", rName)),
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_count"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_time"),
					resource.TestCheckResourceAttrSet(resourceName, "launch_count"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "pattern", pattern),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccEvidentlySegment_disappears(t *testing.T) {
	var segment cloudwatchevidently.Segment

	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	pattern := "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"
	resourceName := "aws_evidently_segment.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, cloudwatchevidently.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSegmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentConfig_basic(rName, pattern),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSegmentExists(resourceName, &segment),
					acctest.CheckResourceDisappears(acctest.Provider, tfcloudwatchevidently.ResourceSegment(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckSegmentDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EvidentlyConn
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_evidently_segment" {
			continue
		}

		_, err := tfcloudwatchevidently.FindSegmentByName(context.Background(), conn, rs.Primary.ID)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("CloudWatch Evidently Segment %s still exists", rs.Primary.ID)
	}

	return nil
}

func testAccCheckSegmentExists(n string, v *cloudwatchevidently.Segment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No CloudWatch Evidently Segment ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EvidentlyConn

		output, err := tfcloudwatchevidently.FindSegmentByName(context.Background(), conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccSegmentConfig_basic(rName, pattern string) string {
	return fmt.Sprintf(`
resource "aws_evidently_segment" "test" {
  name    = %[1]q
  pattern = %[2]q
}
`, rName, pattern)
}
