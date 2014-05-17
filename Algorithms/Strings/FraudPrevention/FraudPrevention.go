// O(nlogn) solution without counting 
// the loop of reading inputs
package main 

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"container/heap"
	"bytes"
)

// heap to efficiently store order id for output
type IntHeap []int
func (h IntHeap) Len() int              { return len(h) }
func (h IntHeap) Less(i, j int) bool    { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{})   { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

// representation of a record
type Record struct {
	orderId    int
	dealId     int
	email      string
	address    string
	city       string
	state      string
	zip        string
	creditCard string
}

// break the input string into pieces 
// and return a struct of record
func createRecord(rec string) Record {
	record := strings.Split(rec, ",")
	orderId, _ := strconv.Atoi(record[0])
	dealId, _ := strconv.Atoi(record[1])
	return Record{orderId, dealId, 
		record[2], record[3], record[4], 
		record[5], record[6], record[7]}
}

// some emails are just the same so they
// must be united to consistent format
func _cleanEmail(email string) string {
	username := strings.Split(email, "@")[0]
	username = strings.Split(username, "+")[0]
	username = strings.Replace(username, ".", "", -1)
	return strings.ToLower(username)
}

// core function to detect frauds with same emails,
// same ids, but different credit cards
func detectFraudsByEmail(record Record, 
	fraudsByEmail map[string][]Record, 
	fraudsByOrderId map[int]bool, 
	fraudsByOrderIdHeap *IntHeap) {

	key := _cleanEmail(record.email) + strconv.Itoa(record.dealId)
	potentialFraudRecords, existed := fraudsByEmail[key]
	if existed {
		for _, rec := range potentialFraudRecords {
			if !strings.EqualFold(record.creditCard, rec.creditCard) {
				// this is a fraud
				if !fraudsByOrderId[record.orderId] {
					fraudsByOrderId[record.orderId] = true
					heap.Push(fraudsByOrderIdHeap, record.orderId)
				}
				if !fraudsByOrderId[rec.orderId] {
					fraudsByOrderId[rec.orderId] = true
					heap.Push(fraudsByOrderIdHeap, rec.orderId)
				}
			}
		}
	}
	fraudsByEmail[key] = append(fraudsByEmail[key], record)
}

// some addresses are just the same so 
// they must be united to consistent format
func _cleanAddress(address string, 
	city string, state string, zip string) string {
	address = strings.ToLower(address)
	city = strings.ToLower(city)
	state = strings.ToLower(state)
	zip = strings.ToLower(zip)
	if strings.HasSuffix(address, "st.") {
		address = address[0 : len(address) - 3] + "street"
	}
	if strings.HasSuffix(address, "rd.") {
		address = address[0 : len(address) - 3] + "road"
	}
	if strings.EqualFold(state, "new york") {
		state = "ny"
	}
	if strings.EqualFold(state, "california") {
		state = "ca"
	}
	if strings.EqualFold(state, "illinois") {
		state = "il"
	}
	return address + city + state + zip
}

// core function to detect frauds with same addresses
// same deal ids, but different credit cards
func detectFraudsByAddress(record Record, 
	fraudsByAddress map[string][]Record,
	fraudsByOrderId map[int]bool,
	fraudsByOrderIdHeap *IntHeap) {

	address := _cleanAddress(record.address, 
		record.city, record.state, record.zip)
	key := address + strconv.Itoa(record.dealId)
	potentialFraudRecords, existed := fraudsByAddress[key]
	if existed {
		for _, rec := range potentialFraudRecords {
			if !strings.EqualFold(record.creditCard, rec.creditCard) {
				// this is a fraud
				if !fraudsByOrderId[record.orderId] {
					fraudsByOrderId[record.orderId] = true
					heap.Push(fraudsByOrderIdHeap, record.orderId)
				}
				if !fraudsByOrderId[rec.orderId] {
					fraudsByOrderId[rec.orderId] = true
					heap.Push(fraudsByOrderIdHeap, rec.orderId)
				}
			}
		}
	}
	fraudsByAddress[key] = append(fraudsByAddress[key], record)
}

func main() {
	var inputSize int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &inputSize)
	// read the last character left on first line
	io.ReadString('\n')

	// generously use data structures for efficient stores
	fraudsByOrderId := make(map[int]bool)
	fraudsByOrderIdHeap := &IntHeap{}
	heap.Init(fraudsByOrderIdHeap)
	fraudsByEmail := make(map[string][]Record)
	fraudsByAddress := make(map[string][]Record)

	for i := 0; i < inputSize; i++ {
		rec, _ := io.ReadString('\n')
		record := createRecord(rec)
		detectFraudsByEmail(record, 
			fraudsByEmail, 
			fraudsByOrderId, 
			fraudsByOrderIdHeap)
		detectFraudsByAddress(record, 
			fraudsByAddress, 
			fraudsByOrderId, 
			fraudsByOrderIdHeap)
	}

	// output to console with comma-separated string
	var strBuffer bytes.Buffer
	for i := 0; fraudsByOrderIdHeap.Len() > 0; i++ {
		orderId := heap.Pop(fraudsByOrderIdHeap).(int)
		strBuffer.WriteString(strconv.Itoa(orderId))
		strBuffer.WriteString(",")
	}
 	output := strBuffer.String()
 	fmt.Println(output[0 : len(output)-1])
}