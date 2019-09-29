package hashing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	k               int8   = 7
	a               int8   = 26
	known64BitPrime uint64 = 17586613600806056593
	// Other 64 bit primes
	// known64BitPrime uint64 = 10324706610870574883
	// known64BitPrime uint64 = 14385965969526276271
	// known64BitPrime uint64 = 15700719402893486197
	// known64BitPrime uint64 = 13390804203280917121
	// known64BitPrime uint64 = 12631952504492069741
	// known64BitPrime uint64 = 14687623246052906689
	// known64BitPrime uint64 = 18235099962527857067
	// known64BitPrime uint64 = 13557970565612484931
)

var (
	bytesread    int    = 0
	offset       int8   = 0
	rollingHash  uint64 = 0
	b                   = make([]byte, 1)
	corpusCounts        = make(map[uint64]int)
)

func checkEOFError(err error) {
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
			fmt.Println("got to end of file")
		}
	}
}

func update(m map[uint64]int, u uint64) {
	i, _ := m[u]
	m[u] = i + 1
}

func printBytes(buffer []byte, bytesread int) {
	fmt.Println("bytes read          : ", bytesread)
	fmt.Println("buffer              : ", buffer)
	fmt.Println("bytes               : ", buffer[:bytesread])
	fmt.Println("buffer to string    : ", string(buffer))
	fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
}

func genHash(currentHash uint64, oldByte, newByte byte) uint64 {
	newHash := currentHash
	o := uint64(oldByte)
	n := uint64(newByte)
	A := uint64(a)

	newHash += known64BitPrime
	newHash -= o * A
	newHash += n * A
	newHash = newHash % known64BitPrime

	return newHash
}

// Indent will format a map for prettier printing
func Indent(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(b)
}

// RollingHashes returns a map of rolling hashes based on
// reading k number of bytes at a time
func RollingHashes(file *os.File) map[uint64]int {
	fileCounts := make(map[uint64]int)

	for i := 0; i < int(k); i++ {
		bytesread, err := file.ReadAt(b, int64(i))
		checkEOFError(err)
		if bytesread == 0 {
			break
		}
		rollingHash = genHash(rollingHash, 0, b[0])
		// log.Printf("byte #%-2d was %3d (%s), rollingHash was: %d\n", i, b[0], string(b[0]), rollingHash)
	}
	update(fileCounts, rollingHash)
	update(corpusCounts, rollingHash)

	for i := int(k); true; i++ {
		bytesread, err := file.ReadAt(b, int64(i))
		checkEOFError(err)
		if bytesread == 0 {
			// log.Println("got to end of file")
			break
		}
		newByte := b[0]

		bytesread, err = file.ReadAt(b, int64(i-int(k)))
		checkEOFError(err)
		oldByte := b[0]

		rollingHash = genHash(rollingHash, oldByte, newByte)
		// fmt.Printf("byte #%-2d was %3d (%s), rollingHash was: %d\n", i, newByte, string(newByte), rollingHash)
		update(fileCounts, rollingHash)
		update(corpusCounts, rollingHash)
	}

	return fileCounts
}

// func main() {
// 	log.SetPrefix("hashing ")
// 	absPath, _ := filepath.Abs("data/filetoread3.txt")
// 	file, err := os.Open(absPath)
// 	if err != nil {
// 		log.Println("Got an unresolved error when trying to open a file")
// 		log.Fatal(err)
// 		return
// 	}
// 	defer file.Close()

// 	fileCounts := RollingHashes(file)
// 	fmt.Println(indent(fileCounts))
// }
