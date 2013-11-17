package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style>
</head><title>Quadratic Equation Solver</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solves equations of the form
a<i>x</i><sup>2</sup> + b<i>x</i> + c</p>`
	form = `<form action="/" method="POST">
<input type="text" name="a" size="5"><i>x</i><sup>2</sup> + 
<input type="text" name="b" size="5"><i>x</i> +
<input type="text" name="c" size="5"><br>
<input type="submit" value="Calculate">
</form><br>`
	result = `%v<i>x</i><sup>2</sup> %s<i>x</i> %s 
<span style="font-size:large">→</span>
%v or %v`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type equation struct {
	a, b, c      float64
	root1, root2 complex128
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if a, b, c, err := processRequest(request); err != nil {
			fmt.Fprintf(writer, anError, err)
		} else {
			eq := solveEquation(a, b, c)
			fmt.Fprint(writer, formatEquation(eq))
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (a, b, c float64, err error) {
	getField := func(field string) (float64, error) {
		vals, present := request.Form[field]
		if !present {
			return 0, nil
		} else {
			return strconv.ParseFloat(vals[0], 64)
		}
	}

	if a, err = getField("a"); err != nil {
		return 0, 0, 0, errors.New("Coefficient a is invalid")
	}
	if b, err = getField("b"); err != nil {
		return 0, 0, 0, errors.New("Coefficient b is invalid")
	}
	if c, err = getField("c"); err != nil {
		return 0, 0, 0, errors.New("Coefficient c is invalid")
	}

	if a == 0 && b == 0 && c == 0 {
		return a, b, c, errors.New("Enter coefficients")
	} else {
		return a, b, c, nil
	}
}

func formatEquation(eq equation) string {
	signed := func(x float64) string {
		if math.Signbit(x) {
			return fmt.Sprintf("- %v", math.Abs(x))
		} else {
			return fmt.Sprintf("+ %v", x)
		}
	}
	return fmt.Sprintf(result, eq.a, signed(eq.b), signed(eq.c),
		eq.root1, eq.root2)
}

func solveEquation(a, b, c float64) (eq equation) {
	eq.a = a
	eq.b = b
	eq.c = c
	disc := complex(b*b-4*a*c, 0)
	eq.root1 = (complex(-b, 0) + cmplx.Sqrt(disc)) / 2 / complex(a, 0)
	eq.root2 = (complex(-b, 0) - cmplx.Sqrt(disc)) / 2 / complex(a, 0)
	return eq
}
