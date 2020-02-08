# Interpreter and Compiler in Golang for Monkey programming language

Monkey has a C-like syntax, supports variable bindings, prefix and infix operators, has first-class and higher-order functions, can handle closures with ease, and has built-in integers, booleans, arrays and hashes.

Based on Thorsten Ball's books interpreterbook.com and compilerbook.com

**By Raymond Gan**

## Part 1: Interpreter
- Building an interpreter for a C-like programming language from scratch, with NO 3rd party libraries
- Building a lexer, a parser and an Abstract Syntax Tree (AST)
- The Pratt parsing technique and a recursive descent parser
- Building a REPL
- Build a tree-walking evaluator

Monkey language:
```
// Bind values to names with let-statements
let version = 1;
let name = "Monkey programming language";
let myArray = [1, 2, 3, 4, 5];
let coolBooleanLiteral = true;

// Use expressions to produce values
let awesomeValue = (10 / 2) * 5 + 30;
let arrayWithValues = [1 + 1, 2 * 2, 3];
```
Monkey also supports function literals. We can use them to bind a function to a name:
```
// Define a `fibonacci` function
let fibonacci = fn(x) {
  if (x == 0) {
    0                // Monkey supports implicit returning of values
  } else {
    if (x == 1) {
      return 1;      // ... and explicit return statements
    } else {
      fibonacci(x - 1) + fibonacci(x - 2); // Recursion! Yay!
    }
  }
};
```
Supported data types: booleans, strings, hashes, integers and arrays. We can combine them!
```
// Here is an array containing two hashes, that use strings as keys and integers
// and strings as values
let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

// Getting elements out of the data types is also supported.
// Here is how we can access array elements by using index expressions:
fibonacci(myArray[4]);
// => 5

// We can also access hash elements with index expressions:
let getName = fn(person) { person["name"]; };

// And here we access array elements and call a function with the element as
// argument:
getName(people[0]); // => "Anna"
getName(people[1]); // => "Bob"
```
In Monkey functions are first-class citizens, treated like any other value. Thus we can use higher-order functions and pass functions around as values:
```
// Define the higher-order function `map`, that calls the given function `f`
// on each element in `arr` and returns an array of the produced values.
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };

  iter(arr, []);
};

// Now let's take the `people` array and the `getName` function from above and
// use them with `map`.
map(people, getName); // => ["Anna", "Bob"]
```
Monkey also supports closures:
```
// newGreeter returns a new function, that greets a `name` with the given
// `greeting`.
let newGreeter = fn(greeting) {
  // `puts` is a built-in function we add to the interpreter
  return fn(name) { puts(greeting + " " + name); }
};

// `hello` is a greeter function that says "Hello"
let hello = newGreeter("Hello");

// Calling it outputs the greeting:
hello("dear, future Reader!"); // => Hello dear, future Reader!
```
## Part 2: Compiler

- Take the lexer, the parser, the AST, the REPL and the object system and use them to build a new, faster implementation of Monkey
- Change its architecture and turn it into a **bytecode compiler and virtual machine**, from scratch.
- Build compiler and VM side-by-side so that we always have a running system to steadily evolve.

- Define our own **bytecode instructions**, specifying their operands and their encoding. Along the way, we also build a mini-disassembler for them.
- Write a **compiler** that takes in a Monkey AST and turns it into bytecode by emitting instructions
- At the same time, build a **stack-based virtual machine** that executes the bytecode in its main loop

- build a **symbol table** and a constant pool
- do **stack arithmetic**
- generate **jump instructions**
- build **frames** into our VM to execute functions with **local bindings and arguments**!
- add **built-in functions** to the VM
- get real **closures** working in the virtual machine and learn why closure-compilation is so tricky
