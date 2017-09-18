package d7024e

import "net"
import "fmt"
import "bufio"
import "encoding/json"
import "time"

type Network struct {
}

func Listen(ip string, port int) {
  // TODO
}

func SendPingMessage(srcContact Contact, finalContact Contact) bool{
  // listen for reply
  input := make(chan string, 1)
  go getInput(input, contactToPing, sourceContact)

  for {
    select {
    case i := <-input:
      var message Message
      json.Unmarshal([]byte(i),&message)
      if(message.MessageType==RESPONSE){
        return true
      }
      case <-time.After(4000 * time.Millisecond):
        fmt.Println("timed out")
        return false
      }
    }
}

func getInput(input chan string, contactToPing Contact, sourceContact Contact) {
    for {
		messageToSend := &Message{sourceContact, PING,contactToPing.ID.String()}
		//fmt.Println("messageToSend to messageToSend server: "+messageToSend.Content )

		conn, conErr := net.Dial("tcp", contactToPing.Address)
		//fmt.Println(conErr)
		if(conErr==nil){
			//fmt.Println("Text to send: ")
			text, err := json.Marshal(messageToSend)
			if (err != nil) {
				fmt.Println("error " )
				fmt.Println(err)
			}
			//fmt.Println("Message to send server: "+string(text))

			// send to socket
			fmt.Fprintf(conn, string(text) + "\n")
			JSONmessage, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
			}
			input <- JSONmessage
		}
    }
}


func SendFindContactMessage(srcContact Contact, sendContact Contact, findContact Contact) {
  messageToSend := &Message{sourceContact, FINDCONTACT,contactToFind.ID.String()}
	//fmt.Print("messageToSend to messageToSend server: "+messageToSend.Content )

	conn, _ := net.Dial("tcp", contactToSend.Address)
	//	  fmt.Print("Text to send: ")
	  text, err := json.Marshal(messageToSend)
	  if err != nil {
		    fmt.Println("error " )
		    fmt.Println(err)
	  }
	  //fmt.Println("Message to send server: "+string(text))

	  // send to socket
	  fmt.Fprintf(conn, string(text) + "\n")
	  // listen for reply
	  JSONmessage, _ := bufio.NewReader(conn).ReadString('\n')
	  var message Message
	  json.Unmarshal([]byte(JSONmessage),&message)
	  var contacts []Contact
	  json.Unmarshal([]byte(message.Content),&contacts)
	  /*for i := range contacts {
	  	fmt.Println("Message from server " +string(i) +" : "+ contacts[i].ID.String())
	  }*/
	  return contacts
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO Sprint 2
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO Sprint 2
}
