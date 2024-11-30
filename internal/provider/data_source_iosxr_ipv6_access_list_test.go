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

func TestAccDataSourceIosxrIPv6AccessList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.sequence_number", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_address", "1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_prefix_length", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_port_range_start", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_port_range_end", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_destination_host", "2::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_destination_port_eq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_nexthop1_ipv6", "3::3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_nexthop2_ipv6", "4::4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_log", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrIPv6AccessListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrIPv6AccessListConfig() string {
	config := `resource "iosxr_ipv6_access_list" "test" {` + "\n"
	config += `	access_list_name = "TEST1"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `		sequence_number = 22` + "\n"
	config += `		permit_protocol = "tcp"` + "\n"
	config += `		permit_source_address = "1::1"` + "\n"
	config += `		permit_source_prefix_length = 64` + "\n"
	config += `		permit_source_port_range_start = "100"` + "\n"
	config += `		permit_source_port_range_end = "200"` + "\n"
	config += `		permit_destination_host = "2::1"` + "\n"
	config += `		permit_destination_port_eq = "10"` + "\n"
	config += `		permit_nexthop1_ipv6 = "3::3"` + "\n"
	config += `		permit_nexthop2_ipv6 = "4::4"` + "\n"
	config += `		permit_log = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_ipv6_access_list" "test" {
			access_list_name = "TEST1"
			depends_on = [iosxr_ipv6_access_list.test]
		}
	`
	return config
}