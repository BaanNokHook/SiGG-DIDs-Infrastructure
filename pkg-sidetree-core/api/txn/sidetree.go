/*
SiGG-DIDs-Infrastructure
*/

package txn

// SidetreeTxn defines info about sidetree transaction.
type SidetreeTxn struct {
	TransactionTime      uint64
	TransactionNumber    uint64
	AnchorString         string
	Namespace            string
	ProtocolVersion      uint64
	CanonicalReference   string
	EquivalentReferences []string
	AlternateSources     []string
}
