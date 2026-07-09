// Package services is the Services content pack for togo: Service resource with features & icons and slug detail pages.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package services

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("services", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
