package clients

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
)

// 创建配置相关的客户端
func CreateConfigClient(properties map[string]interface{}) (iClient config_client.IConfigClient,
	err error) {
	return
}

// 创建服务发现相关的客户端
func CreateServiceClient(properties map[string]interface{}) (iClient naming_client.INamingClient, err error) {
	nacosClient, errSetConfig := setConfig(properties)
	if errSetConfig != nil {
		err = errSetConfig
		return
	}
	nacosClient.SetHttpAgent(&http_agent.HttpAgent{})
	client, errNew := naming_client.NewNamingClient(nacosClient)
	if errNew != nil {
		err = errNew
		return
	}
	iClient = &client
	return
}

func setConfig(properties map[string]interface{}) (iClient nacos_client.INacosClient, err error) {
	client := nacos_client.NacosClient{}
	if clientConfigTmp, exist := properties[constant.KEY_CLIENT_CONFIG]; exist {
		if clientConfig, ok := clientConfigTmp.(constant.ClientConfig); ok {
			err = client.SetClientConfig(clientConfig)
			if err != nil {
				return
			}
		}
	} else {
		_ = client.SetClientConfig(constant.ClientConfig{
			TimeoutMs:      30 * 1000,
			ListenInterval: 10 * 1000,
			BeatInterval:   5 * 1000,
		})
	}
	// 设置 serverConfig
	if serverConfigTmp, exist := properties[constant.KEY_SERVER_CONFIGS]; exist {
		if serverConfigs, ok := serverConfigTmp.([]constant.ServerConfig); ok {
			err = client.SetServerConfig(serverConfigs)
			if err != nil {
				return
			}
		}
	} else {
		err = errors.New("server configs not found in properties")
		return
	}

	iClient = &client

	return
}
