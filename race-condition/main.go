package main

func main() {
	// RACE CONDITION

	/*
		Este problema también es conocido como “Productor-Consumidor”. Donde existen los principales involucrados:

		Los productores: crean tareas y las ponen en un buffer compartido.
		Los consumidores: sacan las tareas que el productor puso en el buffer compartido.
		Buffer compartido: a través de este los productores y consumidores se comunican/envían tareas.
		“Algo” que coordine el acceso al buffer; mejor conocidos como semáforos o mutex. Evitan que ocurra la condición de competencia (Race condition) por acceder al buffer.
		En el ejemplo de Néstor:

		Los productores: son nuestros depósitos.
		Los consumidores: son nuestros retiros.
		El buffer: Será nuestra cuenta bancaria.
		Para evitar que algunos retiros o depósitos no se marquen o se marquen doble, necesitamos “algo” que coordine el acceso a nuestra cuenta bancaria.

	*/

}
