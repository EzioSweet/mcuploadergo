package main

import (
    "cn.eziosweet/mcupdatergo/model"
    "fmt"
    "github.com/alecthomas/kingpin/v2"
    "github.com/go-zoox/fetch"
    "os"
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
		fmt.Printf("使用本地配置文件地址 %s\n",*local)
		fmt.Println("还未实现，敬请期待")
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
	if config.Loader==""{
		fmt.Println("请配置mod加载器类型")
		return
	}
	fmt.Println()
	fmt.Println("1.开始Modrinth Mod下载")
	loader:=config.Loader
    for i, element := range config.Modrinth {
		loader = config.Loader

		if element.Loader != ""{
			loader = element.Loader
		}
		modrinthUrl := fmt.Sprintf(
			"https://api.modrinth.com/v2/project/%s/version?game_versions=[\"%s\"]&loaders=[\"%s\"]",
			element.Name,config.McVersion,loader)
		resp ,err = fetch.Get(modrinthUrl)
		modrinth := *new([]model.ModrinthModel)
		err = resp.UnmarshalJSON(&modrinth)
		if err != nil {
			panic(err)
		}
		if element.Version == ""{
			fmt.Printf("\r[%d/%d] Downloading %s",i+1, len(config.Modrinth),(modrinth)[0].Files[0].FileName)
			if element.Path!=""{
				resp,err = fetch.Download(modrinth[0].Files[0].Url,element.Path)
			}else{
				resp,err = fetch.Download(modrinth[0].Files[0].Url,"./mods/"+modrinth[0].Files[0].FileName)
			}
			if err != nil {
				panic(err)
			}
		}else{
			for _,item :=range modrinth{
                if item.Version == element.Version {
					fmt.Printf("\r[%d/%d] Downloading %s",i+1, len(config.Modrinth),item.Files[0].FileName)
					if element.Path!=""{
						resp,err = fetch.Download(item.Files[0].Url,element.Path)
					}else{
						resp,err = fetch.Download(item.Files[0].Url,"./mods/"+item.Files[0].FileName)
					}
					break
                }
			}
		}
	}
	fmt.Println()
	fmt.Println("OK")
	fmt.Println()
	fmt.Println("2.开始Local Mod或者Config下载")
	for _,item := range config.Local{
		for i,element := range item.List{
			fmt.Printf("\r[%d/%d] Downloading %s",i+1, len(item.List),element.Url)
			resp,err = fetch.Download(item.Prefix + element.Url,element.Path)
			if err !=nil {
				panic(err)
			}
		}
	}
	fmt.Println()
	fmt.Println("OK")
	fmt.Println()
	fmt.Println("3.删除无用Mod或者Config")

	for i,element :=range config.Remove{
		fmt.Printf("\r[%d/%d] Removing %s",i+1, len(config.Remove),element)
        err = os.Remove(element)
        if err != nil {
            panic(err)
        }
	}
	fmt.Println()
	fmt.Println("OK")
	fmt.Println()
	fmt.Println("运行完成")
}