# multix

A fast, minimal Linux-style CLI tool to generate multiplication tables directly from your terminal — built with Go.

Works like native commands such as `cal`.

---

## ✨ Features

* Instant multiplication tables
* Clean formatted output
* Lightweight single binary
* Cross-platform (Linux / macOS / Windows)
* Installable via `go install`

---

## 📦 Installation

Make sure you have Go installed (1.18+ recommended).

```bash
go install github.com/godlykmathews/multi@latest
```

Ensure Go bin is in PATH:

```bash
export PATH=$PATH:~/go/bin
```

---

## 🚀 Usage

```bash
multi 5
```

### Example output

```
 1 x  5 =   5
 2 x  5 =  10
 3 x  5 =  15
 4 x  5 =  20
 5 x  5 =  25
 6 x  5 =  30
 7 x  5 =  35
 8 x  5 =  40
 9 x  5 =  45
10 x  5 =  50
```

---

## 🛠️ Build Locally

```bash
git clone https://github.com/<your-username>/multi.git
cd multi
go build -o multi
./multi 7
```

---

## 📂 Project Structure

```
multi/
 ├── go.mod
 ├── main.go
 └── README.md
```

---

## 🎯 Roadmap

* [ ] Range tables (`multi 2 10`)
* [ ] Reverse tables
* [ ] CSV / file export
* [ ] Colorized output
* [ ] Man page support

---

## 🤝 Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you’d like to improve.

---

## 📜 License

MIT License — free to use, modify, and distribute.

---

## 👨‍💻 Author

Built with Go to learn CLI tooling and Linux command workflows.
