package domain

type PushCode struct {
	PkgId    string `json:"pkg_id"`
	PkgName  string `json:"pkg_name"`
	Platform string `json:"platform"`
	Importer Importer
}

type Importer struct {
	Name  string
	Email string
}
