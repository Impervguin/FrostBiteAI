package main

import (
	"fmt"
	"os"
	"os/exec"
)

var LOGO string = `
###################################################################################################################
###################################################################################################################
##                                                                                                               ##
##   ####### ######    #####    #####    # ##### ######   ###  ###  # ##### #######             ###     ######   ##
##    ##   #  ##  ##  ### ###  ##   ##  ## ## ##  ##  ##   ##  ##  ## ## ##  ##   #            ## ##      ##     ##
##    ##      ##  ##  ##   ##  ##          ##     ##  ##    ####      ##     ##               ##   ##     ##     ##
##    ####    #####   ##   ##   #####      ##     #####      ##       ##     ####             ##   ##     ##     ##
##    ##      ## ##   ##   ##       ##     ##     ##  ##     ##       ##     ##               #######     ##     ##
##    ##      ## ##   ### ###  ##   ##     ##     ##  ##     ##       ##     ##   #           ##   ##     ##     ##
##   ####    #### ##   #####    #####     ####   ######     ####     ####   #######           ##   ##   ######   ##
##                                                                                                               ##
###################################################################################################################
###################################################################################################################
`

var START_MESSAGE string = `В холодном свете луны, над темной деревней, витает древний страх. Легенда о Лунном зеркальце, обитающем в тени лесов, ожила, когда деревня стала ареной жутких событий. Неведомый оборотень, покрытый тайной, превратил каждую ночь в кошмар.

Последний вестник беды был найден на главной площади — тихом свидетеле темных сил, поглотивших мирное население. В глазах его последнего вздоха читается ужас и страх перед нечеловеческим злом. Следы указывают на оборотня, который теперь скрывается в тени, ждущий следующей ночи, чтобы снова принести смерть.

Вы — детектив, посланный в этот отдаленный уголок мира, чтобы раскрыть тайну Лунного зеркальца и остановить кровавый след оборотня. Ваши навыки расследования и смелость будут испытаны в этой смертельной игре, где каждая ночь приносит новую угрозу, а каждый шаг ведет к разгадке, затмеваемой мраком лунного света. Время не ждет, детектив. Наступила ночь, и деревня держится на грани катастрофы.
`

func GameStart() {
	fmt.Println(LOGO);
	fmt.Println("Нажмите enter, чтобы продолжить...")
	fmt.Scanln();

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	fmt.Println(START_MESSAGE);
	fmt.Println("Нажмите enter, чтобы продолжить...")
	fmt.Scanln();
}