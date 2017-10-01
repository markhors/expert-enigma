package main

import (
	"log"
	"bytes"
	"fmt"

	"sync"
	"net"
)
var routinecount sync.WaitGroup
var waitForgoroutine sync.WaitGroup
var message string
var ip string="127.0.0.1"
var port int=5002

func main()  {
	waitForgoroutine.Add(1)
	go deployTcpServer()
	deployUdpserver()
	waitForgoroutine.Wait()

}

func deployTcpServer()  {

	defer waitForgoroutine.Done()



	adr:=net.TCPAddr{
		Port:port,
		IP:net.ParseIP(ip),

	}
	listens,err:=net.ListenTCP("tcp",&adr)
	fmt.Println("Server is listening: ",listens.Addr())
	inspectError(err,"Listen Tcp ")
	// listens for connection
	for{
		connection,err:=listens.Accept()
		inspectError(err,"Accept Tcp ")


		fmt.Println(connection.RemoteAddr()," connected to TCP server")





		go handleTCPClient(connection)

	}


}



func handleTCPClient(con net.Conn)  {
	defer con.Close()

	// read is used to read the data from connection stream
	//we are making a buffer to read data from client for this we have a buffer which read byte by byte
	buf:=make([]byte,1024)
	n, err := con.Read(buf)
	inspectError(err,"Tcp reading form client ")


	mesage:=string(buf[:n]) //
	fmt.Println(con.RemoteAddr()," says: ",mesage)
	// writing reply to client
	_,er:=con.Write([]byte("Hey there"))
	inspectError(er,"Tcp writing to client ")

	fmt.Println("Closing: ",con.RemoteAddr())
	fmt.Println("TCP Server Listening at: ",con.LocalAddr())



}









func deployUdpserver()  {


	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	fmt.Println("Waiting and Listen on ", addr.Port)
	for {
		conect, err := net.ListenUDP("udp", &addr)
		inspectError(err, "Listen UDP  ")
		routinecount.Add(1)
		go readfromUdpclient(conect)
		routinecount.Wait()
	}

}

func readfromUdpclient(udp_connection *net.UDPConn){

	defer routinecount.Done()
	bufferreader:=make([]byte,1024)
	_,adres,err:=udp_connection.ReadFromUDP(bufferreader)
	inspectError(err,"Read From udp client ")
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
	inspectError(err,"Write to udp  ")
	con.Close()

}


func inspectError(err error,descri string){
	if err!=nil{
	log.Fatal(descri+": ",err)
}
}