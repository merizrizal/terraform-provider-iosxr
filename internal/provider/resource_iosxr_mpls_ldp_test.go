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

func TestAccIosxrMPLSLDP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "address_families.0.af_name", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "capabilities_sac_ipv4_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "capabilities_sac_ipv6_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "capabilities_sac_fec128_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "capabilities_sac_fec129_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_logging_notifications", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_address_families.0.name", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_address_families.0.make_before_break_delay", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_address_families.0.forwarding_recursive", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_address_families.0.forwarding_recursive_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "mldp_address_families.0.recursive_fec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_mpls_ldp.test", "session_protection", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrMPLSLDPPrerequisitesConfig + testAccIosxrMPLSLDPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrMPLSLDPPrerequisitesConfig + testAccIosxrMPLSLDPConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_mpls_ldp.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-mpls-ldp-cfg:/mpls/ldp",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrMPLSLDPPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

`

func testAccIosxrMPLSLDPConfig_minimum() string {
	config := `resource "iosxr_mpls_ldp" "test" {` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrMPLSLDPConfig_all() string {
	config := `resource "iosxr_mpls_ldp" "test" {` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	address_families = [{` + "\n"
	config += `		af_name = "ipv4"` + "\n"
	config += `		}]` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		}]` + "\n"
	config += `	capabilities_sac_ipv4_disable = true` + "\n"
	config += `	capabilities_sac_ipv6_disable = true` + "\n"
	config += `	capabilities_sac_fec128_disable = true` + "\n"
	config += `	capabilities_sac_fec129_disable = true` + "\n"
	config += `	mldp_logging_notifications = true` + "\n"
	config += `	mldp_address_families = [{` + "\n"
	config += `		name = "ipv4"` + "\n"
	config += `		make_before_break_delay = 30` + "\n"
	config += `		forwarding_recursive = true` + "\n"
	config += `		forwarding_recursive_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `		recursive_fec = true` + "\n"
	config += `		}]` + "\n"
	config += `	session_protection = true` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
