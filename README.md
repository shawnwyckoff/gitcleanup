# gitcleanup
detect &amp; clean executable programs for your git repo

# install

`go get -u github.com/smcduck/gitcleanup`

`go install github.com/smcduck/gitcleanup`

# usage

`gitcleanup {which-dir-to-detect}`

Executable program files like Windows .EXE / .DLL, Linux/Unix Program / .so / .lib. Shell scripts are NOT the targets to cleanup.

There will be a yes/no option if executable program files detected, select "yes" if you really want that.
