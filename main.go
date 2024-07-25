package main

import (
    "cn.eziosweet/mcupdatergo/model"
    "fmt"
    "github.com/alecthomas/kingpin/v2"
    "github.com/go-zoox/fetch"
)

var (
	remote = kingpin.Flag("remote","使用远程配置文件").Short('r').String()
	local = kingpin.Flag("local","使用本地配置文件").Short('l').String()
)

func main() {
	kingpin.Parse()
	if (*remote == "") && (*local == ""){
		fmt.Println("请提供远程配置文件地址或者本地配置文件地址")
	}else if (*remote != "") && (*local != ""){
		fmt.Println("仅能提供远程配置文件地址和本地配置文件地址中的一个")
	}else if *local == "" {
		fmt.Printf("使用远程配置文件地址 %s\n",*remote)
		getRemoteConfig(*remote)
	}else {
		fmt.Printf("使用远程配置文件地址 %s\n",*local)
	}
}


func getRemoteConfig(url string){
    resp, err := fetch.Get(url)
    if err != nil {
		panic(err)
    }
	config := model.ConfigModel{}
    err = resp.UnmarshalYAML(&config)
    if err != nil {
        panic(err) 
    }
    for _, element := range config.Modrinth {
		if element.Version=="" {
			modrinthUrl := fmt.Sprintf(
				"https://api.modrinth.com/v2/project/%s/version?game_versions=[\"%s\"]&loaders=[\"fabric\"]",
				element.Name,config.McVersion)
			resp ,err = fetch.Get(modrinthUrl)
			modrinth := new([]model.ModrinthModel)
            err = resp.UnmarshalJSON(modrinth)
            if err != nil {
                panic(err) 
            }
			resp,err = fetch.Download((*modrinth)[0].Files[0].Url,"./model/"+(*modrinth)[0].Files[0].FileName)
			if err != nil {
				
			}
		}else{
			
		}
    }
}