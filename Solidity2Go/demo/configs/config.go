package configs

import (
	"path/filepath"
	"os"
	"github.com/BurntSushi/toml"
)

// Config 应用程序配置结构
type Config struct {
	Database struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Name     string `toml:"dbname"`
	} `toml:"Database"`

	Blockchain struct {
		NodeURL       string `toml:"node_url"`
		GroupID       int    `toml:"group_id"`
		ChainID       int    `toml:"chain_id"`
		ContractAddr  string `toml:"contract_address"`
		PrivateKey    string `toml:"private_key"`
		CAFile        string `toml:"ca_file"`
		KeyFile       string `toml:"key_file"`
		CertFile      string `toml:"cert_file"`
	} `toml:"Blockchain"`

	Account struct {
		PrivateKeyFiles map[string]string `toml:"-"`
	} `toml:"-"`

	Services struct {
		SpringBootURL string `toml:"-"`
	} `toml:"-"`
}

// LoadConfig 从指定的 TOML 文件中加载配置
func LoadConfig(filename string) (Config, error) {
	var config Config

	// 解析 TOML 配置文件
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return config, err
	}

	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	// 确保证书文件路径是相对于工作目录的
	config.Blockchain.CAFile = filepath.Join(wd, config.Blockchain.CAFile)
	config.Blockchain.KeyFile = filepath.Join(wd, config.Blockchain.KeyFile)
	config.Blockchain.CertFile = filepath.Join(wd, config.Blockchain.CertFile)

	// 加载私钥文件路径映射
	privateKeyFiles := make(map[string]string)
	privateKeyFiles["Consumer"] = filepath.Join(wd, "account/Consumer_key_0x0b35f558d3ad35f5fc7a430a784a3bc06c8cff36.pem")
	privateKeyFiles["Grower"] = filepath.Join(wd, "account/Grower_key_0x3b6267e282741092a740382ce0f2e89c6c1e2fab.pem")
	privateKeyFiles["Logistic"] = filepath.Join(wd, "account/Logistic_key_0x7cc9da5081981a21107a20e30da23d1529ce266f.pem")
	privateKeyFiles["Factory"] = filepath.Join(wd, "account/Factory_key_0xdb8b35c5f720b6aaf5ff80c79cf9442c1038a947.pem")
	privateKeyFiles["QualityExpert"] = filepath.Join(wd, "account/Expert_key_0x762267ee97bb2f1700cdff8bb849e9ab3b17bd8d.pem")

	config.Account.PrivateKeyFiles = privateKeyFiles

	// 设置默认服务URL（可以根据需要修改）
	config.Services.SpringBootURL = "http://localhost:8081/api/detect" // 默认Spring Boot服务地址

	return config, nil
}