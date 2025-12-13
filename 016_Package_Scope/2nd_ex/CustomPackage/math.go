package custompackage

import "fmt"

var Pi = 3.14159  // exported
var tau = 6.28318 


func SayHello(name string) {	// exported
	fmt.Printf("Hi! %s", name)
}


func add(num1 int, num2 int) int {		// unexported
	return num1 + num2
}


/*

-> Lowercase দিয়ে শুরু (যেমন add) → Unexported: শুধু ওই প্যাকেজের ভিতরেই অ্যাক্সেসযোগ্য।
-> Capital Letter দিয়ে শুরু (যেমন SayHello) → Exported: অন্যান্য প্যাকেজ থেকেও import করে ব্যবহার করা যায়।

*/