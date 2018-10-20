package main

import (
	"log"
	"os"
)

// How we represent sequences occurring in the FASTA alignment
type Sequence struct {

	header string       // The Header line of the FASTA file
	rawSequence string  // The raw data of the sequence, containing gaps
	lenNoGaps int       // The length of the sequence if all gaps were removed
}

func calculateSequenceIdentity(seq1 *Sequence, seq2 *Sequence) float64 {

	totalGaps := 0
	identicalChars := 0

	// We need to compare raw sequenes
	raw1 := seq1.rawSequence
	raw2 := seq2.rawSequence

	lenRaw1 := len(raw1)
	lenRaw2 := len(raw2)

	if lenRaw2!= lenRaw1 {

		log.Fatalf("Not an alignment. Conflicting sequence lengths %d and %d", lenRaw1, lenRaw2)
	}

	for i := range raw1 {

		char1 := raw1[i]
		char2 := raw2[i]

		// Gapped
		if char1 == '-' && char2 == '-' {

			totalGaps++
		} else if char1 == char2 {

			identicalChars++
		}
	}
	return 100.0 * (float64(identicalChars) / float64(lenRaw1 - totalGaps))
}

// Writes the sequence as FASTA to the selected FilePath
func (seq *Sequence) writeTo(filePath string) {

	const width = 80

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	defer file.Close()

	file.WriteString(seq.header + newline)

	raw := seq.rawSequence
	limit := width
	for ; limit < len(seq.rawSequence); limit += width {

		file.WriteString(raw[limit-width:width] + newline)
	}
	file.WriteString(raw[limit-width:] + newline)
}
