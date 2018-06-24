package dual

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	//myCredit := 1
	var config = getConfig()
	fmt.Println("result is :", config.Priamy)
}
