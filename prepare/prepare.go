package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const usage = `
=========================================================
Usage: inputFile minSequenceIdentity (percent) outputFile
=========================================================
inputfile : FastA alignment file
minSequenceIdentity: minimum sequence identity below which a new psipred run is invoked
outputFile : mainlog filename
`

func checkAlignment(sequences []*Sequence) {

	if n := len(sequences); n > 1 {

		len1 := len(sequences[n-1].rawSequence)
		len2 := len(sequences[n-2].rawSequence)

		if len1 != len2 {

			log.Fatalf("Not an alignment. Conflicting sequence length %d and %d", len1, len2)
		}
	}
}

func parse(inputPath string) []*Sequence {

	file, err := os.Open(inputPath)
	check(err)
	defer file.Close()
	
	// The Result slice
	var sequences []*Sequence
	var sequence *Sequence = nil
	var seqBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Get line
		line := strings.TrimSpace(scanner.Text())

		// Switch Sequence
		if strings.HasPrefix(line, ">")  {

			// Commit the last sequence
			if sequence != nil {
				sequence.rawSequence = seqBuilder.String()
				seqBuilder.Reset()
				sequences = append(sequences, sequence)
				checkAlignment(sequences)
			}
			sequence = &Sequence{header: line, lenNoGaps: 0}

		} else {
			seqBuilder.WriteString(line)
			sequence.lenNoGaps += len(strings.Replace(line, "-", "", -1))
		}
	}

	sequence.rawSequence = seqBuilder.String()
	sequences = append(sequences, sequence)
	checkAlignment(sequences)
	return sequences
}

func checkMinSeqId(value float64) {

	// Ensure correct range of minSeqId
	if value < 0 || value > 100 {

		log.Fatal("FATAL: Sequence Identity must be between 0 and 100")
	}

	if value <= 1 {

		log.Println("WARNING: Provided value for sequence identity is below 1. Note that the value has to be provided in percent")
	}
}


func orderedIndexVector(sequences []*Sequence) []int {

	var ordered []int
	for i := 0; i < len(sequences); i++ {

		ordered = append(ordered, i)
	}
	sort.Slice(ordered, func(i, j int) bool {

		return sequences[ordered[i]].lenNoGaps > sequences[ordered[j]].lenNoGaps
	})
	return ordered
}


type Output struct {

	SequenceFile string `json:"sequenceFile"`
	CoveredBy *string `json:"coveredBy"`
	SequenceIdentity *float64 `json:"sequenceIdentity"`
}


func main() {

	if len(os.Args) < 4 {

		log.Fatalf(usage)
	}

	// Command line parsing
	fastaFile := os.Args[1]
	minSeqId, err := strconv.ParseFloat(os.Args[2], 64)
	check(err)
	checkMinSeqId(minSeqId)
	outputFile := os.Args[3]

	// Parse sequences
	sequences := parse(fastaFile)
	if len(sequences) == 0 {
		log.Fatal("FASTA file appears to be empty! Cannot continue")
	}

	// Init index vector that sorts the sequences by no-gap length in descending order
	ordered := orderedIndexVector(sequences)

	// Calculate the sequence cover set, start with the 'largest' sequence
	covered := []int {ordered[0]}
	seqAssignment := make(map[int]int)
	idAssignment := make(map[int]float64)

	for i := 1; i < len(ordered); i++ {

		// Index of the sequence in the sequences slice
		seqIndex1 := ordered[i]

		// Calculate the sequence identity to each sequence in the cover set and try to maximize
		firstSequence := sequences[seqIndex1]
		maxSeqId := 0.0
		argMax := 0

		// Maximize the sequence identity
		for j := 0; j < len(covered); j++ {

			seqIndex2 := covered[j]
			secondSequence := sequences[seqIndex2]
			seqId := calculateSequenceIdentity(firstSequence, secondSequence)

			// New max found
			if seqId > maxSeqId {

				maxSeqId = seqId
				argMax = seqIndex2
			}
		}

		// Max Seq Id is too low, this sequence needs to be covered
		if maxSeqId < minSeqId {

			covered = append(covered, seqIndex1)
		} else {

			seqAssignment[seqIndex1] = argMax
			idAssignment[seqIndex1] = maxSeqId
		}
	}
	var outputs []Output
	for i, seq := range sequences {

		filePath1, err := filepath.Abs(fastaFile + "." + strconv.Itoa(i) + ".fas")
		check(err)
		seq.writeTo(filePath1)

		output := Output{SequenceFile: filePath1}
		if coverer, ok := seqAssignment[i]; ok {

			filePath2, err := filepath.Abs(fastaFile + "." + strconv.Itoa(coverer) + ".fas")
			check(err)
			output.CoveredBy = &filePath2
			seqId := idAssignment[i]
			output.SequenceIdentity = &seqId

		}
		outputs = append(outputs, output)
	}
	marshal, err := json.Marshal(outputs)
	check(err)
	ioutil.WriteFile(outputFile, marshal, 0644)
}
