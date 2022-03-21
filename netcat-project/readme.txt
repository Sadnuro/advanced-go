chat.go | chat.exe:
Es el servidor principal del aplicativo

netcat.go | netcat.exe:
Es el aplicativo cliente, es el chat donde se visualizan los mensajes
y se genera la interaccion

Construir el proyecto:
1. go build net/netcat.go
2. go build net/chat.go

Ejecutar app:
1. ./chat
2. por cada cliente: ./netcat