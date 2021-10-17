package main

// import (
// 	"fmt"
// 	"main/mapreduce"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// // Map function takes a chunk of data from the
// // input file and breaks it into a sequence
// // of key/value pairs
// func Map(value string) []mapreduce.KeyValue {
// 	result := []mapreduce.KeyValue{}

// 	// words := strings.Fields(value)
// 	// words := strings.Split(value, " ")

// 	// this function returns true when it is time to split
// 	// split when we encounter a non-letter char
// 	splitFunc := func(c rune) bool {
// 		return !(unicode.IsLetter(c))
// 	}
// 	words := strings.FieldsFunc(value, splitFunc)

// 	for _, word := range words {
// 		result = append(result, mapreduce.KeyValue{
// 			Key:   word,
// 			Value: "1",
// 		})
// 	}
// 	return result
// }

// // called once for each key generated by Map, with a list
// // of that key's associate values. should return a single
// // output value for that key
// func Reduce(key string, values []string) string {
// 	// return strconv.Itoa(len(values))
// 	totalSum := 0
// 	for _, val := range values {
// 		numVal, error := strconv.Atoi(val)
// 		if error != nil {
// 			fmt.Println(error)
// 			return error.Error()
// 		}
// 		totalSum += numVal
// 	}
// 	return strconv.Itoa(totalSum)
// }

// func main() {
// 	if len(os.Args) != 4 {
// 		fmt.Printf("%s: Invalid invocation\n", os.Args[0])
// 	} else if os.Args[1] == "master" {
// 		if os.Args[3] == "sequential" {
// 			mapreduce.RunSingle(5, 3, os.Args[2], Map, Reduce)
// 		} else {
// 			mr := mapreduce.MakeMapReduce(5, 3, os.Args[2], os.Args[3])
// 			// Wait until MR is done
// 			<-mr.DoneChannel
// 		}
// 	} else if os.Args[1] == "worker" {
// 		mapreduce.RunWorker(os.Args[2], os.Args[3], Map, Reduce, 100)
// 	} else {
// 		fmt.Printf("Unexpected input")
// 	}
// }
