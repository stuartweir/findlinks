# findlinks
findlinks is an exercise from the GoPL textbook, exercise 5.1.

With findlinks, the program recursively prints out the links of a given website. To use, first build both files:

`go install stuartweir/findlinks/fetch/fetch` & `go build stuartweir/findlinks`

then, in terminal, type: `fetch <name of website> | ./findlinks`

for `https://golang.org`, the printout should be:

`0, /
1, /
2, #
3, /doc/
4, /pkg/
5, /project/
6, /help/
7, /blog/
8, http://play.golang.org/
9, #
10, #
11, //tour.golang.org/
12, https://golang.org/dl/
13, //blog.golang.org/
14, https://developers.google.com/site-policies#restrictions
15, /LICENSE
16, /doc/tos.html
17, http://www.google.com/intl/en/policies/privacy/`
