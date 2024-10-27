# Golang Modules

Slide : https://docs.google.com/presentation/d/1nzeNb7w_F4_yA-9C95EYCrggObhQWt9WronMuhIG5LE/edit?usp=sharing

Source Code : 

- https://github.com/ProgrammerZamanNow/go-say-hello

- https://github.com/ProgrammerZamanNow/app-say-hello

# Cara Membuat Module

```bash
$ go mod init github.com/fdjrr/nama-module
$ git tag v1.0.0
$ git push origin v1.0.0
$ git tag -d v1.0.10 # misal mau delete tag di local
$ curl https://sum.golang.org/lookup/github.com/fdjrr/go-say-hello@v1.0.0
```

> Note : Jika ingin melakukan upgrade module, buat tag baru.

# Cara Upgrade Depedency

Ubah di file `go.mod` misal `require github.com/fdjrr/go-say-hello v1.0.2` menjadi `require github.com/fdjrr/go-say-hello v1.0.3` lalu jalankan command

```bash
$ go get
```

# Backward Compatible (Module)

Ubah di file `go.mod` misal `module github.com/fdjrr/go-say-hello` menjadi `module github.com/fdjrr/go-say-hello/v2`