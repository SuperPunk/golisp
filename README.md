参考sicp第四章，lisp解释器的go语言版本

```Lisp
(+ 3 5)
(* 5 (+ 2 3))
(* (- (- 5 3) 2) (+ 2 3))
```

```Lisp
(define (append x y)
  (if (null? x)
    y
    (cons (car x) (append (cdr x) y))))

(append (cons 1 (cons 2 nil)) (cons 3 (cons 4 nil))) 
```

```Lisp
(define (f x y)
  (lambda (m) (m x y)))

(define (fl z)
  (z (lambda (p q) p)))

(define (fr z)
  (z (lambda (p q) q)))
```

```Lisp
(define pair (f 1 100))
(fl pair)
(fr pair)
```