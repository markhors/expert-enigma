package main

import (
	"net/http"

	"fmt"
)




func main() {
	fmt.Println("staring server")


	http.HandleFunc("/", page )
	http.ListenAndServe(":5002",nil)

}
func page(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">

    <link href="../css/contact.css" type="text/css" rel="stylesheet"/>
    <link href="../css/backgrounds.css" type="text/css" rel="stylesheet"/>
    <title>FAQ</title>
</head>
<body class="rfq_background" bgcolor="313131">

<table   >

    <tr bgcolor="#66ccff"  >
        <td >
            <h1><a href="homePgae.xhtml"  >Muhammad Talha</a></h1>
            <h2>Student | Programmer | Network Developer </h2>
        </td>
        <td width="65%">
            <table  >
                <tr>
                    <td class="menuborder"></td>

                    <td class="menuborder"><a href="homePgae.xhtml"  >Home</a></td>

                    <td  class="menuborder"><a href="project.xhtml">Projects</a></td>


                    <td class="menuborder"><a href="contact.xhtml">Contact Us</a></td>
                    <td class="menuborder"><a href="RFQ.html" >RFQ</a></td>

                </tr>
            </table>

        </td>

    </tr>







</table>



<div class="maindiv bodydivpositioning" id="divElement"  >
    <br>

        Name:<br/>
        <input type="text" name="username" id="#" placeholder="First Name" class=""   required="required"  >
        <input type="text" name="username" id="." placeholder="Last Name" class=""  required  >

<br/>
    <br/>

    <input type="email" name="username" id="" placeholder="Email" class=""   required="required"  >



<br/><br/>

    Payment Method:  <br/><br/>
    <input type="radio" name="gender" class="" value="Male" checked="checked" />Master card

    <input type="radio" name="gender" class="" value="Male" />Paypal




<br/><br/>

    How can i help you?<br/><br/>
    <select class="clrbacko">
        <option class="clrbacko">Web Projects</option>
        <option class="clrbacko">Python Projects</option>
        <option class="clrbacko">Golang Projects</option>
        <option class="clrbacko">Block chain network implementation</option>


    </select>

<br/><br/>

    <textarea name="comments" cols="40" rows="10" class="clrbacko">Project Description </textarea>




    <br/><br/>


    <input type="button" name="submit" value="Sign Up" class="clrbacko"/>
    <br/>
    <br/>








    </form>


</div>

</body>
</html>`)

}