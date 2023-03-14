package utils

import (
	"fmt"
	"net/url"
)

var serviceList = map[string]string{
	"twitter": "twitter.com",
	"youtube": "www.youtube.com",
}

func CheckURL(serviceName string, serviceUrl string) error {
	// string形式のURLをnet/url形式に変換
	u, err := url.Parse(serviceUrl)
	if err != nil {
		return err
	}
	// URLのスキームを抽出
	if s := u.Scheme; s != "https" {
		return fmt.Errorf("scheme is not https")
	}
	// URLからHostを抽出
	host := u.Host
	// hostがserviceListのものと一致するか確認
	if s, ok := serviceList[serviceName]; ok && s == host {
		// 一致した場合
		return nil
	} else if s != host {
		// hostが一致しなかった場合
		return fmt.Errorf("host name is unmatched with %s", serviceUrl)
	}

	// Serviceが見つからなかった場合
	return fmt.Errorf("not found service with %s", serviceName)
}
