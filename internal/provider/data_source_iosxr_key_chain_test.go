// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrKeyChain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrKeyChainConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.key_name", "1"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.key_string_password", "00071A150754"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.cryptographic_algorithm", "hmac-md5"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_hour", "11"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_minute", "52"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_second", "55"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_day_of_month", "21"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_month", "january"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_year", "2023"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.accept_lifetime_infinite", "true"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_hour", "8"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_minute", "36"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_second", "22"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_day_of_month", "15"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_month", "january"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_start_time_year", "2023"),
					resource.TestCheckResourceAttr("data.iosxr_key_chain.test", "keys.0.send_lifetime_infinite", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrKeyChainConfig = `

resource "iosxr_key_chain" "test" {
	name = "KEY11"
	keys = [{
		key_name = "1"
		key_string_password = "00071A150754"
		cryptographic_algorithm = "hmac-md5"
		accept_lifetime_start_time_hour = 11
		accept_lifetime_start_time_minute = 52
		accept_lifetime_start_time_second = 55
		accept_lifetime_start_time_day_of_month = 21
		accept_lifetime_start_time_month = "january"
		accept_lifetime_start_time_year = 2023
		accept_lifetime_infinite = true
		send_lifetime_start_time_hour = 8
		send_lifetime_start_time_minute = 36
		send_lifetime_start_time_second = 22
		send_lifetime_start_time_day_of_month = 15
		send_lifetime_start_time_month = "january"
		send_lifetime_start_time_year = 2023
		send_lifetime_infinite = true
	}]
}

data "iosxr_key_chain" "test" {
	name = "KEY11"
	depends_on = [iosxr_key_chain.test]
}
`
