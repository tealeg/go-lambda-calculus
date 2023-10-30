#lang racket/base

(define true (lambda (a) (lambda (b) a)))

(define false (lambda (a) (lambda (b) b)))

(define if-then-else (lambda (a) (lambda (b) (a b))))

(define not (lambda (a)
              ((a (lambda (b)
                    (lambda (c) c)))
               (lambda (d)
                 (lambda (e) d)))))

(define and (lambda (a)
              (lambda (b)
                ((a b) false))))

(define or (lambda (a)
             (lambda (b)
               ((a true) b))))

(define zero (lambda (s) (lambda (z) z)))

(define one (lambda (s) (lambda (z) (s z))))

(define two (lambda (s) (lambda (z) (s (s z)))))

(define succ (lambda (a) (lambda (b) (lambda (c) (b ((a b) c))))))

(define pred (lambda (n)
               (lambda (f)
                 (lambda (x)
                   (((n (lambda (g)
                          (lambda (h) (h (g f)))))
                     (lambda (u) x))
                    (lambda (u) u))))))

(define plus (lambda (a) (lambda (b) (lambda (c) (b ((a b) c))))))

(define mul (lambda (a) (lambda (b) (lambda (c) (a (b c))))))

(define is-zero (lambda (a) (((a false) not) false)))

(define y (lambda (f)
            ((lambda (le) ((lambda (g) (g g))
                           (lambda (h)
                             (le (lambda (x)
                                   ((h h) x)))))))))
