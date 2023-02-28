package profile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/usememos/memos/server/version"
)

// Profile is the configuration to start main server.
type Profile struct {
	// Mode can be "prod" or "dev" or "demo"
	Mode string `json:"mode"`
	// Port is the binding port for server
	Port int `json:"-"`
	// Data is the data directory
	Data string `json:"-"`
	// DBDriver is the db type sqlite/mysql/postgresql
	DBDriver string `json:"-"`
	// DSN points to where Memos stores its own data
	DSN string `json:"-"`
	// Version is the current version of server
	Version string `json:"version"`
}

func checkDSN(dataDir string) (string, error) {
	// Convert to absolute path if relative path is supplied.
	if !filepath.IsAbs(dataDir) {
		absDir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/" + dataDir)
		if err != nil {
			return "", err
		}
		dataDir = absDir
	}

	// Trim trailing / in case user supplies
	dataDir = strings.TrimRight(dataDir, "/")

	if _, err := os.Stat(dataDir); err != nil {
		return "", fmt.Errorf("unable to access data folder %s, err %w", dataDir, err)
	}

	return dataDir, nil
}

// GetProfile will return a profile for dev or prod.
func GetProfile() (*Profile, error) {
	profile := Profile{}
	err := viper.Unmarshal(&profile)
	if err != nil {
		return nil, err
	}

	if profile.Mode != "demo" && profile.Mode != "dev" && profile.Mode != "prod" {
		profile.Mode = "demo"
	}

	if profile.Mode == "prod" && profile.Data == "" {
		profile.Data = "/var/opt/memos"
	}

	dataDir, err := checkDSN(profile.Data)
	if err != nil {
		fmt.Printf("Failed to check dsn: %s, err: %+v\n", dataDir, err)
		return nil, err
	}

	profile.Data = dataDir
	//profile.DBDriver = "sqlite"
	//profile.DSN = fmt.Sprintf("%s/memos_%s.db", dataDir, profile.Mode)
	profile.DBDriver = "mysql"
	profile.DSN = fmt.Sprintf("root:tonghs@tcp(127.0.0.1:13306)/memos_%s?charset=utf8mb4&parseTime=True&loc=Local", profile.Mode)
	profile.Version = version.GetCurrentVersion(profile.Mode)

	return &profile, nil
}
