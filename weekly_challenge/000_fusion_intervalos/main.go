// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

/*
	Definición del Problema: Fusionar Intervalos

	Descripción:

	Dada una matriz de intervalos donde cada intervalo es un par de enteros,
	tu tarea es fusionar todos los intervalos que se solapan y devolver una
	matriz de los intervalos que no se solapan y que cubren todos los intervalos en la entrada.

	Cada intervalo en la matriz de entrada se denota como un par [inicio, fin],
	donde inicio <= fin. Un intervalo [a, b] se solapa con [c, d] si y solo si c <= b.
	Después de fusionar todos los intervalos que se solapan, devuelve la lista de intervalos
	fusionados en orden ascendente por sus tiempos de inicio.

	Entrada:

	    intervalos: una lista de intervalos, donde cada intervalo está representado como una lista [inicio, fin].

	Salida:

	    Devuelve una lista de los intervalos fusionados, ordenados por los tiempos de inicio de los intervalos.

	Ejemplo 1:

	    Entrada: intervalos = [[1,3],[2,6],[8,10],[15,18]]
	    Salida: [[1,6],[8,10],[15,18]]
	    Explicación: Dado que los intervalos [1,3] y [2,6] se solapan, se fusionan en [1,6]. Los otros intervalos no se solapan, por lo que permanecen como están.

	Ejemplo 2:

	    Entrada: intervalos = [[1,4],[4,5]]
	    Salida: [[1,5]]
	    Explicación: Los intervalos [1,4] y [4,5] se consideran solapados ya que 4 == 4, por lo que se fusionan en [1,5].

	Restricciones:

	    1 <= intervalos.length <= 10^4
	    intervalos[i].length == 2
	    0 <= intervalos[i][0] <= intervalos[i][1] <= 10^4
*/

func main() {
	mergedIntervals := mergeIntervals([][]int{
		{8, 10}, // intervalo
		{1, 7},  //
		{2, 3},  // inicial
		{15, 18}})
	fmt.Println(mergedIntervals) // [1,7], [8,10], [15,18]
	// [1, 7] = 1, 2, 3, 4, 5, 6, 7
	// [2, 3] =    2, 3
	// [1, 7]
}

func mergeIntervals(intervalos [][]int) [][]int {
	// Haz tu magia aquí!

	// ordenar intervalos de forma creciente en base a
	// su primer elemento
	sort.Slice(intervalos, func(i, j int) bool {
		return intervalos[i][0] < intervalos[j][0]
	})

	inicial := intervalos[0] // [1, 3]

	/*
		{1, 3},  // inicial,
		{2, 6},  // intervalo // [1, 6] *catch
		{8, 10}, //
		{15, 18}})
	*/

	var mezclados [][]int
	for _, intervalo := range intervalos[1:] {
		// intervalo [2, 6]
		if intervalo[0] <= inicial[1] {
			// traslape, hay que mezclar
			// inicial = [1, 7]
			// intevalo = [2,3]
			if intervalo[1] > inicial[1] {
				inicial[1] = intervalo[1]
			}
		} else {
			mezclados = append(mezclados, inicial)
			inicial = intervalo
		}
	}

	mezclados = append(mezclados, inicial)

	return mezclados
}
