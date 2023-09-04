package main

import (
	pb "clientrpc/dummy/dummypb" // Importa el paquete generado a partir del archivo .proto
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establece la conexión con el servidor gRPC
	cred := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:50051", cred)
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()
	client := pb.NewDummyServiceClient(conn)

	// Llama al método Hello del servicio
	request := &pb.DummyRequest{Name: "Hola, servidor gRPC"}
	response, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("Error al llamar al método Hello: %v", err)
	}

	// Imprime la respuesta del servidor
	fmt.Printf("Respuesta del servidor: %s\n", response.Name)
}
