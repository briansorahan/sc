package sc

// FIXME
// import (
// 	. "github.com/scgolang/sc/types"
// 	. "github.com/scgolang/sc/ugens"
// 	"testing"
// )

// func TestBRF(t *testing.T) {
// 	def := NewSynthdef("BRFExample", func(p Params) Ugen {
// 		line := XLine{C(0.7), C(300), C(20), 0}.Rate(KR)
// 		saw := Saw{C(200)}.Rate(AR).Mul(C(0.5))
// 		sine := FSinOsc{line, C(0)}.Rate(KR).MulAdd(C(3800), C(4000))
// 		bpf := BRF{saw, sine, C(0.3)}.Rate(AR)
// 		return Out{C(0), bpf}.Rate(AR)
// 	})
// 	same, err := def.Compare(`{
// 		var line = XLine.kr(0.7, 300, 20);
// 		var saw = Saw.ar(200, 0.5);
// 		var sine = FSinOsc.kr(line, 0, 3800, 4000);
// 		Out.ar(0, BRF.ar(saw, sine, 0.3));
// 	}`)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !same {
// 		t.Fatalf("synthdef is different from sclang version")
// 	}
// }
