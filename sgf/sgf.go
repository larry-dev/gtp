package sgf

const (
	pat_seq   = `\(|\)|(;(\s*[A-Z]+\s*((\[\])|(\[(.|\s)*?([^\\]\])))+)*)`
	pat_node  = `[A-Z]+\s*((\[\])|(\[(.|\s)*?([^\\]\])))+`
	pat_ident = `[A-Z]+`
	pat_props = `(\[\])|(\[(.|\s)*?([^\\]\]))`
	B         = 1
	W         = -1
	Empty     = 0
	NOTHING   = -2
)

var InfoList = [...]string{"black", "white", "AN", "CP", "DT", "EV", "GN", "GC", "HA", "ON", "OT", "RE", "RO", "RU", "SO", "TM", "US", "PC", "KM",}
