package stack

import (
	"testing"

	dgl "github.com/CrimsonSarah/cto/pkg/server-common/card/digimon"
	dtl "github.com/CrimsonSarah/cto/pkg/server-common/card/digitama"
)

const (
	effect byte = iota
	inheritedEffect
	security
)

func TestStack(t *testing.T) {
	stackteste := Stack{}
	Trigger(dgl.Agumon, &stackteste)
	Trigger(dtl.Koromon, &stackteste)

	Resolve(&stackteste, effect)

	Resolve(&stackteste, inheritedEffect)
}
