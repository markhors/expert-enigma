package main

import (
	"log"
	"bytes"
	"fmt"
	"io"
	"sync"
	"net"
)
var routinecount sync.WaitGroup
var waitForgoroutine sync.WaitGroup
var message string

func main()  {
	waitForgoroutine.Add(1)
	go deployTcpServer()
	deployUdpserver()
	waitForgoroutine.Wait()

}

func deployTcpServer()  {

	defer waitForgoroutine.Done()
	// listens for connection
	listens,error:=net.Listen("tcp","localhost:5002")
	fmt.Println("Server is listening: ",listens.Addr())
	if error!=nil{
		log.Fatal(error)
	}
	// listens for connection
	for{
		connection,error:=listens.Accept()
		if error!=nil {
			log.Fatal(error)
			continue
		}
		fmt.Println(connection.RemoteAddr()," connected to TCP server")





		go replyTCPClient(connection)

	}


}



func replyTCPClient(con net.Conn)  {
	defer con.Close()

	// read is used to read the data from connection stream
	//we are making a buffer to read data from client for this we have a buffer which read byte by byte
	buf:=make([]byte,1024)
	_, err := con.Read(buf)
	if err!=nil{
		log.Fatal(err)

	}

	n := bytes.Index(buf, []byte{0})
	mesage:=string(buf[:n-1]) //
	fmt.Println(con.RemoteAddr()," says: ",mesage)
	// writing reply to client
	_,error1:=io.WriteString(con,"Hey i am server you said: "+mesage)
	if error1!=nil{
		log.Fatal(error1)
		return
	}

	fmt.Println("Closing: ",con.RemoteAddr())
	fmt.Println("TCP Server Listening at: ",con.LocalAddr())



}









func deployUdpserver()  {


	addr := net.UDPAddr{
		Port: 5002,
		IP:   net.ParseIP("127.0.0.1"),
	}
	fmt.Println("Waiting and Listen on ", addr.Port)
	conect, doterror := net.ListenUDP("udp", &addr)
	if doterror != nil {
		log.Fatal(doterror)
	}
	routinecount.Add(1)
	go readfromUdpclient(conect)
	routinecount.Wait()
	deployUdpserver()

}

func readfromUdpclient(udp_connection *net.UDPConn){

	defer routinecount.Done()
	bufferreader:=make([]byte,1024)
	_,adres,erri:=udp_connection.ReadFromUDP(bufferreader)
	if erri!=nil{
		log.Fatal(erri)
	}
	fmt.Println(adres, "is connected to UDP server ")
	n:=bytes.Index(bufferreader,[]byte{0})
	message=string(bufferreader[0:n-1])
	fmt.Println(adres, " is requesting ",message)
	writetoUDPClient(udp_connection,adres)
}

func writetoUDPClient(con *net.UDPConn, adres *net.UDPAddr) {

	m:="respond.html"
	fmt.Println("sending",adres.IP,":",m)


	_, err := con.WriteToUDP([]byte(m), adres)
	if err != nil {
		log.Fatal(err)
	}
	con.Close()

}