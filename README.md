# AtomScript

### Tiny Code, Big Reactions!

## Installation and running

- Make sure you have `go` installed
- Clone the repo.
- Change directory to repo and run

```sh
  go mod tidy
```

- Run the code

```sh
  go run ./main.go
```

- The app will open a repl by default
- To execute a file use `--file` flag followed by file path.

```sh
  go run ./main.go --file ./sampleCode.txt
```

## Sample code

```js
atom a = 1;
atom b = 2;
a + b;
a - b;
a * b;
a / b;

molecule element = {
  "name": "hydrogen",
  "symbol": "H",
  "atomicNumber": 1
};

element["name"];
element["symbol"];
element["atomicNumber"];

reaction getElementDetails(element) {
  puts("Name is", element["name"], "Symbol is", element["symbol"], "Atomic number is", element["atomicNumber"])
}

getElementDetails(element);

puts("Hello, world!");


atom metals = ["iron", "copper", "silver", "gold", "aluminum"];

len(metals);

metals[0];
metals[1];
metals[2];
metals[3];
metals[4];

push(metals, "platinum");

len(metals)

metals[5];

first(metals);

last(metals);

rest(metals);
```
