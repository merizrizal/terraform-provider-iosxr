// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrSegmentRoutingV6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "encapsulation_source_address", "fccc:0:214::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "locators.0.locator_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "locators.0.name", "Locator1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "locators.0.micro_segment_behavior", "unode-psp-usd"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "locators.0.prefix", "fccc:0:214::"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_segment_routing_v6.test", "locators.0.prefix_length", "48"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSegmentRoutingV6Config(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrSegmentRoutingV6Config() string {
	config := `resource "iosxr_segment_routing_v6" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	enable = true` + "\n"
	config += `	encapsulation_source_address = "fccc:0:214::1"` + "\n"
	config += `	locators = [{` + "\n"
	config += `		locator_enable = true` + "\n"
	config += `		name = "Locator1"` + "\n"
	config += `		micro_segment_behavior = "unode-psp-usd"` + "\n"
	config += `		prefix = "fccc:0:214::"` + "\n"
	config += `		prefix_length = 48` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_segment_routing_v6" "test" {
			depends_on = [iosxr_segment_routing_v6.test]
		}
	`
	return config
}