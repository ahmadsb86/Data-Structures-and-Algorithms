# Usage

To use this library in your Go projects, import the DSAlib module in your source file.

```go
import "https://github.com/ahmadsb86/Data-Structures-and-Algorithms/tree/Dev/Golang"
```



# Documentation

Each Data structure has its own struct and methods. Details for each one have been listed here. 

## Linked Lists

Linked lists allow us to append and delete nodes in O(1) with the trade off of O(N) lookup time complexity. DSAlib implements a doubly linked list in its `LinkedList` struct.

To initialize a linked list, pass the type of the elements it will contain

```go
LL := DSAlib.LinkedList[int]{}
```

To append or delete nodes to the **end** of a linked list, use `Push()` and `Pop()` respectively. Both methods run in O(1).  *Note: `Pop()` returns the value stored in the node that will be deleted.*

```go
LL.Push(100)    //O(1)
LL.Push(200)
LL.Push(300)
fmt.Println(LL.Pop())    //removes last node and prints its value


//Output: 300
```

To append or delete nodes to the **start** of a linked list, use `Shift()` and `Slice()` respectively. Both methods run in O(1).  *Note: `Slice()` returns the value stored in the node that will be deleted.*

```go
LL.Push(200)
LL.Shift(100)     //puts 100 behind 200
LL.Push(300)      //puts 300 in front of 200
LL.Shift(000)     //puts 000 behind 100
LL.Slice()        //removes 000


//LL is now 100 -> 200 -> 300
```

To print the values in a linked list, use `Display()`. You may optionally provide a delimiter (string displayed between two the value of two nodes) which otherwise defaults to a space. 

```go
LL.Push(100)
LL.Push(200)

LL.Display() //O(N)
//Output: 100 200

LL.Display(", ")
//Output: 100, 200
```

To traverse a linked list, use `InitTraverse()` which returns the first node's value as well as resets the index pointer for traversing. To get the next node's value, use `Traverse()`. To check if the linked list has been traversed, use the `hasTraversed` variable. *Note: If a linked list has been traversed, any further calls to `Traverse()` will return the zero value for the type of the linked list until `InitTraverse()` has been called again.*

```go
for i:=LL.InitTraverse(); !LL.hasTraversed; i=LL.Traverse(){ //O(N)
    fmt.Printf(i)
}
```

Similarly, to traverse a linked list backwards, use `InitReversedTraverse()` and `ReversedTraverse()` instead.

```go
for i:=LL.InitReversedTraverse(); !LL.hasTraversed; i=LL.ReversedTraverse(){ //O(N)
    fmt.Printf(i)
}
```