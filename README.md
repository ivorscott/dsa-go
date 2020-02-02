# Data Structures and Algorithms in Go

Practice repo for [Data Structures & Algorithms In Go](https://www.amazon.com/Data-Structures-Algorithms-Hemant-Jain/dp/1099552060) by Hemant Jain.

## Chapter 1: Algorithms Analysis

An algorithm is a set of steps to accomplish a task.
Knowledge of algorithms help us get the desired results we want faster. Our task is to apply the appropriate algorithm for a given problem.

The most important properties of an algorithm are:

1. Correctness: The algorithm should be correct. It should be able to process all the given inputs and provide correct output.

2. Efficiency: The algorithm should be efficient in solving problems. Efficiency is measured in two parameters. First, time complexity, in other words, how quick the result is provided by an algorithm. Second, is space complexity, or how much RAM or memory that an algorithm consumes to give desired result.

**Time-Complexity** is represented aby function T(n) - time required versus the input size n.

**Space-Complexity** is represented by function S(n) - memory used versus the input size n.

## Asymptotic analysis

Asymptotic analysis is used to compare the efficiency of algorithm independent of any data set of programming language.

It's known as a method of describing the limiting behavior of algorithms. We only see the performance of an algorithm when we have very large input data.

Generally, we in interested in _the order of growth_ (or the Asymptotic-running time) of a particular algorithm, not the exact time required for running an algorithm. Algorithms need to be able to scale.

The point of asymptotic measures is to conceptualize how a function or graph may behave when its input is very large.

## Big-O Notation

As a computer scientist if you are working on an important piece of software, you will likely need to be able to estimate how fast some algorithm is going to run.

Big-O notation describes the limiting behavior of a function when the argument tends towards a particular value or infinity. Big-O notation is used to classify algorithms according to how their running time and space requirement grows as the input size grows.

Big-O notation characterizes functions according to the growth rates: different functions with the same growth rate may be represented using the same O notation.

The "O" stands for "order of the function" or the growth rate of the function.

Big-O describes the number of operations an algorithm will make in the worse case, so that you can compare algorithms of different growth rates.

Big-O usually only provides an upper bound on the growth rate of the function. Associated with Big-O notation are several other notations to describe other kinds of bounds on asymptotic growth rates: _Theta notation and Omega notation._

## Common Notations

**Big-O Notation - Worst Case Complexity**

- It is the complexity of solving the problem for the worst input size n. It provides the upper bound for the algorithm. This is the most common analysis used.

**Theta Notation - Average Case Complexity**

- It is the complexity of solving the problem on an average. We calculate the time for all the possible inputs and then take an average of it.

**Omega Notation - Best Case Complexity**

- It is the complexity of solving the problem for the best input of size n.

## Big-Omega (Ω) Notation

Similar to Big-O notation, Big-Omega (Ω) notation is used to describe the performance or complexity of an algorithm

### The Difference Between Big-O and Big-Omega (Ω) Notation.

The difference between Big-O notation and Big Ω notation is that Big-O is used to describe the worst case running time for an algorithm but, Big Ω notation on the other hand, is used to describe the best case.

## Big-Theta (Θ) Notation

Big-Omega tells us the lower bound of a function's runtime and Big-O tells us the upper bound. Often times, they are different and we can't put a guarantee on the runtime, it will vary between the two bounds and inputs.

But what happens where they are the same? Then we can give the Theta (Θ) bound. Our function will run in that time, no matter what input we give it. In general, we always want to give a theta bound if possible because it is the most accurate and tightest bound.

If we can't give a theta bound, the next best thing is the Big-O (upper) bound.

## Recursive Function

A recursive function is a function that calls itself.

A recursive method consists of two parts: **Termination condition** and the **Body** (which performs the recursive expansion)

**The Termination Condition**

A recursive method always contains one or more terminating conditions. A condition where the recursive method performs a simple case and does not call itself.

**The Body**

The main logic of the recursive method id contained in the body of the method. It also contains the recursive expansion method that calls itself.

In summary:

1. A recursive algorithm must have a termination condition.

2. A recursive algorithm must change its state, and move towards the termination condition.

3. A recursive algorithm must call itself.

**Note:** The speed of a recursive program is slower because of the stack overheads. If the same operation can be done using an iterative solution (like using loops), then we should prefer an iterative solution in place of recursion to avoid stack overhead. Also, remember that without a termination condition, the recursive method may run forever and will finally consume all the stack memory.

## Benchmarking Time Complexity

`make time`

```
└── timecomplexity
    ├── exponential.go
    ├── exponential_test.go
    ├── linear.go
    ├── linear_test.go
    ├── logarithmic.go
    ├── logarithmic_test.go
    ├── quadratic.go
    └── quadratic_test.go
```

## Benchmarking Array Algorithms

Try any example with:

`make [example]`

### Examples

`binarysearch, binarysearchr, commondivisor, factorial, fibonacci, findone, sort, sumarray`

```
arrays
├── binarysearch
│   ├── binarysearch.go
│   └── binarysearch_test.go
├── binarysearchr
│   ├── binarysearchr.go
│   └── binarysearchr_test.go
├── commondivisor
│   ├── commondivisor.go
│   └── commondivisor_test.go
├── factorial
│   ├── factorial.go
│   └── factorial_test.go
├── fibonacci
│   ├── fibonacci.go
│   └── fibonacci_test.go
├── findone
│   ├── findone.go
│   └── findone_test.go
├── sort
│   ├── sort.go
│   └── sort_test.go
└── sumarray
    ├── sumarray.go
    └── sumarray_test.go
```

## Chapter 2: Approach To Solving Algorithm Design Problems

Theoretical knowledge of algorithms is essential, but not enough.

Five steps for solving algorithm design questions are:

1. Constraints
2. Idea Generation
3. Complexities
4. Coding
5. Testing

### Constraints

Many people make mistakes, as they do not ask clarifying questions about a given problem. They assume many things simultaneously and begin working with that. There is a lot of data missing that you need to collect before beginning to solve a problem.

**Determining the Constraints for an array of numbers**

Questions worth asking.

1. How many elements are there in an array?
2. What is the range of value in each element? What is the min and max value?
3. What is the kind of data in each element? Is it an integer or a floating point?
4. Does the array contain unique data or not?

**Determining the Constraints for an array of strings**

1. How many elements are there in an array?
2. What is the length of each string? What is the min and max length?
3. Does the array contain unique data or not?

**Determining the Constraints for a Graph**

1. How many nodes are there in the graph?
2. How many edges are there in the graph?
3. Is it a weighted graph? What is the range of weights?
4. Is the graph directed or undirected?
5. Is there a loop in the graph?
6. Is there a negative sum loop in the graph?
7. Does the graph have self-loops?

### Idea Generation

Even if you have an idea about a possible solution to the problem, you should navigate you options and consider the trade-offs.

How do you solve an unseen problem? The solution to this problem is that you need to do a lot of practice and the more you will practice the more you will be able to solve any unseen question, which come before you. When you have solved enough problems, you will be able to see a pattern in the questions and be able to solve unseen problems easily.

Steps to Solving an unknown problem:

1. Try to simplify the problem
2. Try a few examples
3. Think of a suitable data-structure
4. Think about similar problems that you have already solved.

### Complexities

Solving a problem is not just finding a correct solution. The solution should be fast and should have reasonable memory requirement. You should be able to to Big-O analysis.

Expect that you should be able to find the time and space complexity of the algorithms you use. You should be able to compute the time and space complexity instantly. Whenever you are solving any problem, you should find the complexity associated with it and in doing this you will be able to choose the best solutions. In some problems there is some trade-offs between Space and Time Complexity, so you should know these trade-offs.

### Coding

At this point, you have already captured all the constraints of the problem, proposed a few solutions, evaluated the complexities and picked a solution to code. Never jump into coding before considering constraints, generating ideas and identifies complexities.

### Testing

Once the code is written, you are not yet done. It is most important that you test your code with several small test cases. It shows that you understand the importance of testing. It also gives you confidence that you are not writing buggy code. Consider normal test cases, edges case, and load testing.

Edge cases are test cases designed to test the boundaries of the code. Edge cases may help to make your code more robust. Just few checks need to be added to the code to take care of these conditions.

Load testing implies, your code will be tested with huge data. This will allow you to test if your code is too slow or memory intensive.

## Chapter 3: Abstract Data Type and Golang Collections
