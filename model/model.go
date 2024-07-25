package model

type ConfigModel struct {
    Version string `yaml:"version"`
	McVersion string `yaml:"mcVersion"`
	Loader string `yaml:"loader"`
	Modrinth []struct{
		Name string `yaml:"name"`
		Path string `yaml:"path"`
		Loader string `yaml:"loader"`
		Version string `yaml:"version"`
	} `yaml:"modrinth"`
	Local []struct{
		Prefix string `yaml:"prefix"`
		List []struct{
			Url string `yaml:"url"`
			Path string `yaml:"path"`
		} `yaml:"list"`
	}`yaml:"local"`
	Remove []string `yaml:"remove"`
	
}

type ModrinthModel struct {
	Version string `json:"version_number"`
    Files []struct{
		Url string `json:"url"`
		FileName string `json:"filename"`
	} `json:"files"`
}