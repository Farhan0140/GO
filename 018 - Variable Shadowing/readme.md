->   ধরুন আপনি একই নামে দুটি ভেরিয়েবল লিখলেন — একটি উপরের (বাইরে) scope-এ,
     একটি ভেতরের (inner) block-এ। Go রানটাইম বা কম্পাইলার inner scope-টিকে মনে রাখবে এবং 
     outer scope-এর ভেরিয়েবলকে “অদৃশ্য” করে দেবে যতক্ষণ inner scope চলছে

->   Go-তে scope ঠিক করে দেয় ভেরিয়েবল কোথায় ডিক্লেয়ার হয়েছে তার ওপর ভিত্তি করে। 
     যেহেতু Go block-scoped ভাষা, উপরোক্ত কারণে একই নামে ভেরিয়েবল আবার ভেতরে লিখলে 
     সেটা inner scope-এর জন্য নতুন ভেরিয়েবল তৈরি করে এবং outer scope-এর ভেরিয়েবলটিকে ঢেকে দেয়



1. if condition-এ Variable Shadowing

func main() {
    x := 10
    fmt.Println("Before if, x =", x)         --> 10

    if x := x + 5; x > 10 {
        fmt.Println("Inside if, x =", x)     --> 15
    }

    fmt.Println("After if, x =", x)         --> 10
}



2. for loop-এ Variable Shadowing

func main() {
    i := 100
    fmt.Println("Before loop, i =", i)       --> 100

    for i := 0; i < 3; i++ {
        fmt.Println("Inside loop, i =", i)   --> 0 1 2
    }

    fmt.Println("After loop, i =", i)      --> 100
}



3. range loop-এ Shadowing

func main() {
    nums := []int{10, 20, 30}

    n := 999

    for _, n := range nums {
        fmt.Println("Inside loop, n =", n)      --> 10 20 30
    }

    fmt.Println("After loop, n =", n)      --> 999
}



4. Function parameter-এর মাধ্যমে Shadowing

var count = 50

func printCount(count int) {
    fmt.Println("Inside function, count =", count)     --> 10
}

func main() {
    printCount(10)
    fmt.Println("Global count =", count)     --> 50
}
