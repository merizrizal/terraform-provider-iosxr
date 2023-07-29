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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrEVPNSegmentRoutingSRv6EVI(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "vpn_id", "1235"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_import_two_byte_as_format.0.as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_import_two_byte_as_format.0.assigned_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_export_ipv4_address_format.0.ipv4_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_export_ipv4_address_format.0.assigned_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "advertise_mac", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn_segment_routing_srv6_evi.test", "locator", "LOC12"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrEVPNSegmentRoutingSRv6EVIConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrEVPNSegmentRoutingSRv6EVIConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_evpn_segment_routing_srv6_evi.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/evpn/evis/segment-routing/srv6/evi[vpn-id=1235]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrEVPNSegmentRoutingSRv6EVIConfig_minimum() string {
	config := `resource "iosxr_evpn_segment_routing_srv6_evi" "test" {` + "\n"
	config += `	vpn_id = 1235` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrEVPNSegmentRoutingSRv6EVIConfig_all() string {
	config := `resource "iosxr_evpn_segment_routing_srv6_evi" "test" {` + "\n"
	config += `	vpn_id = 1235` + "\n"
	config += `	description = "My Description"` + "\n"
	config += `	bgp_route_target_import_two_byte_as_format = [{` + "\n"
	config += `		as_number = 1` + "\n"
	config += `		assigned_number = 1` + "\n"
	config += `		}]` + "\n"
	config += `	bgp_route_target_export_ipv4_address_format = [{` + "\n"
	config += `		ipv4_address = "1.1.1.1"` + "\n"
	config += `		assigned_number = 1` + "\n"
	config += `		}]` + "\n"
	config += `	advertise_mac = true` + "\n"
	config += `	locator = "LOC12"` + "\n"
	config += `}` + "\n"
	return config
}
