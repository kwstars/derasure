package repositories

import (
	"context"
	"flag"
	"testing"
)

var configFile = flag.String("f", "../../../../configs/derasure/config.yaml", "set config file which viper will loading.")

func TestErasureRepostiory_DelKey(t *testing.T) {
	flag.Parse()
	sto, cf, err := CreateErasureRepository(*configFile)
	if err != nil {
		t.Fatalf("CreateErasureRepository error,%+v", err)
	}
	defer cf()

	tests := []struct {
		name     string
		id       string
		expected bool
	}{
		{"删除111111", "111111", true},
		{"删除222222", "222222", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := sto.DelKey(context.Background(), test.id)
			if err != nil {
				t.Fatalf("product service get proudct error,%+v", err)
			}
		})
	}

}
