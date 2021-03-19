// Package strand translates sequences from dna to rna
package strand

var translation = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA returns the rna sequence that matches the given dna sequence
func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		rna += translation[string(dna[i])]
	}
	return rna
}
