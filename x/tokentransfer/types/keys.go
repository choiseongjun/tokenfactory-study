package types

const (
	// ModuleName defines the module name
	ModuleName = "tokentransfer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokentransfer"
)

var (
	ParamsKey = []byte("p_tokentransfer")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
