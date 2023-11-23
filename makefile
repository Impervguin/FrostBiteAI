Compiler := go
GoFlags :=
Creators := creators


main.exe : main.go map.go map_player.go map_object.go
	$(Compiler) $(GoFlags) build -o $@

character_creator.exe : $(Creators)/character_creator.go
	$(Compiler) $(GoFlags) build -o $@ $^