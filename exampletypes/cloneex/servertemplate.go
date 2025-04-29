package cloneex

import (
	"github.com/fengdotdev/golibs-traits/exampletypes/handyfuncs"
	"github.com/fengdotdev/golibs-traits/trait"
)

type SSlConfig struct {
	SSl     bool
	SSlCert string
	SSlKey  string
}

type ServerConfigTemplate struct {
	instanceID    string
	port          int
	ssl           SSlConfig
	assetsFolders []string
	assetsUrl     string
}

var _ = trait.Cloneable[ServerConfigTemplate](&ServerConfigTemplate{})

// Clone implements trait.Cloneable.
// It creates a deep copy of the ServerConfigTemplate instance. if instanceID is not empty, it generates a new ID for the clone.
func (s *ServerConfigTemplate) Clone() ServerConfigTemplate {

	id := ""

	if s.instanceID != "" {
		id = handyfuncs.IdGen() //// only relevant for identification not for behavior
	}

	return ServerConfigTemplate{
		instanceID:    id,
		port:          s.port,
		ssl:           s.ssl,
		assetsFolders: s.assetsFolders,
		assetsUrl:     s.assetsUrl,
	}
}

// contructors

// NewServerConfigDefault creates a new ServerConfigTemplate with default values and no ID ""
func NewServerConfigDefault() ServerConfigTemplate {
	return ServerConfigTemplate{
		instanceID: "", // only relevant for identification not for behavior
		// default values
		port:          8080,
		ssl:           SSlConfig{SSl: false, SSlCert: "", SSlKey: ""},
		assetsFolders: []string{"/assets"},
		assetsUrl:     "/assets",
	}
}

// NewServerConfigDefaultWithID creates a new ServerConfigTemplate with default values and a generated ID
func NewServerConfigDefaultWithID() ServerConfigTemplate {

	config := NewServerConfigDefault()

	config.instanceID = handyfuncs.IdGen()

	return config
}

// getters and setters
func (s *ServerConfigTemplate) GetInstanceID() string {
	return s.instanceID
}

func (s *ServerConfigTemplate) GetPort() int {
	return s.port
}
func (s *ServerConfigTemplate) SetPort(port int) {
	s.port = port
}

func (s *ServerConfigTemplate) GetSSl() SSlConfig {
	return s.ssl
}

func (s *ServerConfigTemplate) SetSSl(ssl SSlConfig) {
	s.ssl = ssl
}

func (s *ServerConfigTemplate) GetAssetsFolders() []string {
	return s.assetsFolders
}

func (s *ServerConfigTemplate) SetAssetsFolders(assetsFolders ...string) {
	s.assetsFolders = assetsFolders
}

func (s *ServerConfigTemplate) GetAssetsUrl() string {
	return s.assetsUrl
}
func (s *ServerConfigTemplate) SetAssetsUrl(assetsUrl string) {
	s.assetsUrl = assetsUrl
}
func (s *ServerConfigTemplate) String() string {
	return "ServerConfigTemplate Instance"
}
