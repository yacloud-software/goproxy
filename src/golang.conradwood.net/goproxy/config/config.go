package config

import (
	"bytes"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/utils"
	"gopkg.in/yaml.v3"
	"sync"
	"time"
)

var (
	default_config = &pb.Config{
		GoGetProxy: "http://172.29.1.11:14231/",
		LocalHosts: []string{
			"golang.conradwood.net",
			"golang.singingcat.net",
			"userprotos.singingcat.net",
			"golang.yacloud.eu",
			"git.conradwood.net",
			"git.singingcat.net",
			"git.yacloud.eu",
		},
		ArtefactResolvers: []*pb.ArtefactDef{
			&pb.ArtefactDef{
				Path:       "golang.conradwood.net/go-easyops",
				ArtefactID: 24,
				Domain:     "conradwood.net",
				Name:       "go-easyops"},
			&pb.ArtefactDef{
				Path:       "golang.singingcat.net/scgolib",
				ArtefactID: 193,
				Domain:     "singingcat.net",
				Name:       "scgolib"},
		},
	}
	config      *pb.Config
	config_lock sync.Mutex
)

func init() {
	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		load_config()
	}()
}
func GetConfig() *pb.Config {
	if config != nil {
		return config
	}
	load_config()
	return config
}
func load_config() {
	config_lock.Lock()
	defer config_lock.Unlock()
	if config != nil {
		return
	}
	fname := "extras/config.yaml"
	b, err := utils.ReadFile(fname)
	if err == nil {
		new_config := &pb.Config{}
		err = yaml.Unmarshal(b, new_config)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal config from %s: %s\n", fname, err))
		}
		config = new_config
		return
	}
	fmt.Printf("Failed to read config file %s: %s, using default config\n", fname, err)
	config = default_config

	//print out config to give user an idea of what needs to be configured
	var buf bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buf)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(config)
	b = buf.Bytes()
	if err != nil {
		panic(fmt.Sprintf("failed to marshal default config (%s)\n", err))
	}
	fmt.Printf("Default Config:\n%s\n", string(b))
	return
}
