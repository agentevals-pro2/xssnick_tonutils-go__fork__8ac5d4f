package tlb

import (
	"testing"
)

func TestAccStatusChangeFrozen_ToCell(t *testing.T) {
	st := AccStatusChange{Type: AccStatusChangeFrozen}
	c, err := st.ToCell()
	if err != nil {
		t.Fatal(err)
	}
	val, err := c.BeginParse().LoadUInt(2)
	if err != nil {
		t.Fatal(err)
	}
	if val != 0b10 {
		t.Fatalf("unexpected bits %b", val)
	}

	var decoded AccStatusChange
	if err := decoded.LoadFromCell(c.BeginParse()); err != nil {
		t.Fatal(err)
	}
	if decoded.Type != AccStatusChangeFrozen {
		t.Fatalf("decoded type %s", decoded.Type)
	}
}
