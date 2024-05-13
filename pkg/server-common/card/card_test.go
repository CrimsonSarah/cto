package card_test

import (
	"fmt"
	"testing"

	dgl "github.com/CrimsonSarah/cto/pkg/server-common/card/digimon"
	dtl "github.com/CrimsonSarah/cto/pkg/server-common/card/digitama"
)

func TestDigimons(t *testing.T) {
	fmt.Println(dtl.Koromon)
}

func TestDigievolucao(t *testing.T) {
	dtl.Koromon.IsInherited = true
	dgl.Agumon.EvolutionLine = append(dgl.Agumon.EvolutionLine, dtl.Koromon)
	fmt.Println(dtl.Koromon)
	fmt.Println(dgl.Agumon)

	dgl.Agumon.IsInherited = true
	dgl.Greymon.EvolutionLine = append(dgl.Greymon.EvolutionLine, dgl.Agumon, dgl.Agumon.EvolutionLine[0])
	fmt.Println(dgl.Agumon)
	fmt.Println(dgl.Greymon)
}
