[![GoDoc](https://godoc.org/github.com/phemmer/gota?status.svg)](https://godoc.org/github.com/phemmer/gota)

gota is a Go implementation of several technical analysis algorithms. There is also the [talib](https://github.com/phemmer/talib) library which provides even more algorithms than gota currently offers. The main differences between gota and [talib](https://github.com/phemmer/talib) is that gota is native Go as opposed to C wrapped in Go, and gota is optimized for streaming & large data sets. Meaning that when a new sample is added to an algorithm, the algorithm only computes the update, and not the whole historical data set.

The gota library is mainly intended for my own use. I provide it in case it may be of use to others. Meaning that you are welcome to use it in your own project, but vendoring is highly recommended as I may make breaking API changes if I think it makes sense to do so.

Note: While the implementation is pure Go, running the tests does require the C version of TA-Lib. This is to ensure the implementation matches the functionality of the widely used [ta-lib](https://ta-lib.org).  
For Fedora, I use the RPM Sphere repo: https://rpmsphere.github.io/
