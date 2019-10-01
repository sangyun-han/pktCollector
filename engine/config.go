package engine

type InterfacesConfig struct {
	Device			string	`config:"device"`
	Type			string	`config:"type"`
	File			string	`config:"file"`
	BpfFilter		bool	`config:"bpf_filter"`
	SnapLen			int		`config:"snaplen"`
	BufferSizeMb	int
	DefaultOpt		bool
}

type TsdbConfig struct {

}