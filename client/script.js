var element = document.querySelector('#result')

const date = getDate()
//let xhr = new XMLHttpRequest()

// open a connection
// xhr.open("POST", 'http://localhost:5050/check', true);
  
// Set the request header i.e. which type of content you are sending
//xhr.setRequestHeader("Content-Type", "application/json");

// Converting JSON data to string
const data = JSON.stringify({currentDate:date});

/* // Sending data with the request
xhr.send(data);

// response handler
xhr.onload = function(){
    console.log(`Received status ${xhr.status} and response ${xhr.response} from server`)
    element.innerHTML = xhr.responseText
} */
fetch("/api", {
    method: "POST", 
    body: data
  }).then(res => {
    return res.text()
  }).then(function(data){
      console.log(data)
      let response = JSON.parse(data)
      element.innerHTML=response.message
  })

function getDate(){
    let yourDate = new Date()
    yourDate.toISOString().split('T')[0]
    const offset = yourDate.getTimezoneOffset()
    yourDate = new Date(yourDate.getTime() - (offset*60*1000))
    return yourDate.toISOString().split('T')[0]
}