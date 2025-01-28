# basic_code_interview_syntax

As for 95% of code interviews there is a typical basic syntax used, I decided to store it in one place

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

```

## Data structures

### Array

#### Go

```
// fixed size
m := len(words)
prefixsum := make([]int, m)

// dynamic
var pref_sum []int
pref_sum = append(pref_sum, 1)
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
