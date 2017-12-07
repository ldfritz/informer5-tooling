package informer5

import (
	"encoding/json"
	"log"
)

const LimitChunkSize = 1048576

func ChunkContent(content []byte, size int) [][]byte {
	var chunks [][]byte

	if len(content) < size {
		chunks = append(chunks, content)
		return chunks
	}

	records := ToRecords(content)

	for len(records) > 0 {
		threshold := EndOfChunk(size, records)
		chunks = append(chunks, ToJSON(records[:threshold]))
		records = records[threshold:]
	}
	return chunks
}

func ToRecords(content []byte) []map[string]interface{} {
	var records []map[string]interface{}
	err := json.Unmarshal(content, &records)
	if err != nil {
		log.Fatalln("unable to parse JSON:", err)
	}
	return records
}

func ToJSON(records []map[string]interface{}) []byte {
	b, err := json.Marshal(records)
	if err != nil {
		log.Fatalln("unable to create JSON:", err)
	}
	return b
}

func SumOfChunks(chunks [][]byte) int {
	sum := 0
	for _, v := range chunks {
		sum += len(v)
	}
	return sum
}

func EndOfChunk(size int, records []map[string]interface{}) int {
	threshold := 0

	i := len(records)

	below := 0
	above := len(records)

	if len(ToJSON(records)) < size {
		threshold = len(records)
	}

	for threshold == 0 {
		i = ((above - below) / 2) + below + 1

		hi := len(ToJSON(records[:i]))
		lo := len(ToJSON(records[:i-1]))

		switch {
		case hi < size:
			below = i
		case lo > size:
			above = i - 1
		default:
			threshold = i - 1
		}
	}
	return threshold
}
