Compiler := go
GoFlags :=
Creators := creators


main.exe : main.go map.go map_player.go map_object.go character_object.go clue_object.go color.go gamestart.go gamefinal.go
	$(Compiler) $(GoFlags) build -o $@

character_creator.exe : $(Creators)/character_creator.go
	$(Compiler) $(GoFlags) build -o $@ $^

clue_creator.exe : $(Creators)/clue_creator.go
	$(Compiler) $(GoFlags) build -o $@ $^

clean :
	rm -f ./*.exe
