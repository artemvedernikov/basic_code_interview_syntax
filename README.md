# basic_code_interview_syntax

As for 95% of code interviews there is a typical basic syntax used, I decided to store it in one place

## Ints

### Math

#### Go

```
math.MinInt32
math.MaxInt32

math.Pow(3, 5) // float64
int(math.Pow(10, float64(i))) // int example
```

## String

### Index at

#### Java

"".chatAt(num)

#### Kotlin

""[num]

#### Go

https://pkg.go.dev/strings

""[num]

### Contains

#### Go

```
strings.contains(w1, w2)
```

### Starts / Ends with

#### Go

```
strings.HasPrefix(str, "pref")
strings.HasSuffix(str, "suf")
```

## Loop

### For-loop

#### Go

```
for i := 0; i < 10; i++ {
   sum += i
}
```

### Foreach-loop

#### Go

```
for index, element := range someSlice {
    // index is the index where we are
    // element is the element from someSlice for where we are
}
```

### While-loop

### Go

```
for sum < 1000 {
		sum += sum
}
```

## Data structures

### Array

#### Go

https://go.dev/blog/slices-intro

```
// fixed size
m := len(words)
prefixsum := make([]int, m)

// dynamic
var pref_sum []int
pref_sum = append(pref_sum, 1)

// sort
ints := []int{7, 2, 4}
slices.Sort(ints)

sort.Sort(sort.Reverse(sort.IntSlice(keys)))

// concat
append([]int{1,2}, []int{3,4}...)
```

### Queue

#### Go

There is no built in queue in go but you can use slice

https://stackoverflow.com/questions/2818852/is-there-a-queue-implementation

```
queue := make([]int, 0)
// Push to the queue
queue = append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")
}
```

### Stack

#### Go

There is no build in Stack so use slice

https://dev.to/jpoly1219/stacks-in-go-54k

```
stack := []*TreeNode{}
// append
stack = append(stack, root)
// remove
stack = stack[:len(stack) - 1]
```

### Map

#### Go

make(map[key-type]val-type)

Example

```
m := make(map[string]int)
m["k1"] = 7
m["k2"] = 13

val, ok := m["foo"]
// If the key exists
if ok {
    // Do something
}

delete(m, "test")
len(m)
```

## OOP

### Go

https://medium.com/@canoguz/object-oriented-programming-in-go-e50f8fe4a620

https://yourbasic.org/golang/structs-explained/

```
type Customer struct {
	Name string
	Age  int
}

func (c *Customer) NewCustomer(name string, age int) {
	c.Name = name
	c.Age = age
}

func main() {
	customer := new(Customer)
	customer.NewCustomer("John", 30)
	fmt.Printf("%s: %d\n", customer.Name, customer.Age)
}

////

type Student struct {
	Name string
	Age  int
}

var a Student    // a == Student{"", 0}
a.Name = "Alice" // a == Student{"Alice", 0}

var pa *Student   // pa == nil
pa = new(Student) // pa == &Student{"", 0}
pa.Name = "Alice" // pa == &Student{"Alice", 0}

b := Student{ // b == Student{"Bob", 0}
	Name: "Bob",
}
    
pb := &Student{ // pb == &Student{"Bob", 8}
	Name: "Bob",
	Age:  8,
}

c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
d := Student{}             // d == Student{"", 0}
```


## Misc

### Go

```
import "reflect"
reflect.TypeOf(x)
```
