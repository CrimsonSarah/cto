package stack

import (
	"testing"

	dgl "github.com/CrimsonSarah/cto/pkg/server-common/card/digimon"
	dtl "github.com/CrimsonSarah/cto/pkg/server-common/card/digitama"
)

const (
	onplay byte = iota
	effect
	inheritedEffect
	security
)

func TestStack(t *testing.T) {
	stackteste := Stack{}
	Trigger(dgl.Agumon, &stackteste, effect)
	Trigger(dtl.Koromon, &stackteste, inheritedEffect)

	Resolve(&stackteste)

	Resolve(&stackteste)
}
