package main

import (
    "fmt"
	"strings"
)
type KeyValue struct {
	Key   string
	Value string
}

func quickSort(array []KeyValue, left int, right int){
	if left >= right{
		return;
	}
	tmp := array[left].Key;
	tmp_value := array[left].Value;
	low := left;
	high := right;
	for low < high{
		for strings.Compare(array[high].Key, tmp) >= 0 && low < high{
			high--;
		}
			array[low].Key = array[high].Key;
			array[low].Value = array[high].Value;
		for strings.Compare(array[low].Key, tmp) <= 0 && low < high{
			low++;
		}
			array[high].Key = array[low].Key;
			array[high].Value = array[low].Value;
	}
	array[low].Key = tmp;
	array[low].Value = tmp_value;
	quickSort(array, left, low - 1);
	quickSort(array, low + 1, right);
}
func main() {
    // To create a map as input
var m []KeyValue;
m =append(m,KeyValue{"g", "a"})
m =append(m,KeyValue{"g", "qq"})
m =append(m,KeyValue{"c", "tt"})
m =append(m,KeyValue{"c", "vv"})
m =append(m,KeyValue{"b", "rr"})
m =append(m,KeyValue{"b", "ww"})
fmt.Println(m);

    // To store the keys in slice in sorted order
    quickSort(m, 0, 5)

    // To perform the opertion you want
    fmt.Println(m);
}