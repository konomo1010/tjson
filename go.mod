module tjson

go 1.15

require (
	cseYaml2Json v0.0.0-00010101000000-000000000000
	github.com/lxn/walk v0.0.0-20210112085537-c389da54e794
	github.com/lxn/win v0.0.0-20201111105847-2a20daff6a55 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace cseYaml2Json => D:/projects/tjson/cseYaml2Json
replace cseYaml2Json => ./cseYaml2Json
