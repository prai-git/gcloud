

NEW WAY - 

http://localhost:8000/retrive
http://localhost:8000/save?Input=abc




// OLD WAY
Post Request:  
	curl -d '{"Name":"Pankaj2", "Description": "user desc2", "Age" : "252"}' -i localhost:8000/

get Request with query parameter: 
	curl http://localhost:8000/?Name=Pankaj5&Age=252&Description=Desc4

get Request: 
	curl localhost:8000/

