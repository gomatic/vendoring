# vendoring

    git clone https://github.com/gomatic/vendoring.git
    cd vendoring
    export GOPATH=${PWD}
    go build -v github.com/gomatic/main_a

## libs

- `lib_d` is "global"
- `lib_b` is a vendor and imports `lib_c` and `lib_d`
- `lib_c` is a vendor and imports `lib_d`

## mains

- `main_a` imports `lib_b` and `lib_d` to show that `main_a` and `lib_b` are accessing the same `D` from the global `lib_d`.
