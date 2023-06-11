# Useful Generic Examples

example use cases where Generics are beneficial

snippets which are taken from articles are referenced below

## Go 1.18 introduced support for Generics

- offer a powerful and simple way to create generic functions, interfaces and struct methods.

- performance and code legibility trade-offs in using generics instead of an interface as a value

- A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, which must be followed by type arguments if the type is generic.

- For a generic function, type arguments may be provided explicitly, or they may be partially or completely inferred. A generic function that is not called requires a type argument list for instantiation; if the list is partial, all remaining type arguments must be inferrable.

- A generic function that is called may provide a (possibly partial) type argument list, or may omit it entirely if the omitted type arguments are inferrable from the ordinary (non-type) function arguments.

- By convention generic parameters take the form of a single uppercase letter, such as **K**

- When declaring a generic parameter we also need to declare its [type constraints](https://go.dev/ref/spec#Type_parameter_declarations) if any

- [golang.org/x/exp/](https://pkg.go.dev/golang.org/x/exp/constraints) defines a set of useful constraints to be used with type parameters

- can also use a custom type as a constraint

> Generics add three new big things to the language:
>
> 1.  Type parameters for function and types.
> 2.  Defining interface types as sets of types, including types that don’t have methods.
> 3.  Type inference, which permits omitting type arguments in many cases when calling a function.

---

[Offical Intro to Generics Post](https://go.dev/blog/intro-generics)

### Generics avoid the drawbacks of previous approach:

- type safe

- typing information preserved

- type safety will remain in the runtime code without requiring any extra validation code

## Before 1.18

- empty `interface{}` type was used in combination with a `switch` to create a generic function.

- Use of `interface{}` type causes the loss of typing information and type safety.

### Snippet References

https://blog.streamelements.com/an-introduction-to-generics-in-go-cc8cdae15ef2?gi=61fe8543b909

https://michaelwhatcott.com/go-generics-in-practice/

https://chrisfrew.in/blog/golang-a-powerful-generic-function-to-make-http-requests/

https://bitfieldconsulting.com/golang/generic-set
