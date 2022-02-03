# Arrays y Slices (https://go.dev/blog/slices-intro)

A diferencia de los Arrays que poseen un tamaño fijo, los Slices representan una secuencia de tamaño variable de elementos del mismo tipo, es decir poseen un tamaño dinámico. 

Los Slices se caracterizan por: 
- Tener un apuntador a un Array
- Su tamaño es dinámico
- Poseer una capacidad

 Un Slice se puede crear a través de la función **make(tipo, tamaño, capacidad)**, la cuál asigna un Array a cero y retorna un Slice que hace referencia a ese Array, sin embargo el parámetro de capacidad es opcional.

 ``` elementos := make([int], 5, 20) ``` 

 También es posible acceder a diversos índices de un slice a la vez por medio de un rango:

```
slice[1:4] //Acceder a los alementos desde 1-4
slice[:4] //Acceder a los alementos desde 0-4
slice[2:] //Acceder a los alementos desde 2-len(slice)
slice[:] //Acceder a todos los elementos
 ```

**¡NOTA IMPORTANTE!:** Si se inicializa un slice como en los ejemplos anteriores, es importante tener en cuenta que lo que se crea realmente es un apuntador al espacio de memoria del slice que se usa como información base; por tanto, si se hace un cambio sobre alguno de ellos, el otro también se verá afectado.

## función append() 
Permite añadir uno o varios elementos nuevos desde el último índice del slice. Si el slice aún tiene sus últimos índices disponibles, simplemente se añade el nuevo elemento a la cola; en cambio, si el último índice ya está utilizado, se aumenta dinámicamente la capacidad del slice al doble y se añade dicho elemento en la última posición.
