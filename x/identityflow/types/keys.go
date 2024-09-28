package types

const (
	// ModuleName defines the module name
	ModuleName = "identityflow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_identityflow"
)

var (
	ParamsKey = []byte("p_identityflow")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
