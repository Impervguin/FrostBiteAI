Compiler := go
GoFlags :=


main.exe : main.go map.go map_player.go map_object.go
	$(Compiler) $(GoFlags) build  -o $@