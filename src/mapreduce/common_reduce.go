package mapreduce

import (
	"fmt"
	"encoding/json"
	"strings"
	"os"
)
// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.


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
func quick(array []KeyValue){
	quickSort(array, 0, len(array) - 1);
}

func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	var array []KeyValue;

	for i:= 0; i < nMap; i++{
		intermidiateName := reduceName(jobName, i, reduceTaskNumber);
		
		f, err := os.Open(intermidiateName);
		if err != nil{
			fmt.Println(err);
		}
		enc := json.NewDecoder(f)
		var kv KeyValue;
		for err := enc.Decode(&kv); err == nil; {
			err = enc.Decode(&kv);
			array = append(array, kv);
		}
	}


		var valueSet []string;
		num := len(array);
		foName := mergeName(jobName, reduceTaskNumber);
		fo, fo_err := os.Create(foName);
		if(fo_err != nil){
			fmt.Println("read file failed");
		}
		outEnc := json.NewEncoder(fo);
		if len(array) < 1{
			fmt.Println("no input at all");
		}
		quick(array);
		
		count := 0;
		// for current := array[count].Key; count < num; count++{
		// 	valueSet = append(valueSet, )
		// 	outEnc.Encode(KeyValue{Key:key, Value:reduceF(key, valueSet)});
		// }
		for count < num {
			current := array[count].Key
			interval := 0;
			key := array[count + interval].Key;
			for next := array[count + interval].Value; strings.Compare(next, current) == 0 ;{
				valueSet = append(valueSet, next)
				interval++;
				if count + interval == num{
					break;
				}
				next = array[count + interval].Value
			}
			count += interval;
			outEnc.Encode(KeyValue{Key:key, Value:reduceF(key, valueSet)});
			valueSet = nil;
		}

		fo.Close();
	
	// TODO:
	// You will need to write this function.
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
}
