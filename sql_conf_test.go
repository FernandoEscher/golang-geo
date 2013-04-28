package geo

import (
	_ "fmt"
	"testing"
)

func TestGetSQLConf(t *testing.T) {
	// GetSQLConf should return the DefaultSQLConf in the abscense of a configuration file.
	confFromEmptyPath, _ := GetSQLConf("")

	if confFromEmptyPath != DefaultSQLConf {
		t.Error("SQL Configuration was expected to b the Default SQL Configuration.")
	}

	confFromFile, _ := GetSQLConf(GOLANG_GEO_CONFIG_PATH)
	if confFromFile == DefaultSQLConf {
		t.Error("SQL Configuration was expected to be different than the Default SQL Configuration.")
	}
}
