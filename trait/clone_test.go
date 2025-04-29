package trait_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/cloneex"
	"github.com/fengdotdev/golibs-traits/trait"
)

type ServerConfigTemplate = cloneex.ServerConfigTemplate

type ClonableConfig = trait.Cloneable[ServerConfigTemplate]

func TestClonability(t *testing.T) {

	var _ trait.Cloneable[ServerConfigTemplate] = &ServerConfigTemplate{}

	var _ ClonableConfig = &ServerConfigTemplate{}

}

func TestConfigClone_identical(t *testing.T) {

	t.Run("Equality", func(t *testing.T) {

		original := cloneex.NewServerConfigDefault()

		clone := original.Clone()

		// Assert - Verify content equality

		eq := reflect.DeepEqual(original, clone)

		assert.TrueWithMessage(t, eq, fmt.Sprintf("expected %v and %v to be equal", original, clone))

		// Verify that they are different objects in memory
		assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

	})

	t.Run("Check Change", func(t *testing.T) {

		original := cloneex.NewServerConfigDefault()

		clone := original.Clone()

		// Change the clone
		clone.SetPort(9090)

		newssl := cloneex.SSlConfig{SSl: true, SSlCert: "cert.pem", SSlKey: "key.pem"}
		clone.SetSSl(newssl)
		clone.SetAssetsFolders("/newassets", "/newassets2")
		clone.SetAssetsUrl("/newassetsurl")

		// Assert - Verify content equality
		eq := reflect.DeepEqual(original, clone)
		assert.FalseWithMessage(t, eq, fmt.Sprintf("expected %v and %v to be different", original, clone))

		// Verify that they are different objects in memory
		assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

		// Verify that the original is unchanged

		//port
		assert.EqualWithMessage(t, original.GetPort(), 8080, fmt.Sprintf("expected %v to be unchanged", original.GetPort()))

		// ssl
		eqssl := reflect.DeepEqual(original.GetSSl(), cloneex.SSlConfig{SSl: false, SSlCert: "", SSlKey: ""})
		assert.TrueWithMessage(t, eqssl, fmt.Sprintf("expected %v to be unchanged", original.GetSSl()))

		// Verify that the clone is changed

		// assets

		lenoriginal := len(original.GetAssetsFolders())
		assert.EqualWithMessage(t, lenoriginal, 1, fmt.Sprintf("expected %v to be unchanged", original.GetAssetsFolders()))

		lenclone := len(clone.GetAssetsFolders())
		assert.EqualWithMessage(t, lenclone, 2, fmt.Sprintf("expected %v to be changed", clone.GetAssetsFolders()))

		// port
		assert.EqualWithMessage(t, clone.GetPort(), 9090, fmt.Sprintf("expected %v to be changed", clone.GetPort()))
		//ssl
		eqssl = reflect.DeepEqual(clone.GetSSl(), newssl)
		assert.TrueWithMessage(t, eqssl, fmt.Sprintf("expected %v to be changed", clone.GetSSl()))

	})

}



func TestConfigClone_Distinct(t *testing.T) {
	t.Run("Distinct", func(t *testing.T) {

		originalWithID := cloneex.NewServerConfigDefaultWithID()
		cloneWithID := originalWithID.Clone()
		// Assert - Verify content equality

		// id
		assert.NotEqualWithMessage(t, originalWithID.GetInstanceID(), cloneWithID.GetInstanceID(), "expected different instance IDs")

		eq := reflect.DeepEqual(originalWithID, cloneWithID)

		assert.FalseWithMessage(t, eq, "expected deep equal was false")
		
		
		//else
		assert.EqualWithMessage(t, originalWithID.GetPort(), cloneWithID.GetPort(), "expected port to be equal")

		// ssl
		eqssl := reflect.DeepEqual(originalWithID.GetSSl(), cloneWithID.GetSSl())
		assert.TrueWithMessage(t, eqssl, "expected deep equal ssl was true")
		// assets
		lenoriginal := len(originalWithID.GetAssetsFolders())
		lenclone := len(cloneWithID.GetAssetsFolders())
		assert.EqualWithMessage(t, lenoriginal, lenclone, "expected equal length of assets folders")

		// Verify that they are different objects in memory

		assert.NotEqualWithMessage(t, &originalWithID, &cloneWithID, "expected different memory addresses")

	})


}