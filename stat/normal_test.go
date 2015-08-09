package stat

import (
	"fmt"
	"testing"
)

func TestNormalInv_CDF(t *testing.T) {
	//Test p=0.975, sigma=1.0
	x := NormalInv_CDF(0.975, 1.0)
	if (x < 1.95996388) || (x > 1.95996408) {
		fmt.Printf("NORMALINV_CDF(0.975,1.0) GOT %1.8f EXPECTED 1.95996398\n", x)
		t.FailNow()
	}
	//Test p=0.20, sigma=3600.0
	x = NormalInv_CDF(0.20, 3600.0)
	if (x < -3029.83644096) || (x > -3029.83644076) {
		fmt.Printf("NORMALINV_CDF(0.20,3600.0) GOT %1.8f EXPECTED -3029.83644086\n", x)
		t.FailNow()
	}
}

func BenchmarkNormalInv_CDF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NormalInv_CDF(0.20, 3600.0)
	}
}
