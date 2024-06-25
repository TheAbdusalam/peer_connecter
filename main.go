package main

import (
	"fmt"
	"net"
)

type Client struct {
	ID         string
	Connection net.Conn
}

type Server struct {
	Clients []Client
}

func (s *Server) start() {
	conn, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	print("Server running at :3000\n")
	for {
		client, err := conn.Accept()
		if err != nil {
			panic(err)
		}

		c := Client{
			ID:         client.RemoteAddr().String(),
			Connection: client,
		}

		s.Clients = append(s.Clients, c)

		// list of all clients
		for _, m := range s.Clients {
			fmt.Println("This is", m.ID)
		}
		
		go c.Attach(s)
	}
}

func (c *Client) Attach(s *Server) {
	for {
		// add the client ID to the message
		data := make([]byte, 1024)
		_, err := c.Connection.Read(data)

		c.Connection.Write([]byte(c.ID + ": "))
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}


		for _, conn := range s.Clients {
			if conn.ID != c.ID {
				conn.Connection.Write([]byte(c.ID))
				conn.Connection.Write([]byte(": "))
				conn.Connection.Write([]byte(data))
			}
		}
	}
}

func main() {
	s := &Server{}
	s.start()

	// make the clients listen to 172.16.0.255:3000
	s.start()
}
