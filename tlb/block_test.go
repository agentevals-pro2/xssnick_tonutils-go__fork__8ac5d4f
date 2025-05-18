package tlb

import (
	"github.com/xssnick/tonutils-go/tvm/cell"
	"os"
	"testing"
)

func TestBlockMaster(t *testing.T) {
	boc, err := os.ReadFile("tlb/testdata/master.boc")
	if err != nil {
		t.Fatal(err)
	}
	c, _ := cell.FromBOC(boc)

	var block Block
	if err := LoadFromCell(&block, c.BeginParse()); err != nil {
		t.Fatal(err)
	}

	parents, err := block.BlockInfo.GetParentBlocks()
	if err != nil {
		t.Fatal(err)
	}
	if len(parents) == 0 {
		t.Fatalf("parsed zero parents")
	}
}

func TestBlockNotMaster(t *testing.T) {
	boc, err := os.ReadFile("tlb/testdata/not_master.boc")
	if err != nil {
		t.Fatal(err)
	}
	c, _ := cell.FromBOC(boc)

	var block Block
	if err := LoadFromCell(&block, c.BeginParse()); err != nil {
		t.Fatal(err)
	}

	parents, err := block.BlockInfo.GetParentBlocks()
	if err != nil {
		t.Fatal(err)
	}

	if len(parents) == 0 {
		t.Fatalf("parsed zero parents")
	}
}
