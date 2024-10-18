Go Workspace Structure :

1) pkg (contains dependencies,package managers, go mod)
2) src (contains main files,applications)
3) bin (contains programs binaries)

A Go workspace consists of three directories: src (source code), pkg (compiled package objects), and bin (compiled executable programs).Code is organized into packages inside the src directory.


Commands : 

1) go --help (to get all go commands)
2) bug  (start a bug report)
3) build (compile packages and dependencies)
4) clean (remove object files and cached files)
5) doc (show documentation for package or symbol)
6) env (print Go environment information)
7) fix (update packages to use new APIs)
8) fmt (gofmt (reformat) package sources)
9) generate (generate Go files by processing source)
10) get (add dependencies to current module and install them)
11) install (compile and install packages and dependencies)
12) list (list packages or modules)
13) mod (module maintenance)
14) work (workspace maintenance)
15) run (compile and run Go program)
16) telemetry (manage telemetry data and settings)
17) test (test packages)
18) tool (run specified go tool)
19) version (print Go version)
20) vet (report likely mistakes in packages)

(go.mod contains used third-party libraries , not a core liberies which used in program / application)


