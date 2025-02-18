# Definición del Problema: Fusionar Intervalos

## Descripción:

Dada una matriz de intervalos donde cada intervalo es un par de enteros,
tu tarea es fusionar todos los intervalos que se solapan y devolver una
matriz de los intervalos que no se solapan y que cubren todos los intervalos en la entrada.

Cada intervalo en la matriz de entrada se denota como un par [inicio, fin],
donde inicio <= fin. Un intervalo [a, b] se solapa con [c, d] si y solo si c <= b.
Después de fusionar todos los intervalos que se solapan, devuelve la lista de intervalos
fusionados en orden ascendente por sus tiempos de inicio.

Entrada:

- intervalos: una lista de intervalos, donde cada intervalo está representado como una lista [inicio, fin].

Salida:

- Devuelve una lista de los intervalos fusionados, ordenados por los tiempos de inicio de los intervalos.

Ejemplo 1:

Entrada: intervalos = [[8,10],[1,3],[2,6],[15,18]]
Salida: [[1,6],[8,10],[15,18]]
Explicación: Dado que los intervalos [1,3] y [2,6] se solapan, se fusionan en [1,6]. Los otros intervalos no se solapan, por lo que permanecen como están.

Ejemplo 2:

Entrada: intervalos = [[1,4],[4,5]]
Salida: [[1,5]]
Explicación: Los intervalos [1,4] y [4,5] se consideran solapados ya que 4 == 4, por lo que se fusionan en [1,5].