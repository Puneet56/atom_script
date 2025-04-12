# AtomScript Frontend

This is a Next.js project showcasing **AtomScript**, a programming language. The project includes an interactive REPL (Read-Eval-Print Loop) and a code editor to experiment with AtomScript.

## Features

- **Interactive REPL**: Execute AtomScript code directly in the browser.
- **Code Editor**: Write, tokenize, parse, and evaluate AtomScript code.
- **Syntax Highlighting**: Powered by Prism.js for AtomScript grammar.
- **Responsive Design**: Optimized for various screen sizes.

## Getting Started

### Environment

```bash
NEXT_PUBLIC_API_URL=<atomscript bacend url>
```

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Puneet56/atom_script.git
   cd atom_script/frontend
   ```

2. Install dependencies:
   ```bash
   yarn install
   ```

3. Run the development server:
   ```bash
   yarn dev
   ```

4. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Usage

- Navigate to the **REPL** to execute AtomScript commands.
- Use the **Code Editor** to write and evaluate AtomScript programs.

## Example Code

```atomscript
atom sodium = "Na";
atom chlorine = "Cl";

reaction getSalt(metal, halogen) {
  produce metal + halogen;
}

molecule salt = getSalt(sodium, chlorine);
puts(salt); // Output: NaCl
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

