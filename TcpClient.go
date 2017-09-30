package main
import (
	"log"
	"net"
	"os"
	"io"
	"time"

)

func main() {
	estab,error:=net.Dial("tcp","127.0.0.1:5002")
	if error!=nil {
		log.Fatal(error)


	}

	defer estab.Close()
	witeToeServer(estab);
	copyMessage(os.Stdout,estab)


}
func witeToeServer(connect net.Conn)  {
	_,error:=io.WriteString(connect,"Client sent it\n")
	if error != nil {
		log.Fatal(error)
	}



}

func copyMessage(des io.Writer,src io.Reader)  {

	_,error:=io.Copy(des,src)
	if error!=nil{
		log.Fatal(error)
	}
	time.Sleep(1*time.Second)

}