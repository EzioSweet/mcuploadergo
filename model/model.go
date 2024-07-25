package model

type ConfigModel struct {
    Version string `yaml:"version"`
	McVersion string `yaml:"mcVersion"`
	Modrinth []struct{
		Name string `yaml:"name"`
		Path string `yaml:"path"`
		Version string `yaml:"version"`
	} `yaml:"modrinth"`
	Local []struct{
		Url string `yaml:"url"`
		Path string `yaml:"path"`
	}
	Remove []string `yaml:"remove"`
	
}

type ModrinthModel struct {
	Version string `json:"version_number"`
    Files []struct{
		Url string `json:"url"`
		FileName string `json:"filename"`
	} `json:"files"`
}