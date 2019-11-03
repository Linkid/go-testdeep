---
title: "Code"
weight: 10
---

```go
func Code(fn interface{}) TestDeep
```

[`Code`]({{< ref "Code" >}}) operator allows to check data using a custom function. So
*fn* is a function that must take one parameter whose type must be
the same as the type of the compared value.

*fn* can return a single `bool` kind value, telling that yes or no
the custom test is successful:
```go
Cmp(t, gotTime,
  Code(func (date time.Time) bool {
      return date.Year() == 2018
    }))
```

or two values (`bool`, `string`) kinds. The `bool` value has the same
meaning as above, and the `string` value is used to describe the
test when it fails:
```go
Cmp(t, gotTime,
  Code(func (date time.Time) (bool, string) {
      if date.Year() == 2018 {
        return true, ""
      }
      return false, "year must be 2018"
    }))
```

or a single [`error`](https://golang.org/pkg/builtin/#error) value. If the returned [`error`](https://golang.org/pkg/builtin/#error) is `nil`, the test
succeeded, else the [`error`](https://golang.org/pkg/builtin/#error) contains the reason of failure:
```go
Cmp(t, gotJsonRawMesg,
  Code(func (b json.RawMessage) error {
      var c map[string]int
      err := json.Unmarshal(b, &c)
      if err != nil {
        return err
      }
      if c["test"] != 42 {
        return fmt.Errorf(`key "test" does not match 42`)
      }
      return nil
    }))
```

This operator allows to handle any specific comparison not handled
by standard operators.

It is not recommended to call Cmp (or any other Cmp*
functions or *T methods) inside the body of *fn*, because of
confusion produced by output in case of failure. When the data
needs to be transformed before being compared again, [`Smuggle`]({{< ref "Smuggle" >}})
operator should be used instead.

[TypeBehind]({{< ref "operators#typebehind-method" >}}) method returns the [`reflect.Type`](https://golang.org/pkg/reflect/#Type) of only parameter of *fn*.


> See also [<i class='fas fa-book'></i> Code godoc](https://godoc.org/github.com/maxatome/go-testdeep#Code).

### Examples

{{%expand "Base example" %}}```go
	t := &testing.T{}

	got := "12"

	ok := Cmp(t, got,
		Code(func(num string) bool {
			n, err := strconv.Atoi(num)
			return err == nil && n > 10 && n < 100
		}),
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason
	ok = Cmp(t, got,
		Code(func(num string) (bool, string) {
			n, err := strconv.Atoi(num)
			if err != nil {
				return false, "not a number"
			}
			if n > 10 && n < 100 {
				return true, ""
			}
			return false, "not in ]10 .. 100["
		}),
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason thanks to error
	ok = Cmp(t, got,
		Code(func(num string) error {
			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}
			if n > 10 && n < 100 {
				return nil
			}
			return fmt.Errorf("%d not in ]10 .. 100[", n)
		}),
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true

```{{% /expand%}}
## CmpCode shortcut

```go
func CmpCode(t TestingT, got interface{}, fn interface{}, args ...interface{}) bool
```

CmpCode is a shortcut for:

```go
Cmp(t, got, Code(fn), args...)
```

See above for details.

Returns true if the test is OK, false if it fails.

*args...* are optional and allow to name the test. This name is
used in case of failure to qualify the test. If `len(args) > 1` and
the first item of *args* is a `string` and contains a '%' `rune` then
[`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) is used to compose the name, else *args* are passed to
[`fmt.Fprint`](https://golang.org/pkg/fmt/#Fprint). Do not forget it is the name of the test, not the
reason of a potential failure.


> See also [<i class='fas fa-book'></i> CmpCode godoc](https://godoc.org/github.com/maxatome/go-testdeep#CmpCode).

### Examples

{{%expand "Base example" %}}```go
	t := &testing.T{}

	got := "12"

	ok := CmpCode(t, got, func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 10 && n < 100
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason
	ok = CmpCode(t, got, func(num string) (bool, string) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return false, "not a number"
		}
		if n > 10 && n < 100 {
			return true, ""
		}
		return false, "not in ]10 .. 100["
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason thanks to error
	ok = CmpCode(t, got, func(num string) error {
		n, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		if n > 10 && n < 100 {
			return nil
		}
		return fmt.Errorf("%d not in ]10 .. 100[", n)
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true

```{{% /expand%}}
## T.Code shortcut

```go
func (t *T) Code(got interface{}, fn interface{}, args ...interface{}) bool
```

[`Code`]({{< ref "Code" >}}) is a shortcut for:

```go
t.Cmp(got, Code(fn), args...)
```

See above for details.

Returns true if the test is OK, false if it fails.

*args...* are optional and allow to name the test. This name is
used in case of failure to qualify the test. If `len(args) > 1` and
the first item of *args* is a `string` and contains a '%' `rune` then
[`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) is used to compose the name, else *args* are passed to
[`fmt.Fprint`](https://golang.org/pkg/fmt/#Fprint). Do not forget it is the name of the test, not the
reason of a potential failure.


> See also [<i class='fas fa-book'></i> T.Code godoc](https://godoc.org/github.com/maxatome/go-testdeep#T.Code).

### Examples

{{%expand "Base example" %}}```go
	t := NewT(&testing.T{})

	got := "12"

	ok := t.Code(got, func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 10 && n < 100
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason
	ok = t.Code(got, func(num string) (bool, string) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return false, "not a number"
		}
		if n > 10 && n < 100 {
			return true, ""
		}
		return false, "not in ]10 .. 100["
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Same with failure reason thanks to error
	ok = t.Code(got, func(num string) error {
		n, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		if n > 10 && n < 100 {
			return nil
		}
		return fmt.Errorf("%d not in ]10 .. 100[", n)
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true

```{{% /expand%}}