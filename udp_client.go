package main

import(
"net"
	"fmt"
	"bytes"
	"log"
	"time"
)

func main()  {
	ipadress:="127.0.0.1"
	network_type:="udp"
	port:="5002"
	localadress,error:=net.ResolveUDPAddr("udp",ipadress+":"+"0")
	if error!=nil{
		log.Fatal(error)

	}

	serverAdress,error1:=net.ResolveUDPAddr("udp",ipadress+":"+port )
	if error1!=nil{
		log.Fatal(error1)

	}

	connection,error2:=net.DialUDP(network_type,localadress,serverAdress)
	if error2!=nil{
		log.Fatal(error2)

	}


	defer  connection.Close()

	go sendMessagetoServer(connection)
	time.Sleep(1*time.Second)





}
func sendMessagetoServer(con net.Conn)  {

	buf:=[]byte("Hello this is client 2")
	_,errorbuf:=con.Write(buf)
	if errorbuf!=nil{
		log.Fatal(errorbuf)

	}

 for {
	 bufferreader := make([]byte, 1024)

	 _, buffError := con.Read(bufferreader)

	 if buffError != nil {
		 log.Fatal(buffError)

	 }

	 n := bytes.Index(bufferreader, []byte{0})
	 message := string(bufferreader[0:n-1])
	 fmt.Println(message)

      }
}

