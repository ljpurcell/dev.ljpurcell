---
title: "The Little Schemer: Chapter 1"
slug: "schemer-ch1"
tags: scheme, "functional programming"
---
Chapter notes from _The Little Schemer_, which I am reading as part of my self-study CS curriculum and exploration into functional programming. In particular, I’m aiming to better understand recursion and symbolic manipulation, key topics in functional languages.

Not everything written here is covered in the chapter—some of it I had to research to make better sense of the material. As such, the notes are not purely a summary but also include supplementary explanations.

## Primitive Data Types

### Atom
An atom is any indivisible value in Scheme, meaning it is not a list or pair. Atoms include numbers, strings, booleans, and symbols:
```scheme  
123  
"word"
#t
x
```  


### List
Lists are collections of symbolic expressions (S-expressions) wrapped in parentheses. Lists can contain other lists as well as atoms:
```scheme  
'(123 word)  
'(protein (carbs fat))
'('(we) '(live) '(in) '(a) '(society))  
```  

This includes the null or empty list `()` . 

Unquoted lists are evaluated: 
```scheme
(+ 1 2 3) ;; 6
```

Quoted lists are treated as data and not evaluated:
```scheme
'(+ 1 2 3) ;; (+ 1 2 3)
```

### Pair
Pairs allow us to combine two Scheme objects into a single entity.

As the GNU reference states:

{.quote}
> Pairs are used to combine two Scheme objects into one compound object. Hence the name: A pair stores a pair of objects.
>
>The data type __pair__ is extremely important in Scheme, just like in any other Lisp dialect. The reason is that pairs are not only used to make two values available as one object, but that pairs are used for constructing lists of values.
>
> {.citation}
> [Pairs](https://www.gnu.org/software/guile/manual/html_node/Pairs.html#:~:text=Pairs%20are%20used,of%20their%20own)

The syntax for pairs uses parentheses and a dot to separate the two objects. There must be whitespace on either side of the dot:
```scheme
'(1 . "a") ;; (1 . "a")
'("hello" . ()) ;; ("hello")
```

We can test if an object is a pair using the `pair?` primitive:
```scheme
(pair? '("The" . "Simpsons")) ;; #t
(pair? "The Simpsons") #f
(pair? '()) #f
```


## Symbolic Expressions
An [S-expression](https://en.wikipedia.org/wiki/S-expression) (short for symbolic expression) is a way of representing data in Scheme, as well as in LISP and computer science in general. Both atoms and lists are S-expressions.

The uniform structure of S-expressions makes them especially convenient for representing recursive data, which will become very useful as we explore more of Scheme's power.

## Procedures
`car` returns the first element of a list and cannot be applied to atoms:
```scheme
(car '(12 24 36)) ;; 12
(car (car '((1 2 3 4) (5 6 7)))) ;; 1
(car '(+ (* 2 2) (/ 3 3))) ;; +
```
__The Law of Car:__ _The primitive_ `car` _is only defined for non-empty lists._

`cdr` ("could-er") removes the first element of the list and returns the rest:
```scheme
(cdr '(x y z)) ;; (y z)
(cdr (cdr '(x y z))) ;; (z)
(cdr '(gone)) ;; ()
```
__The Law of Cdr:__ _The primitive_ `cdr` _is only defined for non-empty lists. The_ `cdr` _of any non-empty list is always another list._

`cons` takes an S-expression (which can be an atom or list) and a list as parameters, appending the S-expression to the front of the list:
```scheme
(cons 'zoo '(lake forest)) ;; (zoo lake forest)
(cons '+ '(1 2)) ;; (+ 1 2)
```
__The Law of Cons:__ _The primitive_ `cons` _takes two arguments. The second argument to_ `cons` _must be a list. The result is a list._

`null?` takes a list and returns `#t` (true) if it is empty, otherwise `#f` (false):
```scheme
(null? '()) ;; #t
(null? (/ 7 1)) ;; #f
(null? 1) ;; #f
```
__The Law of Null?:__ _The primitive_ `null?` _is defined only for lists._

`eq?` takes in two non-numeric atoms and tests them for equality
```scheme
(eq? 'abc 'xyz) ;; #f
(eq? 'same 'same) ;; #t
(eq? (car (car '((first) second third))) (car (cons 'first '()))) ;; #t
```
__The Law of Eq?:__ _The primitive_ `eq?` _takes two arguments. Each must be a non-numeric atom._

__Note:__ In some Scheme implementations, `eq?` may behave unexpectedly with numbers and strings. For numeric equality, use `=`.


## Binding
To bind a symbol to a value, we use the `define` procedure:
```scheme
(define one-hundred 100)
(define name "Lyndon")
(define age one-hundred)
(define empty-list ())
(define pod '("pea" . "another pea"))
```

All the procedures we've seen so far can work with symbols that have been bound to values in this way:
```scheme
(null? empty-list) ;; #t
(cons 1 empty-list) ;; (1)
(eq? name "Lyndon") ;; #t
```

We can also define procedures. The first argument is the procedure signature (name and parameters), and the second argument is the procedure body:
```scheme
(define (square x) (* x x))

(square 2) ;; 4
```

Alternatively, we can use `lambda` to define anonymous procedures:
```scheme
(define add-one
  (lambda (n) (+ n 1)))
  
(add-one 99) ;; 100
```
## Putting it All Together

As a helper function in the book, the authors suggest defining an `atom?` procedure that returns a boolean indicating whether an object is an atom:

```scheme
(define atom?
  (lambda (object)
    (and (not (pair? object)) (not (null? object)))))
```

This definition serves as a useful summary of many of the constructs we've covered, and providing a clearer understanding of what an atom actually is.

Here, the `and` and `not` procedures act as logical operators, functioning as you would expect.
